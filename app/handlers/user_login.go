// File: user_login.go
// Creation: Thu Sep  5 15:33:47 2024
// Time-stamp: <2024-09-16 19:00:04>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/app/validators"
	"github.com/pierre-lecocq/todayornever-api/pkg/auth"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/rs/zerolog/log"
)

func UserLoginHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, "Can not decode JSON body")
			return
		}

		err = validators.ValidateUserForLogin(u)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := models.LoginUser(db, u.Email, u.Password)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusNotFound, "User Not Found")
			return
		}

		expires, err := strconv.Atoi(os.Getenv("AUTH_EXPIRES"))

		if err != nil {
			expires = 1
		}

		token, err := auth.GenerateJWTToken(
			user.ID,
			os.Getenv("AUTH_ISSUER"),
			os.Getenv("AUTH_SECRET"),
			expires,
		)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, "Can not generate token")
			return
		}

		log.Debug().Msgf("User %d logged in", user.ID)

		response.SendJSON(w, http.StatusOK, models.UserToken{
			Message: fmt.Sprintf("Welcome %s! Please use the provided token in your next queries to access your resources.", user.Username),
			Token:   token,
		})
	}
}
