package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
	"github.com/go-chi/chi/v5" //mention
)

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go server with chi!\n")
	})
	server := http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	certPath := os.Getenv("TLS_CERT_PATH")
	keyPath := os.Getenv("TLS_KEY_PATH")
	log.Fatal(server.ListenAndServeTLS(certPath, keyPath)) //mention
}
