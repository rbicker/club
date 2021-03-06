// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	context "context"

	gooserv1 "github.com/rbicker/gooser/api/proto/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// GooserClient is an autogenerated mock type for the GooserClient type
type GooserClient struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) ChangePassword(ctx context.Context, in *gooserv1.ChangePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.ChangePasswordRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.ChangePasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfirmMail provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) ConfirmMail(ctx context.Context, in *gooserv1.ConfirmMailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.ConfirmMailRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.ConfirmMailRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroup provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) CreateGroup(ctx context.Context, in *gooserv1.Group, opts ...grpc.CallOption) (*gooserv1.Group, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.Group
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.Group, ...grpc.CallOption) *gooserv1.Group); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.Group, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) CreateUser(ctx context.Context, in *gooserv1.User, opts ...grpc.CallOption) (*gooserv1.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.User
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.User, ...grpc.CallOption) *gooserv1.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.User, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroup provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) DeleteGroup(ctx context.Context, in *gooserv1.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) DeleteUser(ctx context.Context, in *gooserv1.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgotPassword provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) ForgotPassword(ctx context.Context, in *gooserv1.ForgotPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.ForgotPasswordRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.ForgotPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroup provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) GetGroup(ctx context.Context, in *gooserv1.IdRequest, opts ...grpc.CallOption) (*gooserv1.Group, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.Group
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) *gooserv1.Group); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) GetUser(ctx context.Context, in *gooserv1.IdRequest, opts ...grpc.CallOption) (*gooserv1.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.User
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) *gooserv1.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.IdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGroups provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) ListGroups(ctx context.Context, in *gooserv1.ListRequest, opts ...grpc.CallOption) (*gooserv1.ListGroupsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.ListGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.ListRequest, ...grpc.CallOption) *gooserv1.ListGroupsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.ListGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.ListRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) ListUsers(ctx context.Context, in *gooserv1.ListRequest, opts ...grpc.CallOption) (*gooserv1.ListUsersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.ListUsersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.ListRequest, ...grpc.CallOption) *gooserv1.ListUsersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.ListUsersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.ListRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) ResetPassword(ctx context.Context, in *gooserv1.ResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.ResetPasswordRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.ResetPasswordRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGroup provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) UpdateGroup(ctx context.Context, in *gooserv1.UpdateGroupRequest, opts ...grpc.CallOption) (*gooserv1.Group, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.Group
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.UpdateGroupRequest, ...grpc.CallOption) *gooserv1.Group); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.UpdateGroupRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, in, opts
func (_m *GooserClient) UpdateUser(ctx context.Context, in *gooserv1.UpdateUserRequest, opts ...grpc.CallOption) (*gooserv1.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gooserv1.User
	if rf, ok := ret.Get(0).(func(context.Context, *gooserv1.UpdateUserRequest, ...grpc.CallOption) *gooserv1.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gooserv1.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gooserv1.UpdateUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
