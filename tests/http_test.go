package pprofhandler_test

import (
	"net/http/httptest"
	"testing"

	pprofhandler "github.com/munenari/pprof-handler"
)

func TestHttpHandler(t *testing.T) {
	t.Parallel()
	h := pprofhandler.Handler("/debug/pprof/")
	s := httptest.NewServer(h)
	defer s.Close()
	if success := check200(s.URL + "/debug/pprof/heap?debug=1"); !success {
		t.Fatal("failed to get pprof")
	}
}
