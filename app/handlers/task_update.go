// File: task_update.go
// Creation: Thu Sep  5 15:33:26 2024
// Time-stamp: <2024-09-16 18:59:52>
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
			response.SendJSONError(w, http.StatusBadRequest, "Invalid UserID value in context")
			return
		}

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, "Can not decode JSON body")
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, "Invalid ID parameter in URL")
			return
		}

		err = validators.ValidateTaskForUpdate(t)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err := models.UpdateTask(db, userID, int64(id), t)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, "Can not update task")
			return
		}

		log.Debug().Msgf("Task %d updated", id)

		response.SendJSON(w, http.StatusOK, task)
	}
}
