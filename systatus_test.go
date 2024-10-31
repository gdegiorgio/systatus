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

func TestEnvEndpoint(t *testing.T) {
	// test cases
	tests := []struct {
		name          string
		envVars       map[string]string
		sensitiveKeys []string
		expectedMask  map[string]bool
	}{
		{
			name: "Basic sensitive keys",
			envVars: map[string]string{
				"APP_NAME":     "testapp",
				"DB_PASSWORD":  "secret123",
				"API_KEY":      "abc123",
				"PUBLIC_VALUE": "visible",
			},
			sensitiveKeys: []string{"DB_PASSWORD", "API_KEY"},
			expectedMask: map[string]bool{
				"DB_PASSWORD":  true,
				"API_KEY":      true,
				"PUBLIC_VALUE": false,
				"APP_NAME":     false,
			},
		},
		{
			name: "Pattern matching sensitive keys",
			envVars: map[string]string{
				"MY_PASSWORD":    "secret123",
				"ANOTHER_SECRET": "xyz789",
				"PUBLIC_VALUE":   "visible",
			},
			sensitiveKeys: []string{},
			expectedMask: map[string]bool{
				"MY_PASSWORD":    true,
				"ANOTHER_SECRET": true,
				"PUBLIC_VALUE":   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}
			defer func() {
				for k := range tt.envVars {
					os.Unsetenv(k)
				}
			}()

			mux := http.NewServeMux()
			opts := SystatusOptions{
				Prefix:        "",
				Mux:           mux,
				ExposeEnv:     true,
				SensitiveKeys: tt.sensitiveKeys,
			}

			Enable(opts)
			ts := httptest.NewServer(mux)
			defer ts.Close()

			res, err := http.Get(fmt.Sprintf("%s/env", ts.URL))
			assert.NoError(t, err)
			assert.Equal(t, 200, res.StatusCode)

			var envResponse EnvResponse
			err = json.NewDecoder(res.Body).Decode(&envResponse)
			assert.NoError(t, err)

			// Verify masked valeus
			for k := range tt.envVars {
				value := envResponse.Env[k]
				if tt.expectedMask[k] {
					assert.Equal(t, "[REDACTED]", value, "Expected %s to be masked", k)
				} else {
					assert.Equal(t, tt.envVars[k], value, "Expected %s to be visible", k)
				}
			}
		})
	}
}
