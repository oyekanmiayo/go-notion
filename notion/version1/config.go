package version1

import (
	"bytes"
	"encoding/json"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

const (
	clientIDParam     = "client_id"
	redirectURIParam  = "redirect_uri"
	responseTypeParam = "response_type"
	stateParam        = "state"

	contentType = "Content-Type"
)

type Config oauth2.Config

type AccessTokenRequest struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type AccessTokenResponse struct {
	AccessToken   string `json:"access_token"`
	WorkspaceName string `json:"workspace_name"`
	WorkspaceIcon string `json:"workspace_icon"`
	BotID         string `json:"bot_id"`
}

func AccessToken(c *oauth2.Config, authCode string) (AccessTokenResponse, error) {
	tokenRequest := AccessTokenRequest{
		GrantType:   "authorization_code",
		Code:        authCode,
		RedirectURI: c.RedirectURL,
	}
	tokenRequestJSON, err := json.Marshal(tokenRequest)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	req, err := http.NewRequest("POST", c.Endpoint.TokenURL, bytes.NewBuffer(tokenRequestJSON))
	if err != nil {
		return AccessTokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.ClientID, c.ClientSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	defer resp.Body.Close()

	var tokenResponse AccessTokenResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	return tokenResponse, nil
}
