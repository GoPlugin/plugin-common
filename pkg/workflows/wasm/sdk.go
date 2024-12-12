package wasm

import "github.com/goplugin/plugin-common/pkg/workflows/sdk"

type Runtime struct{}

var _ sdk.Runtime = (*Runtime)(nil)
