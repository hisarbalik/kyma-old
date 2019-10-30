// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	connectorservice "github.com/kyma-project/kyma/components/connectivity-certs-controller/internal/connectorservice"
	mock "github.com/stretchr/testify/mock"
)

// EstablishedConnectionClientProvider is an autogenerated mock type for the EstablishedConnectionClientProvider type
type EstablishedConnectionClientProvider struct {
	mock.Mock
}

// CreateClient provides a mock function with given fields: credentials
func (_m *EstablishedConnectionClientProvider) CreateClient(credentials connectorservice.CertificateCredentials) connectorservice.EstablishedConnectionClient {
	ret := _m.Called(credentials)

	var r0 connectorservice.EstablishedConnectionClient
	if rf, ok := ret.Get(0).(func(connectorservice.CertificateCredentials) connectorservice.EstablishedConnectionClient); ok {
		r0 = rf(credentials)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connectorservice.EstablishedConnectionClient)
		}
	}

	return r0
}
