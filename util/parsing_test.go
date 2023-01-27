package util

import (
	"testing"

	"github.com/placeholder-app/go-fonts/deflated"
	a "github.com/placeholder-app/go-fonts/util/assert"
	"github.com/placeholder-app/go-fonts/util/gen"
)

func TestParseFont(t *testing.T) {
	parsed := ParseFont(deflated.CalSansSemiBold)

	a.Equals(t, parsed.Name, "Cal Sans")
	a.Equals(t, parsed.Weight, "SemiBold")
	a.Equals(t, parsed.Source, "https://github.com/calcom/font/raw/main/fonts/webfonts/CalSans-SemiBold.ttf")

	a.NotEquals(t, parsed.Font, nil)
}

func TestParseFontFails(t *testing.T) {
	a.Throws(t, func() {
		ParseFont(deflated.DeflatedFont{})
	})

	a.Throws(t, func() {
		deflatedBytes, _ := gen.DeflateFont([]byte{0xcc})
		ParseFont(deflated.DeflatedFont{
			Name:  "Invalid Font",
			Bytes: deflatedBytes,
		})
	})
}
