// Automatically generated by MockGen. DO NOT EDIT!
// Source: firebase.go

package mock_firebase

import (
	gomock "code.google.com/p/gomock/gomock"
	. "github.com/JustinTulloss/firebase"
)

// Mock of Api interface
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *_MockApiRecorder
}

// Recorder for MockApi (not exported)
type _MockApiRecorder struct {
	mock *MockApi
}

func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &_MockApiRecorder{mock}
	return mock
}

func (_m *MockApi) EXPECT() *_MockApiRecorder {
	return _m.recorder
}

func (_m *MockApi) Call(method string, path string, auth string, body interface{}, params map[string]string, dest interface{}) error {
	ret := _m.ctrl.Call(_m, "Call", method, path, auth, body, params, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockApiRecorder) Call(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Call", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) String() string {
	ret := _m.ctrl.Call(_m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockClientRecorder) String() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "String")
}

func (_m *MockClient) Key() string {
	ret := _m.ctrl.Call(_m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockClientRecorder) Key() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Key")
}

func (_m *MockClient) Value(destination interface{}) error {
	ret := _m.ctrl.Call(_m, "Value", destination)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Value(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Value", arg0)
}

func (_m *MockClient) Iterator(d Destination) <-chan *KeyedValue {
	ret := _m.ctrl.Call(_m, "Iterator", d)
	ret0, _ := ret[0].(<-chan *KeyedValue)
	return ret0
}

func (_mr *_MockClientRecorder) Iterator(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iterator", arg0)
}

func (_m *MockClient) Shallow() ([]string, error) {
	ret := _m.ctrl.Call(_m, "Shallow")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Shallow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shallow")
}

func (_m *MockClient) Child(path string) Client {
	ret := _m.ctrl.Call(_m, "Child", path)
	ret0, _ := ret[0].(Client)
	return ret0
}

func (_mr *_MockClientRecorder) Child(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Child", arg0)
}

func (_m *MockClient) OrderBy(prop string) Client {
	ret := _m.ctrl.Call(_m, "OrderBy", prop)
	ret0, _ := ret[0].(Client)
	return ret0
}

func (_mr *_MockClientRecorder) OrderBy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OrderBy", arg0)
}

func (_m *MockClient) EqualTo(value string) Client {
	ret := _m.ctrl.Call(_m, "EqualTo", value)
	ret0, _ := ret[0].(Client)
	return ret0
}

func (_mr *_MockClientRecorder) EqualTo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EqualTo", arg0)
}

func (_m *MockClient) StartAt(value string) Client {
	ret := _m.ctrl.Call(_m, "StartAt", value)
	ret0, _ := ret[0].(Client)
	return ret0
}

func (_mr *_MockClientRecorder) StartAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartAt", arg0)
}

func (_m *MockClient) EndAt(value string) Client {
	ret := _m.ctrl.Call(_m, "EndAt", value)
	ret0, _ := ret[0].(Client)
	return ret0
}

func (_mr *_MockClientRecorder) EndAt(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EndAt", arg0)
}

func (_m *MockClient) LimitToFirst(value int) Client {
	ret := _m.ctrl.Call(_m, "LimitToFirst", value)
	ret0, _ := ret[0].(Client)
	return ret0
}

func (_mr *_MockClientRecorder) LimitToFirst(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LimitToFirst", arg0)
}

func (_m *MockClient) LimitToLast(value int) Client {
	ret := _m.ctrl.Call(_m, "LimitToLast", value)
	ret0, _ := ret[0].(Client)
	return ret0
}

func (_mr *_MockClientRecorder) LimitToLast(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LimitToLast", arg0)
}

func (_m *MockClient) Push(value interface{}, params map[string]string) (Client, error) {
	ret := _m.ctrl.Call(_m, "Push", value, params)
	ret0, _ := ret[0].(Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Push(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Push", arg0, arg1)
}

func (_m *MockClient) Set(path string, value interface{}, params map[string]string) (Client, error) {
	ret := _m.ctrl.Call(_m, "Set", path, value, params)
	ret0, _ := ret[0].(Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Set", arg0, arg1, arg2)
}

func (_m *MockClient) Update(path string, value interface{}, params map[string]string) error {
	ret := _m.ctrl.Call(_m, "Update", path, value, params)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1, arg2)
}

func (_m *MockClient) Remove(path string, params map[string]string) error {
	ret := _m.ctrl.Call(_m, "Remove", path, params)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0, arg1)
}

func (_m *MockClient) Rules(params map[string]string) (*Rules, error) {
	ret := _m.ctrl.Call(_m, "Rules", params)
	ret0, _ := ret[0].(*Rules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Rules(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rules", arg0)
}

func (_m *MockClient) SetRules(rules *Rules, params map[string]string) error {
	ret := _m.ctrl.Call(_m, "SetRules", rules, params)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) SetRules(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRules", arg0, arg1)
}
