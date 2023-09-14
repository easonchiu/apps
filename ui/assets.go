package ui

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"
)

//go:embed static/* templates/*.html
var embedFS embed.FS

// Template for all
var Template *template.Template

func GetTemplate() *template.Template {
	tmp := template.Must(template.New("").ParseFS(embedFS, "templates/*.html"))
	// funcMap := sprig.FuncMap()
	// funcMap["timeUnix"] = time.Unix
	// funcMap["unescapeHTML"] = unescapeHTML

	return tmp
}

// resource is an interface that provides static file
type resource struct {
	prefix string
	fs     embed.FS
}

// Open to implement the interface by http.FS required
func (r *resource) Open(name string) (fs.File, error) {
	name = path.Join(r.prefix, name)
	return r.fs.Open(name)
}

// StaticFS static http file system
func StaticFS() http.FileSystem {
	return http.FS(&resource{prefix: "static", fs: embedFS})
}
