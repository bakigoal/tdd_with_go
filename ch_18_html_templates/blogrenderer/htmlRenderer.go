package blogrenderer

import (
	"github.com/bakigoal/tdd_with_go/ch_17_reading_files/blogposts"
	"html/template"
	"io"
)

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	err = templ.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}
