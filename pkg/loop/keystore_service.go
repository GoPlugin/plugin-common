package loop

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/loop/internal/goplugin"
	"github.com/goplugin/plugin-common/pkg/loop/internal/keystore"
)

// KeystoreService is a [types.Service] that maintains an internal [keystore.Keystore].
type KeystoreService struct {
	goplugin.PluginService[*GRPCPluginKeystore, keystore.GRPCService]
}

func NewKeystoreService(lggr logger.Logger, grpcOpts GRPCOpts, cmd func() *exec.Cmd, config []byte) *KeystoreService {
	newService := func(ctx context.Context, instance any) (keystore.GRPCService, error) {
		plug, ok := instance.(*keystore.Client)
		if !ok {
			return nil, fmt.Errorf("expected PluginKeystore but got %T", instance)
		}
		return plug, nil
	}
	stopCh := make(chan struct{})
	lggr = logger.Named(lggr, "KeystoreService")
	var rs KeystoreService
	broker := BrokerConfig{StopCh: stopCh, Logger: lggr, GRPCOpts: grpcOpts}
	rs.Init(PluginKeystoreName, &GRPCPluginKeystore{BrokerConfig: broker}, newService, lggr, cmd, stopCh)
	return &rs
}
