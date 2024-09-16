// File: user_signup.go
// Creation: Fri Sep  6 08:32:12 2024
// Time-stamp: <2024-09-16 19:01:53>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"
	"net/mail"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func ValidateUserForCreation(u models.User) error {
	if len(u.Username) < 3 {
		return fmt.Errorf("Invalid username. It must be a valid string with 3 characters minimum")
	}

	if len(u.Password) < 8 {
		return fmt.Errorf("Invalid password. It must be a valid string with 8 characters minimum")
	}

	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return fmt.Errorf("Invalid email")
	}

	return nil
}
