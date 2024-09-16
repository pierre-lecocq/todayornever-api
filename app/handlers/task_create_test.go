// File: task_create_test.go
// Creation: Mon Sep  9 09:30:29 2024
// Time-stamp: <2024-09-16 18:59:26>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func TestTaskCreateHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			Name:         "Success",
			Handler:      http.HandlerFunc(TaskCreateHandler(db)),
			Method:       "POST",
			Path:         "/tasks",
			ExpectedCode: http.StatusCreated,
			Body:         models.Task{Title: "Created"},
			UserID:       1,
		},
		{
			Name:         "Invalid UserID value in context",
			Handler:      http.HandlerFunc(TaskCreateHandler(db)),
			Method:       "POST",
			Path:         "/tasks",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.Task{Title: "Created"},
			UserID:       0,
		},
		{
			Name:         "Validation error",
			Handler:      http.HandlerFunc(TaskCreateHandler(db)),
			Method:       "POST",
			Path:         "/tasks",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.Task{State: "created"},
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
