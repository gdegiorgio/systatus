package systatus

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

type UptimeHandlerOpts struct {
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
}
type UptimeResponse struct {
	Systime string `json:"systime"`
	Uptime  string `json:"uptime"`
}

func handleUptime(opts UptimeHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var res UptimeResponse
		var err error

		if runtime.GOOS == "windows" {
			res, err = getWinUptime()
		} else {
			res, err = getUptime()
		}

		if err != nil {
			log.Err(err).Msg("Could not execute uptime on this host")
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Err(err).Msg("Failed to encode response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func getWinUptime() (UptimeResponse, error) {
	log.Warn().Msg("If FastBoot is enabled, uptime may be tracked incorrectly")
	return UptimeResponse{}, nil
}

func getUptime() (UptimeResponse, error) {
	log.Debug().Msg("Handling uptime on unix machine")
	cmdoutput, err := exec.Command("/usr/bin/uptime").Output()
	if err != nil {
		return UptimeResponse{}, err
	}
	splitCmd := strings.Split(string(cmdoutput), " ")
	return UptimeResponse{
		Systime: strings.TrimSpace(splitCmd[0]),
		Uptime:  strings.TrimSpace(strings.Split(splitCmd[3], ",")[0]),
	}, nil
}
