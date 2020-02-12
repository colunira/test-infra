// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: ctx, ciphertext
func (_m *API) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	ret := _m.Called(ctx, ciphertext)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte) []byte); ok {
		r0 = rf(ctx, ciphertext)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, ciphertext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: ctx, plaintext
func (_m *API) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	ret := _m.Called(ctx, plaintext)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte) []byte); ok {
		r0 = rf(ctx, plaintext)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, plaintext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}