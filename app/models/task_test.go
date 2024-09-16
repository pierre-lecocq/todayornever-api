// File: task_test.go
// Creation: Mon Sep  9 09:31:38 2024
// Time-stamp: <2024-09-15 00:06:12>
// Copyright (C): 2024 Pierre Lecocq

package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTasks(t *testing.T) {
	type DataProvider struct {
		name   string
		userID int
		page   int
		max    int
		len    int
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			name:   "User 1, page 1, 20 max",
			userID: 1,
			page:   1,
			max:    20,
			len:    2,
		},
		{
			name:   "User 1, page 1, 1 max",
			userID: 1,
			page:   1,
			max:    1,
			len:    1,
		},
		{
			name:   "User 1, page 2, 1 max",
			userID: 1,
			page:   2,
			max:    1,
			len:    1,
		},
		{
			name:   "User 2, page 2, 20 max",
			userID: 2,
			page:   1,
			max:    20,
			len:    1,
		},
		{
			name:   "User 3, page 2, 20 max",
			userID: 3,
			page:   1,
			max:    20,
			len:    0,
		},
	}

	for _, d := range dp {
		t.Log(d.name)
		res, err := ListTasks(db, int64(d.userID), int64(d.page), int64(d.max))
		assert.Nil(t, err)
		assert.Equal(t, d.len, len(res))
	}
}

func TestCountTasks(t *testing.T) {
	type DataProvider struct {
		name   string
		userID int
		len    int
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			name:   "User 1",
			userID: 1,
			len:    2,
		},
		{
			name:   "User 2",
			userID: 2,
			len:    1,
		},
		{
			name:   "User 3",
			userID: 3,
			len:    0,
		},
	}

	for _, d := range dp {
		t.Log(d.name)
		res, err := CountTasks(db, int64(d.userID))
		assert.Nil(t, err)
		assert.Equal(t, d.len, res)

	}
}

func TestFetchTask(t *testing.T) {
	type DataProvider struct {
		name     string
		userID   int64
		ID       int64
		title    string
		position int64
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			name:     "User 1, task 1",
			userID:   1,
			ID:       1,
			title:    "First task",
			position: 1,
			err:      nil,
		},
		{
			name:     "User 1, task 3",
			userID:   1,
			ID:       3,
			title:    "Third task",
			position: 2,
			err:      nil,
		},
		{
			name:     "User 2, task 2",
			userID:   2,
			ID:       2,
			title:    "Second task",
			position: 1,
			err:      nil,
		},
		{
			name:     "User 1, task 2 - should fail",
			userID:   1,
			ID:       2,
			title:    "First task",
			position: 1,
			err:      fmt.Errorf("sql: no rows in result set"),
		},
	}

	for _, d := range dp {
		t.Log(d.name)
		res, err := FetchTask(db, int64(d.userID), int64(d.ID))

		if d.err != nil {
			assert.Equal(t, d.err, err)
		} else {
			assert.Equal(t, d.userID, res.UserID)
			assert.Equal(t, d.ID, res.ID)
			assert.Equal(t, d.title, res.Title)
			assert.Equal(t, d.position, res.Position)
		}
	}
}

func TestCreateTask(t *testing.T) {
	type DataProvider struct {
		name     string
		userID   int
		title    string
		position int
		len      int
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			name:     "User 1, new task",
			userID:   1,
			title:    "Fourth task",
			position: 3,
			len:      3,
			err:      nil,
		},
	}

	for _, d := range dp {
		t.Log(d.name)
		res, err := CreateTask(db, int64(d.userID), Task{Title: d.title})

		if d.err != nil {
			assert.Equal(t, d.err, err)
		} else {
			assert.Equal(t, int64(d.position), res.Position)

			res2, _ := ListTasks(db, int64(d.userID), 1, 20)
			assert.Equal(t, d.len, len(res2))
		}
	}
}

func TestUpdateTask(t *testing.T) {
	type DataProvider struct {
		name   string
		userID int
		ID     int
		title  string
		len    int
		err    error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			name:   "User 1, task 1",
			userID: 1,
			ID:     1,
			title:  "First edited task",
			len:    2,
			err:    nil,
		},
		{
			name:   "User 2, task 2",
			userID: 2,
			ID:     2,
			title:  "Second edited task",
			len:    1,
			err:    nil,
		},
		{
			name:   "User 2, task 1 - should fail",
			userID: 2,
			ID:     1,
			title:  "Badly edited task",
			len:    2,
			err:    fmt.Errorf("sql: no rows in result set"),
		},
	}

	for _, d := range dp {
		t.Log(d.name)
		res, err := UpdateTask(db, int64(d.userID), int64(d.ID), Task{Title: d.title})

		if d.err != nil {
			assert.Equal(t, d.err, err)
		} else {
			assert.Equal(t, d.title, res.Title)

			res2, _ := FetchTask(db, int64(d.userID), int64(d.ID))
			assert.Equal(t, d.title, res2.Title)

			res3, _ := ListTasks(db, int64(d.userID), 1, 20)
			assert.Equal(t, d.len, len(res3))
		}
	}
}

func TestDeleteTask(t *testing.T) {
	type DataProvider struct {
		name   string
		userID int
		ID     int
		title  string
		len    int
		nb     int
		err    error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			name:   "User 1, task 1 - should delete",
			userID: 1,
			ID:     1,
			len:    1,
			nb:     1,
		},
		{
			name:   "User 2, task 1 - should not delete",
			userID: 2,
			ID:     1,
			len:    1,
			nb:     0,
		},
	}

	for _, d := range dp {
		t.Log(d.name)
		res, _ := DeleteTask(db, int64(d.userID), int64(d.ID))

		assert.Equal(t, int64(d.nb), res)

		res2, _ := ListTasks(db, int64(d.userID), 1, 20)
		assert.Equal(t, d.len, len(res2))
	}
}
