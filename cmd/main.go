package main

import (
	"log"
	"os"
	"github.com/TChobert/Go_HTTPS_server/internal/router"
	"github.com/TChobert/Go_HTTPS_server/internal/server"
)

func main() {

	r := router.NewRouter()
	server := server.NewServer(r)

	certPath := os.Getenv("TLS_CERT_PATH")
	keyPath := os.Getenv("TLS_KEY_PATH")
	log.Fatal(server.ListenAndServeTLS(certPath, keyPath))
}
