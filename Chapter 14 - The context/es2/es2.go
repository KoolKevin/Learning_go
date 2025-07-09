package main

import (
	"context"
	"fmt"
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
	ctx, cancelFunc := context.WithCancelCause(ctx)
	defer cancelFunc(nil)

	compute(ctx)
	w.Write([]byte("finito"))
}

func compute(ctx context.Context) {
	sum := 0
	iterations := 0

	for {
		n := rand.Intn(10_000_000)
		sum += n

		if n == 1024 {
			fmt.Println("n:", n, "somma:", sum, "iterazioni:", iterations, "causa terminazione:", "trovato 1234!")
			return
		}

		if err := context.Cause(ctx); err != nil {
			fmt.Println("n:", n, "somma:", sum, "iterazioni:", iterations, "causa terminazione:", err)
			return
		}

		iterations++
	}
}
