package main

import (
	"fmt"
	"net/http"

	"github.com/gdegiorgio/systatus"
)

func customHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"application is healthy"}`)
}

func main() {
	opts := systatus.SystatusOptions{Prefix: "/dev", ExposeEnv: true, Healthcheck: customHealthCheck}
	systatus.Enable(opts)
	fmt.Println("Starting server on :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
