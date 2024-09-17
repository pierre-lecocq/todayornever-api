// File: main.go
// Creation: Thu Sep  5 08:17:00 2024
// Time-stamp: <2024-09-17 18:20:23>
// Copyright (C): 2024 Pierre Lecocq

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	hd "github.com/pierre-lecocq/todayornever-api/app/handlers"
	mw "github.com/pierre-lecocq/todayornever-api/app/middleware"
	"github.com/pierre-lecocq/todayornever-api/pkg/database"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	viper.SetDefault("SERVICE_HOST", "localhost")
	viper.SetDefault("SERVICE_PORT", 8080)

	viper.SetDefault("LOGGER_LEVEL", 3)

	viper.SetDefault("AUTH_ISSUER", "todayornever-api")
	viper.SetDefault("AUTH_EXPIRES", 1)

	viper.SetDefault("DATABASE_ENGINE", "sqlite3")
	viper.SetDefault("DATABASE_DSN", ":memory:")

	return nil
}

func initDatabase() (*sql.DB, error) {
	return database.Connect(&database.Config{
		Engine: viper.GetString("DATABASE_ENGINE"),
		DSN:    viper.GetString("DATABASE_DSN"),
	})
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(viper.GetInt("LOGGER_LEVEL")))
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	// Init

	err := initConfig()

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
		Addr:         fmt.Sprintf("%s:%d", viper.Get("SERVICE_HOST"), viper.GetInt("SERVICE_PORT")),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Info().Msgf("Starting service on port %d...", viper.GetInt("SERVICE_PORT"))
	log.Panic().Err(srv.ListenAndServe())
}
