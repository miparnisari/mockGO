// Code generated by MockGen. DO NOT EDIT.
// Source: go.opentelemetry.io/proto/otlp/collector/trace/v1 (interfaces: TraceServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

// MockTraceServiceServer is a mock of TraceServiceServer interface.
type MockTraceServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTraceServiceServerMockRecorder
}

// MockTraceServiceServerMockRecorder is the mock recorder for MockTraceServiceServer.
type MockTraceServiceServerMockRecorder struct {
	mock *MockTraceServiceServer
}

// NewMockTraceServiceServer creates a new mock instance.
func NewMockTraceServiceServer(ctrl *gomock.Controller) *MockTraceServiceServer {
	mock := &MockTraceServiceServer{ctrl: ctrl}
	mock.recorder = &MockTraceServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTraceServiceServer) EXPECT() *MockTraceServiceServerMockRecorder {
	return m.recorder
}

// Export mocks base method.
func (m *MockTraceServiceServer) Export(arg0 context.Context, arg1 *v1.ExportTraceServiceRequest) (*v1.ExportTraceServiceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", arg0, arg1)
	ret0, _ := ret[0].(*v1.ExportTraceServiceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Export indicates an expected call of Export.
func (mr *MockTraceServiceServerMockRecorder) Export(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockTraceServiceServer)(nil).Export), arg0, arg1)
}

// mustEmbedUnimplementedTraceServiceServer mocks base method.
func (m *MockTraceServiceServer) mustEmbedUnimplementedTraceServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTraceServiceServer")
}

// mustEmbedUnimplementedTraceServiceServer indicates an expected call of mustEmbedUnimplementedTraceServiceServer.
func (mr *MockTraceServiceServerMockRecorder) mustEmbedUnimplementedTraceServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTraceServiceServer", reflect.TypeOf((*MockTraceServiceServer)(nil).mustEmbedUnimplementedTraceServiceServer))
}
