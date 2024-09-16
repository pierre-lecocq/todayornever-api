// File: task_list.go
// Creation: Thu Sep  5 15:33:22 2024
// Time-stamp: <2024-09-16 18:59:44>
// Copyright (C): 2024 Pierre Lecocq

package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/pierre-lecocq/todayornever-api/app/models"
	"github.com/pierre-lecocq/todayornever-api/pkg/response"

	"github.com/rs/zerolog/log"
)

func TaskListHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("UserID").(int64)

		if userID == 0 {
			response.SendJSONError(w, http.StatusBadRequest, "Invalid UserID value in context")
			return
		}

		pageValue := r.URL.Query().Get("page")

		if pageValue == "" {
			pageValue = "1"
		}

		page, err := strconv.Atoi(pageValue)

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, "Invalid page query parameter")
			return
		}

		if page <= 0 {
			page = 1
		}

		tasks, err := models.ListTasks(db, userID, int64(page), int64(50))

		if err != nil {
			log.Debug().Err(err)
			response.SendJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		log.Debug().Msgf("%d tasks listed", len(tasks))

		response.SendJSON(w, http.StatusOK, tasks)
	}
}
