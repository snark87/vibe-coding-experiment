// Package api implements the REST API endpoints for the Quantum Circuit Editor
// service, including route definitions and handler functions.
package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all API routes for the application
func RegisterRoutes(r *mux.Router) {
	// Health check endpoint
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	// API subrouter with version prefixing
	api := r.PathPrefix("/api/v1").Subrouter()

	// Auth endpoints
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", notImplementedHandler).Methods("POST")
	auth.HandleFunc("/logout", notImplementedHandler).Methods("POST")

	// Circuits endpoints
	circuits := api.PathPrefix("/circuits").Subrouter()
	circuits.HandleFunc("", notImplementedHandler).Methods("GET")
	circuits.HandleFunc("", notImplementedHandler).Methods("POST")
	circuits.HandleFunc("/{id}", notImplementedHandler).Methods("GET")
	circuits.HandleFunc("/{id}", notImplementedHandler).Methods("PUT")
	circuits.HandleFunc("/{id}", notImplementedHandler).Methods("DELETE")

	// Simulation endpoints
	api.HandleFunc("/simulate", notImplementedHandler).Methods("POST")

	// Export endpoints
	export := api.PathPrefix("/export").Subrouter()
	export.HandleFunc("/qasm", notImplementedHandler).Methods("POST")
}

// healthCheckHandler returns a simple status response to confirm API availability
func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// notImplementedHandler responds with 501 status for endpoints not yet implemented
func notImplementedHandler(w http.ResponseWriter, _ *http.Request) {
	respondWithJSON(w, http.StatusNotImplemented, map[string]string{
		"error":   "Not implemented yet",
		"message": "This endpoint is planned but not yet available",
	})
}

// respondWithJSON sends a JSON response with the specified status code
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Failed to encode JSON response",
			"error", err,
			"statusCode", statusCode,
		)
	}
}
