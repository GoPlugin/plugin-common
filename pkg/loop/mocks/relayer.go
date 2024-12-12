// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	types "github.com/goplugin/plugin-common/pkg/types"
)

// Relayer is an autogenerated mock type for the Relayer type
type Relayer struct {
	mock.Mock
}

type Relayer_Expecter struct {
	mock *mock.Mock
}

func (_m *Relayer) EXPECT() *Relayer_Expecter {
	return &Relayer_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Relayer) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Relayer_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Relayer_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Relayer_Expecter) Close() *Relayer_Close_Call {
	return &Relayer_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Relayer_Close_Call) Run(run func()) *Relayer_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Relayer_Close_Call) Return(_a0 error) *Relayer_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Relayer_Close_Call) RunAndReturn(run func() error) *Relayer_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetChainStatus provides a mock function with given fields: ctx
func (_m *Relayer) GetChainStatus(ctx context.Context) (types.ChainStatus, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetChainStatus")
	}

	var r0 types.ChainStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.ChainStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.ChainStatus); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.ChainStatus)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relayer_GetChainStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChainStatus'
type Relayer_GetChainStatus_Call struct {
	*mock.Call
}

// GetChainStatus is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Relayer_Expecter) GetChainStatus(ctx interface{}) *Relayer_GetChainStatus_Call {
	return &Relayer_GetChainStatus_Call{Call: _e.mock.On("GetChainStatus", ctx)}
}

func (_c *Relayer_GetChainStatus_Call) Run(run func(ctx context.Context)) *Relayer_GetChainStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Relayer_GetChainStatus_Call) Return(_a0 types.ChainStatus, _a1 error) *Relayer_GetChainStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Relayer_GetChainStatus_Call) RunAndReturn(run func(context.Context) (types.ChainStatus, error)) *Relayer_GetChainStatus_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *Relayer) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

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

// Relayer_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type Relayer_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *Relayer_Expecter) HealthReport() *Relayer_HealthReport_Call {
	return &Relayer_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *Relayer_HealthReport_Call) Run(run func()) *Relayer_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Relayer_HealthReport_Call) Return(_a0 map[string]error) *Relayer_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Relayer_HealthReport_Call) RunAndReturn(run func() map[string]error) *Relayer_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// LatestHead provides a mock function with given fields: ctx
func (_m *Relayer) LatestHead(ctx context.Context) (types.Head, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestHead")
	}

	var r0 types.Head
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.Head, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.Head); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.Head)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relayer_LatestHead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LatestHead'
type Relayer_LatestHead_Call struct {
	*mock.Call
}

// LatestHead is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Relayer_Expecter) LatestHead(ctx interface{}) *Relayer_LatestHead_Call {
	return &Relayer_LatestHead_Call{Call: _e.mock.On("LatestHead", ctx)}
}

func (_c *Relayer_LatestHead_Call) Run(run func(ctx context.Context)) *Relayer_LatestHead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Relayer_LatestHead_Call) Return(_a0 types.Head, _a1 error) *Relayer_LatestHead_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Relayer_LatestHead_Call) RunAndReturn(run func(context.Context) (types.Head, error)) *Relayer_LatestHead_Call {
	_c.Call.Return(run)
	return _c
}

