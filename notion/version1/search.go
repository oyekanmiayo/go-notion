package version1

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

// From my tests it looks like the api checks if page titles contain the query
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

// Value is either "page" or "database"
// Property can only be "object"
type SearchFilter struct {
	Value    string `json:"value,omitempty"`
	Property string `json:"property,omitempty"`
}

// TODO: Find a way to unify responses and make them accessible.
// Use json.RawMessage & https://play.golang.org/p/e6kvxtOeTCc
// One approach: Unmarshal into a random struct to get the type of
// each entry, and use that to Unnmarshal into the right struct.

type SearchPageResponse struct {
	Object     string `json:"object,omitempty"`
	Results    []Page `json:"results,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more,omitempty"`
}

func (s *SearchService) SearchPage(params *SearchBodyParams) (*SearchPageResponse, *http.Response, error) {
	sResponse := new(SearchPageResponse)
	apiError := new(APIError)

	params.Filter = &SearchFilter{
		Property: "object",
		Value:    "page",
	}

	httpResponse, httpError := s.sling.New().Post("").BodyJSON(params).Receive(sResponse, apiError)

	return sResponse, httpResponse, relevantError(httpError, *apiError)
}

type SearchDatabaseResponse struct {
	Object     string     `json:"object,omitempty"`
	Results    []Database `json:"results,omitempty"`
	NextCursor string     `json:"next_cursor,omitempty"`
	HasMore    bool       `json:"has_more,omitempty"`
}

func (s *SearchService) SearchDatabase(params *SearchBodyParams) (*SearchDatabaseResponse, *http.Response, error) {
	sResponse := new(SearchDatabaseResponse)
	apiError := new(APIError)

	params.Filter = &SearchFilter{
		Property: "object",
		Value:    "database",
	}
	httpResponse, httpError := s.sling.New().Post("").BodyJSON(params).Receive(sResponse, apiError)

	return sResponse, httpResponse, relevantError(httpError, *apiError)
}
