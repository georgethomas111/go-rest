package main

import (
	"github.com/rcrowley/go-tigertonic"
)

func main() {

	mux := tigertonic.NewTrieServeMux()
	mux.Handle(
		"POST",
		"/accounts/keys",
		tigertonic.Marshaled(PostKeyHandler),
	)
	mux.Handle(
		"GET",
		"/account/keys/{id}",
		tigertonic.Marshaled(GetKeyHandler),
	)
	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}
