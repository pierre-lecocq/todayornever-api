// File: index.go
// Creation: Thu Sep  5 09:37:40 2024
// Time-stamp: <2024-09-16 18:59:07>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/pkg/response"
)

func IndexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendJSON(w, http.StatusOK, map[string]interface{}{
			"service":     "Today or Never!",
			"description": "Focus on things you can do now!",
		})
	}
}
