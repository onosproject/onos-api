// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/provisioner/aspects.proto

// Package provisioner defines the main device provisioner gRPC interface

package provisioner

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

type ConfigStatus_State int32

const (
	ConfigStatus_PENDING ConfigStatus_State = 0
	ConfigStatus_APPLIED ConfigStatus_State = 3
	ConfigStatus_FAILED  ConfigStatus_State = 4
)

var ConfigStatus_State_name = map[int32]string{
	0: "PENDING",
	3: "APPLIED",
	4: "FAILED",
}

var ConfigStatus_State_value = map[string]int32{
	"PENDING": 0,
	"APPLIED": 3,
	"FAILED":  4,
}

func (x ConfigStatus_State) String() string {
	return proto.EnumName(ConfigStatus_State_name, int32(x))
}

func (ConfigStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa6a5635bb96ef51, []int{3, 0}
}

type Failure_Type int32

const (
	// UNKNOWN unknown failure
	Failure_UNKNOWN Failure_Type = 0
	// CANCELED
	Failure_CANCELED Failure_Type = 1
	// NOT_FOUND
	Failure_NOT_FOUND Failure_Type = 2
	// ALREADY_EXISTS
	Failure_ALREADY_EXISTS Failure_Type = 3
	// UNAUTHORIZED
	Failure_UNAUTHORIZED Failure_Type = 4
	// FORBIDDEN
	Failure_FORBIDDEN Failure_Type = 5
	// CONFLICT
	Failure_CONFLICT Failure_Type = 6
	// INVALID
	Failure_INVALID Failure_Type = 7
	// UNAVAILABLE
	Failure_UNAVAILABLE Failure_Type = 8
	// NOT_SUPPORTED
	Failure_NOT_SUPPORTED Failure_Type = 9
	// TIMEOUT
	Failure_TIMEOUT Failure_Type = 10
	// INTERNAL
	Failure_INTERNAL Failure_Type = 11
)

var Failure_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CANCELED",
	2:  "NOT_FOUND",
	3:  "ALREADY_EXISTS",
	4:  "UNAUTHORIZED",
	5:  "FORBIDDEN",
	6:  "CONFLICT",
	7:  "INVALID",
	8:  "UNAVAILABLE",
	9:  "NOT_SUPPORTED",
	10: "TIMEOUT",
	11: "INTERNAL",
}

var Failure_Type_value = map[string]int32{
	"UNKNOWN":        0,
	"CANCELED":       1,
	"NOT_FOUND":      2,
	"ALREADY_EXISTS": 3,
	"UNAUTHORIZED":   4,
	"FORBIDDEN":      5,
	"CONFLICT":       6,
	"INVALID":        7,
	"UNAVAILABLE":    8,
	"NOT_SUPPORTED":  9,
	"TIMEOUT":        10,
	"INTERNAL":       11,
}

func (x Failure_Type) String() string {
	return proto.EnumName(Failure_Type_name, int32(x))
}

func (Failure_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa6a5635bb96ef51, []int{4, 0}
}

// DeviceConfig is a topology entity aspect used to specify what pipeline and chassis config a device should have applied to it
type DeviceConfig struct {
	PipelineConfigID ConfigID `protobuf:"bytes,1,opt,name=pipeline_config_id,json=pipelineConfigId,proto3,casttype=ConfigID" json:"pipeline_config_id,omitempty"`
	ChassisConfigID  ConfigID `protobuf:"bytes,2,opt,name=chassis_config_id,json=chassisConfigId,proto3,casttype=ConfigID" json:"chassis_config_id,omitempty"`
}

