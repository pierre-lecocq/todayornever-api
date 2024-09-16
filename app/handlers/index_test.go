// File: index_test.go
// Creation: Mon Sep  9 09:30:22 2024
// Time-stamp: <2024-09-15 00:07:00>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	data := []DataProvider{
		{
			Name:         "Success",
			Handler:      http.HandlerFunc(IndexHandler()),
			Method:       "GET",
			Path:         "/",
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
