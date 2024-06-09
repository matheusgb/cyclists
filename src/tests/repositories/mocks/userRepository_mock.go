// Code generated by MockGen. DO NOT EDIT.
// Source: src/models/repositories/user/userRepository.go
//
// Generated by this command:
//
//	mockgen -source=src/models/repositories/user/userRepository.go -destination=src/tests/mocks/userRepository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocksRepository

import (
	reflect "reflect"

	domains "github.com/matheusgb/cyclists/src/models/domains/pagination"
	domains0 "github.com/matheusgb/cyclists/src/models/domains/user"
	entities "github.com/matheusgb/cyclists/src/models/repositories/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockIUser is a mock of IUser interface.
type MockIUser struct {
	ctrl     *gomock.Controller
	recorder *MockIUserMockRecorder
}

// MockIUserMockRecorder is the mock recorder for MockIUser.
type MockIUserMockRecorder struct {
	mock *MockIUser
}

// NewMockIUser creates a new mock instance.
func NewMockIUser(ctrl *gomock.Controller) *MockIUser {
	mock := &MockIUser{ctrl: ctrl}
	mock.recorder = &MockIUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUser) EXPECT() *MockIUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUser) CreateUser(user domains0.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserMockRecorder) CreateUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUser)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockIUser) DeleteUser(user domains0.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", user)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserMockRecorder) DeleteUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUser)(nil).DeleteUser), user)
}

// FindUserByEmail mocks base method.
func (m *MockIUser) FindUserByEmail(user domains0.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", user)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockIUserMockRecorder) FindUserByEmail(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockIUser)(nil).FindUserByEmail), user)
}

// FindUserByEmailAndPassword mocks base method.
func (m *MockIUser) FindUserByEmailAndPassword(user domains0.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailAndPassword", user)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmailAndPassword indicates an expected call of FindUserByEmailAndPassword.
func (mr *MockIUserMockRecorder) FindUserByEmailAndPassword(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailAndPassword", reflect.TypeOf((*MockIUser)(nil).FindUserByEmailAndPassword), user)
}

// GetAllUsers mocks base method.
func (m *MockIUser) GetAllUsers(pag *domains.Pagination, email string) (*domains.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", pag, email)
	ret0, _ := ret[0].(*domains.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserMockRecorder) GetAllUsers(pag, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUser)(nil).GetAllUsers), pag, email)
}

// GetUser mocks base method.
func (m *MockIUser) GetUser(user domains0.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", user)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserMockRecorder) GetUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUser)(nil).GetUser), user)
}

// UpdateUser mocks base method.
func (m *MockIUser) UpdateUser(user domains0.User) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", user)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserMockRecorder) UpdateUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUser)(nil).UpdateUser), user)
}