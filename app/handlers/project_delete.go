// File: project_delete.go
// Creation: Thu Sep 26 14:45:54 2024
// Time-stamp: <2024-09-26 14:46:30>
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

func ProjectDeleteHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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

		nb, err := models.DeleteProject(db, userID, int64(id))

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		if nb == 0 {
			log.Error().Int("UserID", int(userID)).Msgf("Can not delete project %d", id)
			response.SendJSONError(w, http.StatusBadRequest, "Can not delete project")
			return
		}

		log.Info().Int("UserID", int(userID)).Msgf("Project %d deleted", id)

		response.SendJSON(w, http.StatusNoContent, nil)
	}
}