func (m *DeviceConfig) Reset()         { *m = DeviceConfig{} }
func (m *DeviceConfig) String() string { return proto.CompactTextString(m) }
func (*DeviceConfig) ProtoMessage()    {}
func (*DeviceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6a5635bb96ef51, []int{0}
}
func (m *DeviceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceConfig.Merge(m, src)
}
func (m *DeviceConfig) XXX_Size() int {
	return m.Size()
}
func (m *DeviceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceConfig proto.InternalMessageInfo

func (m *DeviceConfig) GetPipelineConfigID() ConfigID {
	if m != nil {
		return m.PipelineConfigID
	}
	return ""
}

func (m *DeviceConfig) GetChassisConfigID() ConfigID {
	if m != nil {
		return m.ChassisConfigID
	}
	return ""
}

// PipelineConfigState is a topology entity aspect used to indicate what pipeline config a device has presently applied to it
type PipelineConfigState struct {
	ConfigID ConfigID     `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3,casttype=ConfigID" json:"config_id,omitempty"`
	Cookie   uint64       `protobuf:"varint,2,opt,name=cookie,proto3" json:"cookie,omitempty"`
	Updated  time.Time    `protobuf:"bytes,3,opt,name=updated,proto3,stdtime" json:"updated"`
	Status   ConfigStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status"`
}

func (m *PipelineConfigState) Reset()         { *m = PipelineConfigState{} }
func (m *PipelineConfigState) String() string { return proto.CompactTextString(m) }
func (*PipelineConfigState) ProtoMessage()    {}
func (*PipelineConfigState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6a5635bb96ef51, []int{1}
}
func (m *PipelineConfigState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PipelineConfigState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PipelineConfigState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PipelineConfigState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineConfigState.Merge(m, src)
}
func (m *PipelineConfigState) XXX_Size() int {
	return m.Size()
}
func (m *PipelineConfigState) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineConfigState.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineConfigState proto.InternalMessageInfo

func (m *PipelineConfigState) GetConfigID() ConfigID {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *PipelineConfigState) GetCookie() uint64 {
	if m != nil {
		return m.Cookie
	}
	return 0
}

func (m *PipelineConfigState) GetUpdated() time.Time {
	if m != nil {
		return m.Updated
	}
	return time.Time{}
}

func (m *PipelineConfigState) GetStatus() ConfigStatus {
	if m != nil {
		return m.Status
	}
	return ConfigStatus{}
}

// ChassisConfigState is a topology entity aspect used to indicate what chassis config a device has presently applied to it
type ChassisConfigState struct {
	ConfigID ConfigID     `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3,casttype=ConfigID" json:"config_id,omitempty"`
	Updated  time.Time    `protobuf:"bytes,2,opt,name=updated,proto3,stdtime" json:"updated"`
	Status   ConfigStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status"`
}

func (m *ChassisConfigState) Reset()         { *m = ChassisConfigState{} }
func (m *ChassisConfigState) String() string { return proto.CompactTextString(m) }
func (*ChassisConfigState) ProtoMessage()    {}
func (*ChassisConfigState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6a5635bb96ef51, []int{2}
}
func (m *ChassisConfigState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChassisConfigState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChassisConfigState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChassisConfigState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChassisConfigState.Merge(m, src)
}
func (m *ChassisConfigState) XXX_Size() int {
	return m.Size()
}
func (m *ChassisConfigState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChassisConfigState.DiscardUnknown(m)
}

var xxx_messageInfo_ChassisConfigState proto.InternalMessageInfo

func (m *ChassisConfigState) GetConfigID() ConfigID {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *ChassisConfigState) GetUpdated() time.Time {
	if m != nil {
		return m.Updated
	}
	return time.Time{}
}

func (m *ChassisConfigState) GetStatus() ConfigStatus {
	if m != nil {
		return m.Status
	}
	return ConfigStatus{}
}

type ConfigStatus struct {
	// 'state' config state
	State ConfigStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=onos.provisioner.ConfigStatus_State" json:"state,omitempty"`
	// 'failure' is the transaction failure (if any)
	Failure *Failure `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (m *ConfigStatus) Reset()         { *m = ConfigStatus{} }
func (m *ConfigStatus) String() string { return proto.CompactTextString(m) }
func (*ConfigStatus) ProtoMessage()    {}
func (*ConfigStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6a5635bb96ef51, []int{3}
}
func (m *ConfigStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigStatus.Merge(m, src)
}
func (m *ConfigStatus) XXX_Size() int {
	return m.Size()
}
func (m *ConfigStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigStatus proto.InternalMessageInfo

func (m *ConfigStatus) GetState() ConfigStatus_State {
	if m != nil {
		return m.State
	}
	return ConfigStatus_PENDING
}

func (m *ConfigStatus) GetFailure() *Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

// Failure config update failure type and description
type Failure struct {
	Type        Failure_Type `protobuf:"varint,1,opt,name=type,proto3,enum=onos.provisioner.Failure_Type" json:"type,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa6a5635bb96ef51, []int{4}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return m.Size()
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetType() Failure_Type {
	if m != nil {
		return m.Type
	}
	return Failure_UNKNOWN
}

func (m *Failure) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("onos.provisioner.ConfigStatus_State", ConfigStatus_State_name, ConfigStatus_State_value)
	proto.RegisterEnum("onos.provisioner.Failure_Type", Failure_Type_name, Failure_Type_value)
	proto.RegisterType((*DeviceConfig)(nil), "onos.provisioner.DeviceConfig")
	proto.RegisterType((*PipelineConfigState)(nil), "onos.provisioner.PipelineConfigState")
	proto.RegisterType((*ChassisConfigState)(nil), "onos.provisioner.ChassisConfigState")
	proto.RegisterType((*ConfigStatus)(nil), "onos.provisioner.ConfigStatus")
	proto.RegisterType((*Failure)(nil), "onos.provisioner.Failure")
}

func init() { proto.RegisterFile("onos/provisioner/aspects.proto", fileDescriptor_fa6a5635bb96ef51) }

var fileDescriptor_fa6a5635bb96ef51 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x8e, 0x93, 0x34, 0x97, 0x93, 0xb4, 0x9d, 0xce, 0xff, 0x0b, 0x85, 0x2e, 0x9c, 0x28, 0x62,
	0xd1, 0x0d, 0x8e, 0xd4, 0x8a, 0x0d, 0x42, 0x48, 0x4e, 0xec, 0x80, 0x55, 0x33, 0x8e, 0x1c, 0xbb,
	0x5c, 0x36, 0x91, 0xeb, 0x4c, 0xc3, 0x88, 0x36, 0x63, 0xc5, 0x4e, 0xa5, 0xbe, 0x45, 0x5f, 0x81,
	0x35, 0x12, 0xaf, 0xc0, 0xb6, 0xec, 0xba, 0x64, 0x15, 0x90, 0xfb, 0x16, 0xac, 0xd0, 0xd8, 0xb1,
	0x48, 0xca, 0x65, 0x41, 0x77, 0x9e, 0xf9, 0x2e, 0xe7, 0xfb, 0x8e, 0x07, 0x64, 0x3e, 0xe5, 0x61,
	0x27, 0x98, 0xf1, 0x73, 0x16, 0x32, 0x3e, 0xa5, 0xb3, 0x8e, 0x17, 0x06, 0xd4, 0x8f, 0x42, 0x25,
	0x98, 0xf1, 0x88, 0x63, 0x24, 0x70, 0x65, 0x05, 0xdf, 0xfd, 0x7f, 0xc2, 0x27, 0x3c, 0x01, 0x3b,
	0xe2, 0x2b, 0xe5, 0xed, 0x36, 0x27, 0x9c, 0x4f, 0x4e, 0x69, 0x27, 0x39, 0x1d, 0xcf, 0x4f, 0x3a,
	0x11, 0x3b, 0xa3, 0x61, 0xe4, 0x9d, 0x05, 0x29, 0xa1, 0xfd, 0x41, 0x82, 0xba, 0x46, 0xcf, 0x99,
	0x4f, 0x7b, 0x7c, 0x7a, 0xc2, 0x26, 0x98, 0x00, 0x0e, 0x58, 0x40, 0x4f, 0xd9, 0x94, 0x8e, 0xfc,
	0xe4, 0x6a, 0xc4, 0xc6, 0x0d, 0xa9, 0x25, 0xed, 0x55, 0xbb, 0xad, 0x78, 0xd1, 0x44, 0x83, 0x25,
	0x9a, 0xf2, 0x0d, 0xed, 0xfb, 0xa2, 0x59, 0xc9, 0xbe, 0x6d, 0x14, 0xac, 0xa3, 0x63, 0x7c, 0x08,
	0x3b, 0xfe, 0x5b, 0x2f, 0x0c, 0x59, 0xb8, 0x62, 0x97, 0x4f, 0xec, 0x9a, 0xf1, 0xa2, 0xb9, 0xdd,
	0x4b, 0xc1, 0xdf, 0xba, 0x6d, 0xfb, 0x6b, 0xe0, 0xb8, 0x1d, 0x4b, 0xf0, 0xdf, 0xfa, 0xfc, 0x61,
	0xe4, 0x45, 0x14, 0x3f, 0x82, 0xea, 0xed, 0xac, 0x8d, 0x78, 0xc5, 0x69, 0xcd, 0xb5, 0xe2, 0x67,
	0xd9, 0xee, 0x41, 0xc9, 0xe7, 0xfc, 0x1d, 0xa3, 0x49, 0xa0, 0xa2, 0xbd, 0x3c, 0xe1, 0xa7, 0x50,
	0x9e, 0x07, 0x63, 0x2f, 0xa2, 0xe3, 0x46, 0xa1, 0x25, 0xed, 0xd5, 0xf6, 0x77, 0x95, 0x74, 0x8f,
	0x4a, 0xb6, 0x47, 0xc5, 0xc9, 0xf6, 0xd8, 0xad, 0x5c, 0x2d, 0x9a, 0xb9, 0xcb, 0xaf, 0x4d, 0xc9,
	0xce, 0x44, 0xf8, 0x09, 0x94, 0xc2, 0xc8, 0x8b, 0xe6, 0x61, 0xa3, 0x98, 0xc8, 0x65, 0xe5, 0xf6,
	0xef, 0x52, 0x7e, 0xa6, 0x9f, 0x87, 0xdd, 0xa2, 0xb0, 0xb0, 0x97, 0x9a, 0xf6, 0x67, 0x09, 0xf0,
	0xda, 0x56, 0xee, 0xd4, 0x71, 0xa5, 0x4b, 0xfe, 0x6e, 0x5d, 0x0a, 0xff, 0xd0, 0xe5, 0xa3, 0x04,
	0xf5, 0x55, 0x18, 0x3f, 0x86, 0x0d, 0x01, 0xd1, 0xa4, 0xc1, 0xd6, 0xfe, 0x83, 0xbf, 0xbb, 0x29,
	0x49, 0x75, 0x3b, 0x95, 0xe0, 0x03, 0x28, 0x9f, 0x78, 0xec, 0x74, 0x3e, 0xa3, 0xcb, 0x2a, 0xf7,
	0x7f, 0x55, 0xf7, 0x53, 0x82, 0x9d, 0x31, 0xdb, 0x0f, 0x61, 0x23, 0xdd, 0x5f, 0x0d, 0xca, 0x03,
	0x9d, 0x68, 0x06, 0x79, 0x86, 0x72, 0xe2, 0xa0, 0x0e, 0x06, 0xa6, 0xa1, 0x6b, 0xa8, 0x80, 0x01,
	0x4a, 0x7d, 0xd5, 0x30, 0x75, 0x0d, 0x15, 0xdb, 0xef, 0xf3, 0x50, 0x5e, 0x7a, 0xe0, 0x7d, 0x28,
	0x46, 0x17, 0x41, 0x16, 0x55, 0xfe, 0xe3, 0x30, 0xc5, 0xb9, 0x08, 0xa8, 0x9d, 0x70, 0x71, 0x0b,
	0x6a, 0x63, 0x1a, 0xfa, 0x33, 0x16, 0x44, 0x8c, 0x4f, 0xd3, 0x87, 0x6e, 0xaf, 0x5e, 0xb5, 0x3f,
	0x49, 0x50, 0x14, 0x02, 0x91, 0xc1, 0x25, 0x87, 0xc4, 0x7a, 0x49, 0x50, 0x0e, 0xd7, 0xa1, 0xd2,
	0x53, 0x49, 0x4f, 0x17, 0x29, 0x24, 0xbc, 0x09, 0x55, 0x62, 0x39, 0xa3, 0xbe, 0xe5, 0x12, 0x0d,
	0xe5, 0x31, 0x86, 0x2d, 0xd5, 0xb4, 0x75, 0x55, 0x7b, 0x3d, 0xd2, 0x5f, 0x19, 0x43, 0x67, 0x88,
	0x0a, 0x18, 0x41, 0xdd, 0x25, 0xaa, 0xeb, 0x3c, 0xb7, 0x6c, 0xe3, 0x8d, 0x88, 0x2e, 0x44, 0x7d,
	0xcb, 0xee, 0x1a, 0x9a, 0xa6, 0x13, 0xb4, 0x91, 0x38, 0x5a, 0xa4, 0x6f, 0x1a, 0x3d, 0x07, 0x95,
	0xc4, 0x30, 0x83, 0x1c, 0xa9, 0xa6, 0xa1, 0xa1, 0x32, 0xde, 0x86, 0x9a, 0x4b, 0xd4, 0x23, 0xd5,
	0x30, 0xd5, 0xae, 0xa9, 0xa3, 0x0a, 0xde, 0x81, 0x4d, 0x31, 0x6f, 0xe8, 0x0e, 0x06, 0x96, 0xed,
	0xe8, 0x1a, 0xaa, 0x0a, 0x81, 0x63, 0xbc, 0xd0, 0x2d, 0xd7, 0x41, 0x20, 0xbc, 0x0c, 0xe2, 0xe8,
	0x36, 0x51, 0x4d, 0x54, 0xeb, 0x36, 0xae, 0x62, 0x59, 0xba, 0x8e, 0x65, 0xe9, 0x5b, 0x2c, 0x4b,
	0x97, 0x37, 0x72, 0xee, 0xfa, 0x46, 0xce, 0x7d, 0xb9, 0x91, 0x73, 0xc7, 0xa5, 0xe4, 0x4d, 0x1d,
	0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0x35, 0xd4, 0xd7, 0x05, 0xbf, 0x04, 0x00, 0x00,
}

