package systatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthEndpointDefaultHandler(t *testing.T) {
	mux := http.NewServeMux()

	opts := SystatusOptions{
		Prefix: "",
		Mux:    mux,
	}
	Enable(opts)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/health", ts.URL))
	assert.Equal(t, 200, res.StatusCode)

	var h HealthResponse
	err := json.NewDecoder(res.Body).Decode(&h)
	assert.NoError(t, err)
	assert.Equal(t, "HEALTHY", h.Status)
}

func TestHealthEndpointCustomHandler(t *testing.T) {
	mux := http.NewServeMux()

	opts := SystatusOptions{
		Prefix: "",
		Mux:    mux,
		HealthHandlerOpts: HealthHandlerOpts{
			Healthcheck: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				json.NewEncoder(w).Encode(HealthResponse{Status: "HEALTHY-CUSTOM"})
			},
		},
	}

	Enable(opts)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/health", ts.URL))
	assert.Equal(t, 200, res.StatusCode)

	var h HealthResponse
	err := json.NewDecoder(res.Body).Decode(&h)
	assert.NoError(t, err)
	assert.Equal(t, "HEALTHY-CUSTOM", h.Status)
}
