package ch_14_context

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(writer, d)
		case <-ctx.Done(): // “a signal when the context is "done" or "cancelled”
			store.Cancel()
		}
	}
}

type Store interface {
	Fetch() string
	Cancel()
}
