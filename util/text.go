package util

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type TextInfo struct {
	// font face used to render the text
	FontFace FontFace
	// text to render
	Text string
	// width of the text
	Width int
	// height of the text
	Height int

	// min height point from the baseline
	HeightMin fixed.Int26_6
	// max height point from the baseline
	HeightMax fixed.Int26_6
}

// get the font face for a specific dpi and font size
//
// uses a previously created font face if one exists for the dpi and size,
// otherwise creates a new one and stores it for future use
func (f *Font) GetFontFace(dpi float64, size float64) FontFace {
	if f.fontFaces[dpi] == nil {
		f.fontFaces[dpi] = make(map[float64]FontFace)
	} else {
		if f.fontFaces[dpi][size] != (FontFace{}) {
			return f.fontFaces[dpi][size]
		}
	}

	fontFace := FontFace{
		Face: truetype.NewFace(f.Font, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
		Size: size,
		DPI:  dpi,
	}

	f.fontFaces[dpi][size] = fontFace
	return fontFace
}

// calculates the text width and height for later use rendering the text
func (f *Font) GetTextData(face FontFace, text string) *TextInfo {
	var textWidth int = 0
	var heightMin, heightMax fixed.Int26_6 = 0, 0

	for _, r := range text {
		bounds, advance, _ := face.Face.GlyphBounds(r)
		textWidth += advance.Floor()

		// to get the total height we need to find to the lowest point and the highest point for the string.
		if bounds.Min.Y < heightMin {
			heightMin = bounds.Min.Y
		}
		if bounds.Max.Y > heightMax {
			heightMax = bounds.Max.Y
		}
	}
	// then we can subtract the two to get the value we can use when rendering the text.
	textHeight := (0 - heightMin - heightMax).Ceil()

	return &TextInfo{
		FontFace: face,
		Text:     text,
		Width:    textWidth,
		Height:   textHeight,

		HeightMin: heightMin,
		HeightMax: heightMax,
	}
}
