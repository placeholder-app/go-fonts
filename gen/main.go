package main

import "github.com/placeholder-app/go-fonts/util/gen"

func main() {
	for _, font := range fonts {
		gen.DownloadFont(font, "deflated")
	}
}
