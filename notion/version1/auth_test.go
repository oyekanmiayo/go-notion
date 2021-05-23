package version1_test

import (
	"fmt"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"net/http"
	"testing"
)

var (
	testAuthorizationURL   = "https://api.notion.com/v1/oauth/authorize?client_id=client_id&redirect_uri=http%3A%2F%2Flocalhost%3A8081&response_type=code"
	testAccessTokenReqJSON = `{"grant_type":"authorization_code","code":"1234","redirect_uri":"http://localhost:8081"}` + "\n"
	testAccessTokenResJSON = `{"access_token": "abc123def", "workspace_name": "workspace", "workspace_icon": "icon", "bot_id": "987"}`
	testAccessTokenRes     = notion.AccessTokenResponse{
		AccessToken:   "abc123def",
		WorkspaceName: "workspace",
		WorkspaceIcon: "icon",
		BotID:         "987",
	}
)

func TestAuthService_AuthorizationURL(t *testing.T) {
	config := oauth2.Config{
		ClientID:     "client_id",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.notion.com/v1/oauth/authorize",
			TokenURL: "https://api.notion.com/v1/oauth/token",
		},
		RedirectURL: "http://localhost:8081",
	}
	client := notion.AuthClient(nil)
	authURL := client.Auth.AuthorizationURL(&config, "")
	assert.Equal(t, testAuthorizationURL, authURL)
}

func TestAuthService_AccessToken(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, testAccessTokenReqJSON, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testAccessTokenResJSON)
	})

	config := oauth2.Config{
		ClientID:     "client_id",
		ClientSecret: "client_sec",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.notion.com/v1/oauth/authorize",
			TokenURL: "https://api.notion.com/v1/oauth/token",
		},
		RedirectURL: "http://localhost:8081",
	}

	client := notion.AuthClient(httpClient)
	resp, _, err := client.Auth.AccessToken(&config, "1234")
	assert.Nil(t, err)
	assert.Equal(t, &testAccessTokenRes, resp)
}
