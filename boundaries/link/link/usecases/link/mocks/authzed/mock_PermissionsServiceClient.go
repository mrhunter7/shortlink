// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// MockPermissionsServiceClient is an autogenerated mock type for the PermissionsServiceClient type
type MockPermissionsServiceClient struct {
	mock.Mock
}

type MockPermissionsServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPermissionsServiceClient) EXPECT() *MockPermissionsServiceClient_Expecter {
	return &MockPermissionsServiceClient_Expecter{mock: &_m.Mock}
}

// CheckPermission provides a mock function with given fields: ctx, in, opts
func (_m *MockPermissionsServiceClient) CheckPermission(ctx context.Context, in *v1.CheckPermissionRequest, opts ...grpc.CallOption) (*v1.CheckPermissionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.CheckPermissionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CheckPermissionRequest, ...grpc.CallOption) (*v1.CheckPermissionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CheckPermissionRequest, ...grpc.CallOption) *v1.CheckPermissionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.CheckPermissionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.CheckPermissionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsServiceClient_CheckPermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckPermission'
type MockPermissionsServiceClient_CheckPermission_Call struct {
	*mock.Call
}

// CheckPermission is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.CheckPermissionRequest
//   - opts ...grpc.CallOption
func (_e *MockPermissionsServiceClient_Expecter) CheckPermission(ctx interface{}, in interface{}, opts ...interface{}) *MockPermissionsServiceClient_CheckPermission_Call {
	return &MockPermissionsServiceClient_CheckPermission_Call{Call: _e.mock.On("CheckPermission",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockPermissionsServiceClient_CheckPermission_Call) Run(run func(ctx context.Context, in *v1.CheckPermissionRequest, opts ...grpc.CallOption)) *MockPermissionsServiceClient_CheckPermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.CheckPermissionRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockPermissionsServiceClient_CheckPermission_Call) Return(_a0 *v1.CheckPermissionResponse, _a1 error) *MockPermissionsServiceClient_CheckPermission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsServiceClient_CheckPermission_Call) RunAndReturn(run func(context.Context, *v1.CheckPermissionRequest, ...grpc.CallOption) (*v1.CheckPermissionResponse, error)) *MockPermissionsServiceClient_CheckPermission_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRelationships provides a mock function with given fields: ctx, in, opts
func (_m *MockPermissionsServiceClient) DeleteRelationships(ctx context.Context, in *v1.DeleteRelationshipsRequest, opts ...grpc.CallOption) (*v1.DeleteRelationshipsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.DeleteRelationshipsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.DeleteRelationshipsRequest, ...grpc.CallOption) (*v1.DeleteRelationshipsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.DeleteRelationshipsRequest, ...grpc.CallOption) *v1.DeleteRelationshipsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.DeleteRelationshipsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.DeleteRelationshipsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsServiceClient_DeleteRelationships_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRelationships'
type MockPermissionsServiceClient_DeleteRelationships_Call struct {
	*mock.Call
}

// DeleteRelationships is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.DeleteRelationshipsRequest
//   - opts ...grpc.CallOption
func (_e *MockPermissionsServiceClient_Expecter) DeleteRelationships(ctx interface{}, in interface{}, opts ...interface{}) *MockPermissionsServiceClient_DeleteRelationships_Call {
	return &MockPermissionsServiceClient_DeleteRelationships_Call{Call: _e.mock.On("DeleteRelationships",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockPermissionsServiceClient_DeleteRelationships_Call) Run(run func(ctx context.Context, in *v1.DeleteRelationshipsRequest, opts ...grpc.CallOption)) *MockPermissionsServiceClient_DeleteRelationships_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.DeleteRelationshipsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockPermissionsServiceClient_DeleteRelationships_Call) Return(_a0 *v1.DeleteRelationshipsResponse, _a1 error) *MockPermissionsServiceClient_DeleteRelationships_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsServiceClient_DeleteRelationships_Call) RunAndReturn(run func(context.Context, *v1.DeleteRelationshipsRequest, ...grpc.CallOption) (*v1.DeleteRelationshipsResponse, error)) *MockPermissionsServiceClient_DeleteRelationships_Call {
	_c.Call.Return(run)
	return _c
}

// ExpandPermissionTree provides a mock function with given fields: ctx, in, opts
func (_m *MockPermissionsServiceClient) ExpandPermissionTree(ctx context.Context, in *v1.ExpandPermissionTreeRequest, opts ...grpc.CallOption) (*v1.ExpandPermissionTreeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ExpandPermissionTreeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ExpandPermissionTreeRequest, ...grpc.CallOption) (*v1.ExpandPermissionTreeResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ExpandPermissionTreeRequest, ...grpc.CallOption) *v1.ExpandPermissionTreeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ExpandPermissionTreeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ExpandPermissionTreeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsServiceClient_ExpandPermissionTree_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpandPermissionTree'
type MockPermissionsServiceClient_ExpandPermissionTree_Call struct {
	*mock.Call
}

