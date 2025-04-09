// Package main is the entry point for the Quantum Circuit Editor backend service.
// It sets up the HTTP server, configures middleware, and manages the lifecycle
// of the application.
package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/snark87/quantum-circuit-editor/backend/api"
	"github.com/snark87/quantum-circuit-editor/backend/internal/middleware"
)

// configureServer creates and configures an HTTP server and router
func configureServer(port string) (*http.Server, *mux.Router) {
	r := mux.NewRouter()

	// Apply logging middleware to all requests
	r.Use(middleware.Logger)

	// Setup API routes
	api.RegisterRoutes(r)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // SECURITY: Update this with specific origins in production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      corsMiddleware.Handler(r),
	}

	return srv, r
}

// startServer starts the server and returns a shutdown function
func startServer(srv *http.Server) func() {
	go func() {
		slog.Info("Server starting", "port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		slog.Info("Shutting down server...")
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Server shutdown failed", "error", err)
		} else {
			slog.Info("Server gracefully stopped")
		}
	}
}

func main() {
	// Configure structured logging
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	slog.SetDefault(slog.New(logHandler))

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		slog.Info("No PORT environment variable found, using default", "port", port)
	}

	// Setup and start server
	srv, _ := configureServer(port)
	shutdownFn := startServer(srv)

	// Wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Graceful shutdown
	shutdownFn()
}
