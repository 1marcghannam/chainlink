// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	chains "github.com/smartcontractkit/chainlink/core/chains"

	evm "github.com/smartcontractkit/chainlink/core/chains/evm"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"

	utils "github.com/smartcontractkit/chainlink/core/utils"
)

// ChainSet is an autogenerated mock type for the ChainSet type
type ChainSet struct {
	mock.Mock
}

// ChainCount provides a mock function with given fields:
func (_m *ChainSet) ChainCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Chains provides a mock function with given fields:
func (_m *ChainSet) Chains() []evm.Chain {
	ret := _m.Called()

	var r0 []evm.Chain
	if rf, ok := ret.Get(0).(func() []evm.Chain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]evm.Chain)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ChainSet) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Default provides a mock function with given fields:
func (_m *ChainSet) Default() (evm.Chain, error) {
	ret := _m.Called()

	var r0 evm.Chain
	var r1 error
	if rf, ok := ret.Get(0).(func() (evm.Chain, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() evm.Chain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evm.Chain)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *ChainSet) Get(id *big.Int) (evm.Chain, error) {
	ret := _m.Called(id)

	var r0 evm.Chain
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int) (evm.Chain, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(*big.Int) evm.Chain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evm.Chain)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodes provides a mock function with given fields: ctx, offset, limit
func (_m *ChainSet) GetNodes(ctx context.Context, offset int, limit int) ([]types.Node, int, error) {
	ret := _m.Called(ctx, offset, limit)

	var r0 []types.Node
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]types.Node, int, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []types.Node); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) int); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, int) error); ok {
		r2 = rf(ctx, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNodesByChainIDs provides a mock function with given fields: ctx, chainIDs
func (_m *ChainSet) GetNodesByChainIDs(ctx context.Context, chainIDs []utils.Big) ([]types.Node, error) {
	ret := _m.Called(ctx, chainIDs)

	var r0 []types.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []utils.Big) ([]types.Node, error)); ok {
		return rf(ctx, chainIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []utils.Big) []types.Node); ok {
		r0 = rf(ctx, chainIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []utils.Big) error); ok {
		r1 = rf(ctx, chainIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodesForChain provides a mock function with given fields: ctx, chainID, offset, limit
func (_m *ChainSet) GetNodesForChain(ctx context.Context, chainID utils.Big, offset int, limit int) ([]types.Node, int, error) {
	ret := _m.Called(ctx, chainID, offset, limit)

	var r0 []types.Node
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.Big, int, int) ([]types.Node, int, error)); ok {
		return rf(ctx, chainID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, utils.Big, int, int) []types.Node); ok {
		r0 = rf(ctx, chainID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, utils.Big, int, int) int); ok {
		r1 = rf(ctx, chainID, offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, utils.Big, int, int) error); ok {
		r2 = rf(ctx, chainID, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// HealthReport provides a mock function with given fields:
func (_m *ChainSet) HealthReport() map[string]error {
	ret := _m.Called()

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *ChainSet) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Index provides a mock function with given fields: offset, limit
func (_m *ChainSet) Index(offset int, limit int) ([]chains.DBChain[utils.Big, *types.ChainCfg], int, error) {
	ret := _m.Called(offset, limit)

	var r0 []chains.DBChain[utils.Big, *types.ChainCfg]
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]chains.DBChain[utils.Big, *types.ChainCfg], int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []chains.DBChain[utils.Big, *types.ChainCfg]); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chains.DBChain[utils.Big, *types.ChainCfg])
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Name provides a mock function with given fields:
func (_m *ChainSet) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ORM provides a mock function with given fields:
func (_m *ChainSet) ORM() types.ORM {
	ret := _m.Called()

	var r0 types.ORM
	if rf, ok := ret.Get(0).(func() types.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ORM)
		}
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *ChainSet) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Show provides a mock function with given fields: id
func (_m *ChainSet) Show(id utils.Big) (chains.DBChain[utils.Big, *types.ChainCfg], error) {
	ret := _m.Called(id)

	var r0 chains.DBChain[utils.Big, *types.ChainCfg]
	var r1 error
	if rf, ok := ret.Get(0).(func(utils.Big) (chains.DBChain[utils.Big, *types.ChainCfg], error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(utils.Big) chains.DBChain[utils.Big, *types.ChainCfg]); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(chains.DBChain[utils.Big, *types.ChainCfg])
	}

	if rf, ok := ret.Get(1).(func(utils.Big) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: _a0
func (_m *ChainSet) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewChainSet interface {
	mock.TestingT
	Cleanup(func())
}

// NewChainSet creates a new instance of ChainSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChainSet(t mockConstructorTestingTNewChainSet) *ChainSet {
	mock := &ChainSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
