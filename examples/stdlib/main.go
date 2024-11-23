package main

import (
	"fmt"
	"net/http"

	"github.com/gdegiorgio/systatus"
	"github.com/gdegiorgio/systatus/middleware"
	"github.com/rs/zerolog/log"
)

func customHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"application is healthy"}`)
}

func main() {
	opts := systatus.SystatusOptions{
		Prefix:    "/dev",
		ExposeEnv: true,
		HealthHandlerOpts: systatus.HealthHandlerOpts{
			Middlewares: []func(next http.HandlerFunc) http.HandlerFunc{middleware.Logger},
			Healthcheck: customHealthCheck,
		},
	}
	systatus.Enable(opts)
	log.Info().Msg("Starting server on :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
