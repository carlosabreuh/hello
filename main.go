package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nytimes/dv-interview-exercise/hello/handler"

	"github.com/namsral/flag"
)

var (
	// version is filled by -ldflags  at compile time
	version        = "no version set"
	displayVersion = flag.Bool("v", false, "Show version and quit")
	listenPort     = flag.String("port", "8080", "the port to listen on, default to 8080")
)

func main() {
	// parse cmd flags and env vars
	flag.Parse()

	if *displayVersion {
		fmt.Printf("Hello version %s\n", version)
		os.Exit(0)
	}

	mux := http.NewServeMux()

	s := handler.NewServer()
	mux.HandleFunc("/", handler.OKHandler)
	mux.HandleFunc("/hello/", s.HelloHandler)
	mux.HandleFunc("/counts", s.CountsHandler)
	mux.HandleFunc("/healthz", handler.HealthHandler(version))

	server := &http.Server{Addr: ":" + *listenPort, Handler: mux}

	// Graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		fmt.Printf("Server listening on %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("http server error:", err)
		}
	}()

	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("error shutting down http server:", err)
	}
}
