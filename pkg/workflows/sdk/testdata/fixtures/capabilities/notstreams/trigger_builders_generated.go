// Code generated by github.com/goplugin/plugin-common/pkg/capabilities/cli, DO NOT EDIT.

package notstreams

import (
	"github.com/goplugin/plugin-common/pkg/capabilities"
	"github.com/goplugin/plugin-common/pkg/workflows/sdk"
)

func (cfg TriggerConfig) New(w *sdk.WorkflowSpecFactory) FeedCap {
	ref := "trigger"
	def := sdk.StepDefinition{
		ID: "notstreams@1.0.0", Ref: ref,
		Inputs: sdk.StepInputs{},
		Config: map[string]any{
			"maxFrequencyMs": cfg.MaxFrequencyMs,
		},
		CapabilityType: capabilities.CapabilityTypeTrigger,
	}

	step := sdk.Step[Feed]{Definition: def}
	return FeedCapFromStep(w, step)
}

type FeedCap interface {
	sdk.CapDefinition[Feed]
	Metadata() SignerMetadataCap
	Payload() FeedReportCap
	Timestamp() sdk.CapDefinition[int64]
	private()
}

// FeedCapFromStep should only be called from generated code to assure type safety
func FeedCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[Feed]) FeedCap {
	raw := step.AddTo(w)
	return &feed{CapDefinition: raw}
}

type feed struct {
	sdk.CapDefinition[Feed]
}

func (*feed) private() {}
func (c *feed) Metadata() SignerMetadataCap {
	return &signerMetadata{CapDefinition: sdk.AccessField[Feed, SignerMetadata](c.CapDefinition, "Metadata")}
}
func (c *feed) Payload() FeedReportCap {
	return &feedReport{CapDefinition: sdk.AccessField[Feed, FeedReport](c.CapDefinition, "Payload")}
}
func (c *feed) Timestamp() sdk.CapDefinition[int64] {
	return sdk.AccessField[Feed, int64](c.CapDefinition, "Timestamp")
}

func NewFeedFromFields(
	metadata SignerMetadataCap,
	payload FeedReportCap,
	timestamp sdk.CapDefinition[int64]) FeedCap {
	return &simpleFeed{
		CapDefinition: sdk.ComponentCapDefinition[Feed]{
			"Metadata":  metadata.Ref(),
			"Payload":   payload.Ref(),
			"Timestamp": timestamp.Ref(),
		},
		metadata:  metadata,
		payload:   payload,
		timestamp: timestamp,
	}
}

type simpleFeed struct {
	sdk.CapDefinition[Feed]
	metadata  SignerMetadataCap
	payload   FeedReportCap
	timestamp sdk.CapDefinition[int64]
}

func (c *simpleFeed) Metadata() SignerMetadataCap {
	return c.metadata
}
func (c *simpleFeed) Payload() FeedReportCap {
	return c.payload
}
func (c *simpleFeed) Timestamp() sdk.CapDefinition[int64] {
	return c.timestamp
}

func (c *simpleFeed) private() {}

type FeedReportCap interface {
	sdk.CapDefinition[FeedReport]
	BuyPrice() sdk.CapDefinition[[]uint8]
	FullReport() sdk.CapDefinition[[]uint8]
	ObservationTimestamp() sdk.CapDefinition[int64]
	ReportContext() sdk.CapDefinition[[]uint8]
	SellPrice() sdk.CapDefinition[[]uint8]
	Signature() sdk.CapDefinition[[]uint8]
	private()
}

// FeedReportCapFromStep should only be called from generated code to assure type safety
func FeedReportCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[FeedReport]) FeedReportCap {
	raw := step.AddTo(w)
	return &feedReport{CapDefinition: raw}
}

type feedReport struct {
	sdk.CapDefinition[FeedReport]
}

