package systatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultMux(t *testing.T) {
	opts := SystatusOptions{
		Prefix: "",
		Mux:    nil,
	}
	Enable(opts)

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/health", ts.URL))
	assert.Equal(t, 200, res.StatusCode)
}

func TestCustomMux(t *testing.T) {

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
}

func TestMemEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	opts := SystatusOptions{
		Prefix: "",
		Mux:    mux,
	}

	Enable(opts)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/mem", ts.URL))
	assert.Equal(t, 200, res.StatusCode)

	var m MemResponse
	err := json.NewDecoder(res.Body).Decode(&m)

	assert.NoError(t, err)
	assert.NotNil(t, m.Alloc)
	assert.NotNil(t, m.TotalAlloc)
	assert.NotNil(t, m.Sys)
}

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
	assert.Equal(t, h.Status, "HEALTHY")
}

func TestHealthEndpointCustomHandler(t *testing.T) {
	mux := http.NewServeMux()

	opts := SystatusOptions{
		Prefix: "",
		Mux:    mux,
		Healthcheck: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(HealthResponse{Status: "HEALTHY-CUSTOM"})
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
	assert.Equal(t, h.Status, "HEALTHY-CUSTOM")
}
