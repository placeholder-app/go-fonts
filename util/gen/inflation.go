package gen

import (
	"bytes"
	"compress/flate"
	"io"
)

// sourced from https://github.com/google/gxui/blob/master/gxfont/mkfont.go
func DeflateFont(src []byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	w, err := flate.NewWriter(buf, flate.DefaultCompression)
	if err != nil {
		return nil, err
	}

	if _, err := w.Write(src); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// sourced from https://github.com/google/gxui/blob/master/gxfont/gxfont.go
func InflateFont(src []byte) []byte {
	b, err := io.ReadAll(flate.NewReader(bytes.NewReader(src)))
	if err != nil {
		panic(err)
	}

	return b
}
