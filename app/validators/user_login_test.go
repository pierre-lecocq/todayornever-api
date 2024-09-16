// File: user_login_test.go
// Creation: Mon Sep  9 09:32:32 2024
// Time-stamp: <2024-09-16 19:01:48>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserForLogin(t *testing.T) {
	type DataProvider struct {
		name string
		user models.User
		err  error
	}

	dp := []DataProvider{
		{
			name: "OK",
			user: models.User{Password: "Password", Email: "example@mail.com"},
			err:  nil,
		},
		{
			name: "KO - Invalid password",
			user: models.User{Password: "pa", Email: "example@mail.com"},
			err:  fmt.Errorf("Invalid password"),
		},
		{
			name: "KO - Invalid email",
			user: models.User{Password: "Password", Email: "example"},
			err:  fmt.Errorf("Invalid email"),
		},
	}

	for _, d := range dp {
		t.Log(d.name)

		res := ValidateUserForLogin(d.user)

		if d.err != nil {
			assert.Error(t, d.err, res)
		} else {
			assert.Nil(t, res)
		}
	}
}
