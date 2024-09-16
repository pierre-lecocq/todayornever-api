// File: task_update_test.go
// Creation: Mon Sep  9 09:30:47 2024
// Time-stamp: <2024-09-16 18:59:59>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func TestTaskUpdateHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			Name:         "Success",
			Handler:      http.HandlerFunc(TaskUpdateHandler(db)),
			Method:       "PATCH",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			Body:         models.Task{Title: "Edited"},
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
		{
			Name:         "Invalid UserID value in context",
			Handler:      http.HandlerFunc(TaskUpdateHandler(db)),
			Method:       "PATCH",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			Body:         models.Task{Title: "Edited"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
		{
			Name:         "Invalid ID parameter in URL",
			Handler:      http.HandlerFunc(TaskUpdateHandler(db)),
			Method:       "PATCH",
			Path:         "/tasks/abc",
			URLVars:      map[string]string{"id": "abc"},
			Body:         models.Task{Title: "Edited"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
