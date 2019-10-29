// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	ignore "go.skia.org/infra/golden/go/ignore"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Store) Create(_a0 context.Context, _a1 *ignore.Rule) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ignore.Rule) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Store) Delete(ctx context.Context, id string) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0
func (_m *Store) List(_a0 context.Context) ([]*ignore.Rule, error) {
	ret := _m.Called(_a0)

	var r0 []*ignore.Rule
	if rf, ok := ret.Get(0).(func(context.Context) []*ignore.Rule); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ignore.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, rule
func (_m *Store) Update(ctx context.Context, id string, rule *ignore.Rule) error {
	ret := _m.Called(ctx, id, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *ignore.Rule) error); ok {
		r0 = rf(ctx, id, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
