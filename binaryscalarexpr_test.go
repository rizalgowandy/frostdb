package frostdb

import (
	"testing"

	"github.com/segmentio/parquet-go"
	"github.com/stretchr/testify/require"

	"github.com/polarsignals/frostdb/query/logicalplan"
)

type FakeColumnChunk struct {
	index *FakeColumnIndex
}

func (f *FakeColumnChunk) Type() parquet.Type               { return nil }
func (f *FakeColumnChunk) Column() int                      { return 0 }
func (f *FakeColumnChunk) Pages() parquet.Pages             { return nil }
func (f *FakeColumnChunk) ColumnIndex() parquet.ColumnIndex { return f.index }
func (f *FakeColumnChunk) OffsetIndex() parquet.OffsetIndex { return nil }
func (f *FakeColumnChunk) BloomFilter() parquet.BloomFilter { return nil }
func (f *FakeColumnChunk) NumValues() int64                 { return 0 }

type FakeColumnIndex struct {
	numPages int
	min      parquet.Value
	max      parquet.Value
}

func (f *FakeColumnIndex) NumPages() int {
	return f.numPages
}
func (f *FakeColumnIndex) NullCount(int) int64        { return 0 }
func (f *FakeColumnIndex) NullPage(int) bool          { return false }
func (f *FakeColumnIndex) MinValue(int) parquet.Value { return f.min }
func (f *FakeColumnIndex) MaxValue(int) parquet.Value { return f.max }
func (f *FakeColumnIndex) IsAscending() bool          { return false }
func (f *FakeColumnIndex) IsDescending() bool         { return false }

// This is a regression test that ensures the Min/Max functions return a null
// value (instead of panicing) should they be passed a column chunk that only
// has null values.
func Test_MinMax_EmptyColumnChunk(t *testing.T) {
	fakeChunk := &FakeColumnChunk{
		index: &FakeColumnIndex{
			numPages: 10,
		},
	}

	v := Min(fakeChunk)
	require.True(t, v.IsNull())

	v = Max(fakeChunk)
	require.True(t, v.IsNull())
}

func TestBinaryScalarOperation(t *testing.T) {
	for _, tc := range []struct {
		name        string
		min         int
		max         int
		right       int
		op          logicalplan.Op
		expectFalse bool
	}{
		{
			name:  "OpEqValueContained",
			min:   1,
			max:   10,
			right: 5,
			op:    logicalplan.OpEq,
		},
		{
			name:        "OpEqValueGt",
			min:         1,
			max:         10,
			right:       11,
			op:          logicalplan.OpEq,
			expectFalse: true,
		},
		{
			name:        "OpEqValueLt",
			min:         1,
			max:         10,
			right:       0,
			op:          logicalplan.OpEq,
			expectFalse: true,
		},
		{
			name:  "OpEqMaxBound",
			min:   1,
			max:   10,
			right: 10,
			op:    logicalplan.OpEq,
		},
		{
			name:  "OpEqMinBound",
			min:   1,
			max:   10,
			right: 1,
			op:    logicalplan.OpEq,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.op == logicalplan.OpUnknown {
				t.Fatal("test programming error: remember to set operator")
			}
			fakeChunk := &FakeColumnChunk{
				index: &FakeColumnIndex{
					numPages: 1,
					min:      parquet.ValueOf(tc.min),
					max:      parquet.ValueOf(tc.max),
				},
			}
			res, err := BinaryScalarOperation(fakeChunk, parquet.ValueOf(tc.right), tc.op)
			require.NoError(t, err)
			require.Equal(t, !tc.expectFalse, res)
		})
	}
}