// File: user_login.go
// Creation: Fri Sep  6 08:32:12 2024
// Time-stamp: <2024-09-16 19:01:39>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"
	"net/mail"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func ValidateUserForLogin(u models.User) error {
	if len(u.Password) == 0 {
		return fmt.Errorf("Invalid password")
	}

	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return fmt.Errorf("Invalid email")
	}

	return nil
}
