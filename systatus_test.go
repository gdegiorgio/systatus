package systatus

import (
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
