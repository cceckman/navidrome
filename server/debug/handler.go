// Package debug sets up HTTP handlers for debugging and detailed analysis,
// beyond Prometheus metrics.
package debug

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

func NewHandler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(resp, "Available debug handlers: /pprof")
	})
	r.HandleFunc("/pprof", pprof.Index)
	return r
}
