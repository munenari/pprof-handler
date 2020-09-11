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
