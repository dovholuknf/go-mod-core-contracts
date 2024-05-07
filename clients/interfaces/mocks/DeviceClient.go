// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
)

// DeviceClient is an autogenerated mock type for the DeviceClient type
type DeviceClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, reqs
func (_m *DeviceClient) Add(ctx context.Context, reqs []requests.AddDeviceRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 []common.BaseWithIdResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddDeviceRequest) ([]common.BaseWithIdResponse, errors.EdgeX)); ok {
		return rf(ctx, reqs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddDeviceRequest) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []requests.AddDeviceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddWithQueryParams provides a mock function with given fields: ctx, reqs, queryParams
func (_m *DeviceClient) AddWithQueryParams(ctx context.Context, reqs []requests.AddDeviceRequest, queryParams map[string]string) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs, queryParams)

	if len(ret) == 0 {
		panic("no return value specified for AddWithQueryParams")
	}

	var r0 []common.BaseWithIdResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddDeviceRequest, map[string]string) ([]common.BaseWithIdResponse, errors.EdgeX)); ok {
		return rf(ctx, reqs, queryParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddDeviceRequest, map[string]string) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs, queryParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []requests.AddDeviceRequest, map[string]string) errors.EdgeX); ok {
		r1 = rf(ctx, reqs, queryParams)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllDevices provides a mock function with given fields: ctx, labels, offset, limit
func (_m *DeviceClient) AllDevices(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDevicesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, labels, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for AllDevices")
	}

	var r0 responses.MultiDevicesResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) (responses.MultiDevicesResponse, errors.EdgeX)); ok {
		return rf(ctx, labels, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) responses.MultiDevicesResponse); ok {
		r0 = rf(ctx, labels, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDevicesResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, labels, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllDevicesWithChildren provides a mock function with given fields: ctx, parent, maxLevels, labels, offset, limit
func (_m *DeviceClient) AllDevicesWithChildren(ctx context.Context, parent string, maxLevels uint, labels []string, offset int, limit int) (responses.MultiDevicesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, parent, maxLevels, labels, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for AllDevicesWithChildren")
	}

	var r0 responses.MultiDevicesResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, uint, []string, int, int) (responses.MultiDevicesResponse, errors.EdgeX)); ok {
		return rf(ctx, parent, maxLevels, labels, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint, []string, int, int) responses.MultiDevicesResponse); ok {
		r0 = rf(ctx, parent, maxLevels, labels, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDevicesResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint, []string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, parent, maxLevels, labels, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteDeviceByName provides a mock function with given fields: ctx, name
func (_m *DeviceClient) DeleteDeviceByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDeviceByName")
	}

	var r0 common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceByName provides a mock function with given fields: ctx, name
func (_m *DeviceClient) DeviceByName(ctx context.Context, name string) (responses.DeviceResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeviceByName")
	}

	var r0 responses.DeviceResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (responses.DeviceResponse, errors.EdgeX)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) responses.DeviceResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(responses.DeviceResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceNameExists provides a mock function with given fields: ctx, name
func (_m *DeviceClient) DeviceNameExists(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeviceNameExists")
	}

	var r0 common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DevicesByProfileName provides a mock function with given fields: ctx, name, offset, limit
func (_m *DeviceClient) DevicesByProfileName(ctx context.Context, name string, offset int, limit int) (responses.MultiDevicesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for DevicesByProfileName")
	}

	var r0 responses.MultiDevicesResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) (responses.MultiDevicesResponse, errors.EdgeX)); ok {
		return rf(ctx, name, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiDevicesResponse); ok {
		r0 = rf(ctx, name, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDevicesResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, name, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DevicesByServiceName provides a mock function with given fields: ctx, name, offset, limit
func (_m *DeviceClient) DevicesByServiceName(ctx context.Context, name string, offset int, limit int) (responses.MultiDevicesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for DevicesByServiceName")
	}

	var r0 responses.MultiDevicesResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) (responses.MultiDevicesResponse, errors.EdgeX)); ok {
		return rf(ctx, name, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiDevicesResponse); ok {
		r0 = rf(ctx, name, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDevicesResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, name, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, reqs
func (_m *DeviceClient) Update(ctx context.Context, reqs []requests.UpdateDeviceRequest) ([]common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 []common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateDeviceRequest) ([]common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, reqs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateDeviceRequest) []common.BaseResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []requests.UpdateDeviceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateWithQueryParams provides a mock function with given fields: ctx, reqs, queryParams
func (_m *DeviceClient) UpdateWithQueryParams(ctx context.Context, reqs []requests.UpdateDeviceRequest, queryParams map[string]string) ([]common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs, queryParams)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWithQueryParams")
	}

	var r0 []common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateDeviceRequest, map[string]string) ([]common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, reqs, queryParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateDeviceRequest, map[string]string) []common.BaseResponse); ok {
		r0 = rf(ctx, reqs, queryParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []requests.UpdateDeviceRequest, map[string]string) errors.EdgeX); ok {
		r1 = rf(ctx, reqs, queryParams)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NewDeviceClient creates a new instance of DeviceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceClient {
	mock := &DeviceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
