package version1

import (
	"github.com/dghubble/sling"
	"net/http"
)

const notionAPI = "https://api.notion.com/v1/"
const notionVersion = "2021-05-13"

type Client struct {
	sling *sling.Sling

	// Notion API Services
	Users     *UserService
	Databases *DatabaseService
	Pages     *PageService
	Blocks    *BlockService
	Search    *SearchService
}

func NewClient(client *http.Client, accessToken string) *Client {
	base := sling.New().Client(client).Base(notionAPI)
	base.Add("Authorization", "Bearer "+accessToken)
	base.Add("Notion-Version", notionVersion)

	return &Client{
		sling:     base,
		Users:     newUserService(base.New()),
		Databases: newDatabaseService(base.New()),
		Pages:     newPageService(base.New()),
		Blocks:    newBlockService(base.New()),
		Search:    newSearchService(base.New()),
	}
}
