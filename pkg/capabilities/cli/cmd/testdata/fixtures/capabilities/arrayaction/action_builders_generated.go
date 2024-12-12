// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package arrayaction

import (
	"github.com/goplugin/plugin-common/pkg/capabilities"
	"github.com/goplugin/plugin-common/pkg/workflows/sdk"
)

func (cfg ActionConfig) New(w *sdk.WorkflowSpecFactory, ref string, input ActionInput) sdk.CapDefinition[[]ActionOutputsElem] {

	def := sdk.StepDefinition{
		ID: "array-test-action@1.0.0", Ref: ref,
		Inputs: input.ToSteps(),
		Config: map[string]any{
			"details": cfg.Details,
		},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	step := sdk.Step[[]ActionOutputsElem]{Definition: def}
	return step.AddTo(w)
}

type ActionOutputsElemCap interface {
	sdk.CapDefinition[ActionOutputsElem]
	Results() ActionOutputsElemResultsCap
	private()
}

// ActionOutputsElemCapFromStep should only be called from generated code to assure type safety
func ActionOutputsElemCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[ActionOutputsElem]) ActionOutputsElemCap {
	raw := step.AddTo(w)
	return &actionOutputsElem{CapDefinition: raw}
}

type actionOutputsElem struct {
	sdk.CapDefinition[ActionOutputsElem]
}

func (*actionOutputsElem) private() {}
func (c *actionOutputsElem) Results() ActionOutputsElemResultsCap {
	return &actionOutputsElemResults{CapDefinition: sdk.AccessField[ActionOutputsElem, ActionOutputsElemResults](c.CapDefinition, "results")}
}

func NewActionOutputsElemFromFields(
	results ActionOutputsElemResultsCap) ActionOutputsElemCap {
	return &simpleActionOutputsElem{
		CapDefinition: sdk.ComponentCapDefinition[ActionOutputsElem]{
			"results": results.Ref(),
		},
		results: results,
	}
}

type simpleActionOutputsElem struct {
	sdk.CapDefinition[ActionOutputsElem]
	results ActionOutputsElemResultsCap
}

func (c *simpleActionOutputsElem) Results() ActionOutputsElemResultsCap {
	return c.results
}

func (c *simpleActionOutputsElem) private() {}

type ActionOutputsElemResultsCap interface {
	sdk.CapDefinition[ActionOutputsElemResults]
	AdaptedThing() sdk.CapDefinition[string]
	private()
}

// ActionOutputsElemResultsCapFromStep should only be called from generated code to assure type safety
func ActionOutputsElemResultsCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[ActionOutputsElemResults]) ActionOutputsElemResultsCap {
	raw := step.AddTo(w)
	return &actionOutputsElemResults{CapDefinition: raw}
}

type actionOutputsElemResults struct {
	sdk.CapDefinition[ActionOutputsElemResults]
}

func (*actionOutputsElemResults) private() {}
func (c *actionOutputsElemResults) AdaptedThing() sdk.CapDefinition[string] {
	return sdk.AccessField[ActionOutputsElemResults, string](c.CapDefinition, "adapted_thing")
}

func NewActionOutputsElemResultsFromFields(
	adaptedThing sdk.CapDefinition[string]) ActionOutputsElemResultsCap {
	return &simpleActionOutputsElemResults{
		CapDefinition: sdk.ComponentCapDefinition[ActionOutputsElemResults]{
			"adapted_thing": adaptedThing.Ref(),
		},
		adaptedThing: adaptedThing,
	}
}

type simpleActionOutputsElemResults struct {
	sdk.CapDefinition[ActionOutputsElemResults]
	adaptedThing sdk.CapDefinition[string]
}

func (c *simpleActionOutputsElemResults) AdaptedThing() sdk.CapDefinition[string] {
	return c.adaptedThing
}

func (c *simpleActionOutputsElemResults) private() {}

type ActionInput struct {
	Metadata sdk.CapDefinition[ActionInputsMetadata]
}

func (input ActionInput) ToSteps() sdk.StepInputs {
	return sdk.StepInputs{
		Mapping: map[string]any{
			"metadata": input.Metadata.Ref(),
		},
	}
}