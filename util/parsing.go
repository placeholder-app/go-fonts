package util

import (
	"github.com/golang/freetype/truetype"
	"github.com/placeholder-app/go-fonts/deflated"
	"github.com/placeholder-app/go-fonts/util/gen"
	"golang.org/x/image/font"
)

type Font struct {
	// name of the font
	Name string
	// weight for this variant of the font
	Weight string
	// source of the deflated font
	Source string
	// truetype parsed font
	Font *truetype.Font

	// internal map of font faces, indexed by font size for each dpi
	//
	// used to avoid creating a new font face for a dpi+size when one has been created before
	//
	// indexed by dpi and then size as dpi is likely to stay the same
	fontFaces map[float64]map[float64]FontFace
}

type FontFace struct {
	// font face used to render the text
	Face font.Face
	// font size
	Size float64
	// dpi
	DPI float64
}

// parse an inflated font file to a Font struct
func ParseFont(deflatedFont deflated.DeflatedFont) *Font {
	f, err := truetype.Parse(gen.InflateFont(deflatedFont.Bytes))
	if err != nil {
		panic(err)
	}

	return (&Font{
		Name:   deflatedFont.Name,
		Weight: deflatedFont.Weight,
		Source: deflatedFont.Source,
		Font:   f,

		fontFaces: make(map[float64]map[float64]FontFace),
	})
}
