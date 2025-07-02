package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func loggingMiddlewareFactory() func(http.Handler) http.Handler {
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlogger := slog.New(handler)

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mySlogger.Debug("client info", "ip_address", r.RemoteAddr)
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Richiesta ricevuta")
		w.Write([]byte("l'ora di adesso è: " + time.Now().Format(time.RFC3339) + "\n"))
	})

	lmf := loggingMiddlewareFactory()
	middleware := lmf(mux)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      middleware,
	}

	fmt.Println("Server avviato")
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
