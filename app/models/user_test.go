// File: user_test.go
// Creation: Mon Sep  9 09:31:42 2024
// Time-stamp: <2024-09-15 00:06:26>
// Copyright (C): 2024 Pierre Lecocq

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	type DataProvider struct {
		name     string
		ID       int64
		username string
		email    string
		password string
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	dp := []DataProvider{
		{
			name:     "User 3",
			ID:       3,
			username: "user3",
			email:    "user3@mail.com",
			password: "user3user3user3",
			err:      nil,
		},
	}

	for _, d := range dp {
		t.Log(d.name)
		res, err := CreateUser(db, User{Username: d.username, Email: d.email, Password: d.password})

		if d.err != nil {
			assert.Equal(t, d.err, err)
		} else {
			assert.Equal(t, int64(d.ID), res.ID)
			assert.Equal(t, d.email, res.Email)
			assert.Equal(t, d.username, res.Username)
		}
	}
}

func TestLoginUser(t *testing.T) {
	type DataProvider struct {
		name     string
		ID       int64
		username string
		email    string
		password string
		err      error
	}

	db, _ := InitTestDatabase()
	defer db.Close()

	res, err := LoginUser(db, "user1@mail.com", "user1user1user1")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), res.ID)

	res2, err2 := LoginUser(db, "user2@mail.com", "user2user2user2")
	assert.Nil(t, err2)
	assert.Equal(t, int64(2), res2.ID)
}
