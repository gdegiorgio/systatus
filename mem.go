package systatus

import (
	"encoding/json"
	"net/http"
	"runtime"
)

type MemHandlerOpts struct {
	Middlewares []func(next http.HandlerFunc) http.HandlerFunc
}

type MemResponse struct {
	TotalAlloc uint64 `json:"total_alloc"`
	Alloc      uint64 `json:"alloc"`
	Sys        uint64 `json:"sys"`
}

func HandleMem(opts MemHandlerOpts) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &MemResponse{}
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)
		res.Sys = stats.Sys
		res.TotalAlloc = stats.TotalAlloc
		res.Alloc = stats.Alloc
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
