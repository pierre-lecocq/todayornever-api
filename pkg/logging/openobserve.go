// File: openobserve.go
// Creation: Fri Sep 20 11:16:09 2024
// Time-stamp: <2024-09-27 18:18:10>
// Copyright (C): 2024 Pierre Lecocq

package logging

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type OpenObserveWriter struct {
	Org      string
	Username string
	Password string
	Stream   string
}

func (hw *OpenObserveWriter) Write(p []byte) (n int, err error) {
	// Messages should be buffered and sent by batch in a routine, instead of being sent one by one over HTTP
	// -> HTTP is slow
	// -> we can DDOS the service easily by producing many requests in a short period of time
	// @see https://blog.mi.hdm-stuttgart.de/index.php/2024/02/29/combining-zerolog-loki/

	go func(p []byte) {
		// Request
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("https://api.openobserve.ai/api/%s/%s/_json", hw.Org, hw.Stream),
			strings.NewReader(string(p)),
		)

		if err != nil {
			return
		}

		req.SetBasicAuth(hw.Username, hw.Password)
		req.Header.Set("Content-Type", "application/json")

		// Context
		ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)

		defer cancel()

		req = req.WithContext(ctx)

		// Response
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			_, err := io.ReadAll(resp.Body)

			if err != nil {
				return
			}
		}
	}(p)

	return len(p), nil
}
