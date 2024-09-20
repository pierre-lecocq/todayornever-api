// File: user_login.go
// Creation: Thu Sep  5 15:33:47 2024
// Time-stamp: <2024-09-20 11:26:18>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/app/validators"
	"github.com/pierre-lecocq/todayornever-api/pkg/auth"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func UserLoginHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not decode JSON body")
			return
		}

		err = validators.ValidateUserForLogin(u)

		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := models.LoginUser(db, u.Email, u.Password)

		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusNotFound, "User Not Found")
			return
		}

		token, err := auth.GenerateJWTToken(
			user.ID,
			viper.GetString("AUTH_ISSUER"),
			viper.GetString("AUTH_SECRET"),
			viper.GetInt("AUTH_EXPIRES"),
		)

		if err != nil {
			log.Error().Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not generate token")
			return
		}

		log.Info().Msgf("User %d logged in", user.ID)

		response.SendJSON(w, http.StatusOK, models.UserToken{
			Message: fmt.Sprintf("Welcome %s! Please use the provided token in your next queries to access your resources.", user.Username),
			Token:   token,
		})
	}
}
