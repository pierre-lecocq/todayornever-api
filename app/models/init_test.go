// File: model_test.go
// Creation: Mon Sep  9 11:11:45 2024
// Time-stamp: <2024-09-16 19:01:03>
// Copyright (C): 2024 Pierre Lecocq

package models

import (
	"database/sql"

	"github.com/pierre-lecocq/todayornever-api/pkg/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitTestDatabase() (*sql.DB, error) {
	db, err := database.Connect(&database.Config{Engine: "sqlite3", DSN: ":memory:"})

	if err != nil {
		panic(err)
	}

	driver, _ := sqlite3.WithInstance(db, &sqlite3.Config{})

	m1, err1 := migrate.NewWithDatabaseInstance("file://../migrations", "sqlite3", driver)

	if err1 != nil {
		panic(err1)
	}

	m1.Up()

	passwordHash, salt, _ := GeneratePasswordHashAndSalt("user1user1user1")
	db.Exec(`INSERT INTO user (username, email, password_hash, salt, state) VALUES (?, ?, ?, ?, ?)`,
		"user1", "user1@mail.com", passwordHash, salt, "active")

	passwordHash, salt, _ = GeneratePasswordHashAndSalt("user2user2user2")
	db.Exec(`INSERT INTO user (username, email, password_hash, salt, state) VALUES (?, ?, ?, ?, ?)`,
		"user2", "user2@mail.com", passwordHash, salt, "active")

	db.Exec(`INSERT INTO task (user_id, title, state, position) VALUES (?, ?, ?, ?)`, 1, "First task", "todo", 1)
	db.Exec(`INSERT INTO task (user_id, title, state, position) VALUES (?, ?, ?, ?)`, 2, "Second task", "todo", 1)
	db.Exec(`INSERT INTO task (user_id, title, state, position) VALUES (?, ?, ?, ?)`, 1, "Third task", "todo", 2)

	return db, err
}
