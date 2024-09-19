// File: index.go
// Creation: Thu Sep  5 09:37:40 2024
// Time-stamp: <2024-09-19 11:16:08>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/pkg/response"
)

func IndexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendJSON(w, http.StatusOK, map[string]interface{}{
			"service":     "todayornever-api",
			"version":     "1.0.0",
			"description": "Focus on things you can do now!",
		})
	}
}
