// File: project_list.go
// Creation: Wed Sep 25 10:36:03 2024
// Time-stamp: <2024-09-25 10:38:13>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"database/sql"
	"net/http"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/rs/zerolog/log"
)

func ProjectListHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("UserID").(int64)

		if userID == 0 {
			log.Error().Msg("Invalid UserID value in context")
			response.SendJSONError(w, http.StatusBadRequest, "Invalid UserID value in context")
			return
		}

		projects, err := models.ListProjects(db, userID)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		log.Info().Int("UserID", int(userID)).Msgf("%d projects listed", len(projects))

		response.SendJSON(w, http.StatusOK, projects)
	}
}
