// Code generated by MockGen. DO NOT EDIT.
// Source: go/onos/config/diags/diags.pb.go

// Package diags is a generated GoMock package.
package diags

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockChangeServiceClient is a mock of ChangeServiceClient interface
type MockChangeServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockChangeServiceClientMockRecorder
}

// MockChangeServiceClientMockRecorder is the mock recorder for MockChangeServiceClient
type MockChangeServiceClientMockRecorder struct {
	mock *MockChangeServiceClient
}

// NewMockChangeServiceClient creates a new mock instance
func NewMockChangeServiceClient(ctrl *gomock.Controller) *MockChangeServiceClient {
	mock := &MockChangeServiceClient{ctrl: ctrl}
	mock.recorder = &MockChangeServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChangeServiceClient) EXPECT() *MockChangeServiceClientMockRecorder {
	return m.recorder
}

// ListNetworkChanges mocks base method
func (m *MockChangeServiceClient) ListNetworkChanges(ctx context.Context, in *ListNetworkChangeRequest, opts ...grpc.CallOption) (ChangeService_ListNetworkChangesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNetworkChanges", varargs...)
	ret0, _ := ret[0].(ChangeService_ListNetworkChangesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetworkChanges indicates an expected call of ListNetworkChanges
func (mr *MockChangeServiceClientMockRecorder) ListNetworkChanges(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetworkChanges", reflect.TypeOf((*MockChangeServiceClient)(nil).ListNetworkChanges), varargs...)
}

// ListDeviceChanges mocks base method
func (m *MockChangeServiceClient) ListDeviceChanges(ctx context.Context, in *ListDeviceChangeRequest, opts ...grpc.CallOption) (ChangeService_ListDeviceChangesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeviceChanges", varargs...)
	ret0, _ := ret[0].(ChangeService_ListDeviceChangesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeviceChanges indicates an expected call of ListDeviceChanges
func (mr *MockChangeServiceClientMockRecorder) ListDeviceChanges(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeviceChanges", reflect.TypeOf((*MockChangeServiceClient)(nil).ListDeviceChanges), varargs...)
}

// MockChangeService_ListNetworkChangesClient is a mock of ChangeService_ListNetworkChangesClient interface
type MockChangeService_ListNetworkChangesClient struct {
	ctrl     *gomock.Controller
	recorder *MockChangeService_ListNetworkChangesClientMockRecorder
}

// MockChangeService_ListNetworkChangesClientMockRecorder is the mock recorder for MockChangeService_ListNetworkChangesClient
type MockChangeService_ListNetworkChangesClientMockRecorder struct {
	mock *MockChangeService_ListNetworkChangesClient
}

// NewMockChangeService_ListNetworkChangesClient creates a new mock instance
func NewMockChangeService_ListNetworkChangesClient(ctrl *gomock.Controller) *MockChangeService_ListNetworkChangesClient {
	mock := &MockChangeService_ListNetworkChangesClient{ctrl: ctrl}
	mock.recorder = &MockChangeService_ListNetworkChangesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChangeService_ListNetworkChangesClient) EXPECT() *MockChangeService_ListNetworkChangesClientMockRecorder {
	return m.recorder
}

// Recv mocks base method
func (m *MockChangeService_ListNetworkChangesClient) Recv() (*ListNetworkChangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ListNetworkChangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockChangeService_ListNetworkChangesClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChangeService_ListNetworkChangesClient)(nil).Recv))
}

// Header mocks base method
func (m *MockChangeService_ListNetworkChangesClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockChangeService_ListNetworkChangesClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockChangeService_ListNetworkChangesClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockChangeService_ListNetworkChangesClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockChangeService_ListNetworkChangesClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockChangeService_ListNetworkChangesClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockChangeService_ListNetworkChangesClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockChangeService_ListNetworkChangesClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockChangeService_ListNetworkChangesClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockChangeService_ListNetworkChangesClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockChangeService_ListNetworkChangesClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChangeService_ListNetworkChangesClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockChangeService_ListNetworkChangesClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockChangeService_ListNetworkChangesClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChangeService_ListNetworkChangesClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockChangeService_ListNetworkChangesClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockChangeService_ListNetworkChangesClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChangeService_ListNetworkChangesClient)(nil).RecvMsg), m)
}

// MockChangeService_ListDeviceChangesClient is a mock of ChangeService_ListDeviceChangesClient interface
type MockChangeService_ListDeviceChangesClient struct {
	ctrl     *gomock.Controller
	recorder *MockChangeService_ListDeviceChangesClientMockRecorder
}

// MockChangeService_ListDeviceChangesClientMockRecorder is the mock recorder for MockChangeService_ListDeviceChangesClient
type MockChangeService_ListDeviceChangesClientMockRecorder struct {
	mock *MockChangeService_ListDeviceChangesClient
}

