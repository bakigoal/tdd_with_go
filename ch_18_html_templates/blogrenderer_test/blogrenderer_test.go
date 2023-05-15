package blogrenderer_test

import (
	"bytes"
	"github.com/bakigoal/tdd_with_go/ch_17_reading_files/blogposts"
	"github.com/bakigoal/tdd_with_go/ch_18_html_templates/blogrenderer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogposts.Post{
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
		want := `<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`
		got := buf.String()
		assert.Equal(t, want, got)
	})

}