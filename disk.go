package systatus

import "net/http"

type DiskHandlerOpts struct {
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
}
type DiskResponse struct{}

func handleDisk(opts DiskHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
