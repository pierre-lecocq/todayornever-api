// File: task_list_test.go
// Creation: Mon Sep  9 09:30:43 2024
// Time-stamp: <2024-09-15 00:07:44>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestTaskListHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			Name:         "Success",
			Handler:      http.HandlerFunc(TaskListHandler(db)),
			Method:       "GET",
			Path:         "/tasks",
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
		{
			Name:         "Invalid UserID value in context",
			Handler:      http.HandlerFunc(TaskListHandler(db)),
			Method:       "GET",
			Path:         "/tasks",
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
		{
			Name:         "Invalid page query parameter",
			Handler:      http.HandlerFunc(TaskListHandler(db)),
			Method:       "GET",
			Path:         "/tasks?page=abc",
			ExpectedCode: http.StatusBadRequest,
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
