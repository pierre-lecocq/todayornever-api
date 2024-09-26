// File: project_update_test.go
// Creation: Thu Sep 26 14:44:03 2024
// Time-stamp: <2024-09-26 14:44:50>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateProjectForUpdate(t *testing.T) {
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
	}

	for _, d := range dp {
		t.Log(d.testName)

		res := ValidateProjectForUpdate(d.project)

		if d.err != nil {
			assert.Error(t, d.err, res)
		} else {
			assert.Nil(t, res)
		}
	}
}
