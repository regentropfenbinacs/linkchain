// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lianxiangcloud/linkchain/wallet/rpc (interfaces: Wallet)

// Package rpc is a generated GoMock package.
package rpc

import (
	gomock "github.com/golang/mock/gomock"
	common "github.com/lianxiangcloud/linkchain/libs/common"
	types "github.com/lianxiangcloud/linkchain/libs/cryptonote/types"
	hexutil "github.com/lianxiangcloud/linkchain/libs/hexutil"
	rpc "github.com/lianxiangcloud/linkchain/libs/rpc"
	types0 "github.com/lianxiangcloud/linkchain/types"
	types1 "github.com/lianxiangcloud/linkchain/wallet/types"
	big "math/big"
	reflect "reflect"
	time "time"
)

// MockWallet is a mock of Wallet interface
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// AutoRefreshBlockchain mocks base method
func (m *MockWallet) AutoRefreshBlockchain(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoRefreshBlockchain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoRefreshBlockchain indicates an expected call of AutoRefreshBlockchain
func (mr *MockWalletMockRecorder) AutoRefreshBlockchain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoRefreshBlockchain", reflect.TypeOf((*MockWallet)(nil).AutoRefreshBlockchain), arg0)
}

// CreateSubAccount mocks base method
func (m *MockWallet) CreateSubAccount(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubAccount indicates an expected call of CreateSubAccount
func (mr *MockWalletMockRecorder) CreateSubAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubAccount", reflect.TypeOf((*MockWallet)(nil).CreateSubAccount), arg0)
}

// CreateUTXOTransaction mocks base method
func (m *MockWallet) CreateUTXOTransaction(arg0 common.Address, arg1 uint64, arg2 []uint64, arg3 []types0.DestEntry, arg4, arg5 common.Address, arg6 []byte) ([]*types0.UTXOTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUTXOTransaction", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]*types0.UTXOTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUTXOTransaction indicates an expected call of CreateUTXOTransaction
func (mr *MockWalletMockRecorder) CreateUTXOTransaction(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUTXOTransaction", reflect.TypeOf((*MockWallet)(nil).CreateUTXOTransaction), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// EthEstimateGas mocks base method
func (m *MockWallet) EthEstimateGas(arg0 types1.CallArgs) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EthEstimateGas", arg0)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EthEstimateGas indicates an expected call of EthEstimateGas
func (mr *MockWalletMockRecorder) EthEstimateGas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EthEstimateGas", reflect.TypeOf((*MockWallet)(nil).EthEstimateGas), arg0)
}

// GetAccountInfo mocks base method
func (m *MockWallet) GetAccountInfo(arg0 *common.Address) (*types1.GetAccountInfoResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountInfo", arg0)
	ret0, _ := ret[0].(*types1.GetAccountInfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountInfo indicates an expected call of GetAccountInfo
func (mr *MockWalletMockRecorder) GetAccountInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountInfo", reflect.TypeOf((*MockWallet)(nil).GetAccountInfo), arg0)
}

// GetAddress mocks base method
func (m *MockWallet) GetAddress(arg0 uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddress indicates an expected call of GetAddress
func (mr *MockWalletMockRecorder) GetAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockWallet)(nil).GetAddress), arg0)
}

// GetBalance mocks base method
func (m *MockWallet) GetBalance(arg0 uint64, arg1 *common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0, arg1)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance
func (mr *MockWalletMockRecorder) GetBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockWallet)(nil).GetBalance), arg0, arg1)
}

