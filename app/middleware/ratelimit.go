// File: ratelimit.go
// Creation: Fri Sep  6 23:43:39 2024
// Time-stamp: <2024-09-16 19:00:57>
// Copyright (C): 2024 Pierre Lecocq

package middleware

import (
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"golang.org/x/time/rate"
)

func Ratelimit(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(100, 200)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			response.SendJSON(w, http.StatusTooManyRequests, map[string]interface{}{
				"error": "Request Failed",
			})

			return
		}

		next.ServeHTTP(w, r)
	})
}
