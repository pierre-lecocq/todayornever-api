// File: task_fetch_test.go
// Creation: Mon Sep  9 09:30:39 2024
// Time-stamp: <2024-09-26 14:49:31>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestTaskFetchHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(TaskFetchHandler(db)),
			Method:       "GET",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
		{
			TestName:     "Invalid UserID value in context",
			Handler:      http.HandlerFunc(TaskFetchHandler(db)),
			Method:       "GET",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
		{
			TestName:     "Invalid ID parameter in URL",
			Handler:      http.HandlerFunc(TaskFetchHandler(db)),
			Method:       "GET",
			Path:         "/tasks/abc",
			URLVars:      map[string]string{"id": "abc"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       1,
		},
		{
			TestName:     "Task Not Found",
			Handler:      http.HandlerFunc(TaskFetchHandler(db)),
			Method:       "GET",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusNotFound,
			UserID:       2,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
