package middleware

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msgf("%s %s", r.Method, r.Pattern)
		next(w, r)
	}
}
