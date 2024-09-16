// File: task_update_test.go
// Creation: Mon Sep  9 09:32:28 2024
// Time-stamp: <2024-09-16 19:01:31>
// Copyright (C): 2024 Pierre Lecocq

package validators

import (
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateTaskForUpdate(t *testing.T) {
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
	}

	for _, d := range dp {
		t.Log(d.name)

		res := ValidateTaskForUpdate(d.task)

		if d.err != nil {
			assert.Error(t, d.err, res)
		} else {
			assert.Nil(t, res)
		}
	}
}
