// File: health_test.go
// Creation: Thu Sep 26 14:54:38 2024
// Time-stamp: <2024-09-26 14:55:03>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(HealthHandler()),
			Method:       "GET",
			Path:         "/health",
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
