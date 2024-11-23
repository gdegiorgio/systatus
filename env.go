package systatus

import (
	"encoding/json"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
)

type EnvHandlerOpts struct {
	Middlewares   []func(next http.HandlerFunc) http.HandlerFunc
	SensitiveKeys []string
}

type EnvResponse struct {
	Env map[string]string `json:"env"`
}

func handleEnv(opts EnvHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := EnvResponse{}
		env := os.Environ()

		res.Env = make(map[string]string, len(env))

		for _, envVar := range env {
			split := strings.Split(envVar, "=")
			key := split[0]
			val := split[1]
			if slices.Contains(opts.SensitiveKeys, key) {
				log.Info().Msgf("%s has been found in SensitiveKeys and will value be hidden", key)
				val = "******************"
			}
			res.Env[key] = val
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
