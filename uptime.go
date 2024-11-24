package systatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
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
	cmdoutput, err := exec.Command("powershell", "-Command", "(get-date) – (gcim Win32_OperatingSystem).LastBootUpTime").Output()
	if err != nil {
		return UptimeResponse{}, err
	}

	splitCmd := strings.Split(strings.ReplaceAll(string(cmdoutput), "\r\n", "\n"), "\n")

	days := strings.Split(splitCmd[2], ":")
	hours := strings.Split(splitCmd[3], ":")
	minutes := strings.Split(splitCmd[4], ":")

	return UptimeResponse{
		Uptime: formatDate(strings.TrimSpace(days[1]), strings.TrimSpace(hours[1]), strings.TrimSpace(minutes[1])),
	}, nil
}

func getUptime() (UptimeResponse, error) {
	cmdoutput, err := exec.Command("/usr/bin/uptime").Output()
	if err != nil {
		return UptimeResponse{}, err
	}
	log.Debug().Msgf("System uptime: %v", strings.Split(string(cmdoutput), " "))

	splitCmd := strings.Split(strings.TrimSpace(string(cmdoutput)), " ")

	systime := strings.TrimSpace(splitCmd[0])
	uptime := strings.TrimSpace(strings.Split(splitCmd[4], ",")[0])

	return UptimeResponse{
		Systime: systime,
		Uptime:  uptime,
	}, nil
}

func formatDate(d string, h string, m string) string {

	days, _ := strconv.Atoi(d)
	hours, _ := strconv.Atoi(h)
	minutes, _ := strconv.Atoi(m)

	if days < 10 {
		d = fmt.Sprintf("0%d", days)
	}
	if hours < 10 {
		h = fmt.Sprintf("0%d", hours)
	}
	if minutes < 10 {
		m = fmt.Sprintf("0%d", minutes)
	}

	return fmt.Sprintf("%s:%s:%s", d, h, m)
}
