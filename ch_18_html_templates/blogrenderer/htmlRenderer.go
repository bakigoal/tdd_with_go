package blogrenderer

import (
	"fmt"
	"github.com/bakigoal/tdd_with_go/ch_17_reading_files/blogposts"
	"io"
)

func Render(writer io.Writer, post blogposts.Post) error {
	htmlTemplate := `<h1>%s</h1>`
	_, err := fmt.Fprintf(writer, htmlTemplate, post.Title)
	return err
}
