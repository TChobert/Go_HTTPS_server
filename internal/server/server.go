package server

import (
	"net/http"
	"time"
	"fmt"
)

func Shutdown() {
	fmt.Println("Server stopped")
}

func RunTLS(srv *http.Server, certPath string, keyPath string) error {
	return srv.ListenAndServeTLS(certPath, keyPath)
}

func NewServer(router http.Handler) *http.Server {

	return &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
}
