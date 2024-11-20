package systatus

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(HealthResponse{Status: "HEALTHY"})
}
