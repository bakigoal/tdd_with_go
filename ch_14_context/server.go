package ch_14_context

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())
		if err != nil {
			log.Println("Error in fetching ...")
			return
		}
		fmt.Fprint(writer, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
