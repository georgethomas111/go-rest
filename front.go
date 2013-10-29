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
	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}
