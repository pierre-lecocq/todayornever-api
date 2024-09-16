// File: task_update.go
// Creation: Fri Sep  6 16:57:11 2024
// Time-stamp: <2024-09-16 19:01:22>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	// "fmt"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func ValidateTaskForUpdate(t models.Task) error {
	// if len(t.Title) <= 3 {
	// 	return fmt.Errorf("Invalid title. It must be a valid string with 3 characters minimum")
	// }

	return nil
}
