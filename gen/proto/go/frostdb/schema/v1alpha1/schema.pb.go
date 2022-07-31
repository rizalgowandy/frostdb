// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: frostdb/schema/v1alpha1/schema.proto

package schemav1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type enum of a column.
type StorageLayout_Type int32

const (
	// Unknown type.
	StorageLayout_TYPE_UNKNOWN_UNSPECIFIED StorageLayout_Type = 0
	// Represents a string type.
	StorageLayout_TYPE_STRING StorageLayout_Type = 1
	// Represents an int64 type.
	StorageLayout_TYPE_INT64 StorageLayout_Type = 2
	// Represents a double type.
	StorageLayout_TYPE_DOUBLE StorageLayout_Type = 3
)

// Enum value maps for StorageLayout_Type.
var (
	StorageLayout_Type_name = map[int32]string{
		0: "TYPE_UNKNOWN_UNSPECIFIED",
		1: "TYPE_STRING",
		2: "TYPE_INT64",
		3: "TYPE_DOUBLE",
	}
	StorageLayout_Type_value = map[string]int32{
		"TYPE_UNKNOWN_UNSPECIFIED": 0,
		"TYPE_STRING":              1,
		"TYPE_INT64":               2,
		"TYPE_DOUBLE":              3,
	}
)

func (x StorageLayout_Type) Enum() *StorageLayout_Type {
	p := new(StorageLayout_Type)
	*p = x
	return p
}

func (x StorageLayout_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageLayout_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_frostdb_schema_v1alpha1_schema_proto_enumTypes[0].Descriptor()
}

func (StorageLayout_Type) Type() protoreflect.EnumType {
	return &file_frostdb_schema_v1alpha1_schema_proto_enumTypes[0]
}

func (x StorageLayout_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageLayout_Type.Descriptor instead.
func (StorageLayout_Type) EnumDescriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{2, 0}
}

// Encoding enum of a column.
type StorageLayout_Encoding int32

const (
	// Plain encoding.
	StorageLayout_ENCODING_PLAIN_UNSPECIFIED StorageLayout_Encoding = 0
	// Dictionary run-length encoding.
	StorageLayout_ENCODING_RLE_DICTIONARY StorageLayout_Encoding = 1
	// Delta binary packed encoding.
	StorageLayout_ENCODING_DELTA_BINARY_PACKED StorageLayout_Encoding = 2
	// Delta Byte Array encoding.
	StorageLayout_ENCODING_DELTA_BYTE_ARRAY StorageLayout_Encoding = 3
	// Delta Length Byte Array encoding.
	StorageLayout_ENCODING_DELTA_LENGTH_BYTE_ARRAY StorageLayout_Encoding = 4
)

// Enum value maps for StorageLayout_Encoding.
var (
	StorageLayout_Encoding_name = map[int32]string{
		0: "ENCODING_PLAIN_UNSPECIFIED",
		1: "ENCODING_RLE_DICTIONARY",
		2: "ENCODING_DELTA_BINARY_PACKED",
		3: "ENCODING_DELTA_BYTE_ARRAY",
		4: "ENCODING_DELTA_LENGTH_BYTE_ARRAY",
	}
	StorageLayout_Encoding_value = map[string]int32{
		"ENCODING_PLAIN_UNSPECIFIED":       0,
		"ENCODING_RLE_DICTIONARY":          1,
		"ENCODING_DELTA_BINARY_PACKED":     2,
		"ENCODING_DELTA_BYTE_ARRAY":        3,
		"ENCODING_DELTA_LENGTH_BYTE_ARRAY": 4,
	}
)

func (x StorageLayout_Encoding) Enum() *StorageLayout_Encoding {
	p := new(StorageLayout_Encoding)
	*p = x
	return p
}

func (x StorageLayout_Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageLayout_Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_frostdb_schema_v1alpha1_schema_proto_enumTypes[1].Descriptor()
}

func (StorageLayout_Encoding) Type() protoreflect.EnumType {
	return &file_frostdb_schema_v1alpha1_schema_proto_enumTypes[1]
}

func (x StorageLayout_Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageLayout_Encoding.Descriptor instead.
func (StorageLayout_Encoding) EnumDescriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{2, 1}
}

// Compression enum of a column.
type StorageLayout_Compression int32

