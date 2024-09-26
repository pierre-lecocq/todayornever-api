// File: user_login_test.go
// Creation: Mon Sep  9 09:30:51 2024
// Time-stamp: <2024-09-26 14:49:48>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func TestUserLoginHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			TestName:     "Success",
			Handler:      http.HandlerFunc(UserLoginHandler(db)),
			Method:       "POST",
			Path:         "/login",
			ExpectedCode: http.StatusOK,
			Body:         models.User{Email: "user1@mail.com", Password: "user1user1user1"},
			UserID:       1,
		},
		{
			TestName:     "User Not Found - Username",
			Handler:      http.HandlerFunc(UserLoginHandler(db)),
			Method:       "POST",
			Path:         "/login",
			ExpectedCode: http.StatusNotFound,
			Body:         models.User{Email: "wrong@mail.com", Password: "user1user1user1"},
			UserID:       1,
		},
		{
			TestName:     "User Not Found - Password",
			Handler:      http.HandlerFunc(UserLoginHandler(db)),
			Method:       "POST",
			Path:         "/login",
			ExpectedCode: http.StatusNotFound,
			Body:         models.User{Email: "user1@mail.com", Password: "wrong"},
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