// ExpandPermissionTree is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.ExpandPermissionTreeRequest
//   - opts ...grpc.CallOption
func (_e *MockPermissionsServiceClient_Expecter) ExpandPermissionTree(ctx interface{}, in interface{}, opts ...interface{}) *MockPermissionsServiceClient_ExpandPermissionTree_Call {
	return &MockPermissionsServiceClient_ExpandPermissionTree_Call{Call: _e.mock.On("ExpandPermissionTree",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockPermissionsServiceClient_ExpandPermissionTree_Call) Run(run func(ctx context.Context, in *v1.ExpandPermissionTreeRequest, opts ...grpc.CallOption)) *MockPermissionsServiceClient_ExpandPermissionTree_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.ExpandPermissionTreeRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockPermissionsServiceClient_ExpandPermissionTree_Call) Return(_a0 *v1.ExpandPermissionTreeResponse, _a1 error) *MockPermissionsServiceClient_ExpandPermissionTree_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsServiceClient_ExpandPermissionTree_Call) RunAndReturn(run func(context.Context, *v1.ExpandPermissionTreeRequest, ...grpc.CallOption) (*v1.ExpandPermissionTreeResponse, error)) *MockPermissionsServiceClient_ExpandPermissionTree_Call {
	_c.Call.Return(run)
	return _c
}

// LookupResources provides a mock function with given fields: ctx, in, opts
func (_m *MockPermissionsServiceClient) LookupResources(ctx context.Context, in *v1.LookupResourcesRequest, opts ...grpc.CallOption) (v1.PermissionsService_LookupResourcesClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 v1.PermissionsService_LookupResourcesClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LookupResourcesRequest, ...grpc.CallOption) (v1.PermissionsService_LookupResourcesClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LookupResourcesRequest, ...grpc.CallOption) v1.PermissionsService_LookupResourcesClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PermissionsService_LookupResourcesClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.LookupResourcesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsServiceClient_LookupResources_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LookupResources'
type MockPermissionsServiceClient_LookupResources_Call struct {
	*mock.Call
}

// LookupResources is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.LookupResourcesRequest
//   - opts ...grpc.CallOption
func (_e *MockPermissionsServiceClient_Expecter) LookupResources(ctx interface{}, in interface{}, opts ...interface{}) *MockPermissionsServiceClient_LookupResources_Call {
	return &MockPermissionsServiceClient_LookupResources_Call{Call: _e.mock.On("LookupResources",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockPermissionsServiceClient_LookupResources_Call) Run(run func(ctx context.Context, in *v1.LookupResourcesRequest, opts ...grpc.CallOption)) *MockPermissionsServiceClient_LookupResources_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.LookupResourcesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockPermissionsServiceClient_LookupResources_Call) Return(_a0 v1.PermissionsService_LookupResourcesClient, _a1 error) *MockPermissionsServiceClient_LookupResources_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsServiceClient_LookupResources_Call) RunAndReturn(run func(context.Context, *v1.LookupResourcesRequest, ...grpc.CallOption) (v1.PermissionsService_LookupResourcesClient, error)) *MockPermissionsServiceClient_LookupResources_Call {
	_c.Call.Return(run)
	return _c
}

// LookupSubjects provides a mock function with given fields: ctx, in, opts
func (_m *MockPermissionsServiceClient) LookupSubjects(ctx context.Context, in *v1.LookupSubjectsRequest, opts ...grpc.CallOption) (v1.PermissionsService_LookupSubjectsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 v1.PermissionsService_LookupSubjectsClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LookupSubjectsRequest, ...grpc.CallOption) (v1.PermissionsService_LookupSubjectsClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LookupSubjectsRequest, ...grpc.CallOption) v1.PermissionsService_LookupSubjectsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PermissionsService_LookupSubjectsClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.LookupSubjectsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsServiceClient_LookupSubjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LookupSubjects'
type MockPermissionsServiceClient_LookupSubjects_Call struct {
	*mock.Call
}

