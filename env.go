package systatus

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type EnvResponse struct {
	Env map[string]string `json:"env"`
}

func HandleEnv(w http.ResponseWriter, r *http.Request) {

	res := EnvResponse{}
	env := os.Environ()

	res.Env = make(map[string]string, len(env))

	for _, val := range env {
		split := strings.Split(val, "=")
		res.Env[split[0]] = split[1]
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
