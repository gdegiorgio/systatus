package systatus

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvSensitiveValues(t *testing.T) {
	mux := http.NewServeMux()

	opts := SystatusOptions{
		Prefix:    "",
		Mux:       mux,
		ExposeEnv: true,
		EnvHandlerOpts: EnvHandlerOpts{
			SensitiveKeys: []string{"PASSWORD"},
		},
	}

	os.Setenv("PASSWORD", "CLEARTEXT-PASSWORD")

	Enable(opts)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/env", ts.URL))
	assert.Equal(t, 200, res.StatusCode)

	var r EnvResponse
	err := json.NewDecoder(res.Body).Decode(&r)
	assert.NoError(t, err)
	assert.NotEqual(t, "CLEARTEXT-PASSWORD", r.Env["PASSWORD"])
	assert.Equal(t, "******************", r.Env["PASSWORD"])
}
