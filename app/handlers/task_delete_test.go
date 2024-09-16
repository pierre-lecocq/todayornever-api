// File: task_delete_test.go
// Creation: Mon Sep  9 09:30:34 2024
// Time-stamp: <2024-09-15 00:07:22>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestTaskDeleteHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			Name:         "Success",
			Handler:      http.HandlerFunc(TaskDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusNoContent,
			UserID:       1,
		},
		{
			Name:         "Invalid UserID value in context",
			Handler:      http.HandlerFunc(TaskDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
		{
			Name:         "Invalid ID parameter in URL",
			Handler:      http.HandlerFunc(TaskDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/tasks/abc",
			URLVars:      map[string]string{"id": "abc"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       1,
		},
		{
			Name:         "Can not delete task",
			Handler:      http.HandlerFunc(TaskDeleteHandler(db)),
			Method:       "DELETE",
			Path:         "/tasks/1",
			URLVars:      map[string]string{"id": "1"},
			ExpectedCode: http.StatusBadRequest,
			UserID:       2,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
