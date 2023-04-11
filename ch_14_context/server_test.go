package ch_14_context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, data, response.Body.String())
		assert.Equal(t, false, store.cancelled)
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		srv := Server(store)

		request := requestCancelledAfter(5 * time.Millisecond)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, true, store.cancelled)
	})
}

func requestCancelledAfter(duration time.Duration) *http.Request {
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	cancellingCtx, cancel := context.WithCancel(request.Context())

	time.AfterFunc(duration, cancel)

	return request.WithContext(cancellingCtx)
}