// GetBlockTransactionCountByHash mocks base method
func (m *MockWallet) GetBlockTransactionCountByHash(arg0 common.Hash) (*hexutil.Uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockTransactionCountByHash", arg0)
	ret0, _ := ret[0].(*hexutil.Uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockTransactionCountByHash indicates an expected call of GetBlockTransactionCountByHash
func (mr *MockWalletMockRecorder) GetBlockTransactionCountByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockTransactionCountByHash", reflect.TypeOf((*MockWallet)(nil).GetBlockTransactionCountByHash), arg0)
}

// GetBlockTransactionCountByNumber mocks base method
func (m *MockWallet) GetBlockTransactionCountByNumber(arg0 rpc.BlockNumber) (*hexutil.Uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockTransactionCountByNumber", arg0)
	ret0, _ := ret[0].(*hexutil.Uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockTransactionCountByNumber indicates an expected call of GetBlockTransactionCountByNumber
func (mr *MockWalletMockRecorder) GetBlockTransactionCountByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockTransactionCountByNumber", reflect.TypeOf((*MockWallet)(nil).GetBlockTransactionCountByNumber), arg0)
}

// GetHeight mocks base method
func (m *MockWallet) GetHeight() (uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeight")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// GetHeight indicates an expected call of GetHeight
func (mr *MockWalletMockRecorder) GetHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeight", reflect.TypeOf((*MockWallet)(nil).GetHeight))
}

// GetMaxOutput mocks base method
func (m *MockWallet) GetMaxOutput(arg0 common.Address) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxOutput", arg0)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxOutput indicates an expected call of GetMaxOutput
func (mr *MockWalletMockRecorder) GetMaxOutput(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxOutput", reflect.TypeOf((*MockWallet)(nil).GetMaxOutput), arg0)
}

// GetRawTransactionByBlockHashAndIndex mocks base method
func (m *MockWallet) GetRawTransactionByBlockHashAndIndex(arg0 common.Hash, arg1 hexutil.Uint) (hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawTransactionByBlockHashAndIndex", arg0, arg1)
	ret0, _ := ret[0].(hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawTransactionByBlockHashAndIndex indicates an expected call of GetRawTransactionByBlockHashAndIndex
func (mr *MockWalletMockRecorder) GetRawTransactionByBlockHashAndIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawTransactionByBlockHashAndIndex", reflect.TypeOf((*MockWallet)(nil).GetRawTransactionByBlockHashAndIndex), arg0, arg1)
}

// GetRawTransactionByBlockNumberAndIndex mocks base method
func (m *MockWallet) GetRawTransactionByBlockNumberAndIndex(arg0 rpc.BlockNumber, arg1 hexutil.Uint) (hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawTransactionByBlockNumberAndIndex", arg0, arg1)
	ret0, _ := ret[0].(hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawTransactionByBlockNumberAndIndex indicates an expected call of GetRawTransactionByBlockNumberAndIndex
func (mr *MockWalletMockRecorder) GetRawTransactionByBlockNumberAndIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawTransactionByBlockNumberAndIndex", reflect.TypeOf((*MockWallet)(nil).GetRawTransactionByBlockNumberAndIndex), arg0, arg1)
}

// GetRawTransactionByHash mocks base method
func (m *MockWallet) GetRawTransactionByHash(arg0 common.Hash) (hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawTransactionByHash", arg0)
	ret0, _ := ret[0].(hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawTransactionByHash indicates an expected call of GetRawTransactionByHash
func (mr *MockWalletMockRecorder) GetRawTransactionByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawTransactionByHash", reflect.TypeOf((*MockWallet)(nil).GetRawTransactionByHash), arg0)
}

// GetTransactionByBlockHashAndIndex mocks base method
func (m *MockWallet) GetTransactionByBlockHashAndIndex(arg0 common.Hash, arg1 hexutil.Uint) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByBlockHashAndIndex", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByBlockHashAndIndex indicates an expected call of GetTransactionByBlockHashAndIndex
func (mr *MockWalletMockRecorder) GetTransactionByBlockHashAndIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByBlockHashAndIndex", reflect.TypeOf((*MockWallet)(nil).GetTransactionByBlockHashAndIndex), arg0, arg1)
}

// GetTransactionByBlockNumberAndIndex mocks base method
func (m *MockWallet) GetTransactionByBlockNumberAndIndex(arg0 rpc.BlockNumber, arg1 hexutil.Uint) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByBlockNumberAndIndex", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByBlockNumberAndIndex indicates an expected call of GetTransactionByBlockNumberAndIndex
func (mr *MockWalletMockRecorder) GetTransactionByBlockNumberAndIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByBlockNumberAndIndex", reflect.TypeOf((*MockWallet)(nil).GetTransactionByBlockNumberAndIndex), arg0, arg1)
}

// GetTransactionByHash mocks base method
func (m *MockWallet) GetTransactionByHash(arg0 common.Hash) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash
func (mr *MockWalletMockRecorder) GetTransactionByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockWallet)(nil).GetTransactionByHash), arg0)
}

