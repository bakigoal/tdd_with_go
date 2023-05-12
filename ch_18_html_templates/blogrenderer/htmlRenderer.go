package blogrenderer

import "io"

func Render(writer io.Writer, post Post) error {
	html := `<h1>Hello world</h1>`
	writer.Write([]byte(html))
	return nil
}
