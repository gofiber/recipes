package template

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var content embed.FS

// Dist returns the SvelteKit build with dist/ as its root.
func Dist() fs.FS {
	sub, err := fs.Sub(content, "dist")
	if err != nil {
		panic(err)
	}
	return sub
}