// NewMockChangeService_ListDeviceChangesClient creates a new mock instance
func NewMockChangeService_ListDeviceChangesClient(ctrl *gomock.Controller) *MockChangeService_ListDeviceChangesClient {
	mock := &MockChangeService_ListDeviceChangesClient{ctrl: ctrl}
	mock.recorder = &MockChangeService_ListDeviceChangesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChangeService_ListDeviceChangesClient) EXPECT() *MockChangeService_ListDeviceChangesClientMockRecorder {
	return m.recorder
}

// Recv mocks base method
func (m *MockChangeService_ListDeviceChangesClient) Recv() (*ListDeviceChangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ListDeviceChangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockChangeService_ListDeviceChangesClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChangeService_ListDeviceChangesClient)(nil).Recv))
}

// Header mocks base method
func (m *MockChangeService_ListDeviceChangesClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockChangeService_ListDeviceChangesClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockChangeService_ListDeviceChangesClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockChangeService_ListDeviceChangesClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockChangeService_ListDeviceChangesClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockChangeService_ListDeviceChangesClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockChangeService_ListDeviceChangesClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockChangeService_ListDeviceChangesClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockChangeService_ListDeviceChangesClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockChangeService_ListDeviceChangesClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockChangeService_ListDeviceChangesClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChangeService_ListDeviceChangesClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockChangeService_ListDeviceChangesClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockChangeService_ListDeviceChangesClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChangeService_ListDeviceChangesClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockChangeService_ListDeviceChangesClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockChangeService_ListDeviceChangesClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChangeService_ListDeviceChangesClient)(nil).RecvMsg), m)
}

// MockChangeServiceServer is a mock of ChangeServiceServer interface
type MockChangeServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockChangeServiceServerMockRecorder
}

// MockChangeServiceServerMockRecorder is the mock recorder for MockChangeServiceServer
type MockChangeServiceServerMockRecorder struct {
	mock *MockChangeServiceServer
}

// NewMockChangeServiceServer creates a new mock instance
func NewMockChangeServiceServer(ctrl *gomock.Controller) *MockChangeServiceServer {
	mock := &MockChangeServiceServer{ctrl: ctrl}
	mock.recorder = &MockChangeServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChangeServiceServer) EXPECT() *MockChangeServiceServerMockRecorder {
	return m.recorder
}

// ListNetworkChanges mocks base method
func (m *MockChangeServiceServer) ListNetworkChanges(arg0 *ListNetworkChangeRequest, arg1 ChangeService_ListNetworkChangesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetworkChanges", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListNetworkChanges indicates an expected call of ListNetworkChanges
func (mr *MockChangeServiceServerMockRecorder) ListNetworkChanges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetworkChanges", reflect.TypeOf((*MockChangeServiceServer)(nil).ListNetworkChanges), arg0, arg1)
}

// ListDeviceChanges mocks base method
func (m *MockChangeServiceServer) ListDeviceChanges(arg0 *ListDeviceChangeRequest, arg1 ChangeService_ListDeviceChangesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeviceChanges", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListDeviceChanges indicates an expected call of ListDeviceChanges
func (mr *MockChangeServiceServerMockRecorder) ListDeviceChanges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeviceChanges", reflect.TypeOf((*MockChangeServiceServer)(nil).ListDeviceChanges), arg0, arg1)
}

// MockChangeService_ListNetworkChangesServer is a mock of ChangeService_ListNetworkChangesServer interface
type MockChangeService_ListNetworkChangesServer struct {
	ctrl     *gomock.Controller
	recorder *MockChangeService_ListNetworkChangesServerMockRecorder
}

// MockChangeService_ListNetworkChangesServerMockRecorder is the mock recorder for MockChangeService_ListNetworkChangesServer
type MockChangeService_ListNetworkChangesServerMockRecorder struct {
	mock *MockChangeService_ListNetworkChangesServer
}

// NewMockChangeService_ListNetworkChangesServer creates a new mock instance
func NewMockChangeService_ListNetworkChangesServer(ctrl *gomock.Controller) *MockChangeService_ListNetworkChangesServer {
	mock := &MockChangeService_ListNetworkChangesServer{ctrl: ctrl}
	mock.recorder = &MockChangeService_ListNetworkChangesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChangeService_ListNetworkChangesServer) EXPECT() *MockChangeService_ListNetworkChangesServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockChangeService_ListNetworkChangesServer) Send(arg0 *ListNetworkChangeResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockChangeService_ListNetworkChangesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChangeService_ListNetworkChangesServer)(nil).Send), arg0)
}

