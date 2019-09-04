// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import pager "github.com/kyma-project/kyma/components/console-backend-service/internal/pager"

import v1alpha1 "github.com/kyma-project/kyma/components/helm-broker/pkg/apis/addons/v1alpha1"

// clusterAddonsCfgLister is an autogenerated mock type for the clusterAddonsCfgLister type
type clusterAddonsCfgLister struct {
	mock.Mock
}

// List provides a mock function with given fields: pagingParams
func (_m *clusterAddonsCfgLister) List(pagingParams pager.PagingParams) ([]*v1alpha1.ClusterAddonsConfiguration, error) {
	ret := _m.Called(pagingParams)

	var r0 []*v1alpha1.ClusterAddonsConfiguration
	if rf, ok := ret.Get(0).(func(pager.PagingParams) []*v1alpha1.ClusterAddonsConfiguration); ok {
		r0 = rf(pagingParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.ClusterAddonsConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pager.PagingParams) error); ok {
		r1 = rf(pagingParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
