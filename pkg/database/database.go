// File: database.go
// Creation: Tue Aug 13 10:18:16 2024
// Time-stamp: <2024-09-15 14:57:49>
// Copyright (C): 2024 Pierre Lecocq

package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	ValidateConfig() error
	Connect() (*sql.DB, error)
}

type Config struct {
	Engine string `mapstructure:"engine"`
	DSN    string `mapstructure:"dsn"`
}

func Connect(c *Config) (*sql.DB, error) {
	return sql.Open(c.Engine, c.DSN)
}