// LookupSubjects is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.LookupSubjectsRequest
//   - opts ...grpc.CallOption
func (_e *MockPermissionsServiceClient_Expecter) LookupSubjects(ctx interface{}, in interface{}, opts ...interface{}) *MockPermissionsServiceClient_LookupSubjects_Call {
	return &MockPermissionsServiceClient_LookupSubjects_Call{Call: _e.mock.On("LookupSubjects",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockPermissionsServiceClient_LookupSubjects_Call) Run(run func(ctx context.Context, in *v1.LookupSubjectsRequest, opts ...grpc.CallOption)) *MockPermissionsServiceClient_LookupSubjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.LookupSubjectsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockPermissionsServiceClient_LookupSubjects_Call) Return(_a0 v1.PermissionsService_LookupSubjectsClient, _a1 error) *MockPermissionsServiceClient_LookupSubjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsServiceClient_LookupSubjects_Call) RunAndReturn(run func(context.Context, *v1.LookupSubjectsRequest, ...grpc.CallOption) (v1.PermissionsService_LookupSubjectsClient, error)) *MockPermissionsServiceClient_LookupSubjects_Call {
	_c.Call.Return(run)
	return _c
}

// ReadRelationships provides a mock function with given fields: ctx, in, opts
func (_m *MockPermissionsServiceClient) ReadRelationships(ctx context.Context, in *v1.ReadRelationshipsRequest, opts ...grpc.CallOption) (v1.PermissionsService_ReadRelationshipsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 v1.PermissionsService_ReadRelationshipsClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReadRelationshipsRequest, ...grpc.CallOption) (v1.PermissionsService_ReadRelationshipsClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReadRelationshipsRequest, ...grpc.CallOption) v1.PermissionsService_ReadRelationshipsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PermissionsService_ReadRelationshipsClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ReadRelationshipsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsServiceClient_ReadRelationships_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadRelationships'
type MockPermissionsServiceClient_ReadRelationships_Call struct {
	*mock.Call
}

// ReadRelationships is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.ReadRelationshipsRequest
//   - opts ...grpc.CallOption
func (_e *MockPermissionsServiceClient_Expecter) ReadRelationships(ctx interface{}, in interface{}, opts ...interface{}) *MockPermissionsServiceClient_ReadRelationships_Call {
	return &MockPermissionsServiceClient_ReadRelationships_Call{Call: _e.mock.On("ReadRelationships",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockPermissionsServiceClient_ReadRelationships_Call) Run(run func(ctx context.Context, in *v1.ReadRelationshipsRequest, opts ...grpc.CallOption)) *MockPermissionsServiceClient_ReadRelationships_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.ReadRelationshipsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockPermissionsServiceClient_ReadRelationships_Call) Return(_a0 v1.PermissionsService_ReadRelationshipsClient, _a1 error) *MockPermissionsServiceClient_ReadRelationships_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsServiceClient_ReadRelationships_Call) RunAndReturn(run func(context.Context, *v1.ReadRelationshipsRequest, ...grpc.CallOption) (v1.PermissionsService_ReadRelationshipsClient, error)) *MockPermissionsServiceClient_ReadRelationships_Call {
	_c.Call.Return(run)
	return _c
}

// WriteRelationships provides a mock function with given fields: ctx, in, opts
func (_m *MockPermissionsServiceClient) WriteRelationships(ctx context.Context, in *v1.WriteRelationshipsRequest, opts ...grpc.CallOption) (*v1.WriteRelationshipsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.WriteRelationshipsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WriteRelationshipsRequest, ...grpc.CallOption) (*v1.WriteRelationshipsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WriteRelationshipsRequest, ...grpc.CallOption) *v1.WriteRelationshipsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.WriteRelationshipsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.WriteRelationshipsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsServiceClient_WriteRelationships_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteRelationships'
type MockPermissionsServiceClient_WriteRelationships_Call struct {
	*mock.Call
}

// WriteRelationships is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.WriteRelationshipsRequest
//   - opts ...grpc.CallOption
func (_e *MockPermissionsServiceClient_Expecter) WriteRelationships(ctx interface{}, in interface{}, opts ...interface{}) *MockPermissionsServiceClient_WriteRelationships_Call {
	return &MockPermissionsServiceClient_WriteRelationships_Call{Call: _e.mock.On("WriteRelationships",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockPermissionsServiceClient_WriteRelationships_Call) Run(run func(ctx context.Context, in *v1.WriteRelationshipsRequest, opts ...grpc.CallOption)) *MockPermissionsServiceClient_WriteRelationships_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.WriteRelationshipsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockPermissionsServiceClient_WriteRelationships_Call) Return(_a0 *v1.WriteRelationshipsResponse, _a1 error) *MockPermissionsServiceClient_WriteRelationships_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsServiceClient_WriteRelationships_Call) RunAndReturn(run func(context.Context, *v1.WriteRelationshipsRequest, ...grpc.CallOption) (*v1.WriteRelationshipsResponse, error)) *MockPermissionsServiceClient_WriteRelationships_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPermissionsServiceClient creates a new instance of MockPermissionsServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPermissionsServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPermissionsServiceClient {
	mock := &MockPermissionsServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}