package loop_test

import (
	"os/exec"

	"github.com/goplugin/plugin-common/pkg/loop/internal/test"
)

type HelperProcessCommand test.HelperProcessCommand

func (h *HelperProcessCommand) New() *exec.Cmd {
	h.CommandLocation = "./internal/test/cmd/main.go"
	return (test.HelperProcessCommand)(*h).New()
}

func NewHelperProcessCommand(command string, staticChecks bool, throwErrorType int) *exec.Cmd {
	h := HelperProcessCommand{
		Command:        command,
		StaticChecks:   staticChecks,
		ThrowErrorType: throwErrorType,
	}
	return h.New()
}
