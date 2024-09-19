// File: negociate.go
// Creation: Thu Sep  5 09:27:52 2024
// Time-stamp: <2024-09-19 11:32:09>
// Copyright (C): 2024 Pierre Lecocq

package middleware

import (
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/pkg/response"
)

func Negociate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")

		if ct != "" && ct != "application/json" {
			response.SendJSON(w, http.StatusBadRequest, map[string]interface{}{
				"error": "Wrong ContentType header value",
			})
			return
		}

		// a := r.Header.Get("Accept")

		// if a != "*/*" && a != "application/json" {
		// 	response.SendJSON(w, http.StatusBadRequest, map[string]interface{}{
		// 		"error": "Wrong Accept header value",
		// 	})
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
