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
	opts := systatus.SystatusOptions{Prefix: "/dev", ExposeEnv: true, Healthcheck: customHealthCheck, Middlewares: []func(next http.HandlerFunc) http.HandlerFunc{middleware.Logger}}
	systatus.Enable(opts)
	log.Info().Msg("Starting server on :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
