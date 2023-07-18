package template

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var content embed.FS

// All returns the content of the all directory.
func Dist() http.FileSystem {
	dist, err := fs.Sub(content, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(dist)
}
