package ch_11_select

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var slowServer = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(20 * time.Millisecond)
	writer.WriteHeader(http.StatusOK)
}))

var fastServer = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}))

func TestRacer(t *testing.T) {
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)
	assert.Equal(t, want, got)
}
