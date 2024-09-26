// File: project_test.go
// Creation: Wed Sep 25 15:36:33 2024
// Time-stamp: <2024-09-26 14:25:47>
// Copyright (C): 2024 Pierre Lecocq

package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProjects(t *testing.T) {
	type DataProvider struct {
		testName string
		userID   int
		page     int
		max      int
		len      int
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			testName: "User 1",
			userID:   1,
			len:      2,
		},
		{
			testName: "User 2",
			userID:   2,
			len:      1,
		},
		{
			testName: "User 3",
			userID:   3,
			len:      0,
		},
	}

	for _, d := range dp {
		t.Log(d.testName)
		res, err := ListProjects(db, int64(d.userID))
		assert.Nil(t, err)
		assert.Equal(t, d.len, len(res))
	}
}

func TestFetchProject(t *testing.T) {
	type DataProvider struct {
		testName string
		userID   int64
		ID       int64
		name     string
		position int64
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			testName: "User 1, project 1",
			userID:   1,
			ID:       1,
			name:     "First project",
			position: 1,
			err:      nil,
		},
		{
			testName: "User 1, project 3",
			userID:   1,
			ID:       3,
			name:     "Third project",
			position: 2,
			err:      nil,
		},
		{
			testName: "User 2, project 2",
			userID:   2,
			ID:       2,
			name:     "Second project",
			position: 1,
			err:      nil,
		},
		{
			testName: "User 1, project 2 - should fail",
			userID:   1,
			ID:       2,
			name:     "First project",
			position: 1,
			err:      fmt.Errorf("sql: no rows in result set"),
		},
	}

	for _, d := range dp {
		t.Log(d.testName)
		res, err := FetchProject(db, int64(d.userID), int64(d.ID))

		if d.err != nil {
			assert.Equal(t, d.err, err)
		} else {
			assert.Equal(t, d.userID, res.UserID)
			assert.Equal(t, d.ID, res.ID)
			assert.Equal(t, d.name, res.Name)
			assert.Equal(t, d.position, res.Position)
		}
	}
}

func TestCreateProject(t *testing.T) {
	type DataProvider struct {
		testName string
		userID   int
		name     string
		position int
		len      int
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			testName: "User 1, new project",
			userID:   1,
			name:     "Fourth project",
			position: 3,
			len:      3,
			err:      nil,
		},
	}

	for _, d := range dp {
		t.Log(d.testName)
		res, err := CreateProject(db, int64(d.userID), Project{Name: d.name})

		if d.err != nil {
			assert.Equal(t, d.err, err)
		} else {
			assert.Equal(t, int64(d.position), res.Position)

			res2, _ := ListProjects(db, int64(d.userID))
			assert.Equal(t, d.len, len(res2))
		}
	}
}

func TestUpdateProject(t *testing.T) {
	type DataProvider struct {
		testName string
		userID   int
		ID       int
		name     string
		len      int
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			testName: "User 1, project 1",
			userID:   1,
			ID:       1,
			name:     "First edited oriject",
			len:      2,
			err:      nil,
		},
		{
			testName: "User 2, project 2",
			userID:   2,
			ID:       2,
			name:     "Second edited project",
			len:      1,
			err:      nil,
		},
		{
			testName: "User 2, project 1 - should fail",
			userID:   2,
			ID:       1,
			name:     "Badly edited project",
			len:      2,
			err:      fmt.Errorf("sql: no rows in result set"),
		},
	}

	for _, d := range dp {
		t.Log(d.testName)
		res, err := UpdateProject(db, int64(d.userID), int64(d.ID), Project{Name: d.name})

		if d.err != nil {
			assert.Equal(t, d.err, err)
		} else {
			assert.Equal(t, d.name, res.Name)

			res2, _ := FetchProject(db, int64(d.userID), int64(d.ID))
			assert.Equal(t, d.name, res2.Name)

			res3, _ := ListProjects(db, int64(d.userID))
			assert.Equal(t, d.len, len(res3))
		}
	}
}

func TestDeleteProject(t *testing.T) {
	type DataProvider struct {
		testName string
		userID   int
		ID       int
		len      int
		nb       int
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			testName: "User 1, project 1 - should delete",
			userID:   1,
			ID:       1,
			len:      1,
			nb:       1,
		},
		{
			testName: "User 2, project 1 - should not delete",
			userID:   2,
			ID:       1,
			len:      1,
			nb:       0,
		},
	}

	for _, d := range dp {
		t.Log(d.testName)
		res, _ := DeleteProject(db, int64(d.userID), int64(d.ID))

		assert.Equal(t, int64(d.nb), res)

		res2, _ := ListProjects(db, int64(d.userID))
		assert.Equal(t, d.len, len(res2))
	}
}
