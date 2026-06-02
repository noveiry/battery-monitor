package internal

import (
	"encoding/json"
	"net/http"
)

func RegisterHandlers(
	mux *http.ServeMux,
	store *Store,
) {

	mux.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write([]byte("ok"))
		},
	)

	mux.HandleFunc(
		"/alerts",
		func(w http.ResponseWriter, r *http.Request) {

			writeJSON(
				w,
				store.All(),
			)
		},
	)

	mux.HandleFunc(
		"/alerts/critical",
		func(w http.ResponseWriter, r *http.Request) {

			writeJSON(
				w,
				store.Critical(),
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
