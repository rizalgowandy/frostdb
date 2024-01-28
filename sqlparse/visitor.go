package sqlparse

import (
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/scalar"
	"github.com/parquet-go/parquet-go"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/opcode"
	"github.com/pingcap/tidb/parser/test_driver"

	"github.com/polarsignals/frostdb/query"
	"github.com/polarsignals/frostdb/query/logicalplan"
)

type astVisitor struct {
	explain     bool
	builder     query.Builder
	dynColNames map[string]struct{}
	err         error
	schema      *parquet.Schema

	exprStack []logicalplan.Expr
	groupBy   []logicalplan.Expr
}

var _ ast.Visitor = &astVisitor{}

func newASTVisitor(
	builder query.Builder,
	dynColNames []string,
	schema *parquet.Schema,
) *astVisitor {
	dynMap := make(map[string]struct{})
	for _, n := range dynColNames {
		dynMap[n] = struct{}{}
	}
	return &astVisitor{
		builder:     builder,
		dynColNames: dynMap,
		schema:      schema,
	}
}

func (v *astVisitor) Enter(n ast.Node) (nRes ast.Node, skipChildren bool) {
	switch expr := n.(type) {
	case *ast.SelectStmt:
		// The SelectStmt is handled in during pre-visit given that it has many
		// clauses we need to handle independently (e.g. a group by with a
		// filter).
		if expr.Where != nil {
			expr.Where.Accept(v)
			lastExpr, newExprs := pop(v.exprStack)
			v.exprStack = newExprs
			v.builder = v.builder.Filter(lastExpr)
		}
		expr.Fields.Accept(v)
		switch {
		case expr.GroupBy != nil:
			// This represents everything before the "group by" clause.
			beforeGroupBy := v.exprStack
			expr.GroupBy.Accept(v)
			// This represents everything after the "group by" clause.
			afterGroupBy := v.exprStack[len(beforeGroupBy):]
			groups := afterGroupBy

			includedPreprojections := make(map[string]struct{})
			preProjections := []logicalplan.Expr{}
			postProjections := []logicalplan.Expr{}
			aggregations := []*logicalplan.AggregationFunction{}

			for _, expr := range beforeGroupBy {
				// Walk the expression tree and separate out what projections
				// need to be done before aggregations, and which can be done
				// after. We don't support nested aggregations, so we can just
				// find aggregation functions and put their expressions into
				// the pre-projections. If we wanted to support nested
				// aggregations, this would need to be more sophisticated.
				aggCollector := &aggregationCollector{}
				expr.Accept(aggCollector)

				// If we have any aggregations nested in the AST like
				// "sum(value) as value_sum_or_anything_else" then the actual
				// query plan nesting looks something like:
				// alias(value_sum_or_anything_else, sum(value))
				// sum(value) / count(value) as value_avg
				// alias(value_avg, div(sum(value), count(value)))
				// and it finds: sum(value), count(value)

				if len(aggCollector.aggregations) > 0 {
					// This is the expression that will be aggregated so we
					// need to ensure that it is in the pre-projection.
					for _, agg := range aggCollector.aggregations {
						if _, ok := includedPreprojections[agg.Expr.Name()]; !ok {
							preProjections = append(preProjections, agg.Expr)
							// The same expression can be used in multiple
							// aggregations, but we only need to project it
							// once.
							includedPreprojections[agg.Expr.Name()] = struct{}{}
						}
						if agg.Func == logicalplan.AggFuncAvg {
							// The avg function is a special case where we need
							// to project the count and sum expressions and
							// then divide them.
							//
							// Essentially this:
							//
							// avg(value)
							//
							// Becomes this:
							//
							// alias("avg(value)", div(sum(value), count(value)))

							// For avg, we need to project the count and sum
							// separately.
							count := &logicalplan.AggregationFunction{
								Func: logicalplan.AggFuncCount,
								Expr: agg.Expr,
							}

							sum := &logicalplan.AggregationFunction{
								Func: logicalplan.AggFuncSum,
								Expr: agg.Expr,
							}
							aggregations = append(aggregations, count, sum)

							t, err := agg.Expr.DataType(v.schema)
							if err != nil {
								v.err = err
								return nil, true
							}

							var typedCount logicalplan.Expr = count
							if !arrow.TypeEqual(t, arrow.PrimitiveTypes.Int64) {
								typedCount = logicalplan.Convert(count, arrow.PrimitiveTypes.Float64)
							}

							var avgProjection logicalplan.Expr = &logicalplan.BinaryExpr{
								Left:  sum,
								Op:    logicalplan.OpDiv,
								Right: typedCount,
							}
							if e, ok := expr.(*logicalplan.AliasExpr); ok {
								avgProjection = &logicalplan.AliasExpr{
									Expr:  avgProjection,
									Alias: e.Alias,
								}
							} else {
								avgProjection = &logicalplan.AliasExpr{
									Expr:  avgProjection,
									Alias: expr.String(),
								}
							}
							postProjections = append(postProjections, avgProjection)
							continue
						}

						aggregations = append(aggregations, agg)
						postProjections = append(postProjections, expr)
					}
				} else {
					preProjections = append(preProjections, expr)
					if _, ok := expr.(*logicalplan.DynamicColumn); ok {
						postProjections = append(postProjections, expr)
					} else {
						postProjections = append(postProjections, logicalplan.Col(expr.Name()))
					}
				}
			}

			// We need to ensure that anything we group by is in the pre-projection.
			for _, expr := range afterGroupBy {
				found := false
				for _, preExpr := range preProjections {
					if expr.Name() == preExpr.Name() {
						found = true
						break
					}
				}
				if !found {
					preProjections = append(preProjections, expr)
				}
			}

			// Insert a projection for any groups that need to be computed
			// before they can be used for an aggregation, for example:
			//
			// SELECT sum(value), (timestamp/1000) as timestamp_bucket group by timestamp_bucket
			v.builder = v.builder.Project(preProjections...)
			v.builder = v.builder.Aggregate(aggregations, groups)

			// Insert a projection for any groups that need to be computed
			// before they can be used for an aggregation, for example:
			//
			// SELECT sum(value)/count(value) value_avg, (timestamp/1000) as timestamp_bucket group by timestamp_bucket
			v.builder = v.builder.Project(postProjections...)
		case expr.Distinct:
			v.builder = v.builder.Distinct(v.exprStack...)
		default:
			v.builder = v.builder.Project(v.exprStack...)
		}
		return n, true
	}
	return n, false
}

