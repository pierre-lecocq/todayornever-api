// File: project_create.go
// Creation: Thu Sep 26 14:37:09 2024
// Time-stamp: <2024-09-26 14:38:16>
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

func ProjectCreateHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Project

		userID := r.Context().Value("UserID").(int64)

		if userID == 0 {
			log.Error().Msg("Invalid UserID value in context")
			response.SendJSONError(w, http.StatusBadRequest, "Invalid UserID value in context")
			return
		}

		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not decode JSON body")
			return
		}

		err = validators.ValidateProjectForCreation(p)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		project, err := models.CreateProject(db, userID, p)

		if err != nil {
			log.Error().Int("UserID", int(userID)).Err(err).Msg(err.Error())
			response.SendJSONError(w, http.StatusBadRequest, "Can not create project")
			return
		}

		log.Info().Int("UserID", int(userID)).Msgf("Project %d created", project.ID)

		response.SendJSON(w, http.StatusCreated, project)
	}
}
