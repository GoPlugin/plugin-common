package v4_test

import (
	"bytes"
	"context"
	"fmt"

	ocr2plus_types "github.com/goplugin/plugin-libocr/offchainreporting2plus/types"

	testtypes "github.com/goplugin/plugin-common/pkg/loop/internal/test/types"
	mercury_v4_types "github.com/goplugin/plugin-common/pkg/types/mercury/v4"
)

var ReportCodec = staticReportCodec{}

type ReportCodecEvaluator interface {
	mercury_v4_types.ReportCodec
	testtypes.Evaluator[mercury_v4_types.ReportCodec]
}

type staticReportCodec struct{}

var _ ReportCodecEvaluator = staticReportCodec{}

func (staticReportCodec) BuildReport(ctx context.Context, fields mercury_v4_types.ReportFields) (ocr2plus_types.Report, error) {
	return Fixtures.Report, nil
}

func (staticReportCodec) MaxReportLength(ctx context.Context, n int) (int, error) {
	return Fixtures.MaxReportLength, nil
}

func (staticReportCodec) ObservationTimestampFromReport(ctx context.Context, report ocr2plus_types.Report) (uint32, error) {
	return Fixtures.ObservationTimestamp, nil
}

func (staticReportCodec) Evaluate(ctx context.Context, other mercury_v4_types.ReportCodec) error {
	gotReport, err := other.BuildReport(ctx, Fixtures.ReportFields)
	if err != nil {
		return fmt.Errorf("failed to BuildReport: %w", err)
	}
	if !bytes.Equal(gotReport, Fixtures.Report) {
		return fmt.Errorf("expected Report %x but got %x", Fixtures.Report, gotReport)
	}
	gotMax, err := other.MaxReportLength(ctx, Fixtures.MaxReportLength)
	if err != nil {
		return fmt.Errorf("failed to get MaxReportLength: %w", err)
	}
	if gotMax != Fixtures.MaxReportLength {
		return fmt.Errorf("expected MaxReportLength %d but got %d", Fixtures.MaxReportLength, gotMax)
	}
	gotObservedTimestamp, err := other.ObservationTimestampFromReport(ctx, gotReport)
	if err != nil {
		return fmt.Errorf("failed to get ObservationTimestampFromReport: %w", err)
	}
	if gotObservedTimestamp != Fixtures.ObservationTimestamp {
		return fmt.Errorf("expected ObservationTimestampFromReport %d but got %d", Fixtures.ObservationTimestamp, gotObservedTimestamp)
	}
	return nil
}
