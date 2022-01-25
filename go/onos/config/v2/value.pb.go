// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/config/v2/value.proto

package v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ValueType is the type for a value
type ValueType int32

const (
	ValueType_EMPTY            ValueType = 0
	ValueType_STRING           ValueType = 1
	ValueType_INT              ValueType = 2
	ValueType_UINT             ValueType = 3
	ValueType_BOOL             ValueType = 4
	ValueType_DECIMAL          ValueType = 5
	ValueType_FLOAT            ValueType = 6
	ValueType_BYTES            ValueType = 7
	ValueType_LEAFLIST_STRING  ValueType = 8
	ValueType_LEAFLIST_INT     ValueType = 9
	ValueType_LEAFLIST_UINT    ValueType = 10
	ValueType_LEAFLIST_BOOL    ValueType = 11
	ValueType_LEAFLIST_DECIMAL ValueType = 12
	ValueType_LEAFLIST_FLOAT   ValueType = 13
	ValueType_LEAFLIST_BYTES   ValueType = 14
)

var ValueType_name = map[int32]string{
	0:  "EMPTY",
	1:  "STRING",
	2:  "INT",
	3:  "UINT",
	4:  "BOOL",
	5:  "DECIMAL",
	6:  "FLOAT",
	7:  "BYTES",
	8:  "LEAFLIST_STRING",
	9:  "LEAFLIST_INT",
	10: "LEAFLIST_UINT",
	11: "LEAFLIST_BOOL",
	12: "LEAFLIST_DECIMAL",
	13: "LEAFLIST_FLOAT",
	14: "LEAFLIST_BYTES",
}

var ValueType_value = map[string]int32{
	"EMPTY":            0,
	"STRING":           1,
	"INT":              2,
	"UINT":             3,
	"BOOL":             4,
	"DECIMAL":          5,
	"FLOAT":            6,
	"BYTES":            7,
	"LEAFLIST_STRING":  8,
	"LEAFLIST_INT":     9,
	"LEAFLIST_UINT":    10,
	"LEAFLIST_BOOL":    11,
	"LEAFLIST_DECIMAL": 12,
	"LEAFLIST_FLOAT":   13,
	"LEAFLIST_BYTES":   14,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8823d92a68f01dab, []int{0}
}

// TypedValue is a value represented as a byte array
type TypedValue struct {
	// 'bytes' is the bytes array
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,json=Bytes,proto3" json:"Bytes,omitempty"`
	// 'type' is the value type
	Type ValueType `protobuf:"varint,2,opt,name=type,json=Type,proto3,enum=onos.config.v2.ValueType" json:"Type,omitempty"`
	// 'type_opts' is a set of type options
	TypeOpts []int32 `protobuf:"varint,3,rep,packed,name=type_opts,json=TypeOpts,proto3" json:"TypeOpts,omitempty"`
}

func (m *TypedValue) Reset()         { *m = TypedValue{} }
func (m *TypedValue) String() string { return proto.CompactTextString(m) }
func (*TypedValue) ProtoMessage()    {}
func (*TypedValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8823d92a68f01dab, []int{0}
}
func (m *TypedValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TypedValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TypedValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TypedValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedValue.Merge(m, src)
}
func (m *TypedValue) XXX_Size() int {
	return m.Size()
}
func (m *TypedValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedValue.DiscardUnknown(m)
}

var xxx_messageInfo_TypedValue proto.InternalMessageInfo

func (m *TypedValue) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *TypedValue) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_EMPTY
}

func (m *TypedValue) GetTypeOpts() []int32 {
	if m != nil {
		return m.TypeOpts
	}
	return nil
}

