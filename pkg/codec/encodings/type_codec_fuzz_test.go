package encodings_test

import (
	"testing"

	"github.com/goplugin/plugin-common/pkg/types/interfacetests"
)

func FuzzTypeCodec(f *testing.F) {
	tester := &bigEndianInterfaceTester{}
	interfacetests.RunCodecInterfaceFuzzTests(f, tester)
}
