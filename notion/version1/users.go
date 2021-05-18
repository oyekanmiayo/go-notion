package version1

import (
	"github.com/dghubble/sling"
	"net/http"
)

type UserService struct {
	sling *sling.Sling
}

// newUserService returns a new UserService.
func newUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path("users/"),
	}
}

// https://developers.notion.com/reference/user
// Updateable properties: ID
// Display-only properties: Object, Type, Name, AvatarURL
// Object is always "user" and it
// Type is either "person" or "bot"
type User struct {
	Object    string      `json:"object,omitempty"`
	ID        string      `json:"id,omitempty"`
	Type      string      `json:"type,omitempty"`
	Name      string      `json:"name,omitempty"`
	AvatarURL string      `json:"avatar_url,omitempty"`
	Person    *Person     `json:"person,omitempty"`
	Bot       interface{} `json:"bot,omitempty"`
}

// https://developers.notion.com/reference/user#people
// Display-only properties: Person, PersonEmail
type Person struct {
	PersonEmail string `json:"email,omitempty"`
}

func (u *UserService) RetrieveUser(userID string) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := u.sling.New().Get(userID).Receive(user, apiError)

	return user, resp, relevantError(err, *apiError)
}

type ListUsersQueryParams struct {
	StartCursor string `url:"start_cursor,omitempty"`
	PageSize    int32  `url:"page_size,omitempty"`
}

type ListUsersResponse struct {
	Object     string `json:"object,omitempty"`
	Results    []User `json:"results"`
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}

// https://developers.notion.com/reference/get-users
// See https://developers.notion.com/reference/pagination to understand
// how to iterate through paginated responses
func (u *UserService) ListUsers(params *ListUsersQueryParams) (*ListUsersResponse, *http.Response, error) {
	response := new(ListUsersResponse)
	apiError := new(APIError)
	resp, err := u.sling.New().Get("").QueryStruct(params).Receive(response, apiError)

	return response, resp, relevantError(err, *apiError)
}
