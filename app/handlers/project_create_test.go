// File: project_create_test.go
// Creation: Thu Sep 26 14:47:37 2024
// Time-stamp: <2024-09-26 14:50:56>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func TestProjectCreateHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(ProjectCreateHandler(db)),
			Method:       "POST",
			Path:         "/projects",
			ExpectedCode: http.StatusCreated,
			Body:         models.Project{Name: "Created"},
			UserID:       1,
		},
		{
			TestName:     "Invalid UserID value in context",
			Handler:      http.HandlerFunc(ProjectCreateHandler(db)),
			Method:       "POST",
			Path:         "/projects",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.Project{Name: "Created"},
			UserID:       0,
		},
		{
			TestName:     "Validation error",
			Handler:      http.HandlerFunc(ProjectCreateHandler(db)),
			Method:       "POST",
			Path:         "/projects",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.Project{Name: "a"},
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
