package main

import (
	"net/http"
	"net/url"

	"github.com/rcrowley/go-tigertonic"
)


type MyResponse struct {
	ID    string      `json:"id"`
	Data interface{} `json:"trial"`
}

func DefaultResp() (*MyResponse) {
	return	&MyResponse {
		ID : "1",
		Data : "Trial",
		}
}

func trialHandler(u *url.URL, head http.Header, _ interface{}) (int, http.Header, *MyResponse, error) {
	defResp := DefaultResp()
	return http.StatusOK, nil, defResp, nil
}

func main() {

	mux := tigertonic.NewTrieServeMux()
	mux.Handle(
		"GET",
		"/trial",
		tigertonic.Marshaled(trialHandler),
	)
	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}