// ListNodeStatuses provides a mock function with given fields: ctx, pageSize, pageToken
func (_m *Relayer) ListNodeStatuses(ctx context.Context, pageSize int32, pageToken string) ([]types.NodeStatus, string, int, error) {
	ret := _m.Called(ctx, pageSize, pageToken)

	if len(ret) == 0 {
		panic("no return value specified for ListNodeStatuses")
	}

	var r0 []types.NodeStatus
	var r1 string
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) ([]types.NodeStatus, string, int, error)); ok {
		return rf(ctx, pageSize, pageToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) []types.NodeStatus); ok {
		r0 = rf(ctx, pageSize, pageToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NodeStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, string) string); ok {
		r1 = rf(ctx, pageSize, pageToken)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int32, string) int); ok {
		r2 = rf(ctx, pageSize, pageToken)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(context.Context, int32, string) error); ok {
		r3 = rf(ctx, pageSize, pageToken)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Relayer_ListNodeStatuses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNodeStatuses'
type Relayer_ListNodeStatuses_Call struct {
	*mock.Call
}

// ListNodeStatuses is a helper method to define mock.On call
//   - ctx context.Context
//   - pageSize int32
//   - pageToken string
func (_e *Relayer_Expecter) ListNodeStatuses(ctx interface{}, pageSize interface{}, pageToken interface{}) *Relayer_ListNodeStatuses_Call {
	return &Relayer_ListNodeStatuses_Call{Call: _e.mock.On("ListNodeStatuses", ctx, pageSize, pageToken)}
}

func (_c *Relayer_ListNodeStatuses_Call) Run(run func(ctx context.Context, pageSize int32, pageToken string)) *Relayer_ListNodeStatuses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(string))
	})
	return _c
}

func (_c *Relayer_ListNodeStatuses_Call) Return(stats []types.NodeStatus, nextPageToken string, total int, err error) *Relayer_ListNodeStatuses_Call {
	_c.Call.Return(stats, nextPageToken, total, err)
	return _c
}

func (_c *Relayer_ListNodeStatuses_Call) RunAndReturn(run func(context.Context, int32, string) ([]types.NodeStatus, string, int, error)) *Relayer_ListNodeStatuses_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Relayer) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Relayer_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Relayer_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Relayer_Expecter) Name() *Relayer_Name_Call {
	return &Relayer_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Relayer_Name_Call) Run(run func()) *Relayer_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Relayer_Name_Call) Return(_a0 string) *Relayer_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Relayer_Name_Call) RunAndReturn(run func() string) *Relayer_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewChainWriter provides a mock function with given fields: ctx, chainWriterConfig
func (_m *Relayer) NewChainWriter(ctx context.Context, chainWriterConfig []byte) (types.ChainWriter, error) {
	ret := _m.Called(ctx, chainWriterConfig)

	if len(ret) == 0 {
		panic("no return value specified for NewChainWriter")
	}

	var r0 types.ChainWriter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) (types.ChainWriter, error)); ok {
		return rf(ctx, chainWriterConfig)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) types.ChainWriter); ok {
		r0 = rf(ctx, chainWriterConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ChainWriter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, chainWriterConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relayer_NewChainWriter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewChainWriter'
type Relayer_NewChainWriter_Call struct {
	*mock.Call
}

// NewChainWriter is a helper method to define mock.On call
//   - ctx context.Context
//   - chainWriterConfig []byte
func (_e *Relayer_Expecter) NewChainWriter(ctx interface{}, chainWriterConfig interface{}) *Relayer_NewChainWriter_Call {
	return &Relayer_NewChainWriter_Call{Call: _e.mock.On("NewChainWriter", ctx, chainWriterConfig)}
}

func (_c *Relayer_NewChainWriter_Call) Run(run func(ctx context.Context, chainWriterConfig []byte)) *Relayer_NewChainWriter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte))
	})
	return _c
}

func (_c *Relayer_NewChainWriter_Call) Return(_a0 types.ChainWriter, _a1 error) *Relayer_NewChainWriter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Relayer_NewChainWriter_Call) RunAndReturn(run func(context.Context, []byte) (types.ChainWriter, error)) *Relayer_NewChainWriter_Call {
	_c.Call.Return(run)
	return _c
}

// NewConfigProvider provides a mock function with given fields: _a0, _a1
func (_m *Relayer) NewConfigProvider(_a0 context.Context, _a1 types.RelayArgs) (types.ConfigProvider, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for NewConfigProvider")
	}

	var r0 types.ConfigProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs) (types.ConfigProvider, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs) types.ConfigProvider); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ConfigProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.RelayArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relayer_NewConfigProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewConfigProvider'
type Relayer_NewConfigProvider_Call struct {
	*mock.Call
}

