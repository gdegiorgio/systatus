package systatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type UptimeHandlerOpts struct {
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
}
type UptimeResponse struct {
	Systime string  `json:"systime"`
	Uptime  float64 `json:"uptime"`
}

func handleUptime(opts UptimeHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		res := UptimeResponse{
			Systime: time.Now().Format(time.RFC3339),
		}

		var err error

		log.Debug().Msgf("Handling uptime on %s machine", runtime.GOOS)

		switch runtime.GOOS {
		case "windows":
			res.Uptime, err = getWinUptime()
			break
		case "darwin":
			res.Uptime, err = getMacOSUptime()
			break
		default:
			res.Uptime, err = getUptime()
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

func getWinUptime() (float64, error) {
	log.Warn().Msg("If FastBoot is enabled, uptime on Windows may be tracked incorrectly")
	cmdoutput, err := exec.Command("powershell", "-Command", "(get-date) – (gcim Win32_OperatingSystem).LastBootUpTime").Output()
	if err != nil {
		return 0, err
	}

	splitCmd := strings.Split(strings.ReplaceAll(string(cmdoutput), "\r\n", "\n"), "\n")

	ms := strings.Split(splitCmd[10], ":")[1]
	uptime, err := strconv.ParseFloat(ms, 64)

	if err != nil {
		return 0, err
	}

	return uptime, nil
}

func getMacOSUptime() (float64, error) {
	cmd := "sysctl -n kern.boottime | cut -c9-18"
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {

		return 0, err
	}

	data := strings.TrimSpace(string(output))
	log.Debug().Msgf("%v", data)

	uptime, err := strconv.ParseFloat(data, 64)
	if err != nil {
		fmt.Printf("%w", err)
		return 0, err
	}
	return uptime, nil
}

func getUptime() (float64, error) {
	buf, err := os.ReadFile("/proc/uptime")

	if err != nil {
		return 0, err
	}

	data := strings.Split(string(buf), " ")[0]
	seconds, err := strconv.ParseFloat(data, 64)

	if err != nil {
		return 0, err
	}

	uptime := seconds * 1000

	return uptime, nil
}
