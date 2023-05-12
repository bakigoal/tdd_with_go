package blogrenderer_test

import (
	"bytes"
	"github.com/bakigoal/tdd_with_go/ch_18_html_templates/blogrenderer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		assert.NoError(t, err)
		want := `<h1>Hello world</h1>`
		got := buf.String()
		assert.Equal(t, want, got)
	})

}
