// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	entities "go-wishlist-api/entities"

	mock "github.com/stretchr/testify/mock"
)

// WishlistUsecase is an autogenerated mock type for the WishlistUsecase type
type WishlistUsecase struct {
	mock.Mock
}

// CreateWishlist provides a mock function with given fields: wishlist
func (_m *WishlistUsecase) CreateWishlist(wishlist *entities.Wishlist) error {
	ret := _m.Called(wishlist)

	if len(ret) == 0 {
		panic("no return value specified for CreateWishlist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Wishlist) error); ok {
		r0 = rf(wishlist)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWishlist provides a mock function with given fields: id
func (_m *WishlistUsecase) DeleteWishlist(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWishlist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetWishlistByID provides a mock function with given fields: id
func (_m *WishlistUsecase) GetWishlistByID(id string) (*entities.Wishlist, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetWishlistByID")
	}

	var r0 *entities.Wishlist
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Wishlist, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Wishlist); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Wishlist)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWishlist provides a mock function with given fields: wishlist
func (_m *WishlistUsecase) UpdateWishlist(wishlist *entities.Wishlist) error {
	ret := _m.Called(wishlist)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWishlist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Wishlist) error); ok {
		r0 = rf(wishlist)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWishlistUsecase creates a new instance of WishlistUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWishlistUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *WishlistUsecase {
	mock := &WishlistUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
