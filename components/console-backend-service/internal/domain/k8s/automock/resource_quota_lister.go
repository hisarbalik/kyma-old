// Code generated by mockery v1.0.0
package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// resourceQuotaLister is an autogenerated mock type for the resourceQuotaLister type
type resourceQuotaLister struct {
	mock.Mock
}

// CreateResourceQuota provides a mock function with given fields: namespace, name, ResourceQuotaInput
func (_m *resourceQuotaLister) CreateResourceQuota(namespace string, name string, ResourceQuotaInput gqlschema.ResourceQuotaInput) (*v1.ResourceQuota, error) {
	ret := _m.Called(namespace, name, ResourceQuotaInput)

	var r0 *v1.ResourceQuota
	if rf, ok := ret.Get(0).(func(string, string, gqlschema.ResourceQuotaInput) *v1.ResourceQuota); ok {
		r0 = rf(namespace, name, ResourceQuotaInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ResourceQuota)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, gqlschema.ResourceQuotaInput) error); ok {
		r1 = rf(namespace, name, ResourceQuotaInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListResourceQuotas provides a mock function with given fields: namespace
func (_m *resourceQuotaLister) ListResourceQuotas(namespace string) ([]*v1.ResourceQuota, error) {
	ret := _m.Called(namespace)

	var r0 []*v1.ResourceQuota
	if rf, ok := ret.Get(0).(func(string) []*v1.ResourceQuota); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.ResourceQuota)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