func (m *DeviceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChassisConfigID) > 0 {
		i -= len(m.ChassisConfigID)
		copy(dAtA[i:], m.ChassisConfigID)
		i = encodeVarintAspects(dAtA, i, uint64(len(m.ChassisConfigID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PipelineConfigID) > 0 {
		i -= len(m.PipelineConfigID)
		copy(dAtA[i:], m.PipelineConfigID)
		i = encodeVarintAspects(dAtA, i, uint64(len(m.PipelineConfigID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PipelineConfigState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PipelineConfigState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PipelineConfigState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAspects(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Updated, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAspects(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if m.Cookie != 0 {
		i = encodeVarintAspects(dAtA, i, uint64(m.Cookie))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ConfigID) > 0 {
		i -= len(m.ConfigID)
		copy(dAtA[i:], m.ConfigID)
		i = encodeVarintAspects(dAtA, i, uint64(len(m.ConfigID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChassisConfigState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChassisConfigState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChassisConfigState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAspects(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Updated, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintAspects(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	if len(m.ConfigID) > 0 {
		i -= len(m.ConfigID)
		copy(dAtA[i:], m.ConfigID)
		i = encodeVarintAspects(dAtA, i, uint64(len(m.ConfigID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Failure != nil {
		{
			size, err := m.Failure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAspects(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.State != 0 {
		i = encodeVarintAspects(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Failure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Failure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Failure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintAspects(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintAspects(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAspects(dAtA []byte, offset int, v uint64) int {
	offset -= sovAspects(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PipelineConfigID)
	if l > 0 {
		n += 1 + l + sovAspects(uint64(l))
	}
	l = len(m.ChassisConfigID)
	if l > 0 {
		n += 1 + l + sovAspects(uint64(l))
	}
	return n
}

func (m *PipelineConfigState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConfigID)
	if l > 0 {
		n += 1 + l + sovAspects(uint64(l))
	}
	if m.Cookie != 0 {
		n += 1 + sovAspects(uint64(m.Cookie))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated)
	n += 1 + l + sovAspects(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovAspects(uint64(l))
	return n
}

func (m *ChassisConfigState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConfigID)
	if l > 0 {
		n += 1 + l + sovAspects(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Updated)
	n += 1 + l + sovAspects(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovAspects(uint64(l))
	return n
}

func (m *ConfigStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovAspects(uint64(m.State))
	}
	if m.Failure != nil {
		l = m.Failure.Size()
		n += 1 + l + sovAspects(uint64(l))
	}
	return n
}

func (m *Failure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovAspects(uint64(m.Type))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovAspects(uint64(l))
	}
	return n
}

func sovAspects(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAspects(x uint64) (n int) {
	return sovAspects(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAspects
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
			return fmt.Errorf("proto: DeviceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PipelineConfigID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PipelineConfigID = ConfigID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChassisConfigID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChassisConfigID = ConfigID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAspects(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAspects
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
func (m *PipelineConfigState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAspects
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
			return fmt.Errorf("proto: PipelineConfigState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PipelineConfigState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfigID = ConfigID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cookie", wireType)
			}
			m.Cookie = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cookie |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Updated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAspects(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAspects
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
func (m *ChassisConfigState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAspects
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
			return fmt.Errorf("proto: ChassisConfigState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChassisConfigState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfigID = ConfigID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Updated, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAspects(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAspects
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
func (m *ConfigStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAspects
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
			return fmt.Errorf("proto: ConfigStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= ConfigStatus_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Failure == nil {
				m.Failure = &Failure{}
			}
			if err := m.Failure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAspects(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAspects
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
func (m *Failure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAspects
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
			return fmt.Errorf("proto: Failure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Failure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Failure_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAspects
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
				return ErrInvalidLengthAspects
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAspects
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAspects(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAspects
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
func skipAspects(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAspects
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
					return 0, ErrIntOverflowAspects
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
					return 0, ErrIntOverflowAspects
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
				return 0, ErrInvalidLengthAspects
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAspects
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAspects
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAspects        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAspects          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAspects = fmt.Errorf("proto: unexpected end of group")
)
