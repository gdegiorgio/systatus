package systatus

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

type CPUHandlerOpts struct {
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
}
type CPUResponse struct {
	LoadAverage string `json:"load_average"`
}

func handleCPU(opts CPUHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var res CPUResponse
		var err error

		switch runtime.GOOS {
		case "windows":
			res.LoadAverage, err = getWinLoadAverage()
			break
		case "darwin":
			res.LoadAverage, err = getMacLoadAverage()
			break
		default:
			res.LoadAverage, err = getLoadAverage()
		}

		if err != nil {
			log.Err(err).Msg("Failed to retrieve CPU info")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		log.Debug().Msg(res.LoadAverage)
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Err(err).Msg("Failed to encode response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func getWinLoadAverage() (string, error) {
	cmdoutput, err := exec.Command("powershell", "-Command", "Get-WmiObject Win32_Processor | Measure-Object -Property LoadPercentage -Average | Select Average\n").Output()
	if err != nil {
		return "", err
	}
	avg := strings.TrimSpace(strings.Split(string(cmdoutput), "-------")[1])
	return avg, nil
}

func getMacLoadAverage() (string, error) {
	return "", nil
}

func getLoadAverage() (string, error) {
	return "", nil
}
