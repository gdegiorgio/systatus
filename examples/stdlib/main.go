package main

import (
	"fmt"
	"net/http"

	"github.com/gdegiorgio/systatus"
)

func main() {
	opts := systatus.SystatusOptions{Prefix: "/dev", ExposeEnv: true}
	systatus.Enable(opts)
	fmt.Println("Starting server on :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
