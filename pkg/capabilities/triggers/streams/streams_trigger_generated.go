// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package streams

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Feed struct {
	// Metadata corresponds to the JSON schema field "Metadata".
	Metadata SignersMetadata `json:"Metadata" yaml:"Metadata" mapstructure:"Metadata"`

	// Payload corresponds to the JSON schema field "Payload".
	Payload []FeedReport `json:"Payload" yaml:"Payload" mapstructure:"Payload"`

	// Timestamp corresponds to the JSON schema field "Timestamp".
	Timestamp int64 `json:"Timestamp" yaml:"Timestamp" mapstructure:"Timestamp"`
}

// The ID of the data feed.
type FeedId string

// UnmarshalJSON implements json.Unmarshaler.
func (j *FeedId) UnmarshalJSON(b []byte) error {
	type Plain FeedId
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if matched, _ := regexp.MatchString("^0x[0-9a-f]{64}$", string(plain)); !matched {
		return fmt.Errorf("field %s pattern match: must match %s", "^0x[0-9a-f]{64}$", "")
	}
	*j = FeedId(plain)
	return nil
}

type FeedReport struct {
	// BenchmarkPrice corresponds to the JSON schema field "BenchmarkPrice".
	BenchmarkPrice []uint8 `json:"BenchmarkPrice" yaml:"BenchmarkPrice" mapstructure:"BenchmarkPrice"`

	// FeedID corresponds to the JSON schema field "FeedID".
	FeedID FeedId `json:"FeedID" yaml:"FeedID" mapstructure:"FeedID"`

	// FullReport corresponds to the JSON schema field "FullReport".
	FullReport []uint8 `json:"FullReport" yaml:"FullReport" mapstructure:"FullReport"`

	// ObservationTimestamp corresponds to the JSON schema field
	// "ObservationTimestamp".
	ObservationTimestamp int64 `json:"ObservationTimestamp" yaml:"ObservationTimestamp" mapstructure:"ObservationTimestamp"`

	// ReportContext corresponds to the JSON schema field "ReportContext".
	ReportContext []uint8 `json:"ReportContext" yaml:"ReportContext" mapstructure:"ReportContext"`

	// Signatures corresponds to the JSON schema field "Signatures".
	Signatures [][]uint8 `json:"Signatures" yaml:"Signatures" mapstructure:"Signatures"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FeedReport) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["BenchmarkPrice"]; raw != nil && !ok {
		return fmt.Errorf("field BenchmarkPrice in FeedReport: required")
	}
	if _, ok := raw["FeedID"]; raw != nil && !ok {
		return fmt.Errorf("field FeedID in FeedReport: required")
	}
	if _, ok := raw["FullReport"]; raw != nil && !ok {
		return fmt.Errorf("field FullReport in FeedReport: required")
	}
	if _, ok := raw["ObservationTimestamp"]; raw != nil && !ok {
		return fmt.Errorf("field ObservationTimestamp in FeedReport: required")
	}
	if _, ok := raw["ReportContext"]; raw != nil && !ok {
		return fmt.Errorf("field ReportContext in FeedReport: required")
	}
	if _, ok := raw["Signatures"]; raw != nil && !ok {
		return fmt.Errorf("field Signatures in FeedReport: required")
	}
	type Plain FeedReport
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FeedReport(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Feed) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Metadata"]; raw != nil && !ok {
		return fmt.Errorf("field Metadata in Feed: required")
	}
	if _, ok := raw["Payload"]; raw != nil && !ok {
		return fmt.Errorf("field Payload in Feed: required")
	}
	if _, ok := raw["Timestamp"]; raw != nil && !ok {
		return fmt.Errorf("field Timestamp in Feed: required")
	}
	type Plain Feed
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Feed(plain)
	return nil
}

type SignersMetadata struct {
	// MinRequiredSignatures corresponds to the JSON schema field
	// "MinRequiredSignatures".
	MinRequiredSignatures int64 `json:"MinRequiredSignatures" yaml:"MinRequiredSignatures" mapstructure:"MinRequiredSignatures"`

	// Signers corresponds to the JSON schema field "Signers".
	Signers []string `json:"Signers" yaml:"Signers" mapstructure:"Signers"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignersMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["MinRequiredSignatures"]; raw != nil && !ok {
		return fmt.Errorf("field MinRequiredSignatures in SignersMetadata: required")
	}
	if _, ok := raw["Signers"]; raw != nil && !ok {
		return fmt.Errorf("field Signers in SignersMetadata: required")
	}
	type Plain SignersMetadata
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SignersMetadata(plain)
	return nil
}

// Streams Trigger
type Trigger struct {
	// Config corresponds to the JSON schema field "config".
	Config TriggerConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs *Feed `json:"outputs,omitempty" yaml:"outputs,omitempty" mapstructure:"outputs,omitempty"`
}

type TriggerConfig struct {
	// The IDs of the data feeds that will have their reports included in the trigger
	// event.
	FeedIds []FeedId `json:"feedIds" yaml:"feedIds" mapstructure:"feedIds"`

	// The interval in seconds after which a new trigger event is generated.
	MaxFrequencyMs uint64 `json:"maxFrequencyMs" yaml:"maxFrequencyMs" mapstructure:"maxFrequencyMs"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["feedIds"]; raw != nil && !ok {
		return fmt.Errorf("field feedIds in TriggerConfig: required")
	}
	if _, ok := raw["maxFrequencyMs"]; raw != nil && !ok {
		return fmt.Errorf("field maxFrequencyMs in TriggerConfig: required")
	}
	type Plain TriggerConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.FeedIds != nil && len(plain.FeedIds) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "feedIds", 1)
	}
	if 1 > plain.MaxFrequencyMs {
		return fmt.Errorf("field %s: must be >= %v", "maxFrequencyMs", 1)
	}
	*j = TriggerConfig(plain)
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