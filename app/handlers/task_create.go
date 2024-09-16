// File: task_create.go
// Creation: Thu Sep  5 15:33:00 2024
// Time-stamp: <2024-09-16 18:59:19>
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
			response.SendJSONError(w, http.StatusBadRequest, "Invalid UserID value in context")
			return
		}

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, "Can not decode JSON body")
			return
		}

		err = validators.ValidateTaskForCreation(t)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err := models.CreateTask(db, userID, t)

		if err != nil {
			log.Debug().Msg("Can not create task")
			response.SendJSONError(w, http.StatusBadRequest, "Can not create task")
			return
		}

		log.Debug().Msgf("Task %d created", task.ID)

		response.SendJSON(w, http.StatusCreated, task)
	}
}
