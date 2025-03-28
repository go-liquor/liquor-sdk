// Code generated by MockGen. DO NOT EDIT.
// Source: secret.go
//
// Generated by this command:
//
//	mockgen -source=secret.go -package=k8s -destination=secret_mock.go
//

// Package k8s is a generated GoMock package.
package k8s

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockSecretHelper is a mock of SecretHelper interface.
type MockSecretHelper struct {
	ctrl     *gomock.Controller
	recorder *MockSecretHelperMockRecorder
	isgomock struct{}
}

// MockSecretHelperMockRecorder is the mock recorder for MockSecretHelper.
type MockSecretHelperMockRecorder struct {
	mock *MockSecretHelper
}

// NewMockSecretHelper creates a new mock instance.
func NewMockSecretHelper(ctrl *gomock.Controller) *MockSecretHelper {
	mock := &MockSecretHelper{ctrl: ctrl}
	mock.recorder = &MockSecretHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretHelper) EXPECT() *MockSecretHelperMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSecretHelper) Create(ctx context.Context, namespace string, secret *v1.Secret) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, namespace, secret)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSecretHelperMockRecorder) Create(ctx, namespace, secret any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSecretHelper)(nil).Create), ctx, namespace, secret)
}

// Delete mocks base method.
func (m *MockSecretHelper) Delete(ctx context.Context, namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSecretHelperMockRecorder) Delete(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSecretHelper)(nil).Delete), ctx, namespace, name)
}

// Get mocks base method.
func (m *MockSecretHelper) Get(ctx context.Context, namespace, name string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, namespace, name)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSecretHelperMockRecorder) Get(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretHelper)(nil).Get), ctx, namespace, name)
}

// List mocks base method.
func (m *MockSecretHelper) List(ctx context.Context, namespace string) (*v1.SecretList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, namespace)
	ret0, _ := ret[0].(*v1.SecretList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSecretHelperMockRecorder) List(ctx, namespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecretHelper)(nil).List), ctx, namespace)
}

// Update mocks base method.
func (m *MockSecretHelper) Update(ctx context.Context, namespace string, secret *v1.Secret) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, namespace, secret)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSecretHelperMockRecorder) Update(ctx, namespace, secret any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSecretHelper)(nil).Update), ctx, namespace, secret)
}
