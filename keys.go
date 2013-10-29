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
	defResp := &KeyResponse {
		PublicKey : req.PublicKey,
	}
	defResp.Populate()
	return http.StatusOK, nil, defResp, nil
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

func (k *KeyResponse) cleanFingerPrint(hexStr string) string {
	var fPrint []string
	for curr := 2; curr <= len(hexStr); curr += 2 {
		fPrint = append(fPrint, hexStr[curr-2:curr])
	}
	return strings.Join(fPrint, ":")
}
