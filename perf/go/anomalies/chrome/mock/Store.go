// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	anomalies "go.skia.org/infra/perf/go/anomalies"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// GetAnomalies provides a mock function with given fields: ctx, traceNames, startCommitPosition, endCommitPosition
func (_m *Store) GetAnomalies(ctx context.Context, traceNames []string, startCommitPosition int, endCommitPosition int) (anomalies.AnomalyMap, error) {
	ret := _m.Called(ctx, traceNames, startCommitPosition, endCommitPosition)

	var r0 anomalies.AnomalyMap
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) anomalies.AnomalyMap); ok {
		r0 = rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(anomalies.AnomalyMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) error); ok {
		r1 = rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStore creates a new instance of Store. It also registers a cleanup function to assert the mocks expectations.
func NewStore(t testing.TB) *Store {
	mock := &Store{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
