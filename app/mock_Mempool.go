// Code generated by mockery v1.0.0. DO NOT EDIT.

package app

import common "github.com/lianxiangcloud/linkchain/libs/common"
import mock "github.com/stretchr/testify/mock"
import types "github.com/lianxiangcloud/linkchain/types"
import lktypes "github.com/lianxiangcloud/linkchain/libs/cryptonote/types"

// MockMempool is an autogenerated mock type for the MockMempool type
type MockMempool struct {
	mock.Mock
}

// KeyImageExists provides a mock function with given fields: key
func (_m *MockMempool) KeyImageExists(key lktypes.Key) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(lktypes.Key) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// KeyImagePush provides a mock function with given fields: key
func (_m *MockMempool) KeyImagePush(key lktypes.Key) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(lktypes.Key) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// KeyImageReset provides a mock function with given fields:
func (_m *MockMempool) KeyImageReset() {
	_m.Called()
}

// Lock provides a mock function with given fields:
func (_m *MockMempool) Lock() {
	_m.Called()
}

// Reap provides a mock function with given fields: maxTxs
func (_m *MockMempool) Reap(maxTxs int) types.Txs {
	ret := _m.Called(maxTxs)

	var r0 types.Txs
	if rf, ok := ret.Get(0).(func(int) types.Txs); ok {
		r0 = rf(maxTxs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Txs)
		}
	}

	return r0
}

// Unlock provides a mock function with given fields:
func (_m *MockMempool) Unlock() {
	_m.Called()
}

// Update provides a mock function with given fields: height, txs, keyImages
func (_m *MockMempool) Update(height uint64, txs types.Txs, keyImages []*lktypes.Key) error {
	ret := _m.Called(height, txs, keyImages)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, types.Txs, []*lktypes.Key) error); ok {
		r0 = rf(height, txs, keyImages)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyTxFromCache provides a mock function with given fields: hash
func (_m *MockMempool) VerifyTxFromCache(hash common.Hash) (*common.Address, bool) {
	ret := _m.Called(hash)

	var r0 *common.Address
	if rf, ok := ret.Get(0).(func(common.Hash) *common.Address); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Address)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(common.Hash) bool); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}
