// Copyright 2022 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: internal/proto/torchserve/inference_grpc.pb.go

// Package mock_torchserve is a generated GoMock package.
package mock_torchserve

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	torchserve "github.com/kserve/modelmesh-runtime-adapter/internal/proto/torchserve"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockInferenceAPIsServiceClient is a mock of InferenceAPIsServiceClient interface.
type MockInferenceAPIsServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockInferenceAPIsServiceClientMockRecorder
}

// MockInferenceAPIsServiceClientMockRecorder is the mock recorder for MockInferenceAPIsServiceClient.
type MockInferenceAPIsServiceClientMockRecorder struct {
	mock *MockInferenceAPIsServiceClient
}

// NewMockInferenceAPIsServiceClient creates a new mock instance.
func NewMockInferenceAPIsServiceClient(ctrl *gomock.Controller) *MockInferenceAPIsServiceClient {
	mock := &MockInferenceAPIsServiceClient{ctrl: ctrl}
	mock.recorder = &MockInferenceAPIsServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInferenceAPIsServiceClient) EXPECT() *MockInferenceAPIsServiceClientMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockInferenceAPIsServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*torchserve.TorchServeHealthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ping", varargs...)
	ret0, _ := ret[0].(*torchserve.TorchServeHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockInferenceAPIsServiceClientMockRecorder) Ping(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockInferenceAPIsServiceClient)(nil).Ping), varargs...)
}

// Predictions mocks base method.
func (m *MockInferenceAPIsServiceClient) Predictions(ctx context.Context, in *torchserve.PredictionsRequest, opts ...grpc.CallOption) (*torchserve.PredictionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Predictions", varargs...)
	ret0, _ := ret[0].(*torchserve.PredictionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Predictions indicates an expected call of Predictions.
func (mr *MockInferenceAPIsServiceClientMockRecorder) Predictions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Predictions", reflect.TypeOf((*MockInferenceAPIsServiceClient)(nil).Predictions), varargs...)
}

// MockInferenceAPIsServiceServer is a mock of InferenceAPIsServiceServer interface.
type MockInferenceAPIsServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockInferenceAPIsServiceServerMockRecorder
}

// MockInferenceAPIsServiceServerMockRecorder is the mock recorder for MockInferenceAPIsServiceServer.
type MockInferenceAPIsServiceServerMockRecorder struct {
	mock *MockInferenceAPIsServiceServer
}

// NewMockInferenceAPIsServiceServer creates a new mock instance.
func NewMockInferenceAPIsServiceServer(ctrl *gomock.Controller) *MockInferenceAPIsServiceServer {
	mock := &MockInferenceAPIsServiceServer{ctrl: ctrl}
	mock.recorder = &MockInferenceAPIsServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInferenceAPIsServiceServer) EXPECT() *MockInferenceAPIsServiceServerMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockInferenceAPIsServiceServer) Ping(arg0 context.Context, arg1 *emptypb.Empty) (*torchserve.TorchServeHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0, arg1)
	ret0, _ := ret[0].(*torchserve.TorchServeHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockInferenceAPIsServiceServerMockRecorder) Ping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockInferenceAPIsServiceServer)(nil).Ping), arg0, arg1)
}

// Predictions mocks base method.
func (m *MockInferenceAPIsServiceServer) Predictions(arg0 context.Context, arg1 *torchserve.PredictionsRequest) (*torchserve.PredictionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Predictions", arg0, arg1)
	ret0, _ := ret[0].(*torchserve.PredictionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Predictions indicates an expected call of Predictions.
func (mr *MockInferenceAPIsServiceServerMockRecorder) Predictions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Predictions", reflect.TypeOf((*MockInferenceAPIsServiceServer)(nil).Predictions), arg0, arg1)
}

// mustEmbedUnimplementedInferenceAPIsServiceServer mocks base method.
func (m *MockInferenceAPIsServiceServer) mustEmbedUnimplementedInferenceAPIsServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInferenceAPIsServiceServer")
}

// mustEmbedUnimplementedInferenceAPIsServiceServer indicates an expected call of mustEmbedUnimplementedInferenceAPIsServiceServer.
func (mr *MockInferenceAPIsServiceServerMockRecorder) mustEmbedUnimplementedInferenceAPIsServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInferenceAPIsServiceServer", reflect.TypeOf((*MockInferenceAPIsServiceServer)(nil).mustEmbedUnimplementedInferenceAPIsServiceServer))
}

// MockUnsafeInferenceAPIsServiceServer is a mock of UnsafeInferenceAPIsServiceServer interface.
type MockUnsafeInferenceAPIsServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeInferenceAPIsServiceServerMockRecorder
}

// MockUnsafeInferenceAPIsServiceServerMockRecorder is the mock recorder for MockUnsafeInferenceAPIsServiceServer.
type MockUnsafeInferenceAPIsServiceServerMockRecorder struct {
	mock *MockUnsafeInferenceAPIsServiceServer
}

// NewMockUnsafeInferenceAPIsServiceServer creates a new mock instance.
func NewMockUnsafeInferenceAPIsServiceServer(ctrl *gomock.Controller) *MockUnsafeInferenceAPIsServiceServer {
	mock := &MockUnsafeInferenceAPIsServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeInferenceAPIsServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeInferenceAPIsServiceServer) EXPECT() *MockUnsafeInferenceAPIsServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedInferenceAPIsServiceServer mocks base method.
func (m *MockUnsafeInferenceAPIsServiceServer) mustEmbedUnimplementedInferenceAPIsServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInferenceAPIsServiceServer")
}

// mustEmbedUnimplementedInferenceAPIsServiceServer indicates an expected call of mustEmbedUnimplementedInferenceAPIsServiceServer.
func (mr *MockUnsafeInferenceAPIsServiceServerMockRecorder) mustEmbedUnimplementedInferenceAPIsServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInferenceAPIsServiceServer", reflect.TypeOf((*MockUnsafeInferenceAPIsServiceServer)(nil).mustEmbedUnimplementedInferenceAPIsServiceServer))
}