const (
	// No compression.
	StorageLayout_COMPRESSION_NONE_UNSPECIFIED StorageLayout_Compression = 0
	// Snappy compression.
	StorageLayout_COMPRESSION_SNAPPY StorageLayout_Compression = 1
	// GZIP compression.
	StorageLayout_COMPRESSION_GZIP StorageLayout_Compression = 2
	// Brotli compression.
	StorageLayout_COMPRESSION_BROTLI StorageLayout_Compression = 3
	// LZ4_RAW compression.
	StorageLayout_COMPRESSION_LZ4_RAW StorageLayout_Compression = 4
	// ZSTD compression.
	StorageLayout_COMPRESSION_ZSTD StorageLayout_Compression = 5
)

// Enum value maps for StorageLayout_Compression.
var (
	StorageLayout_Compression_name = map[int32]string{
		0: "COMPRESSION_NONE_UNSPECIFIED",
		1: "COMPRESSION_SNAPPY",
		2: "COMPRESSION_GZIP",
		3: "COMPRESSION_BROTLI",
		4: "COMPRESSION_LZ4_RAW",
		5: "COMPRESSION_ZSTD",
	}
	StorageLayout_Compression_value = map[string]int32{
		"COMPRESSION_NONE_UNSPECIFIED": 0,
		"COMPRESSION_SNAPPY":           1,
		"COMPRESSION_GZIP":             2,
		"COMPRESSION_BROTLI":           3,
		"COMPRESSION_LZ4_RAW":          4,
		"COMPRESSION_ZSTD":             5,
	}
)

func (x StorageLayout_Compression) Enum() *StorageLayout_Compression {
	p := new(StorageLayout_Compression)
	*p = x
	return p
}

func (x StorageLayout_Compression) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageLayout_Compression) Descriptor() protoreflect.EnumDescriptor {
	return file_frostdb_schema_v1alpha1_schema_proto_enumTypes[2].Descriptor()
}

func (StorageLayout_Compression) Type() protoreflect.EnumType {
	return &file_frostdb_schema_v1alpha1_schema_proto_enumTypes[2]
}

func (x StorageLayout_Compression) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageLayout_Compression.Descriptor instead.
func (StorageLayout_Compression) EnumDescriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{2, 2}
}

// Enum of possible sorting directions.
type SortingColumn_Direction int32

const (
	// Unknown direction.
	SortingColumn_DIRECTION_UNKNOWN_UNSPECIFIED SortingColumn_Direction = 0
	// Sort in ascending order.
	SortingColumn_DIRECTION_ASCENDING SortingColumn_Direction = 1
	// Sort in descending order.
	SortingColumn_DIRECTION_DESCENDING SortingColumn_Direction = 2
)

// Enum value maps for SortingColumn_Direction.
var (
	SortingColumn_Direction_name = map[int32]string{
		0: "DIRECTION_UNKNOWN_UNSPECIFIED",
		1: "DIRECTION_ASCENDING",
		2: "DIRECTION_DESCENDING",
	}
	SortingColumn_Direction_value = map[string]int32{
		"DIRECTION_UNKNOWN_UNSPECIFIED": 0,
		"DIRECTION_ASCENDING":           1,
		"DIRECTION_DESCENDING":          2,
	}
)

func (x SortingColumn_Direction) Enum() *SortingColumn_Direction {
	p := new(SortingColumn_Direction)
	*p = x
	return p
}

func (x SortingColumn_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingColumn_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_frostdb_schema_v1alpha1_schema_proto_enumTypes[3].Descriptor()
}

func (SortingColumn_Direction) Type() protoreflect.EnumType {
	return &file_frostdb_schema_v1alpha1_schema_proto_enumTypes[3]
}

func (x SortingColumn_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingColumn_Direction.Descriptor instead.
func (SortingColumn_Direction) EnumDescriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{3, 0}
}

// Schema definition for a table.
type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the schema.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Columns in the schema.
	Columns []*Column `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	// Columns to sort by in the schema.
	SortingColumns []*SortingColumn `protobuf:"bytes,3,rep,name=sorting_columns,json=sortingColumns,proto3" json:"sorting_columns,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Schema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Schema) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Schema) GetSortingColumns() []*SortingColumn {
	if x != nil {
		return x.SortingColumns
	}
	return nil
}

