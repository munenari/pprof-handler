package pprofhandler

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"
)

var (
	// ErrInvalidPath means that path could not parse to pprof name
	ErrInvalidPath = fmt.Errorf("invalid debug pprof path")
)

// Handler returns an HTTP handler that serves with base path
func Handler(basePath string) http.Handler {
	return handler(basePath)
}

type handler string

func (x handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pprofPath := string(x)
	if pprofPath == "" {
		pprofPath = "/debug/pprof/"
	}
	name, err := splitPprofName(r.URL.Path, pprofPath)
	if err != nil {
		serveError(w, http.StatusBadRequest, err.Error())
		return
	}
	switch name {
	case "":
		pprof.Index(w, r)
	case "cmdline":
		pprof.Cmdline(w, r)
	case "profile":
		pprof.Profile(w, r)
	case "trace":
		pprof.Trace(w, r)
	case "symbol":
		pprof.Symbol(w, r)
	default:
		pprof.Handler(name).ServeHTTP(w, r)
	}
}

// serveError respect for net/http/pprof
func serveError(w http.ResponseWriter, status int, txt string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Go-Pprof", "1")
	w.Header().Del("Content-Disposition")
	w.WriteHeader(status)
	fmt.Fprintln(w, txt)
}

func splitPprofName(pathStr, basePath string) (name string, err error) {
	splitted := strings.SplitN(pathStr, basePath, 2)
	if len(splitted) != 2 {
		err = ErrInvalidPath
		return
	}
	name = splitted[1]
	return
}
