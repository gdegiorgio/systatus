package systatus

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type HealthHandlerOpts struct {
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
	Healthcheck func(w http.ResponseWriter, r *http.Request)
}

type HealthResponse struct {
	Status string `json:"status"`
}

func HandleHealth(opts HealthHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	if opts.Healthcheck != nil {
		log.Debug().Msg("Found a custom healthcheck")
		return opts.Healthcheck
	}
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(HealthResponse{Status: "HEALTHY"})
	}
}
