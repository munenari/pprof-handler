package pprofhandler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	pprofhandler "github.com/munenari/pprof-handler"
)

func TestGinHandler(t *testing.T) {
	t.Parallel()
	r := gin.Default()
	r.Any("/gin-router/debug/pprof/*any", gin.WrapH(pprofhandler.Handler("")))
	s := httptest.NewServer(r)
	defer s.Close()
	if success := check200(s.URL + "/gin-router/debug/pprof/heap?debug=1"); !success {
		t.Fatal("failed to get pprof")
	}
}
