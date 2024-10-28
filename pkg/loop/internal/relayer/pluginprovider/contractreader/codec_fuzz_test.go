package contractreader_test

import (
	"testing"

	chaincomponentstest "github.com/goplugin/plugin-common/pkg/loop/internal/relayer/pluginprovider/contractreader/test"
	"github.com/goplugin/plugin-common/pkg/types/interfacetests"
)

func FuzzCodec(f *testing.F) {
	interfaceTester := chaincomponentstest.WrapCodecTesterForLoop(&fakeCodecInterfaceTester{impl: &fakeCodec{}})
	interfacetests.RunCodecInterfaceFuzzTests(f, interfaceTester)
}
