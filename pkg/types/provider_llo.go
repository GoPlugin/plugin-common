package types

import (
	"github.com/goplugin/plugin-common/pkg/types/llo"
)

type LLOProvider interface {
	ConfigProvider
	ContractTransmitter() llo.Transmitter
	ChannelDefinitionCache() llo.ChannelDefinitionCache
}
