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
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	assert.NoError(t, err)
	assert.Equal(t, len(posts), len(fs))
}
