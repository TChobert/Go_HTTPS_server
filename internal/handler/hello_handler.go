package handler

import (
	"fmt"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintln(w, "Hello!")
	case <-ctx.Done():
		http.Error(w, "Request canceled", http.StatusRequestTimeout)
	}
}
