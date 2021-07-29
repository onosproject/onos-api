// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/uenib/ran.proto

package uenib

import (
	encoding_binary "encoding/binary"
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

// CellConnection represents UE cell connection.
type CellConnection struct {
	ID             ID      `protobuf:"bytes,1,opt,name=id,proto3,casttype=ID" json:"id,omitempty"`
	SignalStrength float64 `protobuf:"fixed64,2,opt,name=signal_strength,json=signalStrength,proto3" json:"signal_strength,omitempty"`
}

func (m *CellConnection) Reset()         { *m = CellConnection{} }
func (m *CellConnection) String() string { return proto.CompactTextString(m) }
func (*CellConnection) ProtoMessage()    {}
func (*CellConnection) Descriptor() ([]byte, []int) {
	return fileDescriptor_176aa06704fa2b20, []int{0}
}
func (m *CellConnection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CellConnection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CellConnection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CellConnection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellConnection.Merge(m, src)
}
func (m *CellConnection) XXX_Size() int {
	return m.Size()
}
func (m *CellConnection) XXX_DiscardUnknown() {
	xxx_messageInfo_CellConnection.DiscardUnknown(m)
}

var xxx_messageInfo_CellConnection proto.InternalMessageInfo

func (m *CellConnection) GetID() ID {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CellConnection) GetSignalStrength() float64 {
	if m != nil {
		return m.SignalStrength
	}
	return 0
}

// CellInfo provides data on serving cell and candidate cells.
type CellInfo struct {
	ServingCell    *CellConnection   `protobuf:"bytes,1,opt,name=serving_cell,json=servingCell,proto3" json:"serving_cell,omitempty"`
	CandidateCells []*CellConnection `protobuf:"bytes,2,rep,name=candidate_cells,json=candidateCells,proto3" json:"candidate_cells,omitempty"`
}

func (m *CellInfo) Reset()         { *m = CellInfo{} }
func (m *CellInfo) String() string { return proto.CompactTextString(m) }
func (*CellInfo) ProtoMessage()    {}
func (*CellInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_176aa06704fa2b20, []int{1}
}
func (m *CellInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CellInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CellInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CellInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellInfo.Merge(m, src)
}
func (m *CellInfo) XXX_Size() int {
	return m.Size()
}
func (m *CellInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CellInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CellInfo proto.InternalMessageInfo

func (m *CellInfo) GetServingCell() *CellConnection {
	if m != nil {
		return m.ServingCell
	}
	return nil
}

func (m *CellInfo) GetCandidateCells() []*CellConnection {
	if m != nil {
		return m.CandidateCells
	}
	return nil
}

func init() {
	proto.RegisterType((*CellConnection)(nil), "onos.uenib.CellConnection")
	proto.RegisterType((*CellInfo)(nil), "onos.uenib.CellInfo")
}

func init() { proto.RegisterFile("onos/uenib/ran.proto", fileDescriptor_176aa06704fa2b20) }

var fileDescriptor_176aa06704fa2b20 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xcf, 0xcb, 0x2f,
	0xd6, 0x2f, 0x4d, 0xcd, 0xcb, 0x4c, 0xd2, 0x2f, 0x4a, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x02, 0x89, 0xea, 0x81, 0x45, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xc2, 0xfa,
	0x20, 0x16, 0x44, 0x85, 0x52, 0x38, 0x17, 0x9f, 0x73, 0x6a, 0x4e, 0x8e, 0x73, 0x7e, 0x5e, 0x5e,
	0x6a, 0x72, 0x49, 0x66, 0x7e, 0x9e, 0x90, 0x0c, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0xa7, 0x13, 0xcf, 0xa3, 0x7b, 0xf2, 0x4c, 0x9e, 0x2e, 0xbf, 0xc0, 0x64, 0x10, 0x53, 0x66,
	0x8a, 0x90, 0x3a, 0x17, 0x7f, 0x71, 0x66, 0x7a, 0x5e, 0x62, 0x4e, 0x7c, 0x71, 0x49, 0x51, 0x6a,
	0x5e, 0x7a, 0x49, 0x86, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x1f, 0x44, 0x38, 0x18, 0x2a,
	0xaa, 0xd4, 0xc7, 0xc8, 0xc5, 0x01, 0x32, 0xd9, 0x33, 0x2f, 0x2d, 0x5f, 0xc8, 0x96, 0x8b, 0xa7,
	0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x2f, 0x3d, 0x3e, 0x39, 0x35, 0x27, 0x07, 0x6c, 0x3a, 0xb7, 0x91,
	0x94, 0x1e, 0xc2, 0x79, 0x7a, 0xa8, 0xae, 0x08, 0xe2, 0x86, 0xaa, 0x07, 0x09, 0x0b, 0x39, 0x73,
	0xf1, 0x27, 0x27, 0xe6, 0xa5, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0x82, 0x0d, 0x28, 0x96, 0x60, 0x52,
	0x60, 0x26, 0x60, 0x02, 0x1f, 0x5c, 0x0b, 0x48, 0xa2, 0xd8, 0x49, 0xe2, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x41, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x13, 0xc9, 0x82, 0xb9, 0x44, 0x01, 0x00, 0x00,
}

func (m *CellConnection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CellConnection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CellConnection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignalStrength != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SignalStrength))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintRan(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CellInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CellInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CellInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CandidateCells) > 0 {
		for iNdEx := len(m.CandidateCells) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CandidateCells[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRan(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ServingCell != nil {
		{
			size, err := m.ServingCell.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRan(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRan(dAtA []byte, offset int, v uint64) int {
	offset -= sovRan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CellConnection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovRan(uint64(l))
	}
	if m.SignalStrength != 0 {
		n += 9
	}
	return n
}

func (m *CellInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServingCell != nil {
		l = m.ServingCell.Size()
		n += 1 + l + sovRan(uint64(l))
	}
	if len(m.CandidateCells) > 0 {
		for _, e := range m.CandidateCells {
			l = e.Size()
			n += 1 + l + sovRan(uint64(l))
		}
	}
	return n
}

func sovRan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRan(x uint64) (n int) {
	return sovRan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CellConnection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRan
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
			return fmt.Errorf("proto: CellConnection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CellConnection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRan
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
				return ErrInvalidLengthRan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = ID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignalStrength", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SignalStrength = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipRan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRan
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
func (m *CellInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRan
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
			return fmt.Errorf("proto: CellInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CellInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServingCell", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRan
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
				return ErrInvalidLengthRan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ServingCell == nil {
				m.ServingCell = &CellConnection{}
			}
			if err := m.ServingCell.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CandidateCells", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRan
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
				return ErrInvalidLengthRan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CandidateCells = append(m.CandidateCells, &CellConnection{})
			if err := m.CandidateCells[len(m.CandidateCells)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRan
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
func skipRan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRan
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
					return 0, ErrIntOverflowRan
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
					return 0, ErrIntOverflowRan
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
				return 0, ErrInvalidLengthRan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRan = fmt.Errorf("proto: unexpected end of group")
)