// NewConfigProvider is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 types.RelayArgs
func (_e *Relayer_Expecter) NewConfigProvider(_a0 interface{}, _a1 interface{}) *Relayer_NewConfigProvider_Call {
	return &Relayer_NewConfigProvider_Call{Call: _e.mock.On("NewConfigProvider", _a0, _a1)}
}

func (_c *Relayer_NewConfigProvider_Call) Run(run func(_a0 context.Context, _a1 types.RelayArgs)) *Relayer_NewConfigProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.RelayArgs))
	})
	return _c
}

func (_c *Relayer_NewConfigProvider_Call) Return(_a0 types.ConfigProvider, _a1 error) *Relayer_NewConfigProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Relayer_NewConfigProvider_Call) RunAndReturn(run func(context.Context, types.RelayArgs) (types.ConfigProvider, error)) *Relayer_NewConfigProvider_Call {
	_c.Call.Return(run)
	return _c
}

// NewContractReader provides a mock function with given fields: ctx, contractReaderConfig
func (_m *Relayer) NewContractReader(ctx context.Context, contractReaderConfig []byte) (types.ContractReader, error) {
	ret := _m.Called(ctx, contractReaderConfig)

	if len(ret) == 0 {
		panic("no return value specified for NewContractReader")
	}

	var r0 types.ContractReader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) (types.ContractReader, error)); ok {
		return rf(ctx, contractReaderConfig)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) types.ContractReader); ok {
		r0 = rf(ctx, contractReaderConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ContractReader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, contractReaderConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relayer_NewContractReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewContractReader'
type Relayer_NewContractReader_Call struct {
	*mock.Call
}

// NewContractReader is a helper method to define mock.On call
//   - ctx context.Context
//   - contractReaderConfig []byte
func (_e *Relayer_Expecter) NewContractReader(ctx interface{}, contractReaderConfig interface{}) *Relayer_NewContractReader_Call {
	return &Relayer_NewContractReader_Call{Call: _e.mock.On("NewContractReader", ctx, contractReaderConfig)}
}

func (_c *Relayer_NewContractReader_Call) Run(run func(ctx context.Context, contractReaderConfig []byte)) *Relayer_NewContractReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte))
	})
	return _c
}

func (_c *Relayer_NewContractReader_Call) Return(_a0 types.ContractReader, _a1 error) *Relayer_NewContractReader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Relayer_NewContractReader_Call) RunAndReturn(run func(context.Context, []byte) (types.ContractReader, error)) *Relayer_NewContractReader_Call {
	_c.Call.Return(run)
	return _c
}

// NewLLOProvider provides a mock function with given fields: _a0, _a1, _a2
func (_m *Relayer) NewLLOProvider(_a0 context.Context, _a1 types.RelayArgs, _a2 types.PluginArgs) (types.LLOProvider, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for NewLLOProvider")
	}

	var r0 types.LLOProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs, types.PluginArgs) (types.LLOProvider, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs, types.PluginArgs) types.LLOProvider); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.LLOProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.RelayArgs, types.PluginArgs) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relayer_NewLLOProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewLLOProvider'
type Relayer_NewLLOProvider_Call struct {
	*mock.Call
}

// NewLLOProvider is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 types.RelayArgs
//   - _a2 types.PluginArgs
func (_e *Relayer_Expecter) NewLLOProvider(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Relayer_NewLLOProvider_Call {
	return &Relayer_NewLLOProvider_Call{Call: _e.mock.On("NewLLOProvider", _a0, _a1, _a2)}
}

func (_c *Relayer_NewLLOProvider_Call) Run(run func(_a0 context.Context, _a1 types.RelayArgs, _a2 types.PluginArgs)) *Relayer_NewLLOProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.RelayArgs), args[2].(types.PluginArgs))
	})
	return _c
}

func (_c *Relayer_NewLLOProvider_Call) Return(_a0 types.LLOProvider, _a1 error) *Relayer_NewLLOProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Relayer_NewLLOProvider_Call) RunAndReturn(run func(context.Context, types.RelayArgs, types.PluginArgs) (types.LLOProvider, error)) *Relayer_NewLLOProvider_Call {
	_c.Call.Return(run)
	return _c
}

