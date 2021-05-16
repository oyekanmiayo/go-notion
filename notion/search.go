package notion

import "github.com/dghubble/sling"

type SearchService struct {
	sling *sling.Sling
}

// newSearchService returns a new SearchService.
func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("search/"),
	}
}
