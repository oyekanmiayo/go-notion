package version1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	testUserJSON = `{"object":"user","id":"123abc","type":"person","name":"John Doe","avatar_url":"https://test.com/test","person":{"email":"test@gmail.com"}}` + "\n"
	testUser     = &User{
		Object:    "user",
		ID:        "123abc",
		Type:      "person",
		Name:      "John Doe",
		AvatarURL: "https://test.com/test",
		Person: &Person{
			PersonEmail: "test@gmail.com",
		},
	}

	testBot = &User{
		Object: "user",
		ID:     "456def",
		Type:   "bot",
		Name:   "Test Integration",
		Bot:    map[string]interface{}{},
	}

	testListUsersJSON = `{"object":"list","results":[{"object":"user","id":"123abc","type":"person","name":"John Doe","avatar_url":"https://test.com/test","person":{"email":"test@gmail.com"}},{"object":"user","id":"456def","type":"bot","name":"Test Integration","bot":{}}]}` + "\n"
	testListUsers     = &ListUsersResponse{
		Object: "list",
		Results: []User{
			*testUser,
			*testBot,
		},
	}
)

func TestUserService_RetrieveUser(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/users/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testUserJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Users.RetrieveUser("123")
	assert.Nil(t, err)
	assert.Equal(t, testUser, resp)
}

func TestUserService_ListUsers(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/users/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testListUsersJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Users.ListUsers(&ListUsersQueryParams{})
	assert.Nil(t, err)
	assert.Equal(t, testListUsers, resp)
}
