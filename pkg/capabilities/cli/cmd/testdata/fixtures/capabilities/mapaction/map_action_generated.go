// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package mapaction

import (
	"encoding/json"
	"fmt"
)

// Map Consumer Test Action
type Action struct {
	// Config corresponds to the JSON schema field "config".
	Config ActionConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs ActionInputs `json:"inputs" yaml:"inputs" mapstructure:"inputs"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs ActionOutputs `json:"outputs" yaml:"outputs" mapstructure:"outputs"`
}

type ActionConfig map[string]interface{}

type ActionInputs struct {
	// Payload corresponds to the JSON schema field "payload".
	Payload ActionInputsPayload `json:"payload" yaml:"payload" mapstructure:"payload"`
}

type ActionInputsPayload map[string]string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ActionInputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["payload"]; raw != nil && !ok {
		return fmt.Errorf("field payload in ActionInputs: required")
	}
	type Plain ActionInputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ActionInputs(plain)
	return nil
}

type ActionOutputs struct {
	// Payload corresponds to the JSON schema field "payload".
	Payload ActionOutputsPayload `json:"payload" yaml:"payload" mapstructure:"payload"`
}

type ActionOutputsPayload map[string]string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ActionOutputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["payload"]; raw != nil && !ok {
		return fmt.Errorf("field payload in ActionOutputs: required")
	}
	type Plain ActionOutputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ActionOutputs(plain)
	return nil
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
	if _, ok := raw["inputs"]; raw != nil && !ok {
		return fmt.Errorf("field inputs in Action: required")
	}
	if _, ok := raw["outputs"]; raw != nil && !ok {
		return fmt.Errorf("field outputs in Action: required")
	}
	type Plain Action
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Action(plain)
	return nil
}
