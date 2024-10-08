package ocr3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/types"
	"github.com/goplugin/plugin-common/pkg/types/core"
	"github.com/goplugin/plugin-common/pkg/types/core/mocks"
	"github.com/goplugin/plugin-common/pkg/utils/tests"
)

func TestOCR3_ReportingFactoryAddsCapability(t *testing.T) {
	ctx := tests.Context(t)

	cfg := Config{
		EncoderFactory: mockEncoderFactory,
		Logger:         logger.Test(t),
	}
	o := NewOCR3(cfg)
	require.NoError(t, o.Start(ctx))

	var p types.PluginProvider
	var pr core.PipelineRunnerService
	var tc core.TelemetryClient
	var el core.ErrorLog
	var kv core.KeyValueStore
	var rs core.RelayerSet
	r := mocks.NewCapabilitiesRegistry(t)
	r.On("Add", mock.Anything, o.config.capability).Return(nil)

	_, err := o.NewReportingPluginFactory(ctx, core.ReportingPluginServiceConfig{}, p, pr, tc, el, r, kv, rs)
	require.NoError(t, err)
}

func TestOCR3_ReportingFactoryIsAService(t *testing.T) {
	ctx := tests.Context(t)

	cfg := Config{
		EncoderFactory: mockEncoderFactory,
		Logger:         logger.Test(t),
	}
	o := NewOCR3(cfg)
	require.NoError(t, o.Start(ctx))

	var p types.PluginProvider
	var pr core.PipelineRunnerService
	var tc core.TelemetryClient
	var el core.ErrorLog
	var kv core.KeyValueStore
	var rs core.RelayerSet
	r := mocks.NewCapabilitiesRegistry(t)
	r.On("Add", mock.Anything, o.config.capability).Return(nil)

	factory, err := o.NewReportingPluginFactory(ctx, core.ReportingPluginServiceConfig{}, p, pr, tc, el, r, kv, rs)
	require.NoError(t, err)

	require.NoError(t, factory.Start(ctx))

	assert.Nil(t, factory.Ready())
}
