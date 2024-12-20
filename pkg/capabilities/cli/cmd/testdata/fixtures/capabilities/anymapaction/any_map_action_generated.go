// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package anymapaction

import (
	"encoding/json"
	"fmt"
)

// Map Consumer Test Action
type MapAction struct {
	// Config corresponds to the JSON schema field "config".
	Config MapActionConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs MapActionInputs `json:"inputs" yaml:"inputs" mapstructure:"inputs"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs MapActionOutputs `json:"outputs" yaml:"outputs" mapstructure:"outputs"`
}

type MapActionConfig map[string]interface{}

type MapActionInputs struct {
	// Payload corresponds to the JSON schema field "payload".
	Payload MapActionInputsPayload `json:"payload" yaml:"payload" mapstructure:"payload"`
}

type MapActionInputsPayload map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MapActionInputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["payload"]; raw != nil && !ok {
		return fmt.Errorf("field payload in MapActionInputs: required")
	}
	type Plain MapActionInputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MapActionInputs(plain)
	return nil
}

type MapActionOutputs struct {
	// Payload corresponds to the JSON schema field "payload".
	Payload MapActionOutputsPayload `json:"payload" yaml:"payload" mapstructure:"payload"`
}

type MapActionOutputsPayload map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MapActionOutputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["payload"]; raw != nil && !ok {
		return fmt.Errorf("field payload in MapActionOutputs: required")
	}
	type Plain MapActionOutputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MapActionOutputs(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MapAction) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in MapAction: required")
	}
	if _, ok := raw["inputs"]; raw != nil && !ok {
		return fmt.Errorf("field inputs in MapAction: required")
	}
	if _, ok := raw["outputs"]; raw != nil && !ok {
		return fmt.Errorf("field outputs in MapAction: required")
	}
	type Plain MapAction
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MapAction(plain)
	return nil
}
