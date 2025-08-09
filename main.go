package main

import (
	"embed"
	"fmt"
	"github.com/Meschkov/htmx-playground/pkg/handlers"
	"github.com/Meschkov/htmx-playground/pkg/middleware"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"time"
)

//go:generate npm run build

//go:embed web/static/*
var static embed.FS

func main() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slog.SetDefault(slog.New(textHandler))

	port := "8080"
	addr := ":" + port
	mux := http.NewServeMux()
	chain := &middleware.Chain{}
	chain.Use(middleware.RecoverMiddleware)
	chain.Use(middleware.LogMiddleware)
	wrappedMux := chain.Then(mux)

	fs, err := fs.Sub(static, "web/static")
	if err != nil {
		slog.Error("Failed to create sub filesystem", "error", err)
		return
	}

	// Serve files from the embedded /web/static directory at /static
	fileServer := http.FileServer(http.FS(fs))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		data, err := static.ReadFile("web/static/img/favicon.ico")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "image/x-icon")
		_, err = w.Write(data)
		if err != nil {
			slog.Error("Failed to write on 'GET /favicon.ico'", "error", err)
			return
		}
	})

	mux.HandleFunc("GET /robots.txt", func(w http.ResponseWriter, r *http.Request) {
		data, err := static.ReadFile("web/static/robots.txt")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		_, err = w.Write(data)
		if err != nil {
			slog.Error("Failed to write on 'GET /robots.txt'", "error", err)
			return
		}
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte(`OK`))
		if err != nil {
			slog.Error("Failed to write on 'GET /health'", "error", err)
			return
		}
	})

	mux.HandleFunc("GET /", handlers.RootHandler())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: wrappedMux,
		// Recommended timeouts from
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info("Server listening", "addr", addr)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
