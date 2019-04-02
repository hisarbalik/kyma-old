// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// gqlSecretConverter is an autogenerated mock type for the gqlSecretConverter type
type gqlSecretConverter struct {
	mock.Mock
}

// GQLJSONToSecret provides a mock function with given fields: in
func (_m *gqlSecretConverter) GQLJSONToSecret(in gqlschema.JSON) (v1.Secret, error) {
	ret := _m.Called(in)

	var r0 v1.Secret
	if rf, ok := ret.Get(0).(func(gqlschema.JSON) v1.Secret); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(v1.Secret)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gqlschema.JSON) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlSecretConverter) ToGQL(in *v1.Secret) (*gqlschema.Secret, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.Secret
	if rf, ok := ret.Get(0).(func(*v1.Secret) *gqlschema.Secret); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Secret) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlSecretConverter) ToGQLs(in []*v1.Secret) ([]gqlschema.Secret, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.Secret
	if rf, ok := ret.Get(0).(func([]*v1.Secret) []gqlschema.Secret); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1.Secret) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
