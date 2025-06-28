package middleware

import (
	"context"
	"net/http"
	"hospital/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Read the token from the header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// 2. Validate the token using jwt.go function
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// 3. Store claims in context for later use
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "role", claims.Role)

		// 4. Call the actual handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
