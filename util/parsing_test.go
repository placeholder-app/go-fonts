package util_test

import (
	"testing"

	a "github.com/james-elicx/go-utils/assert"
	"github.com/placeholder-app/go-fonts/deflated"
	"github.com/placeholder-app/go-fonts/util"
	"github.com/placeholder-app/go-fonts/util/gen"
)

func TestParseFont(t *testing.T) {
	parsed := util.ParseFont(deflated.CalSansSemiBold)

	a.Equals(t, parsed.Name, "Cal Sans")
	a.Equals(t, parsed.Weight, "SemiBold")
	a.Equals(t, parsed.Source, "https://github.com/calcom/font/raw/main/fonts/webfonts/CalSans-SemiBold.ttf")

	a.NotEquals(t, parsed.Font, nil)
}

func TestParseFontFails(t *testing.T) {
	a.Throws(t, func() {
		util.ParseFont(deflated.DeflatedFont{})
	})

	a.Throws(t, func() {
		deflatedBytes, _ := gen.DeflateFont([]byte{0xcc})
		util.ParseFont(deflated.DeflatedFont{
			Name:  "Invalid Font",
			Bytes: deflatedBytes,
		})
	})
}
