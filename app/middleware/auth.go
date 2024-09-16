// File: auth.go
// Creation: Thu Sep  5 09:43:49 2024
// Time-stamp: <2024-09-16 19:00:45>
// Copyright (C): 2024 Pierre Lecocq

package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/pierre-lecocq/todayornever-api/pkg/auth"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")

		if len(h) == 0 {
			response.SendJSON(w, http.StatusUnauthorized, map[string]interface{}{
				"error": "Unauthorized",
			})
			return
		}

		chunks := strings.Split(h, " ")

		if len(chunks) != 2 || chunks[0] != "Bearer" || chunks[1] == "" {
			response.SendJSON(w, http.StatusUnauthorized, map[string]interface{}{
				"error": "Invalid Authorization header value",
			})
			return
		}

		if len(chunks) != 2 || chunks[0] != "Bearer" {
			response.SendJSON(w, http.StatusUnauthorized, map[string]interface{}{
				"error": "Invalid Authorization header value",
			})
			return
		}

		claimedUserID, err := auth.ValidateAndDecodeJWTToken(chunks[1], os.Getenv("AUTH_SECRET"))

		if err != nil {
			response.SendJSON(w, http.StatusUnauthorized, map[string]interface{}{
				"error": "Unauthorized",
			})
			return
		}

		ctx := context.WithValue(r.Context(), "UserID", claimedUserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
