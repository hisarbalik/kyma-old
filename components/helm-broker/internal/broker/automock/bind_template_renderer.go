// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import services "k8s.io/helm/pkg/proto/hapi/services"
import bind "github.com/kyma-project/kyma/components/helm-broker/internal/bind"
import internal "github.com/kyma-project/kyma/components/helm-broker/internal"

// BindTemplateRenderer is an autogenerated mock type for the BindTemplateRenderer type
type BindTemplateRenderer struct {
	mock.Mock
}

// Render provides a mock function with given fields: bindTemplate, resp
func (_m *BindTemplateRenderer) Render(bindTemplate internal.AddonPlanBindTemplate, resp *services.InstallReleaseResponse) (bind.RenderedBindYAML, error) {
	ret := _m.Called(bindTemplate, resp)

	var r0 bind.RenderedBindYAML
	if rf, ok := ret.Get(0).(func(internal.AddonPlanBindTemplate, *services.InstallReleaseResponse) bind.RenderedBindYAML); ok {
		r0 = rf(bindTemplate, resp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bind.RenderedBindYAML)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.AddonPlanBindTemplate, *services.InstallReleaseResponse) error); ok {
		r1 = rf(bindTemplate, resp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
