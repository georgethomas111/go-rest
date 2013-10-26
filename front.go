package main

import (
	"crypto/md5"
	//	"io"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"code.google.com/p/go.crypto/ssh"
	"github.com/rcrowley/go-tigertonic"
)

type keyCreateResp struct {
	CreatedAt   string `json:"created_at"`
	Email       string `json:"email"`
	FingerPrint string `json:"fingerprint"`
	ID          string `json:"id"` // To be decided
	PublicKey   string `json:"public_key"`
	UpdatedAt   string `json:"updated_at"`
}

func getKeyCreateResp(publicKey string) *keyCreateResp {
	// TODO Do action to retrieve the data.
	// How? 1. Create the server.
	// 2. Check the rpc library.
	md5Hash := md5.New()
	keyBytes := []byte(publicKey)[1:(len([]byte(publicKey)) - 1)]
	fmt.Printf("\n\nKeyBytes = %s\n\n", keyBytes)
	pubKeyIns, comment, _, _, ok := ssh.ParseAuthorizedKey(keyBytes)
	if ok {
		md5Hash.Write(ssh.MarshalAuthorizedKey(pubKeyIns))
		fmt.Printf("\nHash = %x\n", md5Hash.Sum(nil))
		fmt.Println("\n\nComment\n\n" + comment)
	}

	return &keyCreateResp{
		Email:       comment,
		FingerPrint: fmt.Sprintf("%x", md5Hash.Sum(nil)),
		PublicKey:   publicKey,
		UpdatedAt:   fmt.Sprintln(time.Now()),
		CreatedAt:   fmt.Sprintln(time.Now()), // Same as of now.
	}
}

func keyCreateHandler(u *url.URL, head http.Header, _ interface{}) (int, http.Header, *keyCreateResp, error) {
	data := u.Query().Get("public_key")
	defResp := getKeyCreateResp(data)
	return http.StatusOK, nil, defResp, nil
}

func main() {

	mux := tigertonic.NewTrieServeMux()
	mux.Handle(
		"GET",
		"/accounts/keys",
		tigertonic.Marshaled(keyCreateHandler),
	)
	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}