func (*feedReport) private() {}
func (c *feedReport) BuyPrice() sdk.CapDefinition[[]uint8] {
	return sdk.AccessField[FeedReport, []uint8](c.CapDefinition, "BuyPrice")
}
func (c *feedReport) FullReport() sdk.CapDefinition[[]uint8] {
	return sdk.AccessField[FeedReport, []uint8](c.CapDefinition, "FullReport")
}
func (c *feedReport) ObservationTimestamp() sdk.CapDefinition[int64] {
	return sdk.AccessField[FeedReport, int64](c.CapDefinition, "ObservationTimestamp")
}
func (c *feedReport) ReportContext() sdk.CapDefinition[[]uint8] {
	return sdk.AccessField[FeedReport, []uint8](c.CapDefinition, "ReportContext")
}
func (c *feedReport) SellPrice() sdk.CapDefinition[[]uint8] {
	return sdk.AccessField[FeedReport, []uint8](c.CapDefinition, "SellPrice")
}
func (c *feedReport) Signature() sdk.CapDefinition[[]uint8] {
	return sdk.AccessField[FeedReport, []uint8](c.CapDefinition, "Signature")
}

func NewFeedReportFromFields(
	buyPrice sdk.CapDefinition[[]uint8],
	fullReport sdk.CapDefinition[[]uint8],
	observationTimestamp sdk.CapDefinition[int64],
	reportContext sdk.CapDefinition[[]uint8],
	sellPrice sdk.CapDefinition[[]uint8],
	signature sdk.CapDefinition[[]uint8]) FeedReportCap {
	return &simpleFeedReport{
		CapDefinition: sdk.ComponentCapDefinition[FeedReport]{
			"BuyPrice":             buyPrice.Ref(),
			"FullReport":           fullReport.Ref(),
			"ObservationTimestamp": observationTimestamp.Ref(),
			"ReportContext":        reportContext.Ref(),
			"SellPrice":            sellPrice.Ref(),
			"Signature":            signature.Ref(),
		},
		buyPrice:             buyPrice,
		fullReport:           fullReport,
		observationTimestamp: observationTimestamp,
		reportContext:        reportContext,
		sellPrice:            sellPrice,
		signature:            signature,
	}
}

type simpleFeedReport struct {
	sdk.CapDefinition[FeedReport]
	buyPrice             sdk.CapDefinition[[]uint8]
	fullReport           sdk.CapDefinition[[]uint8]
	observationTimestamp sdk.CapDefinition[int64]
	reportContext        sdk.CapDefinition[[]uint8]
	sellPrice            sdk.CapDefinition[[]uint8]
	signature            sdk.CapDefinition[[]uint8]
}

func (c *simpleFeedReport) BuyPrice() sdk.CapDefinition[[]uint8] {
	return c.buyPrice
}
func (c *simpleFeedReport) FullReport() sdk.CapDefinition[[]uint8] {
	return c.fullReport
}
func (c *simpleFeedReport) ObservationTimestamp() sdk.CapDefinition[int64] {
	return c.observationTimestamp
}
func (c *simpleFeedReport) ReportContext() sdk.CapDefinition[[]uint8] {
	return c.reportContext
}
func (c *simpleFeedReport) SellPrice() sdk.CapDefinition[[]uint8] {
	return c.sellPrice
}
func (c *simpleFeedReport) Signature() sdk.CapDefinition[[]uint8] {
	return c.signature
}

func (c *simpleFeedReport) private() {}

type SignerMetadataCap interface {
	sdk.CapDefinition[SignerMetadata]
	Signer() sdk.CapDefinition[string]
	private()
}

// SignerMetadataCapFromStep should only be called from generated code to assure type safety
func SignerMetadataCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[SignerMetadata]) SignerMetadataCap {
	raw := step.AddTo(w)
	return &signerMetadata{CapDefinition: raw}
}

type signerMetadata struct {
	sdk.CapDefinition[SignerMetadata]
}

func (*signerMetadata) private() {}
func (c *signerMetadata) Signer() sdk.CapDefinition[string] {
	return sdk.AccessField[SignerMetadata, string](c.CapDefinition, "Signer")
}

func NewSignerMetadataFromFields(
	signer sdk.CapDefinition[string]) SignerMetadataCap {
	return &simpleSignerMetadata{
		CapDefinition: sdk.ComponentCapDefinition[SignerMetadata]{
			"Signer": signer.Ref(),
		},
		signer: signer,
	}
}

type simpleSignerMetadata struct {
	sdk.CapDefinition[SignerMetadata]
	signer sdk.CapDefinition[string]
}

func (c *simpleSignerMetadata) Signer() sdk.CapDefinition[string] {
	return c.signer
}

func (c *simpleSignerMetadata) private() {}
