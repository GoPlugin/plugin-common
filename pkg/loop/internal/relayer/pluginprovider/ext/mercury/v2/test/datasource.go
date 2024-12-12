package v2_test

import (
	"context"
	"fmt"

	ocrtypes "github.com/goplugin/plugin-libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"

	testtypes "github.com/goplugin/plugin-common/pkg/loop/internal/test/types"
	mercury_v2_types "github.com/goplugin/plugin-common/pkg/types/mercury/v2"
)

var DataSource = staticDataSource{}

type DataSourceEvaluator interface {
	mercury_v2_types.DataSource
	testtypes.Evaluator[mercury_v2_types.DataSource]
}

type staticDataSource struct{}

var _ mercury_v2_types.DataSource = staticDataSource{}

func (staticDataSource) Observe(ctx context.Context, repts ocrtypes.ReportTimestamp, fetchMaxFinalizedTimestamp bool) (mercury_v2_types.Observation, error) {
	return Fixtures.Observation, nil
}

func (staticDataSource) Evaluate(ctx context.Context, other mercury_v2_types.DataSource) error {
	gotVal, err := other.Observe(ctx, Fixtures.ReportTimestamp, false)
	if err != nil {
		return fmt.Errorf("failed to observe dataSource: %w", err)
	}
	if !assert.ObjectsAreEqual(Fixtures.Observation, gotVal) {
		return fmt.Errorf("expected Value %v but got %v", Fixtures.Observation, gotVal)
	}
	return nil
}