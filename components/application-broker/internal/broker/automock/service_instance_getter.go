// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"

// serviceInstanceGetter is an autogenerated mock type for the serviceInstanceGetter type
type serviceInstanceGetter struct {
	mock.Mock
}

// GetByNamespaceAndExternalID provides a mock function with given fields: namespace, extID
func (_m *serviceInstanceGetter) GetByNamespaceAndExternalID(namespace string, extID string) (*v1beta1.ServiceInstance, error) {
	ret := _m.Called(namespace, extID)

	var r0 *v1beta1.ServiceInstance
	if rf, ok := ret.Get(0).(func(string, string) *v1beta1.ServiceInstance); ok {
		r0 = rf(namespace, extID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, extID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
