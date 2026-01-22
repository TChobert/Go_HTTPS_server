package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
)

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go!\n")
}

func main() {

	server := http.Server{
		Addr: ":8080",
		Handler: &myHandler{},
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	certPath := os.Getenv("TLS_CERT_PATH")
	keyPath := os.Getenv("TLS_KEY_PATH")
	log.Fatal(server.ListenAndServeTLS(certPath, keyPath))
}
