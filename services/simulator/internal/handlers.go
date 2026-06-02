package internal

import (
	"encoding/json"
	"net/http"
	"strings"
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
		"/batteries",
		func(w http.ResponseWriter, r *http.Request) {

			writeJSON(
				w,
				store.List(),
			)
		},
	)

	mux.HandleFunc(
		"/stats",
		func(w http.ResponseWriter, r *http.Request) {

			writeJSON(
				w,
				store.Stats(),
			)
		},
	)

	mux.HandleFunc(
		"/batteries/",
		func(w http.ResponseWriter, r *http.Request) {

			id := strings.TrimPrefix(
				r.URL.Path,
				"/batteries/",
			)

			battery, ok := store.Get(id)

			if !ok {

				http.NotFound(
					w,
					r,
				)

				return
			}

			writeJSON(
				w,
				battery,
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
