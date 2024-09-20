// File: task_create.go
// Creation: Thu Sep  5 15:33:00 2024
// Time-stamp: <2024-09-20 11:25:01>
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

func TaskCreateHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var t models.Task

		userID := r.Context().Value("UserID").(int64)

		if userID == 0 {
			log.Error().Msg("Invalid UserID value in context")
			response.SendJSONError(w, http.StatusBadRequest, "Invalid UserID value in context")
			return
		}

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not decode JSON body")
			return
		}

		err = validators.ValidateTaskForCreation(t)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err := models.CreateTask(db, userID, t)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not create task")
			return
		}

		log.Info().Int("UserID", int(userID)).Msgf("Task %d created", task.ID)

		response.SendJSON(w, http.StatusCreated, task)
	}
}
