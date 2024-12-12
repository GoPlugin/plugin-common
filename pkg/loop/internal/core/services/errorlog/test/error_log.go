package test

import (
	"context"
	"fmt"

	testtypes "github.com/goplugin/plugin-common/pkg/loop/internal/test/types"
	"github.com/goplugin/plugin-common/pkg/types/core"
)

var ErrorLog = StaticErrorLog{errMsg: "an error"}

type StaticErrorLog struct {
	errMsg string
}

var _ testtypes.ErrorLogEvaluator = StaticErrorLog{}

func (s StaticErrorLog) SaveError(ctx context.Context, msg string) error {
	if msg != s.errMsg {
		return fmt.Errorf("expected %q but got %q", s.errMsg, msg)
	}
	return nil
}

func (s StaticErrorLog) Evaluate(ctx context.Context, other core.ErrorLog) error {
	return s.SaveError(ctx, s.errMsg)
}
