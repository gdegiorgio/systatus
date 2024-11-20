package systatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemEndpoint(t *testing.T) {

	mux := http.NewServeMux()
	opts := SystatusOptions{
		Prefix: "",
		Mux:    mux,
	}

	Enable(opts)

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	res, _ := http.Get(fmt.Sprintf("%s/health", testServer.URL))
	assert.Equal(t, 200, res.StatusCode)

	var h HealthResponse
	err := json.NewDecoder(res.Body).Decode(&h)
	assert.NoError(t, err)
	assert.Equal(t, h.Status, "HEALTHY")
}
