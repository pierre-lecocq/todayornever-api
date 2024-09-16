// File: handler_test.go
// Creation: Tue Sep 10 08:42:53 2024
// Time-stamp: <2024-09-16 18:59:13>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/pkg/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type DataProvider struct {
	Name         string
	Handler      http.HandlerFunc
	Method       string
	Path         string
	URLVars      map[string]string
	ExpectedCode int
	Body         any
	Headers      http.Header
	Middleware   func(next http.Handler) http.Handler
	UserID       int64
}

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

	passwordHash, salt, _ := models.GeneratePasswordHashAndSalt("user1user1user1")
	db.Exec(`INSERT INTO user (username, email, password_hash, salt, state) VALUES (?, ?, ?, ?, ?)`,
		"user1", "user1@mail.com", passwordHash, salt, "active")

	passwordHash, salt, _ = models.GeneratePasswordHashAndSalt("user2user2user2")
	db.Exec(`INSERT INTO user (username, email, password_hash, salt, state) VALUES (?, ?, ?, ?, ?)`,
		"user2", "user2@mail.com", passwordHash, salt, "active")

	db.Exec(`INSERT INTO task (user_id, title, state, position) VALUES (?, ?, ?, ?)`, 1, "First task", "todo", 1)
	db.Exec(`INSERT INTO task (user_id, title, state, position) VALUES (?, ?, ?, ?)`, 2, "Second task", "todo", 1)
	db.Exec(`INSERT INTO task (user_id, title, state, position) VALUES (?, ?, ?, ?)`, 1, "Third task", "todo", 2)

	return db, err
}

func RequestTest(t *testing.T, dp DataProvider) {
	t.Logf("%s\n", dp.Name)

	var req *http.Request
	var err error

	if dp.Body != nil {
		j, err := json.Marshal(dp.Body)

		if err != nil {
			t.Fatal(err)
		}

		req, err = http.NewRequest(dp.Method, dp.Path, bytes.NewBuffer(j))
	} else {
		req, err = http.NewRequest(dp.Method, dp.Path, nil)
	}

	if err != nil {
		t.Fatal(err)
	}

	if len(dp.URLVars) > 0 {
		req = mux.SetURLVars(req, dp.URLVars)
	}

	if dp.Headers != nil {
		req.Header = dp.Headers
	}

	rec := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "UserID", dp.UserID)

	dp.Handler.ServeHTTP(rec, req.WithContext(ctx))

	assert.Equal(t, dp.ExpectedCode, rec.Code)
}
