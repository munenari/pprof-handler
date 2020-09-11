package pprofhandler

import (
	"fmt"
	"testing"
)

func TestSplitPprofName(t *testing.T) {
	t.Parallel()
	values := []struct {
		base string
		url  string
		name string
		err  error
	}{
		{
			base: "/debug/pprof/",
			url:  "stage/api/debug/pprof/",
			name: "",
			err:  nil,
		},
		{
			base: "/debug/pprof/",
			url:  "stage/api/debug/pprof/heap",
			name: "heap",
			err:  nil,
		},
		{
			base: "/debug/pprof/",
			url:  "/stage/api/debug/pprof/",
			name: "",
			err:  nil,
		},
		{
			base: "/debug/pprof/",
			url:  "/stage/api/debug/pprof/cmdline",
			name: "cmdline",
			err:  nil,
		},
		{
			base: "/debug/pprof",
			url:  "stage/api/debug/pprof/",
			name: "/",
			err:  nil,
		},
		{
			base: "/debug/pprof/test",
			url:  "stage/api/debug/pprof/",
			name: "",
			err:  ErrInvalidPath,
		},
	}
	for i, v := range values {
		vv := v
		caption := fmt.Sprintf("[%d] %s:%s", i, vv.url, vv.base)
		t.Run(caption, func(tt *testing.T) {
			tt.Parallel()
			n, err := splitPprofName(vv.url, vv.base)
			if err != vv.err {
				tt.Fatal("error was not match (err, vv.err)", err, vv.err)
			}
			if n != vv.name {
				tt.Fatal("name was not match (n, vv.name)", n, vv.name)
			}
		})
	}
}

func TestBuildHandlerPath(t *testing.T) {
	t.Parallel()
	values := []struct {
		vars   []string
		result string
	}{
		{
			vars:   []string{"debug", "pprof"},
			result: "/debug/pprof/",
		},
		{
			vars:   []string{},
			result: "",
		},
		{
			vars:   []string{""},
			result: "",
		},
		{
			vars:   []string{"/debug/pprof/"},
			result: "/debug/pprof/",
		},
		{
			vars:   []string{"debug/pprof"},
			result: "/debug/pprof/",
		},
		{
			vars:   []string{"debug", "pprof", "testing-path"},
			result: "/debug/pprof/testing-path/",
		},
	}
	for i, v := range values {
		vv := v
		caption := fmt.Sprintf("[%d] %s:%s", i, vv.vars, vv.result)
		t.Run(caption, func(tt *testing.T) {
			tt.Parallel()
			s := buildHandlerPath(vv.vars)
			if s != vv.result {
				tt.Fatal("path was not match (s, vv.result)", s, vv.result)
			}
		})
	}
}