// Column definition.
type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the column.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Storage layout of the column.
	StorageLayout *StorageLayout `protobuf:"bytes,2,opt,name=storage_layout,json=storageLayout,proto3" json:"storage_layout,omitempty"`
	// Whether the column can dynamically expand.
	Dynamic bool `protobuf:"varint,3,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetStorageLayout() *StorageLayout {
	if x != nil {
		return x.StorageLayout
	}
	return nil
}

func (x *Column) GetDynamic() bool {
	if x != nil {
		return x.Dynamic
	}
	return false
}

// Storage layout describes the physical storage properties of a column.
type StorageLayout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the column.
	Type StorageLayout_Type `protobuf:"varint,1,opt,name=type,proto3,enum=frostdb.schema.v1alpha1.StorageLayout_Type" json:"type,omitempty"`
	// Encoding of the column.
	Encoding StorageLayout_Encoding `protobuf:"varint,2,opt,name=encoding,proto3,enum=frostdb.schema.v1alpha1.StorageLayout_Encoding" json:"encoding,omitempty"`
	// Compression of the column.
	Compression StorageLayout_Compression `protobuf:"varint,3,opt,name=compression,proto3,enum=frostdb.schema.v1alpha1.StorageLayout_Compression" json:"compression,omitempty"`
	// Wether values in the column are allowed to be null.
	Nullable bool `protobuf:"varint,4,opt,name=nullable,proto3" json:"nullable,omitempty"`
}

func (x *StorageLayout) Reset() {
	*x = StorageLayout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageLayout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageLayout) ProtoMessage() {}

func (x *StorageLayout) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageLayout.ProtoReflect.Descriptor instead.
func (*StorageLayout) Descriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{2}
}

func (x *StorageLayout) GetType() StorageLayout_Type {
	if x != nil {
		return x.Type
	}
	return StorageLayout_TYPE_UNKNOWN_UNSPECIFIED
}

func (x *StorageLayout) GetEncoding() StorageLayout_Encoding {
	if x != nil {
		return x.Encoding
	}
	return StorageLayout_ENCODING_PLAIN_UNSPECIFIED
}

func (x *StorageLayout) GetCompression() StorageLayout_Compression {
	if x != nil {
		return x.Compression
	}
	return StorageLayout_COMPRESSION_NONE_UNSPECIFIED
}

func (x *StorageLayout) GetNullable() bool {
	if x != nil {
		return x.Nullable
	}
	return false
}

// SortingColumn definition.
type SortingColumn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the column to sort by.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Direction of the sorting.
	Direction SortingColumn_Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=frostdb.schema.v1alpha1.SortingColumn_Direction" json:"direction,omitempty"`
	// Whether nulls are the smallest or largest values.
	NullsFirst bool `protobuf:"varint,3,opt,name=nulls_first,json=nullsFirst,proto3" json:"nulls_first,omitempty"`
}

func (x *SortingColumn) Reset() {
	*x = SortingColumn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortingColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortingColumn) ProtoMessage() {}

func (x *SortingColumn) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_schema_v1alpha1_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortingColumn.ProtoReflect.Descriptor instead.
func (*SortingColumn) Descriptor() ([]byte, []int) {
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP(), []int{3}
}

func (x *SortingColumn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SortingColumn) GetDirection() SortingColumn_Direction {
	if x != nil {
		return x.Direction
	}
	return SortingColumn_DIRECTION_UNKNOWN_UNSPECIFIED
}

func (x *SortingColumn) GetNullsFirst() bool {
	if x != nil {
		return x.NullsFirst
	}
	return false
}

var File_frostdb_schema_v1alpha1_schema_proto protoreflect.FileDescriptor

var file_frostdb_schema_v1alpha1_schema_proto_rawDesc = []byte{
	0x0a, 0x24, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22,
	0xa8, 0x01, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x0e, 0x73, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x06, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x22, 0xbf, 0x05, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64,
	0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x54, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64,
	0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x22, 0x56, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x22, 0xae, 0x01, 0x0a,
	0x08, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x4e, 0x43,
	0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x4e, 0x43,
	0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x4c, 0x45, 0x5f, 0x44, 0x49, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x44, 0x45, 0x4c, 0x54, 0x41, 0x5f, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x5f,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x4e, 0x43, 0x4f,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45, 0x4c, 0x54, 0x41, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x5f,
	0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x4e, 0x43, 0x4f, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45, 0x4c, 0x54, 0x41, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48,
	0x5f, 0x42, 0x59, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x04, 0x22, 0xa4, 0x01,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x1c, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x4e, 0x41, 0x50, 0x50, 0x59, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4d, 0x50, 0x52,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x02, 0x12, 0x16, 0x0a,
	0x12, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x52, 0x4f,
	0x54, 0x4c, 0x49, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x5a, 0x34, 0x5f, 0x52, 0x41, 0x57, 0x10, 0x04, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x5a, 0x53,
	0x54, 0x44, 0x10, 0x05, 0x22, 0xf7, 0x01, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75,
	0x6c, 0x6c, 0x73, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0xfd,
	0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x72, 0x6f, 0x73,
	0x74, 0x64, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x53, 0x58, 0xaa, 0x02, 0x17, 0x46, 0x72, 0x6f, 0x73, 0x74,
	0x64, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x17, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x5c, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x23, 0x46,
	0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x5c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x19, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x3a, 0x3a, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_frostdb_schema_v1alpha1_schema_proto_rawDescOnce sync.Once
	file_frostdb_schema_v1alpha1_schema_proto_rawDescData = file_frostdb_schema_v1alpha1_schema_proto_rawDesc
)

