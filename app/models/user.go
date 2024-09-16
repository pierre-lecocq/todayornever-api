// File: user.go
// Creation: Thu Sep  5 08:30:07 2024
// Time-stamp: <2024-09-14 18:39:06>
// Copyright (C): 2024 Pierre Lecocq

package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64      `sql:"id" json:"id"`
	Username     string     `sql:"username" json:"username"`
	Email        string     `sql:"email" json:"email"`
	Password     string     `sql:"-" json:"password"`
	PasswordHash string     `sql:"password_hash" json:"password_hash"`
	Salt         string     `sql:"salt" json:"salt"`
	State        string     `sql:"state" json:"state"`
	CreatedAt    *time.Time `sql:"created_at" json:"created_at"`
}

type UserToken struct {
	Message string `sql:"-" json:"message"`
	Token   string `sql:"-" json:"token"`
}

func LoginUser(db *sql.DB, email string, password string) (User, error) {
	u := User{}

	err := db.QueryRow(
		"SELECT id, username, password_hash, salt FROM user WHERE email = ? AND state='active'",
		email,
	).Scan(
		&u.ID,
		&u.Username,
		&u.PasswordHash,
		&u.Salt,
	)

	if err != nil {
		return u, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password+u.Salt))

	if err != nil {
		return u, err
	}

	return u, nil
}

func GeneratePasswordHashAndSalt(password string) (string, string, error) {
	saltb := make([]byte, 24)
	_, err := rand.Read(saltb)

	if err != nil {
		return "", "", fmt.Errorf("Error generating salt")
	}

	salt := base64.StdEncoding.EncodeToString(saltb)

	// https://gowebexamples.com/password-hashing/
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password+salt), 14)

	if err != nil {
		return "", "", fmt.Errorf("Error hashing password")
	}

	return string(passwordHash), salt, err
}

func CreateUser(db *sql.DB, u User) (User, error) {
	passwordHash, salt, err := GeneratePasswordHashAndSalt(u.Password)

	stmt, err := db.Prepare(`INSERT INTO user (username, email, password_hash, salt, state) VALUES (?, ?, ?, ?, 'active')
      RETURNING id, username, email, state, created_at`)

	if err != nil {
		return u, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		u.Username,
		u.Email,
		passwordHash,
		salt,
	).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.State,
		&u.CreatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}
