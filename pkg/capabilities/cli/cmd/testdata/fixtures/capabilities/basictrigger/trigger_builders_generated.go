// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package basictrigger

import (
	"github.com/goplugin/plugin-common/pkg/capabilities"
	"github.com/goplugin/plugin-common/pkg/workflows/sdk"
)

func (cfg TriggerConfig) New(w *sdk.WorkflowSpecFactory) TriggerOutputsCap {
	ref := "trigger"
	def := sdk.StepDefinition{
		ID: "basic-test-trigger@1.0.0", Ref: ref,
		Inputs: sdk.StepInputs{},
		Config: map[string]any{
			"name":   cfg.Name,
			"number": cfg.Number,
		},
		CapabilityType: capabilities.CapabilityTypeTrigger,
	}

	step := sdk.Step[TriggerOutputs]{Definition: def}
	return TriggerOutputsCapFromStep(w, step)
}

type TriggerOutputsCap interface {
	sdk.CapDefinition[TriggerOutputs]
	CoolOutput() sdk.CapDefinition[string]
	private()
}

// TriggerOutputsCapFromStep should only be called from generated code to assure type safety
func TriggerOutputsCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[TriggerOutputs]) TriggerOutputsCap {
	raw := step.AddTo(w)
	return &triggerOutputs{CapDefinition: raw}
}

type triggerOutputs struct {
	sdk.CapDefinition[TriggerOutputs]
}

func (*triggerOutputs) private() {}
func (c *triggerOutputs) CoolOutput() sdk.CapDefinition[string] {
	return sdk.AccessField[TriggerOutputs, string](c.CapDefinition, "cool_output")
}

func NewTriggerOutputsFromFields(
	coolOutput sdk.CapDefinition[string]) TriggerOutputsCap {
	return &simpleTriggerOutputs{
		CapDefinition: sdk.ComponentCapDefinition[TriggerOutputs]{
			"cool_output": coolOutput.Ref(),
		},
		coolOutput: coolOutput,
	}
}

type simpleTriggerOutputs struct {
	sdk.CapDefinition[TriggerOutputs]
	coolOutput sdk.CapDefinition[string]
}

func (c *simpleTriggerOutputs) CoolOutput() sdk.CapDefinition[string] {
	return c.coolOutput
}

func (c *simpleTriggerOutputs) private() {}
