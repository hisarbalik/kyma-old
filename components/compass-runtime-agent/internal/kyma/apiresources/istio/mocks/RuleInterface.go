// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import v1alpha2 "github.com/kyma-project/kyma/components/application-registry/pkg/apis/istio/v1alpha2"

// RuleInterface is an autogenerated mock type for the RuleInterface type
type RuleInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *RuleInterface) Create(_a0 *v1alpha2.Rule) (*v1alpha2.Rule, error) {
	ret := _m.Called(_a0)

	var r0 *v1alpha2.Rule
	if rf, ok := ret.Get(0).(func(*v1alpha2.Rule) *v1alpha2.Rule); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha2.Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha2.Rule) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, options
func (_m *RuleInterface) Delete(name string, options *v1.DeleteOptions) error {
	ret := _m.Called(name, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1.DeleteOptions) error); ok {
		r0 = rf(name, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
