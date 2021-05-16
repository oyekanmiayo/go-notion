package notion

import "github.com/dghubble/sling"

type UserService struct {
	sling *sling.Sling
}

// newUserService returns a new UserService.
func newUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path("users/"),
	}
}
