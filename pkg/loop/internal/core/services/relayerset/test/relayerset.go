package test

import (
	"context"

	"github.com/goplugin/plugin-common/pkg/types"
	"github.com/goplugin/plugin-common/pkg/types/core"
)

type RelayerSet struct {
}

func (s RelayerSet) Get(ctx context.Context, relayID types.RelayID) (core.Relayer, error) {
	//TODO implement me
	panic("implement me")
}

func (s RelayerSet) List(ctx context.Context, relayIDs ...types.RelayID) (map[types.RelayID]core.Relayer, error) {
	//TODO implement me
	panic("implement me")
}
