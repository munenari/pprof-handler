package pprofhandler_test

import (
	"io"
	"io/ioutil"
	"net/http"
)

func check200(u string) bool {
	resp, err := http.Get(u)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	// io.Copy(os.Stderr, resp.Body)
	return resp.StatusCode == 200
}
