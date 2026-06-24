package middleware

import (
	"context"
	"net/http"
	"strings"

	"vue-go-backend/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		// Check if it's a Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid authorization format. Use Bearer token")
			return
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), utils.UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, utils.UsernameKey, claims.Username)
		ctx = context.WithValue(ctx, utils.UserRoleKey, claims.Role)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
