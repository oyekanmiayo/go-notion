package version1

import (
	"encoding/base64"
	"github.com/dghubble/sling"
	"golang.org/x/oauth2"
	"net/http"
)

type AuthService struct {
	sling *sling.Sling
}

// newAuthService returns a new AuthService.
func newAuthService(sling *sling.Sling) *AuthService {
	return &AuthService{
		sling: sling.Path("oauth/"),
	}
}

func (a *AuthService) AuthorizationURL(c *oauth2.Config, state string) string {
	return c.AuthCodeURL("")
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

func (a *AuthService) AccessToken(c *oauth2.Config, authCode string) (*AccessTokenResponse, *http.Response, error) {

	tokenResponse := new(AccessTokenResponse)
	apiError := new(APIError)

	tokenRequest := AccessTokenRequest{
		GrantType:   "authorization_code",
		Code:        authCode,
		RedirectURI: c.RedirectURL,
	}

	resp, err := a.sling.New().Post("token").
		BodyJSON(tokenRequest).Add("Content-Type", "application/json").
		Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.ClientID+":"+c.ClientSecret))).
		Receive(tokenResponse, apiError)

	return tokenResponse, resp, relevantError(err, *apiError)
}
