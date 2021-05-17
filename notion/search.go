package notion

import (
	"github.com/dghubble/sling"
	"net/http"
)

type SearchService struct {
	sling *sling.Sling
}

// newSearchService returns a new SearchService.
func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("search/"),
	}
}

type SearchBodyParams struct {
	Query       string        `json:"query,omitempty"`
	Sort        *Sort         `json:"sort,omitempty"`
	Filter      *SearchFilter `json:"filter,omitempty"`
	StartCursor string        `json:"start_cursor,omitempty"`
	PageSize    int32         `json:"page_size,omitempty"`
}

// Direction is either "ascending" or "descending"
// Timestamp is always "last_edited_time"
type Sort struct {
	Direction string `json:"direction,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// This is was CONFUSING, I'm not sure about this one (:
// Value is either "page" or "database"
type SearchFilter struct {
	Value    string `json:"value,omitempty"`
	Property string `json:"property,omitempty"`
}

// From what I understand atm, Results can contains an array with Page and Database objects
type SearchResponse struct {
	Object     string      `json:"object,omitempty"`
	Results    interface{} `json:"results,omitempty"`
	NextCursor string      `json:"next_cursor,omitempty"`
	HasMore    bool        `json:"has_more,omitempty"`
}

// https://developers.notion.com/reference/post-search
func (s *SearchService) Search(params *SearchBodyParams) (*SearchResponse, *http.Response, error) {
	sResponse := new(SearchResponse)
	apiError := new(APIError)
	httpResponse, httpError := s.sling.New().Post("").BodyJSON(params).Receive(sResponse, apiError)

	return sResponse, httpResponse, relevantError(httpError, *apiError)
}
