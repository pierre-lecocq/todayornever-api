// File: task_create.go
// Creation: Fri Sep  6 15:32:28 2024
// Time-stamp: <2024-09-16 19:01:09>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func ValidateTaskForCreation(t models.Task) error {
	if len(t.Title) < 3 {
		return fmt.Errorf("Invalid title. It must be a valid string with 3 characters minimum")
	}

	return nil
}
