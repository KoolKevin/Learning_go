package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Mytime struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func buildText(now time.Time) string {
	return now.Format(time.RFC3339)
}

// === stringify
func buildJSON(now time.Time) string {
	timeOut := Mytime{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}
	out, _ := json.Marshal(timeOut)
	return string(out)
}

func newLoggingMiddlewareFactory() func(http.Handler) http.Handler {
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
		fmt.Println("richiesta ricevuta")
		now := time.Now()
		var out string

		if r.Header.Get("Accept") == "application/json" {
			out = buildJSON(now)
		} else {
			out = buildText(now)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("l'ora di adesso è: " + out + "\n"))
	})

	lmf := newLoggingMiddlewareFactory()
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
