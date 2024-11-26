package testutils

import (
	"github.com/goplugin/plugin-common/pkg/workflows/sdk"
)

type runtime struct{}

var _ sdk.Runtime = &runtime{}
