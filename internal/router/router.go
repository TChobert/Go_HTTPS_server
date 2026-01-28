package router

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/TChobert/Go_HTTPS_server/internal/handler"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/", handler.HelloHandler)
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from test location!\n")
	})
	return r
}
