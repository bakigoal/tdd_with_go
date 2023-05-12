package blogrenderer

import (
	"github.com/bakigoal/tdd_with_go/ch_17_reading_files/blogposts"
	"io"
)

func Render(writer io.Writer, post blogposts.Post) error {
	html := `<h1>Hello world</h1>`
	writer.Write([]byte(html))
	return nil
}
