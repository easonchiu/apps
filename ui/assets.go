package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var embedFS embed.FS

// GetDistFS returns the embedded dist file system for serving React build output
func GetDistFS() http.FileSystem {
	distFS, err := fs.Sub(embedFS, "dist")
	if err != nil {
		panic(err)
	}
	return http.FS(distFS)
}
