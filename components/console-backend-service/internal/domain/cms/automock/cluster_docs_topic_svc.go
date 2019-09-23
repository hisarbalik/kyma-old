// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import resource "github.com/kyma-project/kyma/components/console-backend-service/pkg/resource"
import v1alpha1 "github.com/kyma-project/kyma/components/cms-controller-manager/pkg/apis/cms/v1alpha1"

// clusterDocsTopicSvc is an autogenerated mock type for the clusterDocsTopicSvc type
type clusterDocsTopicSvc struct {
	mock.Mock
}

// Find provides a mock function with given fields: name
func (_m *clusterDocsTopicSvc) Find(name string) (*v1alpha1.ClusterDocsTopic, error) {
	ret := _m.Called(name)

	var r0 *v1alpha1.ClusterDocsTopic
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.ClusterDocsTopic); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ClusterDocsTopic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: viewContext, groupName
func (_m *clusterDocsTopicSvc) List(viewContext *string, groupName *string) ([]*v1alpha1.ClusterDocsTopic, error) {
	ret := _m.Called(viewContext, groupName)

	var r0 []*v1alpha1.ClusterDocsTopic
	if rf, ok := ret.Get(0).(func(*string, *string) []*v1alpha1.ClusterDocsTopic); ok {
		r0 = rf(viewContext, groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.ClusterDocsTopic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(viewContext, groupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: listener
func (_m *clusterDocsTopicSvc) Subscribe(listener resource.Listener) {
	_m.Called(listener)
}

// Unsubscribe provides a mock function with given fields: listener
func (_m *clusterDocsTopicSvc) Unsubscribe(listener resource.Listener) {
	_m.Called(listener)
}
