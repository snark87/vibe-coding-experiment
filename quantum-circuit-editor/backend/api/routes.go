package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all API routes
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

// healthCheckHandler responds with a simple health status
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// notImplementedHandler is a placeholder for endpoints not yet implemented
func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{"error": "Not implemented yet"})
}