type aggregationCollector struct {
	aggregations []*logicalplan.AggregationFunction
}

func (a *aggregationCollector) PreVisit(e logicalplan.Expr) bool {
	return true
}

func (a *aggregationCollector) Visit(e logicalplan.Expr) bool {
	if agg, ok := e.(*logicalplan.AggregationFunction); ok {
		a.aggregations = append(a.aggregations, agg)
	}
	return true
}

func (a *aggregationCollector) PostVisit(e logicalplan.Expr) bool {
	return true
}

func (v *astVisitor) Leave(n ast.Node) (nRes ast.Node, ok bool) {
	if err := v.leaveImpl(n); err != nil {
		v.err = err
		return n, false
	}
	return n, true
}

func (v *astVisitor) leaveImpl(n ast.Node) error {
	switch expr := n.(type) {
	case *ast.SelectStmt:
		// Handled in Enter.
		return nil
	case *ast.ExplainStmt:
		v.explain = true
		return nil
	case *ast.AggregateFuncExpr:
		// At this point, the child node is the column name, so it has just been
		// added to exprs.
		lastExpr := len(v.exprStack) - 1
		switch strings.ToLower(expr.F) {
		case "count":
			v.exprStack[lastExpr] = logicalplan.Count(v.exprStack[lastExpr])
		case "sum":
			v.exprStack[lastExpr] = logicalplan.Sum(v.exprStack[lastExpr])
		case "min":
			v.exprStack[lastExpr] = logicalplan.Min(v.exprStack[lastExpr])
		case "max":
			v.exprStack[lastExpr] = logicalplan.Max(v.exprStack[lastExpr])
		case "avg":
			v.exprStack[lastExpr] = logicalplan.Avg(v.exprStack[lastExpr])
		default:
			return fmt.Errorf("unhandled aggregate function %s", expr.F)
		}
	case *ast.BinaryOperationExpr:
		// Note that we're resolving exprs as a stack, so the last two
		// expressions are the leaf expressions.
		rightExpr, newExprs := pop(v.exprStack)
		leftExpr, newExprs := pop(newExprs)
		v.exprStack = newExprs

		var frostDBOp logicalplan.Op
		switch expr.Op {
		case opcode.GT:
			frostDBOp = logicalplan.OpGt
		case opcode.LT:
			frostDBOp = logicalplan.OpLt
		case opcode.GE:
			frostDBOp = logicalplan.OpGtEq
		case opcode.LE:
			frostDBOp = logicalplan.OpLtEq
		case opcode.EQ:
			frostDBOp = logicalplan.OpEq
		case opcode.NE:
			frostDBOp = logicalplan.OpNotEq
		case opcode.Plus:
			frostDBOp = logicalplan.OpAdd
		case opcode.Minus:
			frostDBOp = logicalplan.OpSub
		case opcode.Mul:
			frostDBOp = logicalplan.OpMul
		case opcode.Div:
			frostDBOp = logicalplan.OpDiv
		case opcode.LogicAnd:
			v.exprStack = append(v.exprStack, logicalplan.And(leftExpr, rightExpr))
			return nil
		case opcode.LogicOr:
			v.exprStack = append(v.exprStack, logicalplan.Or(leftExpr, rightExpr))
			return nil
		}
		v.exprStack = append(v.exprStack, &logicalplan.BinaryExpr{
			Left:  leftExpr,
			Op:    frostDBOp,
			Right: rightExpr,
		})
	case *ast.ColumnName:
		colName := columnNameToString(expr)
		var col logicalplan.Expr
		if _, ok := v.dynColNames[colName]; ok {
			col = logicalplan.DynCol(colName)
		} else {
			col = logicalplan.Col(colName)
		}
		v.exprStack = append(v.exprStack, col)
	case *test_driver.ValueExpr:
		switch logicalplan.Literal(expr.GetValue()).Name() { // NOTE: special case for boolean fields since the mysql parser doesn't support booleans as a type
		case "true":
			v.exprStack = append(v.exprStack, logicalplan.Literal(true))
		case "false":
			v.exprStack = append(v.exprStack, logicalplan.Literal(false))
		default:
			v.exprStack = append(v.exprStack, logicalplan.Literal(expr.GetValue()))
		}
	case *ast.SelectField:
		if as := expr.AsName.String(); as != "" {
			lastExpr := len(v.exprStack) - 1
			switch e := v.exprStack[lastExpr].(type) {
			case *logicalplan.AggregationFunction:
				v.exprStack[lastExpr] = e.Alias(as)
			case *logicalplan.BinaryExpr:
				v.exprStack[lastExpr] = e.Alias(as)
			default:
				return fmt.Errorf("unhandled select field %s", as)
			}
		}
	case *ast.PatternRegexpExpr:
		rightExpr, newExprs := pop(v.exprStack)
		leftExpr, newExprs := pop(newExprs)
		v.exprStack = newExprs

		e := &logicalplan.BinaryExpr{
			Left:  logicalplan.Col(leftExpr.Name()),
			Op:    logicalplan.OpRegexMatch,
			Right: rightExpr,
		}
		if expr.Not {
			e.Op = logicalplan.OpRegexNotMatch
		}
		v.exprStack = append(v.exprStack, e)
	case *ast.GroupByClause:
	case *ast.FieldList, *ast.ColumnNameExpr, *ast.ByItem, *ast.RowExpr,
		*ast.ParenthesesExpr:
		// Deliberate pass-through nodes.
	case *ast.FuncCallExpr:
		switch expr.FnName.String() {
		case ast.Second:
			// This is pretty hacky and only fine because it's in the test only.
			left, right := pop(v.exprStack)
			var exprStack []logicalplan.Expr
			exprStack = append(exprStack, right...)
			switch l := left.(type) {
			case *logicalplan.LiteralExpr:
				val := l.Value.(*scalar.Int64)
				duration := time.Duration(val.Value) * time.Second
				exprStack = append(exprStack, logicalplan.Duration(duration))
				v.exprStack = exprStack
			}
		default:
			return fmt.Errorf("unhandled func call: %s", expr.FnName.String())
		}
	default:
		return fmt.Errorf("unhandled ast node %T", expr)
	}
	return nil
}

func columnNameToString(c *ast.ColumnName) string {
	// Note that in SQL labels.label2 is interpreted as referencing
	// the label2 column of a table called labels. In our case,
	// these are dynamic columns, which is why the table name is
	// accessed here.
	colName := ""
	if c.Table.String() != "" {
		colName = c.Table.String() + "."
	}
	colName += c.Name.String()
	return colName
}

func pop[T any](s []T) (T, []T) {
	lastIdx := len(s) - 1
	return s[lastIdx], s[:lastIdx]
}
