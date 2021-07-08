// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "user-service/internal/app/model"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthService is a mock of AuthService interface
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockAuthService) Decode(tokenStr string) (*model.JwtCustomClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", tokenStr)
	ret0, _ := ret[0].(*model.JwtCustomClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockAuthServiceMockRecorder) Decode(tokenStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockAuthService)(nil).Decode), tokenStr)
}

// Encode mocks base method
func (m *MockAuthService) Encode(user model.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode
func (mr *MockAuthServiceMockRecorder) Encode(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockAuthService)(nil).Encode), user)
}

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetUserByPhoneNumberAndPassword mocks base method
func (m *MockUserService) GetUserByPhoneNumberAndPassword(ctx context.Context, phone_number, password string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumberAndPassword", ctx, phone_number, password)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumberAndPassword indicates an expected call of GetUserByPhoneNumberAndPassword
func (mr *MockUserServiceMockRecorder) GetUserByPhoneNumberAndPassword(ctx, phone_number, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumberAndPassword", reflect.TypeOf((*MockUserService)(nil).GetUserByPhoneNumberAndPassword), ctx, phone_number, password)
}

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserByPhoneNumber mocks base method
func (m *MockUserRepository) GetUserByPhoneNumber(ctx context.Context, phone_number string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumber", ctx, phone_number)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumber indicates an expected call of GetUserByPhoneNumber
func (mr *MockUserRepositoryMockRecorder) GetUserByPhoneNumber(ctx, phone_number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumber", reflect.TypeOf((*MockUserRepository)(nil).GetUserByPhoneNumber), ctx, phone_number)
}
