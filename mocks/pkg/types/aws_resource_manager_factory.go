// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/aws/aws-service-operator-k8s/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// AWSResourceManagerFactory is an autogenerated mock type for the AWSResourceManagerFactory type
type AWSResourceManagerFactory struct {
	mock.Mock
}

// GroupKind provides a mock function with given fields:
func (_m *AWSResourceManagerFactory) GroupKind() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ManagerFor provides a mock function with given fields: _a0
func (_m *AWSResourceManagerFactory) ManagerFor(_a0 types.AWSAccountID) (types.AWSResourceManager, error) {
	ret := _m.Called(_a0)

	var r0 types.AWSResourceManager
	if rf, ok := ret.Get(0).(func(types.AWSAccountID) types.AWSResourceManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResourceManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.AWSAccountID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceFactory provides a mock function with given fields:
func (_m *AWSResourceManagerFactory) ResourceFactory() types.AWSResourceFactory {
	ret := _m.Called()

	var r0 types.AWSResourceFactory
	if rf, ok := ret.Get(0).(func() types.AWSResourceFactory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResourceFactory)
		}
	}

	return r0
}