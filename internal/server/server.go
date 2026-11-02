package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

func Shutdown(ctx context.Context, srv *http.Server) error {
	err := srv.Shutdown(ctx)
	if err != nil {
		return err
	}
	log.Println("Server stopped")
	return nil
}

func RunTLS(srv *http.Server, certPath string, keyPath string) error {
	return srv.ListenAndServeTLS(certPath, keyPath)
}

func NewServer(router http.Handler) *http.Server {

	return &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
}
