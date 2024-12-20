package ocr3

import (
	"context"
	"fmt"
	"os/exec"

	"google.golang.org/grpc"

	"github.com/goplugin/plugin-libocr/offchainreporting2plus/ocr3types"

	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/loop"
	"github.com/goplugin/plugin-common/pkg/loop/internal/goplugin"
	"github.com/goplugin/plugin-common/pkg/loop/internal/net"
	"github.com/goplugin/plugin-common/pkg/loop/reportingplugins"
	"github.com/goplugin/plugin-common/pkg/types"
	"github.com/goplugin/plugin-common/pkg/types/core"
)

type LOOPPService struct {
	goplugin.PluginService[*GRPCService[types.PluginProvider], core.OCR3ReportingPluginFactory]
}

func NewLOOPPService(
	lggr logger.Logger,
	grpcOpts loop.GRPCOpts,
	cmd func() *exec.Cmd,
	config core.ReportingPluginServiceConfig,
	providerConn grpc.ClientConnInterface,
	pipelineRunner core.PipelineRunnerService,
	telemetryService core.TelemetryService,
	errorLog core.ErrorLog,
	capRegistry core.CapabilitiesRegistry,
	keyValueStore core.KeyValueStore,
	relayerSet core.RelayerSet,
) *LOOPPService {
	newService := func(ctx context.Context, instance any) (core.OCR3ReportingPluginFactory, error) {
		plug, ok := instance.(core.OCR3ReportingPluginClient)
		if !ok {
			return nil, fmt.Errorf("expected OCR3ReportingPluginClient but got %T", instance)
		}
		return plug.NewReportingPluginFactory(ctx, config, providerConn, pipelineRunner, telemetryService, errorLog, capRegistry, keyValueStore, relayerSet)
	}

	stopCh := make(chan struct{})
	lggr = logger.Named(lggr, "OCR3GenericService")
	var ps LOOPPService
	broker := net.BrokerConfig{StopCh: stopCh, Logger: lggr, GRPCOpts: grpcOpts}
	ps.Init(reportingplugins.PluginServiceName, &GRPCService[types.PluginProvider]{BrokerConfig: broker}, newService, lggr, cmd, stopCh)
	return &ps
}

func (g *LOOPPService) NewReportingPlugin(ctx context.Context, config ocr3types.ReportingPluginConfig) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	if err := g.WaitCtx(ctx); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, err
	}
	return g.Service.NewReportingPlugin(ctx, config)
}

func NewLOOPPServiceValidation(
	lggr logger.Logger,
	grpcOpts loop.GRPCOpts,
	cmd func() *exec.Cmd,
) *reportingplugins.LOOPPServiceValidation {
	newService := func(ctx context.Context, instance any) (core.ValidationService, error) {
		plug, ok := instance.(core.OCR3ReportingPluginClient)
		if !ok {
			return nil, fmt.Errorf("expected ValidationServiceClient but got %T", instance)
		}
		return plug.NewValidationService(ctx)
	}
	stopCh := make(chan struct{})
	lggr = logger.Named(lggr, "GenericService")
	var ps reportingplugins.LOOPPServiceValidation
	broker := net.BrokerConfig{StopCh: stopCh, Logger: lggr, GRPCOpts: grpcOpts}
	ps.Init(PluginServiceName, &reportingplugins.GRPCService[types.PluginProvider]{BrokerConfig: broker}, newService, lggr, cmd, stopCh)
	return &ps
}
