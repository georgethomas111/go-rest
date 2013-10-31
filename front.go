package main

import (
	"github.com/rcrowley/go-tigertonic"
)

func main() {

	mux := tigertonic.NewTrieServeMux()
	mux.Handle("POST", "/accounts/keys", tigertonic.Marshaled(PostKeyHandler))
	mux.Handle("GET", "/account/keys/{id}", tigertonic.Marshaled(GetKeyHandler))
	mux.Handle("DELETE", "/account/keys/{id}", tigertonic.Marshaled(DeleteKeyHandler))
	mux.Handle("GET", "/account/keys", tigertonic.Marshaled(GetKeysHandler))

	mux.Handle("POST", "/apps", tigertonic.Marshaled(PostAppsHandler))
	mux.Handle("DELETE", "/apps/{id}", tigertonic.Marshaled(DeleteAppsHandler))
	mux.Handle("GET", "/apps/{id}", tigertonic.Marshaled(GetAppHandler))
	mux.Handle("GET", "/apps", tigertonic.Marshaled(GetAppsHandler))

	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}
