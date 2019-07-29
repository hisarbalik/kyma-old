// Code generated by mockery v1.0.0
package automock

import addon "github.com/kyma-project/kyma/components/helm-broker/internal/addon"
import mock "github.com/stretchr/testify/mock"

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// ProvideAddons provides a mock function with given fields: URL
func (_m *Provider) ProvideAddons(URL string) ([]addon.CompleteAddon, error) {
	ret := _m.Called(URL)

	var r0 []addon.CompleteAddon
	if rf, ok := ret.Get(0).(func(string) []addon.CompleteAddon); ok {
		r0 = rf(URL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]addon.CompleteAddon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(URL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
