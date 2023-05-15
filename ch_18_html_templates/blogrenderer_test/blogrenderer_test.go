package blogrenderer_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"github.com/bakigoal/tdd_with_go/ch_17_reading_files/blogposts"
	"github.com/bakigoal/tdd_with_go/ch_18_html_templates/blogrenderer"
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

		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

}
