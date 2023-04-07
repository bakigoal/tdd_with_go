package ch_10_concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, urls)
		// blocking: 		2093125417 ns/op
		// non-blocking: 	21251202 ns/op
	}
}
