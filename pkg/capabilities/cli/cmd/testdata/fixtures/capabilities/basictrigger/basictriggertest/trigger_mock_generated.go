// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package basictriggertest

import (
	"github.com/goplugin/plugin-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/basictrigger"
	"github.com/goplugin/plugin-common/pkg/workflows/sdk/testutils"
)

// Trigger registers a new capability mock with the runner
func Trigger(runner *testutils.Runner, fn func() (basictrigger.TriggerOutputs, error)) *testutils.TriggerMock[basictrigger.TriggerOutputs] {
	mock := testutils.MockTrigger[basictrigger.TriggerOutputs]("basic-test-trigger@1.0.0", fn)
	runner.MockCapability("basic-test-trigger@1.0.0", nil, mock)
	return mock
}