package loop

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/goplugin/plugin-libocr/offchainreporting2/reportingplugin/median"
	ocrtypes "github.com/goplugin/plugin-libocr/offchainreporting2plus/types"

	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/loop/internal/goplugin"
	"github.com/goplugin/plugin-common/pkg/types"
	"github.com/goplugin/plugin-common/pkg/types/core"
)

var _ ocrtypes.ReportingPluginFactory = (*MedianService)(nil)

// MedianService is a [types.Service] that maintains an internal [types.PluginMedian].
type MedianService struct {
	goplugin.PluginService[*GRPCPluginMedian, types.ReportingPluginFactory]
}

// NewMedianService returns a new [*MedianService].
// cmd must return a new exec.Cmd each time it is called.
func NewMedianService(lggr logger.Logger, grpcOpts GRPCOpts, cmd func() *exec.Cmd, provider types.MedianProvider, contractAddress string, dataSource, juelsPerFeeCoin, gasPriceSubunits median.DataSource, errorLog core.ErrorLog) *MedianService {
	newService := func(ctx context.Context, instance any) (types.ReportingPluginFactory, error) {
		plug, ok := instance.(core.PluginMedian)
		if !ok {
			return nil, fmt.Errorf("expected PluginMedian but got %T", instance)
		}
		return plug.NewMedianFactory(ctx, provider, contractAddress, dataSource, juelsPerFeeCoin, gasPriceSubunits, errorLog)
	}
	stopCh := make(chan struct{})
	lggr = logger.Named(lggr, "MedianService")
	var ms MedianService
	broker := BrokerConfig{StopCh: stopCh, Logger: lggr, GRPCOpts: grpcOpts}
	ms.Init(PluginMedianName, &GRPCPluginMedian{BrokerConfig: broker}, newService, lggr, cmd, stopCh)
	return &ms
}

func (m *MedianService) NewReportingPlugin(ctx context.Context, config ocrtypes.ReportingPluginConfig) (ocrtypes.ReportingPlugin, ocrtypes.ReportingPluginInfo, error) {
	if err := m.Wait(); err != nil {
		return nil, ocrtypes.ReportingPluginInfo{}, err
	}
	return m.Service.NewReportingPlugin(ctx, config)
}