package main

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

func NewTimeoutMiddleware(ms time.Duration) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler { // middleware
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx, cancelFunc := context.WithTimeout(ctx, ms*time.Millisecond)
			defer cancelFunc()
			req = req.WithContext(ctx)

			handler.ServeHTTP(rw, req)
		})
	}
}

func main() {
	middleware := NewTimeoutMiddleware(100)
	server := http.Server{
		Handler: middleware(http.HandlerFunc(sleepy)),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}

func sleepy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message, err := doThing(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(message))
}

func doThing(ctx context.Context) (string, error) {
	wait := rand.Intn(200)
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return "Done!", nil
	case <-ctx.Done():
		return "Too slow!", ctx.Err()
	}
}
