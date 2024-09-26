// File: project_delete_test.go
// Creation: Thu Sep 26 14:51:42 2024
// Time-stamp: <2024-09-26 14:52:13>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestProjectDeleteHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(ProjectDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusNoContent,
			UserID:       1,
		},
		{
			TestName:     "Invalid UserID value in context",
			Handler:      http.HandlerFunc(ProjectDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
		{
			TestName:     "Invalid ID parameter in URL",
			Handler:      http.HandlerFunc(ProjectDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/projects/abc",
			URLVars:      map[string]string{"id": "abc"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       1,
		},
		{
			TestName:     "Can not delete project",
			Handler:      http.HandlerFunc(ProjectDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       2,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
