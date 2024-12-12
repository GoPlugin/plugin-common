//go:build wasip1

package main

import (
	"os"

	"github.com/goplugin/plugin-common/pkg/workflows/wasm"

	"github.com/goplugin/plugin-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/basictrigger"
	"github.com/goplugin/plugin-common/pkg/workflows/sdk"
)

func BuildWorkflow(config []byte) *sdk.WorkflowSpecFactory {
	workflow := sdk.NewWorkflowSpecFactory(
		sdk.NewWorkflowParams{
			Name:  "tester",
			Owner: "ryan",
		},
	)

	triggerCfg := basictrigger.TriggerConfig{Name: "trigger", Number: 100}
	trigger := triggerCfg.New(workflow)

	sdk.Compute1[basictrigger.TriggerOutputs, bool](
		workflow,
		"transform",
		sdk.Compute1Inputs[basictrigger.TriggerOutputs]{Arg0: trigger},
		func(sdk sdk.Runtime, outputs basictrigger.TriggerOutputs) (bool, error) {
			_, err := os.ReadFile("/tmp/file")
			return false, err
		})

	return workflow
}

func main() {
	runner := wasm.NewRunner()
	workflow := BuildWorkflow(runner.Config())
	runner.Run(workflow)
}
