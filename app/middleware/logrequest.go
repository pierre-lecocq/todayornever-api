// File: logrequest.go
// Creation: Fri Sep  6 22:47:10 2024
// Time-stamp: <2024-09-20 12:10:46>
// Copyright (C): 2024 Pierre Lecocq

package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type ResponseWriterRecorder struct {
	http.ResponseWriter
	Status int
	Start  time.Time
}

func (rws *ResponseWriterRecorder) WriteHeader(code int) {
	rws.Status = code
	rws.ResponseWriter.WriteHeader(code)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rws := ResponseWriterRecorder{w, 200, time.Now()}

		next.ServeHTTP(&rws, r)

		if r.URL.Path != "/health" {
			log.Info().Msgf("[%s] %s %s (%db) - %s %d",
				r.RemoteAddr,
				r.Method,
				r.URL.Path,
				r.ContentLength,
				time.Since(rws.Start).String(),
				rws.Status,
			)
		}
	})
}
