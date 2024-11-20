package systatus

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

type UptimeResponse struct {
	Systime string `json:"systime"`
	Uptime  string `json:"uptime"`
}

func HandleUptime(w http.ResponseWriter, r *http.Request) {
	res := UptimeResponse{}
	cmdoutput, err := exec.Command("/bin/uptime").Output()
	if err != nil {
		http.Error(w, "Could not exec uptime command on this machine", http.StatusInternalServerError)
		return
	}
	split := strings.Split(string(cmdoutput), " ")

	res.Systime = split[1]
	// Remove comma e.g 3:05,
	res.Uptime = strings.Split(split[4], ",")[0]

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
