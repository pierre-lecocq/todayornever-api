// File: response.go
// Creation: Thu Sep  5 09:22:24 2024
// Time-stamp: <2024-09-11 11:24:03>
// Copyright (C): 2024 Pierre Lecocq

package response

import (
	"encoding/json"
	"net/http"
)

func SendJSONError(w http.ResponseWriter, status int, message string) {
	SendJSON(w, status, map[string]interface{}{
		"error": message,
	})
}

func SendJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	payload, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		payload, _ = json.Marshal(map[string]interface{}{
			"error": "Can not encore JSON payload",
		})
	} else {
		w.WriteHeader(status)
	}

	w.Write(payload)
}