// GetTransactionCount mocks base method
func (m *MockWallet) GetTransactionCount(arg0 common.Address, arg1 rpc.BlockNumber) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionCount", arg0, arg1)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionCount indicates an expected call of GetTransactionCount
func (mr *MockWalletMockRecorder) GetTransactionCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCount", reflect.TypeOf((*MockWallet)(nil).GetTransactionCount), arg0, arg1)
}

// GetTransactionReceipt mocks base method
func (m *MockWallet) GetTransactionReceipt(arg0 common.Hash) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionReceipt", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionReceipt indicates an expected call of GetTransactionReceipt
func (mr *MockWalletMockRecorder) GetTransactionReceipt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceipt", reflect.TypeOf((*MockWallet)(nil).GetTransactionReceipt), arg0)
}

// GetTxKey mocks base method
func (m *MockWallet) GetTxKey(arg0 *common.Hash) (*types.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxKey", arg0)
	ret0, _ := ret[0].(*types.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxKey indicates an expected call of GetTxKey
func (mr *MockWalletMockRecorder) GetTxKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxKey", reflect.TypeOf((*MockWallet)(nil).GetTxKey), arg0)
}

// GetUTXOTx mocks base method
func (m *MockWallet) GetUTXOTx(arg0 common.Hash) (*types0.UTXOTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTXOTx", arg0)
	ret0, _ := ret[0].(*types0.UTXOTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUTXOTx indicates an expected call of GetUTXOTx
func (mr *MockWalletMockRecorder) GetUTXOTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTXOTx", reflect.TypeOf((*MockWallet)(nil).GetUTXOTx), arg0)
}

// GetWalletEthAddress mocks base method
func (m *MockWallet) GetWalletEthAddress() (*common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletEthAddress")
	ret0, _ := ret[0].(*common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletEthAddress indicates an expected call of GetWalletEthAddress
func (mr *MockWalletMockRecorder) GetWalletEthAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletEthAddress", reflect.TypeOf((*MockWallet)(nil).GetWalletEthAddress))
}

// LockAccount mocks base method
func (m *MockWallet) LockAccount(arg0 common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockAccount indicates an expected call of LockAccount
func (mr *MockWalletMockRecorder) LockAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockAccount", reflect.TypeOf((*MockWallet)(nil).LockAccount), arg0)
}

// OpenWallet mocks base method
func (m *MockWallet) OpenWallet(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenWallet indicates an expected call of OpenWallet
func (mr *MockWalletMockRecorder) OpenWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenWallet", reflect.TypeOf((*MockWallet)(nil).OpenWallet), arg0, arg1)
}

// RescanBlockchain mocks base method
func (m *MockWallet) RescanBlockchain() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RescanBlockchain")
	ret0, _ := ret[0].(error)
	return ret0
}

// RescanBlockchain indicates an expected call of RescanBlockchain
func (mr *MockWalletMockRecorder) RescanBlockchain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RescanBlockchain", reflect.TypeOf((*MockWallet)(nil).RescanBlockchain))
}

// SelectAddress mocks base method
func (m *MockWallet) SelectAddress(arg0 common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAddress", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectAddress indicates an expected call of SelectAddress
func (mr *MockWalletMockRecorder) SelectAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAddress", reflect.TypeOf((*MockWallet)(nil).SelectAddress), arg0)
}

// SendRawTransaction mocks base method
func (m *MockWallet) SendRawTransaction(arg0 hexutil.Bytes) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTransaction", arg0)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction
func (mr *MockWalletMockRecorder) SendRawTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockWallet)(nil).SendRawTransaction), arg0)
}

// SetRefreshBlockInterval mocks base method
func (m *MockWallet) SetRefreshBlockInterval(arg0 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRefreshBlockInterval", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRefreshBlockInterval indicates an expected call of SetRefreshBlockInterval
func (mr *MockWalletMockRecorder) SetRefreshBlockInterval(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRefreshBlockInterval", reflect.TypeOf((*MockWallet)(nil).SetRefreshBlockInterval), arg0)
}

// Status mocks base method
func (m *MockWallet) Status() *types1.StatusResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(*types1.StatusResult)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockWalletMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockWallet)(nil).Status))
}

// Transfer mocks base method
func (m *MockWallet) Transfer(arg0 []string) []types1.SendTxRet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", arg0)
	ret0, _ := ret[0].([]types1.SendTxRet)
	return ret0
}

// Transfer indicates an expected call of Transfer
func (mr *MockWalletMockRecorder) Transfer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockWallet)(nil).Transfer), arg0)
}
