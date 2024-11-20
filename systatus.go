package systatus

import (
	"fmt"
	"net/http"
)

type SystatusOptions struct {
	Mux         *http.ServeMux
	Prefix      string
	ExposeEnv   bool
	Healthcheck func(w http.ResponseWriter, r *http.Request)
}

func Enable(opts SystatusOptions) {

	var healthcheck = HandleHealth

	mux := http.DefaultServeMux

	if opts.Mux != nil {
		mux = opts.Mux
	}

	if opts.Healthcheck != nil {
		healthcheck = opts.Healthcheck
	}

	mux.HandleFunc(fmt.Sprintf("%s/health", opts.Prefix), healthcheck)
	mux.HandleFunc(fmt.Sprintf("%s/uptime", opts.Prefix), HandleUptime)
	mux.HandleFunc(fmt.Sprintf("%s/cpu", opts.Prefix), HandleCPU)
	mux.HandleFunc(fmt.Sprintf("%s/mem", opts.Prefix), HandleMem)
	mux.HandleFunc(fmt.Sprintf("%s/disk", opts.Prefix), HandleDisk)

	if opts.ExposeEnv {
		mux.HandleFunc(fmt.Sprintf("%s/env", opts.Prefix), HandleEnv)
	}
}
