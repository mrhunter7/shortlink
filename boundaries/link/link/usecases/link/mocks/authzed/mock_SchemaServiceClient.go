// Code generated by mockery v2.36.1. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// MockSchemaServiceClient is an autogenerated mock type for the SchemaServiceClient type
type MockSchemaServiceClient struct {
	mock.Mock
}

type MockSchemaServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSchemaServiceClient) EXPECT() *MockSchemaServiceClient_Expecter {
	return &MockSchemaServiceClient_Expecter{mock: &_m.Mock}
}

// ReadSchema provides a mock function with given fields: ctx, in, opts
func (_m *MockSchemaServiceClient) ReadSchema(ctx context.Context, in *v1.ReadSchemaRequest, opts ...grpc.CallOption) (*v1.ReadSchemaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ReadSchemaResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReadSchemaRequest, ...grpc.CallOption) (*v1.ReadSchemaResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReadSchemaRequest, ...grpc.CallOption) *v1.ReadSchemaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ReadSchemaResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ReadSchemaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSchemaServiceClient_ReadSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadSchema'
type MockSchemaServiceClient_ReadSchema_Call struct {
	*mock.Call
}

// ReadSchema is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.ReadSchemaRequest
//   - opts ...grpc.CallOption
func (_e *MockSchemaServiceClient_Expecter) ReadSchema(ctx interface{}, in interface{}, opts ...interface{}) *MockSchemaServiceClient_ReadSchema_Call {
	return &MockSchemaServiceClient_ReadSchema_Call{Call: _e.mock.On("ReadSchema",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockSchemaServiceClient_ReadSchema_Call) Run(run func(ctx context.Context, in *v1.ReadSchemaRequest, opts ...grpc.CallOption)) *MockSchemaServiceClient_ReadSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.ReadSchemaRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockSchemaServiceClient_ReadSchema_Call) Return(_a0 *v1.ReadSchemaResponse, _a1 error) *MockSchemaServiceClient_ReadSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSchemaServiceClient_ReadSchema_Call) RunAndReturn(run func(context.Context, *v1.ReadSchemaRequest, ...grpc.CallOption) (*v1.ReadSchemaResponse, error)) *MockSchemaServiceClient_ReadSchema_Call {
	_c.Call.Return(run)
	return _c
}

// WriteSchema provides a mock function with given fields: ctx, in, opts
func (_m *MockSchemaServiceClient) WriteSchema(ctx context.Context, in *v1.WriteSchemaRequest, opts ...grpc.CallOption) (*v1.WriteSchemaResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.WriteSchemaResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WriteSchemaRequest, ...grpc.CallOption) (*v1.WriteSchemaResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.WriteSchemaRequest, ...grpc.CallOption) *v1.WriteSchemaResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.WriteSchemaResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.WriteSchemaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSchemaServiceClient_WriteSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteSchema'
type MockSchemaServiceClient_WriteSchema_Call struct {
	*mock.Call
}

// WriteSchema is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.WriteSchemaRequest
//   - opts ...grpc.CallOption
func (_e *MockSchemaServiceClient_Expecter) WriteSchema(ctx interface{}, in interface{}, opts ...interface{}) *MockSchemaServiceClient_WriteSchema_Call {
	return &MockSchemaServiceClient_WriteSchema_Call{Call: _e.mock.On("WriteSchema",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockSchemaServiceClient_WriteSchema_Call) Run(run func(ctx context.Context, in *v1.WriteSchemaRequest, opts ...grpc.CallOption)) *MockSchemaServiceClient_WriteSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.WriteSchemaRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockSchemaServiceClient_WriteSchema_Call) Return(_a0 *v1.WriteSchemaResponse, _a1 error) *MockSchemaServiceClient_WriteSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSchemaServiceClient_WriteSchema_Call) RunAndReturn(run func(context.Context, *v1.WriteSchemaRequest, ...grpc.CallOption) (*v1.WriteSchemaResponse, error)) *MockSchemaServiceClient_WriteSchema_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSchemaServiceClient creates a new instance of MockSchemaServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSchemaServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSchemaServiceClient {
	mock := &MockSchemaServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}