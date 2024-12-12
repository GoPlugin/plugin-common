// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package basictrigger

import (
	"encoding/json"
	"fmt"
)

// Basic Test Trigger
type Trigger struct {
	// Config corresponds to the JSON schema field "config".
	Config TriggerConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs *TriggerOutputs `json:"outputs,omitempty" yaml:"outputs,omitempty" mapstructure:"outputs,omitempty"`
}

type TriggerConfig struct {
	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// The interval in seconds after which a new trigger event is generated.
	Number uint64 `json:"number" yaml:"number" mapstructure:"number"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in TriggerConfig: required")
	}
	if _, ok := raw["number"]; raw != nil && !ok {
		return fmt.Errorf("field number in TriggerConfig: required")
	}
	type Plain TriggerConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if 1 > plain.Number {
		return fmt.Errorf("field %s: must be >= %v", "number", 1)
	}
	*j = TriggerConfig(plain)
	return nil
}

type TriggerOutputs struct {
	// CoolOutput corresponds to the JSON schema field "cool_output".
	CoolOutput string `json:"cool_output" yaml:"cool_output" mapstructure:"cool_output"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerOutputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["cool_output"]; raw != nil && !ok {
		return fmt.Errorf("field cool_output in TriggerOutputs: required")
	}
	type Plain TriggerOutputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TriggerOutputs(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Trigger) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in Trigger: required")
	}
	type Plain Trigger
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Trigger(plain)
	return nil
}
