// Code generated by mockery v1.0.0
package automock

import internal "github.com/kyma-project/kyma/components/helm-broker/internal"
import mock "github.com/stretchr/testify/mock"
import semver "github.com/Masterminds/semver"

// AddonStorage is an autogenerated mock type for the AddonStorage type
type AddonStorage struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: _a0
func (_m *AddonStorage) FindAll(_a0 internal.Namespace) ([]*internal.Addon, error) {
	ret := _m.Called(_a0)

	var r0 []*internal.Addon
	if rf, ok := ret.Get(0).(func(internal.Namespace) []*internal.Addon); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*internal.Addon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.Namespace) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1, _a2
func (_m *AddonStorage) Get(_a0 internal.Namespace, _a1 internal.AddonName, _a2 semver.Version) (*internal.Addon, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *internal.Addon
	if rf, ok := ret.Get(0).(func(internal.Namespace, internal.AddonName, semver.Version) *internal.Addon); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.Addon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.Namespace, internal.AddonName, semver.Version) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: _a0, _a1, _a2
func (_m *AddonStorage) Remove(_a0 internal.Namespace, _a1 internal.AddonName, _a2 semver.Version) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(internal.Namespace, internal.AddonName, semver.Version) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *AddonStorage) Upsert(_a0 internal.Namespace, _a1 *internal.Addon) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(internal.Namespace, *internal.Addon) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.Namespace, *internal.Addon) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
