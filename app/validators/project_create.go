// File: project_create.go
// Creation: Thu Sep 26 14:38:27 2024
// Time-stamp: <2024-09-26 14:38:53>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"

	"github.com/pierre-lecocq/todayornever-api/app/models"
)

func ValidateProjectForCreation(p models.Project) error {
	if len(p.Name) < 3 {
		return fmt.Errorf("Invalid name. It must be a valid string with 3 characters minimum")
	}

	return nil
}