// Value is the state of a value in the configuration tree
type Value struct {
	// 'value' is the change value
	Value TypedValue `protobuf:"bytes,2,opt,name=value,json=Value,proto3" json:"Value,omitempty"`
	// 'deleted' indicates whether this is a delete
	Deleted bool `protobuf:"varint,3,opt,name=deleted,proto3" json:"Deleted,omitempty"`
	// 'index'
	Index Index `protobuf:"varint,4,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_8823d92a68f01dab, []int{1}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Value.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return m.Size()
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetValue() TypedValue {
	if m != nil {
		return m.Value
	}
	return TypedValue{}
}

func (m *Value) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Value) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

// PathValue is the state of a path/value in the configuration tree
type PathValue struct {
	// 'path' is the path to change
	Path string `protobuf:"bytes,1,opt,name=path,json=Path,proto3" json:"Path,omitempty"`
	// 'value' is the change value
	Value `protobuf:"bytes,2,opt,name=value,json=Value,proto3,embedded=value" json:"value"`
}

func (m *PathValue) Reset()         { *m = PathValue{} }
func (m *PathValue) String() string { return proto.CompactTextString(m) }
func (*PathValue) ProtoMessage()    {}
func (*PathValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8823d92a68f01dab, []int{2}
}
func (m *PathValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PathValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PathValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PathValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathValue.Merge(m, src)
}
func (m *PathValue) XXX_Size() int {
	return m.Size()
}
func (m *PathValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PathValue.DiscardUnknown(m)
}

var xxx_messageInfo_PathValue proto.InternalMessageInfo

func (m *PathValue) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterEnum("onos.config.v2.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*TypedValue)(nil), "onos.config.v2.TypedValue")
	proto.RegisterType((*Value)(nil), "onos.config.v2.Value")
	proto.RegisterType((*PathValue)(nil), "onos.config.v2.PathValue")
}

func init() { proto.RegisterFile("onos/config/v2/value.proto", fileDescriptor_8823d92a68f01dab) }

var fileDescriptor_8823d92a68f01dab = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xbd, 0xb5, 0x9d, 0xd8, 0x93, 0xd4, 0x75, 0xb7, 0x2d, 0x98, 0x1c, 0x6c, 0x2b, 0x07,
	0x64, 0x10, 0xb2, 0xa5, 0x54, 0x3c, 0x40, 0x96, 0xa6, 0xc8, 0x92, 0xdb, 0x54, 0x89, 0x41, 0xea,
	0xa9, 0x4a, 0x89, 0x49, 0x83, 0xda, 0xd8, 0x22, 0x4b, 0x44, 0xde, 0x82, 0x87, 0xe0, 0xce, 0x6b,
	0xf4, 0x98, 0x23, 0x27, 0x0b, 0x25, 0xb7, 0x48, 0xbc, 0x00, 0x27, 0xb4, 0xbb, 0x8d, 0x71, 0x50,
	0x6f, 0xb3, 0xdf, 0x8c, 0xff, 0xf9, 0xff, 0x91, 0xa1, 0x91, 0x4e, 0xd2, 0x69, 0xf0, 0x21, 0x9d,
	0x7c, 0x1c, 0x8f, 0x82, 0x59, 0x2b, 0x98, 0x0d, 0x6e, 0xbf, 0x24, 0x7e, 0xf6, 0x39, 0xa5, 0x29,
	0x36, 0x58, 0xcf, 0x17, 0x3d, 0x7f, 0xd6, 0x6a, 0x1c, 0x8e, 0xd2, 0x51, 0xca, 0x5b, 0x01, 0xab,
	0xc4, 0x54, 0xf3, 0x07, 0x02, 0x88, 0xe7, 0x59, 0x32, 0x7c, 0xcf, 0x3e, 0xc5, 0x2f, 0x40, 0xbd,
	0x9e, 0xd3, 0x64, 0x6a, 0x21, 0x17, 0x79, 0x75, 0x72, 0xb0, 0xce, 0x9d, 0x3d, 0xc2, 0xc0, 0xab,
	0xf4, 0x6e, 0x4c, 0x93, 0xbb, 0x8c, 0xce, 0x7b, 0x2a, 0x07, 0xb8, 0x0d, 0x0a, 0x9d, 0x67, 0x89,
	0xb5, 0xe3, 0x22, 0xcf, 0x68, 0x3d, 0xf3, 0xb7, 0xd7, 0xf9, 0x5c, 0x8f, 0x29, 0x13, 0xbc, 0xce,
	0x1d, 0x83, 0x55, 0x25, 0x0d, 0x85, 0xbd, 0xf1, 0x31, 0xe8, 0x4c, 0xe2, 0x2a, 0xcd, 0xe8, 0xd4,
	0x92, 0x5d, 0xd9, 0x53, 0xc9, 0x93, 0x75, 0xee, 0x60, 0xd6, 0xec, 0x66, 0xb4, 0xbc, 0x54, 0xdb,
	0xb0, 0xe6, 0x77, 0x04, 0xaa, 0x30, 0x1b, 0x82, 0xca, 0x03, 0x73, 0x0b, 0xb5, 0x56, 0xe3, 0x7f,
	0x0b, 0xff, 0x72, 0x91, 0xa7, 0xf7, 0xb9, 0x23, 0xb1, 0x30, 0xfc, 0x59, 0x0e, 0x23, 0xa4, 0x02,
	0xa8, 0x0e, 0x93, 0xdb, 0x84, 0x26, 0x43, 0x4b, 0x76, 0x91, 0xa7, 0x91, 0xa3, 0x75, 0xee, 0xec,
	0x9f, 0x08, 0x54, 0x1a, 0xdf, 0x4c, 0x61, 0x07, 0xd4, 0xf1, 0x64, 0x98, 0x7c, 0xb5, 0x14, 0x17,
	0x79, 0x0a, 0xd1, 0xff, 0xe4, 0x8e, 0x1a, 0x32, 0xd0, 0x13, 0xbc, 0xf9, 0x09, 0xf4, 0x8b, 0x01,
	0xbd, 0x11, 0xf2, 0xcf, 0x41, 0xc9, 0x06, 0xf4, 0x86, 0x5f, 0x55, 0x17, 0x07, 0x61, 0xcd, 0xf2,
	0x41, 0xd8, 0x1b, 0xbf, 0xde, 0x4e, 0x74, 0xf4, 0xe8, 0x51, 0x89, 0xc6, 0xc2, 0x2c, 0x72, 0x07,
	0x3d, 0xb8, 0x7f, 0xf9, 0x1b, 0x81, 0x5e, 0xdc, 0x1b, 0xeb, 0xa0, 0x76, 0xce, 0x2e, 0xe2, 0x4b,
	0x53, 0xc2, 0x00, 0x95, 0x7e, 0xdc, 0x0b, 0xcf, 0xdf, 0x9a, 0x08, 0x57, 0x41, 0x0e, 0xcf, 0x63,
	0x73, 0x07, 0x6b, 0xa0, 0xbc, 0x63, 0x95, 0xcc, 0x2a, 0xd2, 0xed, 0x46, 0xa6, 0x82, 0x6b, 0x50,
	0x3d, 0xe9, 0xbc, 0x09, 0xcf, 0xda, 0x91, 0xa9, 0x32, 0x81, 0xd3, 0xa8, 0xdb, 0x8e, 0xcd, 0x0a,
	0x2b, 0xc9, 0x65, 0xdc, 0xe9, 0x9b, 0x55, 0x7c, 0x00, 0x7b, 0x51, 0xa7, 0x7d, 0x1a, 0x85, 0xfd,
	0xf8, 0xea, 0x41, 0x54, 0xc3, 0x26, 0xd4, 0x0b, 0xc8, 0x34, 0x75, 0xbc, 0x0f, 0xbb, 0x05, 0xe1,
	0x6b, 0x60, 0x0b, 0xf1, 0x7d, 0x35, 0x7c, 0x08, 0x66, 0x81, 0x36, 0x8b, 0xeb, 0x18, 0x83, 0x51,
	0x50, 0xe1, 0x60, 0x77, 0x8b, 0x09, 0x2b, 0x06, 0xb1, 0xee, 0x97, 0x36, 0x5a, 0x2c, 0x6d, 0xf4,
	0x6b, 0x69, 0xa3, 0x6f, 0x2b, 0x5b, 0x5a, 0xac, 0x6c, 0xe9, 0xe7, 0xca, 0x96, 0xae, 0x2b, 0xfc,
	0xaf, 0x3e, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x55, 0x67, 0xf3, 0x4b, 0x19, 0x03, 0x00, 0x00,
}

func (m *TypedValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TypedValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TypedValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TypeOpts) > 0 {
		dAtA2 := make([]byte, len(m.TypeOpts)*10)
		var j1 int
		for _, num1 := range m.TypeOpts {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintValue(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintValue(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Bytes) > 0 {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = encodeVarintValue(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Value) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintValue(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if m.Deleted {
		i--
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintValue(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *PathValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PathValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PathValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintValue(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintValue(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValue(dAtA []byte, offset int, v uint64) int {
	offset -= sovValue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TypedValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bytes)
	if l > 0 {
		n += 1 + l + sovValue(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovValue(uint64(m.Type))
	}
	if len(m.TypeOpts) > 0 {
		l = 0
		for _, e := range m.TypeOpts {
			l += sovValue(uint64(e))
		}
		n += 1 + sovValue(uint64(l)) + l
	}
	return n
}

func (m *Value) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Value.Size()
	n += 1 + l + sovValue(uint64(l))
	if m.Deleted {
		n += 2
	}
	if m.Index != 0 {
		n += 1 + sovValue(uint64(m.Index))
	}
	return n
}

func (m *PathValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovValue(uint64(l))
	}
	l = m.Value.Size()
	n += 1 + l + sovValue(uint64(l))
	return n
}

func sovValue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValue(x uint64) (n int) {
	return sovValue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TypedValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TypedValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypedValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = append(m.Bytes[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes == nil {
				m.Bytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ValueType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowValue
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TypeOpts = append(m.TypeOpts, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowValue
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthValue
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthValue
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.TypeOpts) == 0 {
					m.TypeOpts = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowValue
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TypeOpts = append(m.TypeOpts, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeOpts", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Value) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= Index(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PathValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PathValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipValue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValue
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowValue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowValue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthValue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValue = fmt.Errorf("proto: unexpected end of group")
)
