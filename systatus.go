package systatus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type SystatusOptions struct {
	Mux           *http.ServeMux
	Prefix        string
	ExposeEnv     bool
	Healthcheck   func(w http.ResponseWriter, r *http.Request)
	SensitiveKeys []string
}

type HealthResponse struct{}
type UptimeResponse struct {
	Systime string `json:"systime"`
	Uptime  string `json:"uptime"`
}
type CPURepsponse struct{}
type MemResponse struct {
	TotalAlloc uint64
	Alloc      uint64
	Sys        uint64
}
type EnvResponse struct {
	Env map[string]string `json:"env"`
}

type contextKey string

const sensitiveKeysContextKey contextKey = "sensitiveKeys"

func Enable(opts SystatusOptions) {

	mux := http.DefaultServeMux

	if opts.Mux != nil {
		mux = opts.Mux
	}

	if opts.Healthcheck == nil {
		mux.HandleFunc(fmt.Sprintf("%s/health", opts.Prefix), handleHealth)
	} else {
		mux.HandleFunc(fmt.Sprintf("%s/health", opts.Prefix), opts.Healthcheck)
	}

	mux.HandleFunc(fmt.Sprintf("%s/uptime", opts.Prefix), handleUptime)
	mux.HandleFunc(fmt.Sprintf("%s/cpu", opts.Prefix), handleCPU)
	mux.HandleFunc(fmt.Sprintf("%s/mem", opts.Prefix), handleMem)
	mux.HandleFunc(fmt.Sprintf("%s/disk", opts.Prefix), handleDisk)

	if opts.ExposeEnv {
		envHandler := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), sensitiveKeysContextKey, opts.SensitiveKeys)
			handleEnv(w, r.WithContext(ctx))
		}
		mux.HandleFunc(fmt.Sprintf("%s/env", opts.Prefix), envHandler)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {

}

func handleUptime(w http.ResponseWriter, r *http.Request) {

	res := UptimeResponse{}

	if runtime.GOOS == "windows" {
		// TODO Implement windows uptime
	} else {
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
}
func handleCPU(w http.ResponseWriter, r *http.Request) {

}
func handleMem(w http.ResponseWriter, r *http.Request) {
	res := &MemResponse{}

	var stats runtime.MemStats

	runtime.ReadMemStats(&stats)

	res.Sys = stats.Sys
	res.TotalAlloc = stats.TotalAlloc
	res.Alloc = stats.Alloc

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
func handleDisk(w http.ResponseWriter, r *http.Request) {

}
func handleEnv(w http.ResponseWriter, r *http.Request) {
	res := EnvResponse{}
	env := os.Environ()

	// Get sensitive keys from context (set by Enable function)
	sensitiveKeys := r.Context().Value(sensitiveKeysContextKey).([]string)

	res.Env = make(map[string]string, len(env))

	for _, val := range env {
		split := strings.Split(val, "=")
		key, value := split[0], split[1]

		// Check if this key should be masked
		if containsSensitiveKey(key, sensitiveKeys) {
			res.Env[key] = "[REDACTED]"
		} else {
			res.Env[key] = value
		}
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Helper function to check if a key is sensitive
func containsSensitiveKey(key string, sensitiveKeys []string) bool {
	key = strings.ToLower(key)
	for _, sensitive := range sensitiveKeys {
		if strings.ToLower(sensitive) == key {
			return true
		}
		// Also check if key contains common sensitive patterns
		if strings.Contains(key, "password") ||
			strings.Contains(key, "secret") ||
			strings.Contains(key, "token") ||
			strings.Contains(key, "key") {
			return true
		}
	}
	return false
}
