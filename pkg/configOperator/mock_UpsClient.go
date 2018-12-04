// Code generated by mockery v1.0.0. DO NOT EDIT.

package configOperator

import mock "github.com/stretchr/testify/mock"

// MockUpsClient is an autogenerated mock type for the UpsClient type
type MockUpsClient struct {
	mock.Mock
}

// createAndroidVariant provides a mock function with given fields: variant
func (_m *MockUpsClient) createAndroidVariant(variant *AndroidVariant) (bool, *AndroidVariant) {
	ret := _m.Called(variant)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*AndroidVariant) bool); ok {
		r0 = rf(variant)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *AndroidVariant
	if rf, ok := ret.Get(1).(func(*AndroidVariant) *AndroidVariant); ok {
		r1 = rf(variant)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*AndroidVariant)
		}
	}

	return r0, r1
}

// createIOSVariant provides a mock function with given fields: variant
func (_m *MockUpsClient) createIOSVariant(variant *IOSVariant) (bool, *IOSVariant) {
	ret := _m.Called(variant)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*IOSVariant) bool); ok {
		r0 = rf(variant)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *IOSVariant
	if rf, ok := ret.Get(1).(func(*IOSVariant) *IOSVariant); ok {
		r1 = rf(variant)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*IOSVariant)
		}
	}

	return r0, r1
}

// deleteVariant provides a mock function with given fields: platform, variantId
func (_m *MockUpsClient) deleteVariant(platform string, variantId string) bool {
	ret := _m.Called(platform, variantId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(platform, variantId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// getApplicationId provides a mock function with given fields:
func (_m *MockUpsClient) getApplicationId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getBaseUrl provides a mock function with given fields:
func (_m *MockUpsClient) getBaseUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getPushApplicationName provides a mock function with given fields:
func (_m *MockUpsClient) getPushApplicationName() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getServiceInstanceId provides a mock function with given fields:
func (_m *MockUpsClient) getServiceInstanceId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getVariants provides a mock function with given fields:
func (_m *MockUpsClient) getVariants() ([]Variant, error) {
	ret := _m.Called()

	var r0 []Variant
	if rf, ok := ret.Get(0).(func() []Variant); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Variant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// hasAndroidVariant provides a mock function with given fields: key
func (_m *MockUpsClient) hasAndroidVariant(key string) *AndroidVariant {
	ret := _m.Called(key)

	var r0 *AndroidVariant
	if rf, ok := ret.Get(0).(func(string) *AndroidVariant); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AndroidVariant)
		}
	}

	return r0
}
