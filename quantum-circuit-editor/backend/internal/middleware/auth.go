// Package middleware provides HTTP middleware functions for the Quantum Circuit Editor backend,
// including authentication, logging, and other request processing functionality.
package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"
)

// User represents the authenticated user information
type User struct {
	ID    string
	Email string
}

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

// UserKey is the context key for storing user information
const UserKey contextKey = "user"

var (
	// ErrNoAuthHeader is returned when no Authorization header is present in the request
	ErrNoAuthHeader = errors.New("no authorization header provided")
	// ErrInvalidAuthFormat is returned when the Authorization header format is incorrect
	ErrInvalidAuthFormat = errors.New("invalid authorization header format")
	// ErrInvalidToken is returned when the provided authentication token is invalid
	ErrInvalidToken = errors.New("invalid authentication token")
)

// Authenticate validates JWT tokens from the Authorization header
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handleAuthError(w, r, ErrNoAuthHeader, http.StatusUnauthorized)
			return
		}

		// Check Authorization header format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			handleAuthError(w, r, ErrInvalidAuthFormat, http.StatusUnauthorized)
			return
		}

		token := parts[1]

		// TODO: Implement actual token validation with JWT
		user, err := validateToken(token)
		if err != nil {
			handleAuthError(w, r, err, http.StatusUnauthorized)
			return
		}

		// Add user to request context
		ctx := context.WithValue(r.Context(), UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext extracts the authenticated user from the request context
func GetUserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(UserKey).(*User)
	if !ok {
		return nil
	}
	return user
}

// validateToken is a placeholder for actual JWT token validation logic
// TODO: Implement proper JWT validation with a library like golang-jwt
func validateToken(token string) (*User, error) {
	// For development purposes only - REMOVE in production
	if token == "test-token" {
		return &User{
			ID:    "test-user-id",
			Email: "test@example.com",
		}, nil
	}

	return nil, ErrInvalidToken
}

// handleAuthError responds with the appropriate error status and logs the failure
func handleAuthError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	slog.Info("Authentication failed",
		"error", err.Error(),
		"path", r.URL.Path,
		"method", r.Method,
		"remote_addr", r.RemoteAddr,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// Fix: Check error when writing the response
	_, writeErr := w.Write([]byte(`{"error": "Authentication required"}`))
	if writeErr != nil {
		slog.Error("Failed to write error response", "error", writeErr)
	}
}
