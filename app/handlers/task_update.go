// File: task_update.go
// Creation: Thu Sep  5 15:33:26 2024
// Time-stamp: <2024-09-20 11:25:56>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/app/validators"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func TaskUpdateHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Invalid ID parameter in URL")
			return
		}

		err = validators.ValidateTaskForUpdate(t)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err := models.UpdateTask(db, userID, int64(id), t)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not update task")
			return
		}

		log.Info().Int("UserID", int(userID)).Msgf("Task %d updated", task.ID)

		response.SendJSON(w, http.StatusOK, task)
	}
}