func file_frostdb_schema_v1alpha1_schema_proto_rawDescGZIP() []byte {
	file_frostdb_schema_v1alpha1_schema_proto_rawDescOnce.Do(func() {
		file_frostdb_schema_v1alpha1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_frostdb_schema_v1alpha1_schema_proto_rawDescData)
	})
	return file_frostdb_schema_v1alpha1_schema_proto_rawDescData
}

var file_frostdb_schema_v1alpha1_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_frostdb_schema_v1alpha1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_frostdb_schema_v1alpha1_schema_proto_goTypes = []interface{}{
	(StorageLayout_Type)(0),        // 0: frostdb.schema.v1alpha1.StorageLayout.Type
	(StorageLayout_Encoding)(0),    // 1: frostdb.schema.v1alpha1.StorageLayout.Encoding
	(StorageLayout_Compression)(0), // 2: frostdb.schema.v1alpha1.StorageLayout.Compression
	(SortingColumn_Direction)(0),   // 3: frostdb.schema.v1alpha1.SortingColumn.Direction
	(*Schema)(nil),                 // 4: frostdb.schema.v1alpha1.Schema
	(*Column)(nil),                 // 5: frostdb.schema.v1alpha1.Column
	(*StorageLayout)(nil),          // 6: frostdb.schema.v1alpha1.StorageLayout
	(*SortingColumn)(nil),          // 7: frostdb.schema.v1alpha1.SortingColumn
}
var file_frostdb_schema_v1alpha1_schema_proto_depIdxs = []int32{
	5, // 0: frostdb.schema.v1alpha1.Schema.columns:type_name -> frostdb.schema.v1alpha1.Column
	7, // 1: frostdb.schema.v1alpha1.Schema.sorting_columns:type_name -> frostdb.schema.v1alpha1.SortingColumn
	6, // 2: frostdb.schema.v1alpha1.Column.storage_layout:type_name -> frostdb.schema.v1alpha1.StorageLayout
	0, // 3: frostdb.schema.v1alpha1.StorageLayout.type:type_name -> frostdb.schema.v1alpha1.StorageLayout.Type
	1, // 4: frostdb.schema.v1alpha1.StorageLayout.encoding:type_name -> frostdb.schema.v1alpha1.StorageLayout.Encoding
	2, // 5: frostdb.schema.v1alpha1.StorageLayout.compression:type_name -> frostdb.schema.v1alpha1.StorageLayout.Compression
	3, // 6: frostdb.schema.v1alpha1.SortingColumn.direction:type_name -> frostdb.schema.v1alpha1.SortingColumn.Direction
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_frostdb_schema_v1alpha1_schema_proto_init() }
func file_frostdb_schema_v1alpha1_schema_proto_init() {
	if File_frostdb_schema_v1alpha1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frostdb_schema_v1alpha1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_frostdb_schema_v1alpha1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_frostdb_schema_v1alpha1_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageLayout); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_frostdb_schema_v1alpha1_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortingColumn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_frostdb_schema_v1alpha1_schema_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frostdb_schema_v1alpha1_schema_proto_goTypes,
		DependencyIndexes: file_frostdb_schema_v1alpha1_schema_proto_depIdxs,
		EnumInfos:         file_frostdb_schema_v1alpha1_schema_proto_enumTypes,
		MessageInfos:      file_frostdb_schema_v1alpha1_schema_proto_msgTypes,
	}.Build()
	File_frostdb_schema_v1alpha1_schema_proto = out.File
	file_frostdb_schema_v1alpha1_schema_proto_rawDesc = nil
	file_frostdb_schema_v1alpha1_schema_proto_goTypes = nil
	file_frostdb_schema_v1alpha1_schema_proto_depIdxs = nil
}
