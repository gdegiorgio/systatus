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
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
}

func useMiddlewares(handler func(w http.ResponseWriter, r *http.Request), middlewares []func(next http.HandlerFunc) http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	for _, mw := range middlewares {
		handler = mw(handler)
	}
	return handler
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

	mux.HandleFunc(fmt.Sprintf("%s/health", opts.Prefix), useMiddlewares(healthcheck, opts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/uptime", opts.Prefix), useMiddlewares(HandleUptime, opts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/cpu", opts.Prefix), useMiddlewares(HandleCPU, opts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/mem", opts.Prefix), useMiddlewares(HandleCPU, opts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/disk", opts.Prefix), useMiddlewares(HandleDisk, opts.Middlewares))

	if opts.ExposeEnv {
		mux.HandleFunc(fmt.Sprintf("%s/env", opts.Prefix), HandleEnv)
	}
}
