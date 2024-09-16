// File: task_create_test.go
// Creation: Mon Sep  9 09:32:23 2024
// Time-stamp: <2024-09-16 19:01:15>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"fmt"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateTaskForCreation(t *testing.T) {
	type DataProvider struct {
		name string
		task models.Task
		err  error
	}

	dp := []DataProvider{
		{
			name: "OK",
			task: models.Task{Title: "My task"},
			err:  nil,
		},
		{
			name: "KO - Invalid title",
			task: models.Task{Title: "a"},
			err:  fmt.Errorf("Invalid title. It must be a valid string with 3 characters minimum"),
		},
	}

	for _, d := range dp {
		t.Log(d.name)

		res := ValidateTaskForCreation(d.task)

		if d.err != nil {
			assert.Error(t, d.err, res)
		} else {
			assert.Nil(t, res)
		}
	}
}
