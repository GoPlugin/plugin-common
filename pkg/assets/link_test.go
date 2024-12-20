package assets_test

import (
	"encoding/json"
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/goplugin/plugin-common/pkg/assets"
)

func TestAssets_NewLinkAndString(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(0)

	assert.Equal(t, "0", link.String())

	link.SetInt64(1)
	assert.Equal(t, "1", link.String())

	link.SetString("900000000000000000", 10)
	assert.Equal(t, "900000000000000000", link.String())

	link.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	assert.Equal(t, "115792089237316195423570985008687907853269984665640564039457584007913129639935", link.String())

	var nilLink *assets.Link
	assert.Equal(t, "0", nilLink.String())
}

func TestAssets_NewLinkAndLink(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(0)

	assert.Equal(t, "0.000000000000000000", link.Link())

	link.SetInt64(1)
	assert.Equal(t, "0.000000000000000001", link.Link())

	link.SetString("900000000000000000", 10)
	assert.Equal(t, "0.900000000000000000", link.Link())

	link.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	assert.Equal(t, "115792089237316195423570985008687907853269984665640564039457.584007913129639935", link.Link())

	link.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639936", 10)
	assert.Equal(t, "115792089237316195423570985008687907853269984665640564039457.584007913129639936", link.Link())

	var nilLink *assets.Link
	assert.Equal(t, "0", nilLink.Link())
}

