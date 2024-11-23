package systatus

import (
	"fmt"
	"net/http"
)

type SystatusOptions struct {
	Mux               *http.ServeMux
	Prefix            string
	ExposeEnv         bool
	HealthHandlerOpts HealthHandlerOpts
	CPUHandlerOpts    CPUHandlerOpts
	EnvHandlerOpts    EnvHandlerOpts
	DiskHandlerOpts   DiskHandlerOpts
	UptimeHandlerOpts UptimeHandlerOpts
	MemHandlerOpts    MemHandlerOpts
}

func useMiddlewares(handler func(w http.ResponseWriter, r *http.Request), middlewares []func(next http.HandlerFunc) http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	for _, mw := range middlewares {
		handler = mw(handler)
	}
	return handler
}

func Enable(opts SystatusOptions) {

	mux := http.DefaultServeMux

	if opts.Mux != nil {
		mux = opts.Mux
	}
	mux.HandleFunc(fmt.Sprintf("%s/health", opts.Prefix), useMiddlewares(HandleHealth(opts.HealthHandlerOpts), opts.CPUHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/uptime", opts.Prefix), useMiddlewares(HandleUptime(opts.UptimeHandlerOpts), opts.UptimeHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/cpu", opts.Prefix), useMiddlewares(HandleCPU(opts.CPUHandlerOpts), opts.CPUHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/mem", opts.Prefix), useMiddlewares(HandleMem(opts.MemHandlerOpts), opts.MemHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/disk", opts.Prefix), useMiddlewares(HandleDisk(opts.DiskHandlerOpts), opts.DiskHandlerOpts.Middlewares))

	if opts.ExposeEnv {
		mux.HandleFunc(fmt.Sprintf("%s/env", opts.Prefix), useMiddlewares(HandleEnv(opts.EnvHandlerOpts), opts.EnvHandlerOpts.Middlewares))
	}
}
