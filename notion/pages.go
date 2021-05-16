package notion

import "github.com/dghubble/sling"

type PageService struct {
	sling *sling.Sling
}

// newPageService returns a new PageService.
func newPageService(sling *sling.Sling) *PageService {
	return &PageService{
		sling: sling.Path("pages/"),
	}
}
