// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/config/v2/object.proto

package v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ObjectMeta struct {
	Key      string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Version  uint64     `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Revision Revision   `protobuf:"varint,3,opt,name=revision,proto3,casttype=Revision" json:"revision,omitempty"`
	Created  time.Time  `protobuf:"bytes,4,opt,name=created,proto3,stdtime" json:"created"`
	Updated  time.Time  `protobuf:"bytes,5,opt,name=updated,proto3,stdtime" json:"updated"`
	Deleted  *time.Time `protobuf:"bytes,6,opt,name=deleted,proto3,stdtime" json:"deleted,omitempty"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be24dfbef090a8f, []int{0}
}
func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(m, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return m.Size()
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ObjectMeta) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ObjectMeta) GetRevision() Revision {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *ObjectMeta) GetCreated() time.Time {
	if m != nil {
		return m.Created
	}
	return time.Time{}
}

func (m *ObjectMeta) GetUpdated() time.Time {
	if m != nil {
		return m.Updated
	}
	return time.Time{}
}

func (m *ObjectMeta) GetDeleted() *time.Time {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func init() {
	proto.RegisterType((*ObjectMeta)(nil), "onos.config.v2.ObjectMeta")
}

func init() { proto.RegisterFile("onos/config/v2/object.proto", fileDescriptor_6be24dfbef090a8f) }

var fileDescriptor_6be24dfbef090a8f = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xce, 0xbf, 0x4e, 0xac, 0x40,
	0x14, 0x06, 0x70, 0x66, 0x97, 0xbb, 0x70, 0x47, 0x63, 0xcc, 0xc4, 0x82, 0x60, 0x32, 0x10, 0x2b,
	0xaa, 0x99, 0x04, 0x3b, 0x0b, 0x0b, 0x7a, 0x63, 0x42, 0x7c, 0x01, 0xfe, 0xcc, 0x12, 0x74, 0x97,
	0x43, 0x60, 0x96, 0xc4, 0xb7, 0xd8, 0x27, 0xf0, 0x79, 0xb6, 0xdc, 0xd2, 0x6a, 0x35, 0xf0, 0x16,
	0x56, 0x66, 0x06, 0xb0, 0x35, 0x76, 0xe7, 0xcc, 0x77, 0x7e, 0x93, 0x0f, 0x5f, 0x43, 0x05, 0x2d,
	0xcf, 0xa0, 0x5a, 0x97, 0x05, 0xef, 0x42, 0x0e, 0xe9, 0xb3, 0xc8, 0x24, 0xab, 0x1b, 0x90, 0x40,
	0x2e, 0x54, 0xc8, 0xc6, 0x90, 0x75, 0xa1, 0x7b, 0x55, 0x40, 0x01, 0x3a, 0xe2, 0x6a, 0x1a, 0xaf,
	0x5c, 0xaf, 0x00, 0x28, 0x36, 0x82, 0xeb, 0x2d, 0xdd, 0xad, 0xb9, 0x2c, 0xb7, 0xa2, 0x95, 0xc9,
	0xb6, 0x1e, 0x0f, 0x6e, 0xde, 0x16, 0x18, 0x3f, 0xea, 0x7f, 0x1f, 0x84, 0x4c, 0xc8, 0x25, 0x5e,
	0xbe, 0x88, 0x57, 0x07, 0xf9, 0x28, 0xf8, 0x1f, 0xab, 0x91, 0x38, 0xd8, 0xea, 0x44, 0xd3, 0x96,
	0x50, 0x39, 0x0b, 0x1f, 0x05, 0x66, 0x3c, 0xaf, 0x24, 0xc0, 0x76, 0x23, 0xba, 0x52, 0x47, 0x4b,
	0x15, 0x45, 0xe7, 0x5f, 0x27, 0xcf, 0x8e, 0xa7, 0xb7, 0xf8, 0x27, 0x25, 0xf7, 0xd8, 0xca, 0x1a,
	0x91, 0x48, 0x91, 0x3b, 0xa6, 0x8f, 0x82, 0xb3, 0xd0, 0x65, 0x63, 0x2f, 0x36, 0xf7, 0x62, 0x4f,
	0x73, 0xaf, 0xc8, 0x3e, 0x9c, 0x3c, 0x63, 0xff, 0xe1, 0xa1, 0x78, 0x46, 0xca, 0xef, 0xea, 0x5c,
	0xfb, 0x7f, 0x7f, 0xf1, 0x13, 0x22, 0x77, 0xd8, 0xca, 0xc5, 0x46, 0x28, 0xbf, 0xfa, 0xd5, 0x9b,
	0xa3, 0x9d, 0x40, 0xe4, 0x1c, 0x7a, 0x8a, 0x8e, 0x3d, 0x45, 0x9f, 0x3d, 0x45, 0xfb, 0x81, 0x1a,
	0xc7, 0x81, 0x1a, 0xef, 0x03, 0x35, 0xd2, 0x95, 0xc6, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x33, 0x97, 0xd8, 0x6f, 0xa7, 0x01, 0x00, 0x00,
}

func (m *ObjectMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObjectMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObjectMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deleted != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Deleted, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Deleted):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintObject(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x32
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Updated, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintObject(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Created, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Created):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintObject(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if m.Revision != 0 {
		i = encodeVarintObject(dAtA, i, uint64(m.Revision))
		i--
		dAtA[i] = 0x18
	}
	if m.Version != 0 {
		i = encodeVarintObject(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintObject(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	offset -= sovObject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObjectMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovObject(uint64(m.Version))
	}
	if m.Revision != 0 {
		n += 1 + sovObject(uint64(m.Revision))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Created)
	n += 1 + l + sovObject(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated)
	n += 1 + l + sovObject(uint64(l))
	if m.Deleted != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Deleted)
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func sovObject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObjectMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: ObjectMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjectMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= Revision(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Created, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Updated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deleted == nil {
				m.Deleted = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Deleted, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObject
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
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObject = fmt.Errorf("proto: unexpected end of group")
)
