package notion

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Database struct {
	Object         string      `json:"object"`
	ID             string      `json:"id"`
	CreatedTime    string      `json:"created_time"`
	LastEditedTime string      `json:"last_edited_time"`
	Title          []RichText  `json:"title"`
	Properties     interface{} `json:"properties"`
}

type RichText struct {
}

type DatabaseService struct {
	sling *sling.Sling
}

// newDatabaseService returns a new DatabaseService.
func newDatabaseService(sling *sling.Sling) *DatabaseService {
	return &DatabaseService{
		sling: sling.Path("databases/"),
	}
}

func (d *DatabaseService) RetrieveDatabase(databaseID string) (*Database, *http.Response, error) {
	database := new(Database)
	apiError := new(APIError)
	resp, err := d.sling.New().Get(databaseID).Receive(database, apiError)

	return database, resp, relevantError(err, *apiError)
}
