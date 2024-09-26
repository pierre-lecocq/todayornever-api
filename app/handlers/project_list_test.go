// File: project_list_test.go
// Creation: Wed Sep 25 14:51:49 2024
// Time-stamp: <2024-09-26 14:49:07>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"
)

func TestProjectListHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(ProjectListHandler(db)),
			Method:       "GET",
			Path:         "/projects",
			ExpectedCode: http.StatusOK,
			UserID:       1,
		},
		{
			TestName:     "Invalid UserID value in context",
			Handler:      http.HandlerFunc(ProjectListHandler(db)),
			Method:       "GET",
			Path:         "/projects",
			ExpectedCode: http.StatusBadRequest,
			UserID:       0,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
