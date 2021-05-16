package notion

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	clientIDParam     = "client_id"
	redirectURIParam  = "redirect_uri"
	responseTypeParam = "response_type"
	stateParam        = "state"

	contentType = "Content-Type"
)

type Config struct {
	Config oauth2.Config
}

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

func (c *Config) AuthorizationURL(state string) (*url.URL, error) {
	authorizationURL, err := url.Parse(c.Config.Endpoint.AuthURL)
	if err != nil {
		return nil, err
	}

	values := authorizationURL.Query()
	values.Add(clientIDParam, c.Config.ClientID)
	values.Add(redirectURIParam, c.Config.RedirectURL)
	values.Add(responseTypeParam, "code")
	values.Add(stateParam, state)

	authorizationURL.RawQuery = values.Encode()

	return authorizationURL, nil
}

func (c *Config) AccessToken(authCode string) (AccessTokenResponse, error) {
	tokenRequest := AccessTokenRequest{
		GrantType:   "authorization_code",
		Code:        authCode,
		RedirectURI: c.Config.RedirectURL,
	}
	tokenRequestJSON, err := json.Marshal(tokenRequest)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	req, err := http.NewRequest("POST", c.Config.Endpoint.TokenURL, bytes.NewBuffer(tokenRequestJSON))
	if err != nil {
		return AccessTokenResponse{}, err
	}

	req.Header.Set(contentType, "application/json")
	req.SetBasicAuth(c.Config.ClientID, c.Config.ClientSecret)

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
