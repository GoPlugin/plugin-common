// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package referenceaction

import (
	"encoding/json"
	"fmt"
)

// Basic Test Action
type Action struct {
	// Config corresponds to the JSON schema field "config".
	Config SomeConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs *SomeInputs `json:"inputs,omitempty" yaml:"inputs,omitempty" mapstructure:"inputs,omitempty"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs *SomeOutputs `json:"outputs,omitempty" yaml:"outputs,omitempty" mapstructure:"outputs,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Action) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in Action: required")
	}
	type Plain Action
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Action(plain)
	return nil
}

type SomeConfig struct {
	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// The interval in seconds after which a new trigger event is generated.
	Number uint64 `json:"number" yaml:"number" mapstructure:"number"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SomeConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in SomeConfig: required")
	}
	if _, ok := raw["number"]; raw != nil && !ok {
		return fmt.Errorf("field number in SomeConfig: required")
	}
	type Plain SomeConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if 1 > plain.Number {
		return fmt.Errorf("field %s: must be >= %v", "number", 1)
	}
	*j = SomeConfig(plain)
	return nil
}

type SomeInputs struct {
	// InputThing corresponds to the JSON schema field "input_thing".
	InputThing bool `json:"input_thing" yaml:"input_thing" mapstructure:"input_thing"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SomeInputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["input_thing"]; raw != nil && !ok {
		return fmt.Errorf("field input_thing in SomeInputs: required")
	}
	type Plain SomeInputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SomeInputs(plain)
	return nil
}

type SomeOutputs struct {
	// AdaptedThing corresponds to the JSON schema field "adapted_thing".
	AdaptedThing string `json:"adapted_thing" yaml:"adapted_thing" mapstructure:"adapted_thing"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SomeOutputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["adapted_thing"]; raw != nil && !ok {
		return fmt.Errorf("field adapted_thing in SomeOutputs: required")
	}
	type Plain SomeOutputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SomeOutputs(plain)
	return nil
}