// SetHeader mocks base method
func (m *MockChangeService_ListNetworkChangesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockChangeService_ListNetworkChangesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockChangeService_ListNetworkChangesServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockChangeService_ListNetworkChangesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockChangeService_ListNetworkChangesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockChangeService_ListNetworkChangesServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockChangeService_ListNetworkChangesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockChangeService_ListNetworkChangesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockChangeService_ListNetworkChangesServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockChangeService_ListNetworkChangesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockChangeService_ListNetworkChangesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChangeService_ListNetworkChangesServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockChangeService_ListNetworkChangesServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockChangeService_ListNetworkChangesServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChangeService_ListNetworkChangesServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockChangeService_ListNetworkChangesServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockChangeService_ListNetworkChangesServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChangeService_ListNetworkChangesServer)(nil).RecvMsg), m)
}

// MockChangeService_ListDeviceChangesServer is a mock of ChangeService_ListDeviceChangesServer interface
type MockChangeService_ListDeviceChangesServer struct {
	ctrl     *gomock.Controller
	recorder *MockChangeService_ListDeviceChangesServerMockRecorder
}

// MockChangeService_ListDeviceChangesServerMockRecorder is the mock recorder for MockChangeService_ListDeviceChangesServer
type MockChangeService_ListDeviceChangesServerMockRecorder struct {
	mock *MockChangeService_ListDeviceChangesServer
}

// NewMockChangeService_ListDeviceChangesServer creates a new mock instance
func NewMockChangeService_ListDeviceChangesServer(ctrl *gomock.Controller) *MockChangeService_ListDeviceChangesServer {
	mock := &MockChangeService_ListDeviceChangesServer{ctrl: ctrl}
	mock.recorder = &MockChangeService_ListDeviceChangesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChangeService_ListDeviceChangesServer) EXPECT() *MockChangeService_ListDeviceChangesServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockChangeService_ListDeviceChangesServer) Send(arg0 *ListDeviceChangeResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockChangeService_ListDeviceChangesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChangeService_ListDeviceChangesServer)(nil).Send), arg0)
}

// SetHeader mocks base method
func (m *MockChangeService_ListDeviceChangesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockChangeService_ListDeviceChangesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockChangeService_ListDeviceChangesServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockChangeService_ListDeviceChangesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockChangeService_ListDeviceChangesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockChangeService_ListDeviceChangesServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockChangeService_ListDeviceChangesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockChangeService_ListDeviceChangesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockChangeService_ListDeviceChangesServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockChangeService_ListDeviceChangesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockChangeService_ListDeviceChangesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChangeService_ListDeviceChangesServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockChangeService_ListDeviceChangesServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockChangeService_ListDeviceChangesServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChangeService_ListDeviceChangesServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockChangeService_ListDeviceChangesServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockChangeService_ListDeviceChangesServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChangeService_ListDeviceChangesServer)(nil).RecvMsg), m)
}

// MockOpStateDiagsClient is a mock of OpStateDiagsClient interface
type MockOpStateDiagsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpStateDiagsClientMockRecorder
}

// MockOpStateDiagsClientMockRecorder is the mock recorder for MockOpStateDiagsClient
type MockOpStateDiagsClientMockRecorder struct {
	mock *MockOpStateDiagsClient
}

// NewMockOpStateDiagsClient creates a new mock instance
func NewMockOpStateDiagsClient(ctrl *gomock.Controller) *MockOpStateDiagsClient {
	mock := &MockOpStateDiagsClient{ctrl: ctrl}
	mock.recorder = &MockOpStateDiagsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpStateDiagsClient) EXPECT() *MockOpStateDiagsClientMockRecorder {
	return m.recorder
}

// GetOpState mocks base method
func (m *MockOpStateDiagsClient) GetOpState(ctx context.Context, in *OpStateRequest, opts ...grpc.CallOption) (OpStateDiags_GetOpStateClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOpState", varargs...)
	ret0, _ := ret[0].(OpStateDiags_GetOpStateClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpState indicates an expected call of GetOpState
func (mr *MockOpStateDiagsClientMockRecorder) GetOpState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpState", reflect.TypeOf((*MockOpStateDiagsClient)(nil).GetOpState), varargs...)
}

// MockOpStateDiags_GetOpStateClient is a mock of OpStateDiags_GetOpStateClient interface
type MockOpStateDiags_GetOpStateClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpStateDiags_GetOpStateClientMockRecorder
}

// MockOpStateDiags_GetOpStateClientMockRecorder is the mock recorder for MockOpStateDiags_GetOpStateClient
type MockOpStateDiags_GetOpStateClientMockRecorder struct {
	mock *MockOpStateDiags_GetOpStateClient
}

