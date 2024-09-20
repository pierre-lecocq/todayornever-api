// File: index.go
// Creation: Thu Sep  5 09:37:40 2024
// Time-stamp: <2024-09-20 10:12:54>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/spf13/viper"
)

func IndexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendJSON(w, http.StatusOK, map[string]interface{}{
			"name":        viper.GetString("APP_NAME"),
			"version":     viper.GetString("APP_VERSION"),
			"description": "Focus on things you can do now!",
			"environment": viper.GetString("APP_ENVIRONMENT"),
		})
	}
}
