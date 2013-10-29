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
func keyCreateHandler(u *url.URL, head http.Header, req *keyReq) (int, http.Header, *keyCreateResp, error) {
	data := req.PublicKey
	defResp := NewKeyResp(data)
	defResp.Populate()
	return http.StatusOK, nil, defResp, nil
}

type keyReq struct {
	PublicKey string `json:"public_key"`
}

type keyCreateResp struct {
	CreatedAt   string `json:"created_at"`
	Email       string `json:"email"`
	FingerPrint string `json:"fingerprint"`
	ID          string `json:"id"` // To be decided
	PublicKey   string `json:"public_key"`
	UpdatedAt   string `json:"updated_at"`
}

func NewKeyResp(pubKey string) *keyCreateResp {

	return &keyCreateResp{
		PublicKey: pubKey,
	}
}

// Populate ... Populates all the data in keyCreateResp
func (k *keyCreateResp) Populate() {

	k.PopCreatedAt()
	k.PopEmail()
	k.PopFingerPrint()
	k.PopID()
	k.PopPublicKey()
	k.PopUpdatedAt()
}

func (k *keyCreateResp) PopCreatedAt() string {

	// TODO Correct in future
	k.CreatedAt = fmt.Sprintln(time.Now())
	return k.CreatedAt
}

func (k *keyCreateResp) PopEmail() string {

	keyBytes := []byte(k.PublicKey)
	_, email, _, _, ok := ssh.ParseAuthorizedKey(keyBytes)
	if ok {
		k.Email = email
	}
	return k.Email
}

func (k *keyCreateResp) PopFingerPrint() string {

	keyBytes := []byte(k.PublicKey)
	pubKeyIns, _, _, _, ok := ssh.ParseAuthorizedKey(keyBytes)
	if ok {
		data := ssh.MarshalPublicKey(pubKeyIns)
		md5Hash := md5.New()
		md5Hash.Write(data)
		k.FingerPrint = k.cleanFingerPrint(fmt.Sprintf("%x", md5Hash.Sum(nil)))
	}
	return k.FingerPrint
}

func (k *keyCreateResp) PopID() string {

	//TODO Populate the id properly.
	k.ID = "-1"
	return k.ID
}

func (k *keyCreateResp) PopPublicKey() string {

	return k.PublicKey
}

func (k *keyCreateResp) PopUpdatedAt() string {

	k.UpdatedAt = fmt.Sprintln(time.Now())
	return k.UpdatedAt
}

func (k *keyCreateResp) cleanFingerPrint(hexStr string) string {
	var fPrint []string
	for curr := 2; curr <= len(hexStr); curr += 2 {
		fPrint = append(fPrint, hexStr[curr-2:curr])
	}
	return strings.Join(fPrint, ":")
}
