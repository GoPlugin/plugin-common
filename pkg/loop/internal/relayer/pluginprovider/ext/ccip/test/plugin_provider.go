package test

import (
	"github.com/google/uuid"

	"github.com/goplugin/plugin-common/pkg/types"
)

var (
	ExecutionRelayArgs = types.RelayArgs{
		ExternalJobID: uuid.MustParse("12348153-1234-5678-9012-fd0985d00000"),
		JobID:         42,
		ContractID:    "exec-testcontract",
		New:           true,
		RelayConfig:   []byte{1: 4, 36: 101},
		ProviderType:  string(types.CCIPExecution),
	}
	ExecutionPluginArgs = types.PluginArgs{
		TransmitterID: "exec-testtransmitter",
		PluginConfig:  []byte{133: 79},
	}
	CommitRelayArgs = types.RelayArgs{
		ExternalJobID: uuid.MustParse("12348153-1234-5678-9012-fd0985d00000"),
		JobID:         42,
		ContractID:    "commit-testcontract",
		New:           true,
		RelayConfig:   []byte{1: 4, 36: 101},
		ProviderType:  string(types.CCIPCommit),
	}
	CommitPluginArgs = types.PluginArgs{
		TransmitterID: "commit-testtransmitter",
		PluginConfig:  []byte{133: 79},
	}
)
