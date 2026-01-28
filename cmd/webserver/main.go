package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TChobert/Go_HTTPS_server/internal/router"
	"github.com/TChobert/Go_HTTPS_server/internal/server"
)

func main() {

	r := router.NewRouter()
	srv := server.NewServer(r)
	signalCtx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop() // this function is defered -> it will be called when main returns

	certPath := os.Getenv("TLS_CERT_PATH")
	keyPath := os.Getenv("TLS_KEY_PATH")
	go func() {
		if err := server.RunTLS(srv, certPath, keyPath); err != nil {
			log.Fatal(err)
		}
	}()
	<-signalCtx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	server.Shutdown(shutdownCtx, srv)
}
