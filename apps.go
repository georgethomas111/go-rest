package main

import (
	"net/http"
	"net/url"
	"time"
)

type AppResponse struct {
	ArchivedAt        time.Time `json:"archived_at"`
	BuildPackProvider string    `json:"buildpack_provided_description"`
	CreatedAt         time.Time `json:"created_at"`
	GitUrl            string    `json:"git_url"`
	ID                string    `json:"id"`
	Maintainence      bool      `json:"maintenance"`
	Name              string    `json:"name"`
	Email             string    `json:"owner:email"`
	OwnerId           string    `json:"owner:id"`
	RegionId          string    `json:"region:id"`
	RegionName        string    `json:"region:name"`
	ReleasedAt        string    `json:"released_at"`
	RepoSize          string    `json:"repo_size"`
	SlugSize          string    `json:"slug_size"`
	StackID           string    `json:"stack:id"`
	StackName         string    `json:"stack:name"`
	UpdatedAt         time.Time `json:"updated_at"`
	WebUrl            string    `json:"web_url"`
}

type PostAppRequest struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Stack  string `json:"stack"`
}

func PostAppsHandler(u *url.URL, head http.Header, req *PostAppRequest) (int, http.Header, *AppResponse, error) {
	resp := &AppResponse{}
	resp.UpdatedAt = time.Now()
	return http.StatusOK, nil, resp, nil
}

func DeleteAppsHandler(u *url.URL, head http.Header, _ interface{}) (int, http.Header, *AppResponse, error) {
	resp := &AppResponse{}
	return http.StatusOK, nil, resp, nil
}

func GetAppHandler(u *url.URL, head http.Header, _ interface{}) (int, http.Header, *AppResponse, error) {
	data := u.Query().Get("id")
	resp := &AppResponse{}
	resp.ID = data
	return http.StatusOK, nil, resp, nil
}

func GetAppsHandler(u *url.URL, head http.Header, _ interface{}) (int, http.Header, []*AppResponse, error) {
	var respList []*AppResponse
	resp := &AppResponse{}
	respList = append(respList, resp)
	return http.StatusOK, nil, respList, nil
}
