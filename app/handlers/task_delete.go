// File: task_delete.go
// Creation: Thu Sep  5 15:33:07 2024
// Time-stamp: <2024-09-20 11:25:18>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func TaskDeleteHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("UserID").(int64)

		if userID == 0 {
			log.Error().Msg("Invalid UserID value in context")
			response.SendJSONError(w, http.StatusBadRequest, "Invalid UserID value in context")
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Invalid ID parameter in URL")
			return
		}

		nb, err := models.DeleteTask(db, userID, int64(id))

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		if nb == 0 {
			log.Error().Int("UserID", int(userID)).Msgf("Can not delete task %d", id)
			response.SendJSONError(w, http.StatusBadRequest, "Can not delete task")
			return
		}

		log.Info().Int("UserID", int(userID)).Msgf("Task %d deleted", id)

		response.SendJSON(w, http.StatusNoContent, nil)
	}
}
