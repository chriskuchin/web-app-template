package controller

import "net/http"

type (
	HealthCheck struct {
	}
)

func (hc HealthCheck) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", healthcheck)
	})

	return r
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
