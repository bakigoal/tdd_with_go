package ch_10_concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "www.wrong.site" {
		return false
	}
	return true
}

func TestWebsites(t *testing.T) {
	websites := []string{
		"www.good.com",
		"www.blog.ru",
		"www.wrong.site",
	}
	want := map[string]bool{
		"www.good.com":   true,
		"www.blog.ru":    true,
		"www.wrong.site": false,
	}

	got := CheckWebsite(mockWebsiteChecker, websites)

	assert.Equal(t, want, got)
}
