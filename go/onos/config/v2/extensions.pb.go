// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/config/v2/extensions.proto

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

type TransactionalCommand int32

const (
	TransactionalCommand_TRANSACTIONAL_PREPARE  TransactionalCommand = 0
	TransactionalCommand_TRANSACTIONAL_COMMIT   TransactionalCommand = 1
	TransactionalCommand_TRANSACTIONAL_ROLLBACK TransactionalCommand = 2
)

var TransactionalCommand_name = map[int32]string{
	0: "TRANSACTIONAL_PREPARE",
	1: "TRANSACTIONAL_COMMIT",
	2: "TRANSACTIONAL_ROLLBACK",
}

var TransactionalCommand_value = map[string]int32{
	"TRANSACTIONAL_PREPARE":  0,
	"TRANSACTIONAL_COMMIT":   1,
	"TRANSACTIONAL_ROLLBACK": 2,
}

func (x TransactionalCommand) String() string {
	return proto.EnumName(TransactionalCommand_name, int32(x))
}

func (TransactionalCommand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac6be9aa3eaafbdc, []int{0}
}

// TransactionInfo is an extension providing information about the transaction
// to clients in responses.
type TransactionInfo struct {
	Index Index `protobuf:"varint,1,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
}

func (m *TransactionInfo) Reset()         { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()    {}
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6be9aa3eaafbdc, []int{0}
}
func (m *TransactionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInfo.Merge(m, src)
}
func (m *TransactionInfo) XXX_Size() int {
	return m.Size()
}
func (m *TransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInfo proto.InternalMessageInfo

func (m *TransactionInfo) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

// TransactionMode is an extension for constraining the execution of a transaction for
// stronger consistency guarantees.
type TransactionMode struct {
	Sync   bool `protobuf:"varint,1,opt,name=sync,proto3" json:"sync,omitempty"`
	Atomic bool `protobuf:"varint,2,opt,name=atomic,proto3" json:"atomic,omitempty"`
}

func (m *TransactionMode) Reset()         { *m = TransactionMode{} }
func (m *TransactionMode) String() string { return proto.CompactTextString(m) }
func (*TransactionMode) ProtoMessage()    {}
func (*TransactionMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6be9aa3eaafbdc, []int{1}
}
func (m *TransactionMode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionMode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionMode.Merge(m, src)
}
func (m *TransactionMode) XXX_Size() int {
	return m.Size()
}
func (m *TransactionMode) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionMode.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionMode proto.InternalMessageInfo

func (m *TransactionMode) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

func (m *TransactionMode) GetAtomic() bool {
	if m != nil {
		return m.Atomic
	}
	return false
}

// Transactional is a extension that if supported by targets enables atomic
// transactions across multiple targets.
type Transactional struct {
	ID      TransactionID        `protobuf:"bytes,1,opt,name=id,proto3,casttype=TransactionID" json:"id,omitempty"`
	Index   Index                `protobuf:"varint,2,opt,name=index,proto3,casttype=Index" json:"index,omitempty"`
	Command TransactionalCommand `protobuf:"varint,3,opt,name=command,proto3,enum=onos.config.v2.TransactionalCommand" json:"command,omitempty"`
}

func (m *Transactional) Reset()         { *m = Transactional{} }
func (m *Transactional) String() string { return proto.CompactTextString(m) }
func (*Transactional) ProtoMessage()    {}
func (*Transactional) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac6be9aa3eaafbdc, []int{2}
}
func (m *Transactional) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transactional) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transactional.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transactional) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transactional.Merge(m, src)
}
func (m *Transactional) XXX_Size() int {
	return m.Size()
}
func (m *Transactional) XXX_DiscardUnknown() {
	xxx_messageInfo_Transactional.DiscardUnknown(m)
}

var xxx_messageInfo_Transactional proto.InternalMessageInfo

func (m *Transactional) GetID() TransactionID {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Transactional) GetIndex() Index {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Transactional) GetCommand() TransactionalCommand {
	if m != nil {
		return m.Command
	}
	return TransactionalCommand_TRANSACTIONAL_PREPARE
}

func init() {
	proto.RegisterEnum("onos.config.v2.TransactionalCommand", TransactionalCommand_name, TransactionalCommand_value)
	proto.RegisterType((*TransactionInfo)(nil), "onos.config.v2.TransactionInfo")
	proto.RegisterType((*TransactionMode)(nil), "onos.config.v2.TransactionMode")
	proto.RegisterType((*Transactional)(nil), "onos.config.v2.Transactional")
}

func init() { proto.RegisterFile("onos/config/v2/extensions.proto", fileDescriptor_ac6be9aa3eaafbdc) }

var fileDescriptor_ac6be9aa3eaafbdc = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x73, 0xb1, 0xad, 0xf6, 0xc0, 0x5a, 0x8e, 0x5a, 0x63, 0x87, 0xa4, 0x14, 0xc1, 0xe2,
	0x90, 0x40, 0x9c, 0x15, 0x92, 0xb4, 0x43, 0xb0, 0xff, 0x38, 0xb3, 0x4b, 0x4c, 0xd2, 0x70, 0x60,
	0xef, 0x95, 0x26, 0x94, 0xfa, 0x2d, 0x5c, 0xfd, 0x46, 0x8e, 0x1d, 0x9d, 0x8a, 0xa4, 0xdf, 0xa2,
	0x93, 0xf4, 0xa2, 0xd8, 0x40, 0xb7, 0xf7, 0x9e, 0xe7, 0xf7, 0xe3, 0x5e, 0xee, 0xb0, 0x06, 0x1c,
	0x12, 0x23, 0x00, 0x3e, 0x65, 0xb1, 0xb1, 0x30, 0x8d, 0x68, 0x99, 0x46, 0x3c, 0x61, 0xc0, 0x13,
	0xfd, 0x75, 0x0e, 0x29, 0x90, 0xda, 0x0e, 0xd0, 0x73, 0x40, 0x5f, 0x98, 0xad, 0x46, 0x0c, 0x31,
	0x88, 0xca, 0xd8, 0x4d, 0x39, 0xd5, 0x31, 0xf1, 0x99, 0x37, 0xf7, 0x79, 0xe2, 0x07, 0x29, 0x03,
	0xee, 0xf2, 0x29, 0x10, 0x0d, 0x97, 0x19, 0x0f, 0xa3, 0xa5, 0x82, 0xda, 0xa8, 0x5b, 0xb2, 0xab,
	0xdb, 0xb5, 0x56, 0x76, 0x77, 0x01, 0xcd, 0xf3, 0xce, 0x5d, 0xc1, 0x19, 0x42, 0x18, 0x11, 0x82,
	0x4b, 0xc9, 0x1b, 0x0f, 0x84, 0x72, 0x42, 0xc5, 0x4c, 0x9a, 0xb8, 0xe2, 0xa7, 0x30, 0x63, 0x81,
	0x22, 0x8b, 0xf4, 0xf7, 0xd4, 0xf9, 0x40, 0xf8, 0x74, 0xcf, 0xf7, 0x5f, 0xc8, 0x35, 0x96, 0x59,
	0x28, 0xdc, 0xaa, 0x7d, 0x91, 0xad, 0x35, 0xd9, 0xed, 0x6d, 0xd7, 0xda, 0x3e, 0xe4, 0xf6, 0xa8,
	0xcc, 0xc2, 0xff, 0xd5, 0xe4, 0xc3, 0xab, 0x91, 0x7b, 0x7c, 0x1c, 0xc0, 0x6c, 0xe6, 0xf3, 0x50,
	0x39, 0x6a, 0xa3, 0x6e, 0xcd, 0xbc, 0xd2, 0x8b, 0xcf, 0xa0, 0x17, 0x6e, 0x76, 0x72, 0x96, 0xfe,
	0x49, 0x37, 0x31, 0x6e, 0x1c, 0x02, 0xc8, 0x25, 0x3e, 0xf7, 0xa8, 0x35, 0x7a, 0xb4, 0x1c, 0xcf,
	0x1d, 0x8f, 0xac, 0xc1, 0xd3, 0x84, 0xf6, 0x27, 0x16, 0xed, 0xd7, 0x25, 0xa2, 0xe0, 0x46, 0xb1,
	0x72, 0xc6, 0xc3, 0xa1, 0xeb, 0xd5, 0x11, 0x69, 0xe1, 0x66, 0xb1, 0xa1, 0xe3, 0xc1, 0xc0, 0xb6,
	0x9c, 0x87, 0xba, 0x6c, 0x2b, 0x9f, 0x99, 0x8a, 0x56, 0x99, 0x8a, 0xbe, 0x33, 0x15, 0xbd, 0x6f,
	0x54, 0x69, 0xb5, 0x51, 0xa5, 0xaf, 0x8d, 0x2a, 0x3d, 0x57, 0xc4, 0xc7, 0xdc, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x4b, 0xf4, 0x54, 0x68, 0xe1, 0x01, 0x00, 0x00,
}

func (m *TransactionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintExtensions(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TransactionMode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionMode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionMode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Atomic {
		i--
		if m.Atomic {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Sync {
		i--
		if m.Sync {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Transactional) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transactional) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transactional) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Command != 0 {
		i = encodeVarintExtensions(dAtA, i, uint64(m.Command))
		i--
		dAtA[i] = 0x18
	}
	if m.Index != 0 {
		i = encodeVarintExtensions(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintExtensions(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtensions(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtensions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovExtensions(uint64(m.Index))
	}
	return n
}

func (m *TransactionMode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sync {
		n += 2
	}
	if m.Atomic {
		n += 2
	}
	return n
}

func (m *Transactional) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovExtensions(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovExtensions(uint64(m.Index))
	}
	if m.Command != 0 {
		n += 1 + sovExtensions(uint64(m.Command))
	}
	return n
}

func sovExtensions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtensions(x uint64) (n int) {
	return sovExtensions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtensions
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
			return fmt.Errorf("proto: TransactionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtensions
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
			skippy, err := skipExtensions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtensions
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
func (m *TransactionMode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtensions
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
			return fmt.Errorf("proto: TransactionMode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionMode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sync", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtensions
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
			m.Sync = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Atomic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtensions
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
			m.Atomic = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipExtensions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtensions
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
func (m *Transactional) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtensions
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
			return fmt.Errorf("proto: Transactional: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transactional: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtensions
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
				return ErrInvalidLengthExtensions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtensions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = TransactionID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtensions
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			m.Command = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtensions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Command |= TransactionalCommand(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExtensions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtensions
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
func skipExtensions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtensions
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
					return 0, ErrIntOverflowExtensions
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
					return 0, ErrIntOverflowExtensions
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
				return 0, ErrInvalidLengthExtensions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtensions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtensions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtensions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtensions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtensions = fmt.Errorf("proto: unexpected end of group")
)