// NewMockOpStateDiags_GetOpStateClient creates a new mock instance
func NewMockOpStateDiags_GetOpStateClient(ctrl *gomock.Controller) *MockOpStateDiags_GetOpStateClient {
	mock := &MockOpStateDiags_GetOpStateClient{ctrl: ctrl}
	mock.recorder = &MockOpStateDiags_GetOpStateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpStateDiags_GetOpStateClient) EXPECT() *MockOpStateDiags_GetOpStateClientMockRecorder {
	return m.recorder
}

// Recv mocks base method
func (m *MockOpStateDiags_GetOpStateClient) Recv() (*OpStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*OpStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockOpStateDiags_GetOpStateClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockOpStateDiags_GetOpStateClient)(nil).Recv))
}

// Header mocks base method
func (m *MockOpStateDiags_GetOpStateClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockOpStateDiags_GetOpStateClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockOpStateDiags_GetOpStateClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockOpStateDiags_GetOpStateClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockOpStateDiags_GetOpStateClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockOpStateDiags_GetOpStateClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockOpStateDiags_GetOpStateClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockOpStateDiags_GetOpStateClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockOpStateDiags_GetOpStateClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockOpStateDiags_GetOpStateClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockOpStateDiags_GetOpStateClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOpStateDiags_GetOpStateClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockOpStateDiags_GetOpStateClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockOpStateDiags_GetOpStateClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOpStateDiags_GetOpStateClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockOpStateDiags_GetOpStateClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockOpStateDiags_GetOpStateClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOpStateDiags_GetOpStateClient)(nil).RecvMsg), m)
}

// MockOpStateDiagsServer is a mock of OpStateDiagsServer interface
type MockOpStateDiagsServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpStateDiagsServerMockRecorder
}

// MockOpStateDiagsServerMockRecorder is the mock recorder for MockOpStateDiagsServer
type MockOpStateDiagsServerMockRecorder struct {
	mock *MockOpStateDiagsServer
}

// NewMockOpStateDiagsServer creates a new mock instance
func NewMockOpStateDiagsServer(ctrl *gomock.Controller) *MockOpStateDiagsServer {
	mock := &MockOpStateDiagsServer{ctrl: ctrl}
	mock.recorder = &MockOpStateDiagsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpStateDiagsServer) EXPECT() *MockOpStateDiagsServerMockRecorder {
	return m.recorder
}

// GetOpState mocks base method
func (m *MockOpStateDiagsServer) GetOpState(arg0 *OpStateRequest, arg1 OpStateDiags_GetOpStateServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOpState indicates an expected call of GetOpState
func (mr *MockOpStateDiagsServerMockRecorder) GetOpState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpState", reflect.TypeOf((*MockOpStateDiagsServer)(nil).GetOpState), arg0, arg1)
}

// MockOpStateDiags_GetOpStateServer is a mock of OpStateDiags_GetOpStateServer interface
type MockOpStateDiags_GetOpStateServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpStateDiags_GetOpStateServerMockRecorder
}

// MockOpStateDiags_GetOpStateServerMockRecorder is the mock recorder for MockOpStateDiags_GetOpStateServer
type MockOpStateDiags_GetOpStateServerMockRecorder struct {
	mock *MockOpStateDiags_GetOpStateServer
}

// NewMockOpStateDiags_GetOpStateServer creates a new mock instance
func NewMockOpStateDiags_GetOpStateServer(ctrl *gomock.Controller) *MockOpStateDiags_GetOpStateServer {
	mock := &MockOpStateDiags_GetOpStateServer{ctrl: ctrl}
	mock.recorder = &MockOpStateDiags_GetOpStateServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpStateDiags_GetOpStateServer) EXPECT() *MockOpStateDiags_GetOpStateServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockOpStateDiags_GetOpStateServer) Send(arg0 *OpStateResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockOpStateDiags_GetOpStateServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOpStateDiags_GetOpStateServer)(nil).Send), arg0)
}

// SetHeader mocks base method
func (m *MockOpStateDiags_GetOpStateServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockOpStateDiags_GetOpStateServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockOpStateDiags_GetOpStateServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockOpStateDiags_GetOpStateServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockOpStateDiags_GetOpStateServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockOpStateDiags_GetOpStateServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockOpStateDiags_GetOpStateServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockOpStateDiags_GetOpStateServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockOpStateDiags_GetOpStateServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockOpStateDiags_GetOpStateServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockOpStateDiags_GetOpStateServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockOpStateDiags_GetOpStateServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockOpStateDiags_GetOpStateServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockOpStateDiags_GetOpStateServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockOpStateDiags_GetOpStateServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockOpStateDiags_GetOpStateServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockOpStateDiags_GetOpStateServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockOpStateDiags_GetOpStateServer)(nil).RecvMsg), m)
}
