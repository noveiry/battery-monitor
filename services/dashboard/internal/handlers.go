package internal

import (
	"encoding/json"
	"net/http"
)

func RegisterHandlers(
	mux *http.ServeMux,
	client *Client,
) {

	mux.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write([]byte("ok"))
		},
	)

	mux.HandleFunc(
		"/fleet",
		func(w http.ResponseWriter, r *http.Request) {

			batteries, err := client.Batteries()
			if err != nil {
				http.Error(
					w,
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}

			stats, err := client.Stats()
			if err != nil {
				http.Error(
					w,
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}

			alerts, err := client.Alerts()
			if err != nil {
				http.Error(
					w,
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}

			writeJSON(
				w,
				FleetResponse{
					Stats:     stats,
					Batteries: batteries,
					Alerts:    alerts,
				},
			)
		},
	)

	mux.HandleFunc(
		"/fleet/stats",
		func(w http.ResponseWriter, r *http.Request) {

			stats, err := client.Stats()
			if err != nil {

				http.Error(
					w,
					err.Error(),
					http.StatusInternalServerError,
				)

				return
			}

			writeJSON(
				w,
				stats,
			)
		},
	)

	mux.HandleFunc(
		"/fleet/alerts",
		func(w http.ResponseWriter, r *http.Request) {

			alerts, err := client.Alerts()
			if err != nil {

				http.Error(
					w,
					err.Error(),
					http.StatusInternalServerError,
				)

				return
			}

			writeJSON(
				w,
				alerts,
			)
		},
	)
}

func writeJSON(
	w http.ResponseWriter,
	v any,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).Encode(v)
}
