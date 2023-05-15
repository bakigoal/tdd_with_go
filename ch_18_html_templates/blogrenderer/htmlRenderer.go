package blogrenderer

import (
	"embed"
	"github.com/bakigoal/tdd_with_go/ch_17_reading_files/blogposts"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	err = templ.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}
