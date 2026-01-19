package template

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var content embed.FS

// All returns the content of the all directory.
func Dist() fs.FS {
	return content
}
