package gen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type FontConfig struct {
	// name of the font
	FontName string
	// weight for this variant of the font
	FontWeight string
	// url to download the font from
	Url string
}

// sourced/adapted from https://github.com/google/gxui/blob/master/gxfont/mkfont.go
func DownloadFont(cfg FontConfig, directory string) {
	fileName := fmt.Sprintf("%s_%s.go", strings.ReplaceAll(strings.ToLower(cfg.FontName), " ", "_"), strings.ReplaceAll(strings.ToLower(cfg.FontWeight), " ", "_"))
	variableName := fmt.Sprintf("%s%s", strings.ReplaceAll(cfg.FontName, " ", ""), strings.ReplaceAll(cfg.FontWeight, " ", ""))

	path := path.Join(directory, fileName)

	if file, err := os.Stat(path); os.IsExist(err) || file != nil {
		fmt.Printf("file for %s already exists\n", variableName)
		return
	}

	resp, err := http.Get(cfg.Url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading %q: %v", variableName, err)
	}

	data, err = DeflateFont(data)
	if err != nil {
		log.Fatalf("error deflating %q: %v", variableName, err)
	}

	w := new(bytes.Buffer)
	fmt.Fprint(w, "// GENERATED BY adaption of google/gxui/gxfont\n\n")
	fmt.Fprintf(w, "package deflated\n\n")
	fmt.Fprintf(w, "var %s DeflatedFont = DeflatedFont{\n", variableName)
	fmt.Fprintf(w, "\tName:   %q,\n", cfg.FontName)
	fmt.Fprintf(w, "\tWeight: %q,\n", cfg.FontWeight)
	fmt.Fprintf(w, "\tSource: %q,\n", cfg.Url)
	fmt.Fprintf(w, "\tBytes:  []byte{\n")

	for len(data) > 0 {
		n := 16
		if n > len(data) {
			n = len(data)
		}
		for _, c := range data[:n] {
			fmt.Fprintf(w, "0x%02x,", c)
		}
		fmt.Fprintf(w, "\n")
		data = data[n:]
	}
	fmt.Fprint(w, "\t},\n")
	fmt.Fprint(w, "}\n")
	wbytes := w.Bytes()

	b, err := format.Source(wbytes)
	if err != nil {
		os.Stderr.Write(wbytes)
		log.Fatalf("error formatting: %v", err)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(b); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
