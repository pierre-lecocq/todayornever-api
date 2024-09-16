// File: main.go
// Creation: Thu Sep  5 08:17:00 2024
// Time-stamp: <2024-09-16 19:02:13>
// Copyright (C): 2024 Pierre Lecocq

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	hd "github.com/pierre-lecocq/todayornever-api/app/handlers"
	mw "github.com/pierre-lecocq/todayornever-api/app/middleware"
	"github.com/pierre-lecocq/todayornever-api/pkg/database"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func validateConfig() error {
	keys := []string{
		"SERVICE_HOST",
		"SERVICE_PORT",
		"AUTH_ISSUER",
		"AUTH_SECRET",
		"AUTH_EXPIRES",
		"DATABASE_ENGINE",
		"DATABASE_DSN",
	}

	for k := range keys {
		_, ok := os.LookupEnv(keys[k])

		if !ok {
			return fmt.Errorf("Missing environment variable %s", keys[k])
		}
	}

	return nil
}

func initDatabase() (*sql.DB, error) {
	return database.Connect(&database.Config{
		Engine: os.Getenv("DATABASE_ENGINE"),
		DSN:    os.Getenv("DATABASE_DSN"),
	})
}

func initLogger() {
	lvl, err := strconv.Atoi(os.Getenv("LOGGER_LEVEL"))

	if err != nil {
		lvl = 3
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(lvl))
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	// Init

	err := validateConfig()

	if err != nil {
		panic(err)
	}

	initLogger()

	db, err := initDatabase()

	if err != nil {
		log.Fatal().Err(err)
	}

	// Setup

	r := mux.NewRouter()

	r.Use(mw.LogRequest)
	r.Use(mw.Ratelimit)
	r.Use(mw.Negociate)

	r.Handle("/", http.HandlerFunc(hd.IndexHandler())).Methods(http.MethodGet)

	r.Handle("/login", http.HandlerFunc(hd.UserLoginHandler(db))).Methods(http.MethodPost)
	r.Handle("/signup", http.HandlerFunc(hd.UserSignupHandler(db))).Methods(http.MethodPost)

	r.Handle("/tasks", mw.Auth(http.HandlerFunc(hd.TaskListHandler(db)))).Methods(http.MethodGet)
	r.Handle("/tasks", mw.Auth(http.HandlerFunc(hd.TaskCreateHandler(db)))).Methods(http.MethodPost)
	r.Handle("/tasks/{id}", mw.Auth(http.HandlerFunc(hd.TaskFetchHandler(db)))).Methods(http.MethodGet)
	r.Handle("/tasks/{id}", mw.Auth(http.HandlerFunc(hd.TaskUpdateHandler(db)))).Methods(http.MethodPatch, http.MethodPut)
	r.Handle("/tasks/{id}", mw.Auth(http.HandlerFunc(hd.TaskDeleteHandler(db)))).Methods(http.MethodDelete)

	// Start

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT")),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Info().Msgf("Starting service on port %s...", os.Getenv("SERVICE_PORT"))
	log.Panic().Err(srv.ListenAndServe())
}
