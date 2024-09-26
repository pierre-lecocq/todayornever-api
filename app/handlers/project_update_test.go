// File: project_update_test.go
// Creation: Thu Sep 26 14:53:01 2024
// Time-stamp: <2024-09-26 14:53:39>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func TestProjectUpdateHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(ProjectUpdateHandler(db)),
			Method:       "PATCH",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			Body:         models.Project{Name: "Edited"},
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
		{
			TestName:     "Invalid UserID value in context",
			Handler:      http.HandlerFunc(ProjectUpdateHandler(db)),
			Method:       "PATCH",
			Path:         "/projects/1",
			URLVars:      map[string]string{"id": "1"},
			Body:         models.Project{Name: "Edited"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
		{
			TestName:     "Invalid ID parameter in URL",
			Handler:      http.HandlerFunc(ProjectUpdateHandler(db)),
			Method:       "PATCH",
			Path:         "/projects/abc",
			URLVars:      map[string]string{"id": "abc"},
			Body:         models.Project{Name: "Edited"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
