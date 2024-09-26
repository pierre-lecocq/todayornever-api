// File: project_fetch_test.go
// Creation: Thu Sep 26 14:52:34 2024
// Time-stamp: <2024-09-26 14:52:46>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestProjectFetchHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(ProjectFetchHandler(db)),
			Method:       "GET",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
		{
			TestName:     "Invalid UserID value in context",
			Handler:      http.HandlerFunc(ProjectFetchHandler(db)),
			Method:       "GET",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
		{
			TestName:     "Invalid ID parameter in URL",
			Handler:      http.HandlerFunc(ProjectFetchHandler(db)),
			Method:       "GET",
			Path:         "/projects/abc",
			URLVars:      map[string]string{"id": "abc"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       1,
		},
		{
			TestName:     "Project Not Found",
			Handler:      http.HandlerFunc(ProjectFetchHandler(db)),
			Method:       "GET",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusNotFound,
			UserID:       2,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
