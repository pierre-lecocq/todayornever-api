// File: user_signup_test.go
// Creation: Mon Sep  9 09:32:35 2024
// Time-stamp: <2024-09-16 19:01:59>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserForCreation(t *testing.T) {
	type DataProvider struct {
		name string
		user models.User
		err  error
	}

	dp := []DataProvider{
		{
			name: "OK",
			user: models.User{Username: "example", Password: "Password", Email: "example@mail.com"},
			err:  nil,
		},
		{
			name: "KO - Invalid username",
			user: models.User{Username: "me", Password: "Password", Email: "example@mail.com"},
			err:  fmt.Errorf("Invalid username. It must be a valid string with 3 characters minimum"),
		},
		{
			name: "KO - Invalid password",
			user: models.User{Username: "example", Password: "pa", Email: "example@mail.com"},
			err:  fmt.Errorf("Invalid password. It must be a valid string with 8 characters minimum"),
		},
		{
			name: "KO - Invalid email",
			user: models.User{Username: "example", Password: "Password", Email: "example"},
			err:  fmt.Errorf("Invalid email"),
		},
	}

	for _, d := range dp {
		t.Log(d.name)

		res := ValidateUserForCreation(d.user)

		if d.err != nil {
			assert.Error(t, d.err, res)
		} else {
			assert.Nil(t, res)
		}
	}
}
