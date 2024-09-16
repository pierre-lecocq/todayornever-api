// File: user_signup_test.go
// Creation: Mon Sep  9 09:30:58 2024
// Time-stamp: <2024-09-16 19:00:36>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"net/http"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func TestUserSignupHandler(t *testing.T) {
	db, _ := InitTestDatabase()
	defer db.Close()

	data := []DataProvider{
		{
			Name:         "Success",
			Handler:      http.HandlerFunc(UserSignupHandler(db)),
			Method:       "POST",
			Path:         "/signup",
			ExpectedCode: http.StatusCreated,
			Body:         models.User{Email: "user3@mail.com", Username: "user3", Password: "user3user3user3"},
			UserID:       1,
		},
		{
			Name:         "Duplicate",
			Handler:      http.HandlerFunc(UserSignupHandler(db)),
			Method:       "POST",
			Path:         "/signup",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.User{Email: "user3@mail.com", Username: "user3", Password: "user3user3user3"},
			UserID:       1,
		},
		{
			Name:         "Invalid email",
			Handler:      http.HandlerFunc(UserSignupHandler(db)),
			Method:       "POST",
			Path:         "/signup",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.User{Email: "user", Username: "user3", Password: "user3user3user3"},
			UserID:       1,
		},
		{
			Name:         "Invalid username",
			Handler:      http.HandlerFunc(UserSignupHandler(db)),
			Method:       "POST",
			Path:         "/signup",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.User{Email: "user3@mail.com", Username: "u", Password: "user3user3user3"},
			UserID:       1,
		},
		{
			Name:         "Invalid password",
			Handler:      http.HandlerFunc(UserSignupHandler(db)),
			Method:       "POST",
			Path:         "/signup",
			ExpectedCode: http.StatusBadRequest,
			Body:         models.User{Email: "user3@mail.com", Username: "user3", Password: "u"},
			UserID:       1,
		},
	}

	for _, dp := range data {
		RequestTest(t, dp)
	}
}