func TestAssets_Link_MarshalJson(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(1)

	b, err := json.Marshal(link)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"1"`), b)
}

func TestAssets_Link_UnmarshalJsonOk(t *testing.T) {
	t.Parallel()

	link := assets.Link{}

	err := json.Unmarshal([]byte(`"1"`), &link)
	assert.NoError(t, err)
	assert.Equal(t, "0.000000000000000001", link.Link())
}

func TestAssets_Link_UnmarshalJsonError(t *testing.T) {
	t.Parallel()

	link := assets.Link{}

	err := json.Unmarshal([]byte(`"x"`), &link)
	assert.EqualError(t, err, "assets: cannot unmarshal \"x\" into a *assets.Link")

	err = json.Unmarshal([]byte(`1`), &link)
	assert.Equal(t, assets.ErrNoQuotesForCurrency, err)
}

func TestAssets_LinkToInt(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(0)
	assert.Equal(t, big.NewInt(0), link.ToInt())

	link = assets.NewLinkFromJuels(123)
	assert.Equal(t, big.NewInt(123), link.ToInt())
}

func TestAssets_LinkSetLink(t *testing.T) {
	t.Parallel()

	link1 := assets.NewLinkFromJuels(123)
	link2 := assets.NewLinkFromJuels(321)
	link3 := link1.Set(link2)
	assert.Equal(t, link3, link2)
}

func TestAssets_LinkCmpLink(t *testing.T) {
	t.Parallel()

	link1 := assets.NewLinkFromJuels(123)
	link2 := assets.NewLinkFromJuels(321)
	assert.NotZero(t, link1.Cmp(link2))

	link3 := assets.NewLinkFromJuels(321)
	assert.Zero(t, link3.Cmp(link2))
}

func TestAssets_LinkAddLink(t *testing.T) {
	t.Parallel()

	link1 := assets.NewLinkFromJuels(123)
	link2 := assets.NewLinkFromJuels(321)
	sum := assets.NewLinkFromJuels(123 + 321)
	assert.Equal(t, sum, link1.Add(link1, link2))
}

func TestAssets_LinkText(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(123)
	assert.Equal(t, "123", link.Text(10))
	assert.Equal(t, "7b", link.Text(16))
}

func TestAssets_LinkIsZero(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(123)
	assert.False(t, link.IsZero())

	link = assets.NewLinkFromJuels(0)
	assert.True(t, link.IsZero())
}

func TestAssets_LinkSymbol(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(123)
	assert.Equal(t, "PLI", link.Symbol())
}

func TestAssets_LinkScanValue(t *testing.T) {
	t.Parallel()

	link := assets.NewLinkFromJuels(123)
	v, err := link.Value()
	assert.NoError(t, err)

	link2 := assets.NewLinkFromJuels(0)
	err = link2.Scan(v)
	assert.NoError(t, err)
	assert.Equal(t, link2, link)

	err = link2.Scan("123")
	assert.NoError(t, err)
	assert.Equal(t, link2, link)

	err = link2.Scan([]uint8{'1', '2', '3'})
	assert.NoError(t, err)
	assert.Equal(t, link2, link)

	assert.ErrorContains(t, link2.Scan([]uint8{'x'}), "unable to set string")
	assert.ErrorContains(t, link2.Scan("123.56"), "unable to set string")
	assert.ErrorContains(t, link2.Scan(1.5), "unable to convert")
	assert.ErrorContains(t, link2.Scan(int64(123)), "unable to convert")
}

func TestLink(t *testing.T) {
	for _, tt := range []struct {
		input string
		exp   string
	}{
		{"0", "0"},
		{"1", "1"},
		{"1 juels", "1"},
		{"100000000000", "100000000000"},
		{"0.0000001 pli", "100000000000"},
		{"1000000000000", "0.000001 pli"},
		{"100000000000000", "0.0001 pli"},
		{"0.0001 pli", "0.0001 pli"},
		{"10000000000000000", "0.01 pli"},
		{"0.01 pli", "0.01 pli"},
		{"100000000000000000", "0.1 pli"},
		{"0.1 pli", "0.1 pli"},
		{"1.0 pli", "1 pli"},
		{"1000000000000000000", "1 pli"},
		{"1000000000000000000 juels", "1 pli"},
		{"1100000000000000000", "1.1 pli"},
		{"1.1pli", "1.1 pli"},
		{"1.1 pli", "1.1 pli"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			var l assets.Link
			err := l.UnmarshalText([]byte(tt.input))
			require.NoError(t, err)
			b, err := l.MarshalText()
			require.NoError(t, err)
			assert.Equal(t, tt.exp, string(b))
		})
	}
}

func FuzzLink(f *testing.F) {
	f.Add("1")
	f.Add("1 pli")
	f.Add("1.1pli")
	f.Add("2.3")
	f.Add("2.3 pli")
	f.Add("00005 pli")
	f.Add("0.0005pli")
	f.Add("1100000000000000000000000000000")
	f.Add("1100000000000000000000000000000 juels")
	f.Fuzz(func(t *testing.T, v string) {
		if len(v) > 1_000 {
			t.Skipf("too many characters: %d", len(v))
		}
		if e := tryParseExp(v); -1000 > e || e > 1000 {
			t.Skipf("exponent too large: %d", e)
		}
		var l assets.Link
		err := l.UnmarshalText([]byte(v))
		if err != nil {
			t.Skip()
		}

		b, err := l.MarshalText()
		require.NoErrorf(t, err, "failed to marshal %v after unmarshaling from %q", l, v)

		var l2 assets.Link
		err = l2.UnmarshalText(b)
		require.NoErrorf(t, err, "failed to unmarshal %s after marshaling from %v", string(b), l)
		require.Equal(t, l, l2, "unequal values after marshal/unmarshal")
	})
}

func tryParseExp(v string) int64 {
	i := strings.IndexAny(v, "Ee")
	if i == -1 {
		return -1
	}
	v = v[i+1:]
	if i := strings.IndexFunc(v, func(r rune) bool {
		switch {
		case r == '-' || r == '+':
			return false
		case r < '0' || '9' < r:
			return true
		}
		return false
	}); i > -1 {
		v = v[:i]
	}
	e, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return -1
	}
	return e
}

func Test_tryParseExp(t *testing.T) {
	got := tryParseExp("000000000E0000000060000000juels")
	assert.Equal(t, int64(60000000), got)
	got = tryParseExp("0e-80000800")
	assert.Equal(t, int64(-80000800), got)
	got = tryParseExp("0e+802444440")
	assert.Equal(t, int64(802444440), got)
}
