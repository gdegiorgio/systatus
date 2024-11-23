package systatus

import "net/http"

type CPUHandlerOpts struct {
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
}
type CPUResponse struct{}

func handleCPU(opts CPUHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
