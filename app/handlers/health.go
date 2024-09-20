// File: health.go
// Creation: Fri Sep 20 11:59:46 2024
// Time-stamp: <2024-09-20 12:00:19>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/pkg/response"
)

func HealthHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendJSON(w, http.StatusOK, map[string]interface{}{
			"healthy": true,
		})
	}
}
