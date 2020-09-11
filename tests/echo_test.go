package pprofhandler_test

import (
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	pprofhandler "github.com/munenari/pprof-handler"
)

func TestEchoHandler(t *testing.T) {
	t.Parallel()
	e := echo.New()
	e.Any("/stage/api/debug/pprof/*", echo.WrapHandler(pprofhandler.Handler("/debug/pprof/")))
	defer e.Close()
	go e.Start("")
	for e.Listener == nil {
		time.Sleep(1)
	}
	u := "http://" + e.Listener.Addr().String() + "/stage/api/debug/pprof/goroutine?debug=1"
	if success := check200(u); !success {
		t.Fatal("failed to get pprof")
	}
}
