package blogposts_test

import (
	"errors"
	"github.com/bakigoal/blogposts"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestFailingStub(t *testing.T) {
	_, err := blogposts.NewPostsFromFS(StubFailingFS{})
	assert.Error(t, err)
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	assert.NoError(t, err)
	assert.Equal(t, len(posts), len(fs))

	assert.Equal(t, blogposts.Post{Title: "Post 1"}, posts[0])
	assert.Equal(t, blogposts.Post{Title: "Post 2"}, posts[1])
}
