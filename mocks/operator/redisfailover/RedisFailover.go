// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import redisfailoverv1 "github.com/spotahome/redis-operator/api/redisfailover/v1"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import watch "k8s.io/apimachinery/pkg/watch"

// RedisFailover is an autogenerated mock type for the RedisFailover type
type RedisFailover struct {
	mock.Mock
}

// ListRedisFailovers provides a mock function with given fields: namespace, opts
func (_m *RedisFailover) ListRedisFailovers(namespace string, opts v1.ListOptions) (*redisfailoverv1.RedisFailoverList, error) {
	ret := _m.Called(namespace, opts)

	var r0 *redisfailoverv1.RedisFailoverList
	if rf, ok := ret.Get(0).(func(string, v1.ListOptions) *redisfailoverv1.RedisFailoverList); ok {
		r0 = rf(namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redisfailoverv1.RedisFailoverList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, v1.ListOptions) error); ok {
		r1 = rf(namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRedisFailovers provides a mock function with given fields: namespace, opts
func (_m *RedisFailover) WatchRedisFailovers(namespace string, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(namespace, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(string, v1.ListOptions) watch.Interface); ok {
		r0 = rf(namespace, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, v1.ListOptions) error); ok {
		r1 = rf(namespace, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
