package systatus

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SystatusOptions struct {
	Mux               *http.ServeMux
	Prefix            string
	ExposeEnv         bool
	PrettyLogger      bool
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

	if opts.PrettyLogger {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	mux := http.DefaultServeMux

	if opts.Mux != nil {
		mux = opts.Mux
	}
	mux.HandleFunc(fmt.Sprintf("%s/health", opts.Prefix), useMiddlewares(handleHealth(opts.HealthHandlerOpts), opts.CPUHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/uptime", opts.Prefix), useMiddlewares(handleUptime(opts.UptimeHandlerOpts), opts.UptimeHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/cpu", opts.Prefix), useMiddlewares(handleCPU(opts.CPUHandlerOpts), opts.CPUHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/mem", opts.Prefix), useMiddlewares(HandleMem(opts.MemHandlerOpts), opts.MemHandlerOpts.Middlewares))
	mux.HandleFunc(fmt.Sprintf("%s/disk", opts.Prefix), useMiddlewares(handleDisk(opts.DiskHandlerOpts), opts.DiskHandlerOpts.Middlewares))

	if opts.ExposeEnv {
		mux.HandleFunc(fmt.Sprintf("%s/env", opts.Prefix), useMiddlewares(handleEnv(opts.EnvHandlerOpts), opts.EnvHandlerOpts.Middlewares))
	}
}