// NewPluginProvider provides a mock function with given fields: _a0, _a1, _a2
func (_m *Relayer) NewPluginProvider(_a0 context.Context, _a1 types.RelayArgs, _a2 types.PluginArgs) (types.PluginProvider, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for NewPluginProvider")
	}

	var r0 types.PluginProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs, types.PluginArgs) (types.PluginProvider, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.RelayArgs, types.PluginArgs) types.PluginProvider); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.PluginProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.RelayArgs, types.PluginArgs) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relayer_NewPluginProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewPluginProvider'
type Relayer_NewPluginProvider_Call struct {
	*mock.Call
}

// NewPluginProvider is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 types.RelayArgs
//   - _a2 types.PluginArgs
func (_e *Relayer_Expecter) NewPluginProvider(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Relayer_NewPluginProvider_Call {
	return &Relayer_NewPluginProvider_Call{Call: _e.mock.On("NewPluginProvider", _a0, _a1, _a2)}
}

func (_c *Relayer_NewPluginProvider_Call) Run(run func(_a0 context.Context, _a1 types.RelayArgs, _a2 types.PluginArgs)) *Relayer_NewPluginProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.RelayArgs), args[2].(types.PluginArgs))
	})
	return _c
}

func (_c *Relayer_NewPluginProvider_Call) Return(_a0 types.PluginProvider, _a1 error) *Relayer_NewPluginProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Relayer_NewPluginProvider_Call) RunAndReturn(run func(context.Context, types.RelayArgs, types.PluginArgs) (types.PluginProvider, error)) *Relayer_NewPluginProvider_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *Relayer) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Relayer_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type Relayer_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *Relayer_Expecter) Ready() *Relayer_Ready_Call {
	return &Relayer_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *Relayer_Ready_Call) Run(run func()) *Relayer_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Relayer_Ready_Call) Return(_a0 error) *Relayer_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Relayer_Ready_Call) RunAndReturn(run func() error) *Relayer_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *Relayer) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Relayer_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Relayer_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Relayer_Expecter) Start(_a0 interface{}) *Relayer_Start_Call {
	return &Relayer_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *Relayer_Start_Call) Run(run func(_a0 context.Context)) *Relayer_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Relayer_Start_Call) Return(_a0 error) *Relayer_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Relayer_Start_Call) RunAndReturn(run func(context.Context) error) *Relayer_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Transact provides a mock function with given fields: ctx, from, to, amount, balanceCheck
func (_m *Relayer) Transact(ctx context.Context, from string, to string, amount *big.Int, balanceCheck bool) error {
	ret := _m.Called(ctx, from, to, amount, balanceCheck)

	if len(ret) == 0 {
		panic("no return value specified for Transact")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *big.Int, bool) error); ok {
		r0 = rf(ctx, from, to, amount, balanceCheck)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Relayer_Transact_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transact'
type Relayer_Transact_Call struct {
	*mock.Call
}

// Transact is a helper method to define mock.On call
//   - ctx context.Context
//   - from string
//   - to string
//   - amount *big.Int
//   - balanceCheck bool
func (_e *Relayer_Expecter) Transact(ctx interface{}, from interface{}, to interface{}, amount interface{}, balanceCheck interface{}) *Relayer_Transact_Call {
	return &Relayer_Transact_Call{Call: _e.mock.On("Transact", ctx, from, to, amount, balanceCheck)}
}

func (_c *Relayer_Transact_Call) Run(run func(ctx context.Context, from string, to string, amount *big.Int, balanceCheck bool)) *Relayer_Transact_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*big.Int), args[4].(bool))
	})
	return _c
}

func (_c *Relayer_Transact_Call) Return(_a0 error) *Relayer_Transact_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Relayer_Transact_Call) RunAndReturn(run func(context.Context, string, string, *big.Int, bool) error) *Relayer_Transact_Call {
	_c.Call.Return(run)
	return _c
}

// NewRelayer creates a new instance of Relayer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRelayer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Relayer {
	mock := &Relayer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}