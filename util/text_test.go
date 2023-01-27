package util

import (
	"testing"

	"github.com/placeholder-app/go-fonts/deflated"
	a "github.com/placeholder-app/go-fonts/util/assert"
)

func TestGetFontFace(t *testing.T) {
	parsed := ParseFont(deflated.CalSansSemiBold)
	face := parsed.GetFontFace(12, 72)
	faceFromStorage := parsed.GetFontFace(12, 72)

	a.NotEquals(t, face, nil)
	a.NotEquals(t, faceFromStorage, nil)
}

func TestGetTextData(t *testing.T) {
	parsed := ParseFont(deflated.CalSansSemiBold)
	face := parsed.GetFontFace(12, 72)

	textData := parsed.GetTextData(face, "Hello World")

	a.Equals(t, textData.FontFace, face)
	a.Equals(t, textData.Text, "Hello World")
	a.Equals(t, textData.Width, 64)
	a.Equals(t, textData.Height, 10)

	a.Equals(t, (0 - textData.HeightMin).Ceil(), 10)
	a.Equals(t, (textData.HeightMax).Ceil(), 0)

	textDataApple := parsed.GetTextData(face, "apple")

	a.Equals(t, textDataApple.FontFace, face)
	a.Equals(t, textDataApple.Text, "apple")
	a.Equals(t, textDataApple.Width, 31)
	a.Equals(t, textDataApple.Height, 7)

	a.Equals(t, (0 - textDataApple.HeightMin).Ceil(), 10)
	a.Equals(t, (textDataApple.HeightMax).Ceil(), 3)
}
