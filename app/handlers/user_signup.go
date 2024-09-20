// File: user_signup.go
// Creation: Thu Sep  5 15:33:36 2024
// Time-stamp: <2024-09-20 11:26:31>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/app/validators"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/rs/zerolog/log"
)

func UserSignupHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not decode JSON body")
			return
		}

		err = validators.ValidateUserForCreation(u)

		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := models.CreateUser(db, u)

		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not create user")
			return
		}

		log.Info().Msgf("User %d created", user.ID)

		response.SendJSON(w, http.StatusCreated, user)
	}
}
