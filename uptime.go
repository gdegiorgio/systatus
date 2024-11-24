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

		res := UptimeResponse{}
		var err error

		log.Debug().Msgf("Handling uptime on %s machine", runtime.GOOS)

		if runtime.GOOS == "windows" {
			res, err = getWinUptime()
		} else {
			res, err = getUptime()
		}

		if err != nil {
			log.Err(err).Msg("Could not execute uptime on this host")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Err(err).Msg("Failed to encode response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func getWinUptime() (UptimeResponse, error) {
	log.Warn().Msg("If FastBoot is enabled, uptime on Windows may be tracked incorrectly")
	cmdoutput, err := exec.Command("systeminfo | find \"System Boot Time\"").Output()
	if err != nil {
		return UptimeResponse{}, err
	}
	log.Debug().Msg(string(cmdoutput))
	return UptimeResponse{}, nil
}

func getUptime() (UptimeResponse, error) {
	cmdoutput, err := exec.Command("/usr/bin/uptime").Output()
	if err != nil {
		return UptimeResponse{}, err
	}
	splitCmd := strings.Split(string(cmdoutput), " ")
	return UptimeResponse{
		Systime: strings.TrimSpace(splitCmd[1]),
		Uptime:  strings.TrimSpace(strings.Split(splitCmd[4], ",")[0]),
	}, nil
}
