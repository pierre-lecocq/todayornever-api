// File: project_create_test.go
// Creation: Thu Sep 26 14:39:06 2024
// Time-stamp: <2024-09-26 14:40:20>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateProjectForCreation(t *testing.T) {
	type DataProvider struct {
		testName string
		project  models.Project
		err      error
	}

	dp := []DataProvider{
		{
			testName: "OK",
			project:  models.Project{Name: "My project"},
			err:      nil,
		},
		{
			testName: "KO - Invalid name",
			project:  models.Project{Name: "a"},
			err:      fmt.Errorf("Invalid name. It must be a valid string with 3 characters minimum"),
		},
	}

	for _, d := range dp {
		t.Log(d.testName)

		res := ValidateProjectForCreation(d.project)

		if d.err != nil {
			assert.Error(t, d.err, res)
		} else {
			assert.Nil(t, res)
		}
	}
}
