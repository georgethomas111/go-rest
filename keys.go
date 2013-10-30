package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"code.google.com/p/go.crypto/ssh"
)

// keyCreateHandler ... This is the handler for go-tigertonic for url
// /accounts/keys
func PostKeyHandler(u *url.URL, head http.Header, req *PostKeyRequest) (int, http.Header, *KeyResponse, error) {
	defResp := &KeyResponse{
		PublicKey: req.PublicKey,
	}
	defResp.Populate()
	return http.StatusOK, nil, defResp, nil
}

func GetKeyHandler(u *url.URL, head http.Header, _ interface{}) (int, http.Header, *KeyResponse, error) {
	resp := &KeyResponse{}
	data := u.Query().Get("id")
	if len(strings.Split(data, ":")) > 1 {
		resp.FingerPrint = data
		resp.PopulateWithFingerPrint()
	} else {
		resp.ID = data
		resp.PopulateWithID()
	}

	return http.StatusOK, nil, resp, nil
}

type PostKeyRequest struct {
	PublicKey string `json:"public_key"`
}

type KeyResponse struct {
	CreatedAt   string `json:"created_at"`
	Email       string `json:"email"`
	FingerPrint string `json:"fingerprint"`
	ID          string `json:"id"` // To be decided
	PublicKey   string `json:"public_key"`
	UpdatedAt   string `json:"updated_at"`
}

// Populate ... Populates all the data in keyCreateResp
func (k *KeyResponse) Populate() {

	k.CreatedAt = fmt.Sprintln(time.Now())

	keyBytes := []byte(k.PublicKey)
	pubKeyIns, email, _, _, ok := ssh.ParseAuthorizedKey(keyBytes)
	if ok {
		k.Email = email
	}

	if ok {
		data := ssh.MarshalPublicKey(pubKeyIns)
		md5Hash := md5.New()
		md5Hash.Write(data)
		k.FingerPrint = k.cleanFingerPrint(fmt.Sprintf("%x", md5Hash.Sum(nil)))
	}

	k.ID = "-1"

	k.UpdatedAt = fmt.Sprintln(time.Now())
}

func (k *KeyResponse) PopulateWithID() {
	k.CreatedAt = fmt.Sprintln(time.Now())
	k.UpdatedAt = fmt.Sprintln(time.Now())
	k.Email = "user@server"
	k.PublicKey = "pubkey as in ~/.ssh/id_rsa.pub"
	k.FingerPrint = "aa:bb:cc:dd:ee:ff:gg:hh:ii:kk:ll:mm:nn:oo:pp:qq"
}

func (k *KeyResponse) PopulateWithFingerPrint() {
	// Do the same as Populate With ID for now.
	k.PopulateWithID()
}

func (k *KeyResponse) cleanFingerPrint(hexStr string) string {
	var fPrint []string
	for curr := 2; curr <= len(hexStr); curr += 2 {
		fPrint = append(fPrint, hexStr[curr-2:curr])
	}
	return strings.Join(fPrint, ":")
}
