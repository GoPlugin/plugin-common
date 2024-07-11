package chainreader_test

import (
	"testing"

	chainreadertest "github.com/goplugin/plugin-common/pkg/loop/internal/relayer/pluginprovider/chainreader/test"
	"github.com/goplugin/plugin-common/pkg/types/interfacetests"
)

func FuzzCodec(f *testing.F) {
	interfaceTester := chainreadertest.WrapCodecTesterForLoop(&fakeCodecInterfaceTester{impl: &fakeCodec{}})
	interfacetests.RunCodecInterfaceFuzzTests(f, interfaceTester)
}
