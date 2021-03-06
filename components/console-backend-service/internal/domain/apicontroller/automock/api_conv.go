// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"
import v1alpha2 "github.com/kyma-project/kyma/components/api-controller/pkg/apis/gateway.kyma-project.io/v1alpha2"

// apiConv is an autogenerated mock type for the apiConv type
type apiConv struct {
	mock.Mock
}

// ToApi provides a mock function with given fields: name, namespace, in
func (_m *apiConv) ToApi(name string, namespace string, in gqlschema.APIInput) *v1alpha2.Api {
	ret := _m.Called(name, namespace, in)

	var r0 *v1alpha2.Api
	if rf, ok := ret.Get(0).(func(string, string, gqlschema.APIInput) *v1alpha2.Api); ok {
		r0 = rf(name, namespace, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha2.Api)
		}
	}

	return r0
}

// ToGQL provides a mock function with given fields: in
func (_m *apiConv) ToGQL(in *v1alpha2.Api) *gqlschema.API {
	ret := _m.Called(in)

	var r0 *gqlschema.API
	if rf, ok := ret.Get(0).(func(*v1alpha2.Api) *gqlschema.API); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.API)
		}
	}

	return r0
}

// ToGQLs provides a mock function with given fields: in
func (_m *apiConv) ToGQLs(in []*v1alpha2.Api) []gqlschema.API {
	ret := _m.Called(in)

	var r0 []gqlschema.API
	if rf, ok := ret.Get(0).(func([]*v1alpha2.Api) []gqlschema.API); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.API)
		}
	}

	return r0
}
