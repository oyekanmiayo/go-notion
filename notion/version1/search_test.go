package version1_test

import (
	"fmt"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	testSearchPageBodyJSON = `{"query":"Jamboree","sort":{"direction":"descending","timestamp":"last_edited_time"},"filter":{"value":"page","property":"object"}}` + "\n"
	testSearchBody         = &notion.SearchBodyParams{
		Query: "Jamboree",
		Sort: &notion.Sort{
			Direction: "descending",
			Timestamp: "last_edited_time",
		},
	}
	testSearchPageResJSON = `{"object":"list","results":[{"object":"page","id":"5678","created_time":"2021-05-14T01:06:32.845Z","last_edited_time":"2021-05-23T08:02:00.000Z","parent":{"database_id":"38923","type":"database_id"},"properties":{"Name":{"id":"title","type":"title","title":[{"plain_text":"Jamboree"}]},"Recommended":{"id":"EZMA","type":"checkbox","checkbox":true},"Tags":{"id":"VSvn","type":"multi_select","multi_select":[{"id":"44645","name":"TagTest","color":"purple"}]}}}]}` + "\n"
	testSearchPageRes     = &notion.SearchPageResponse{
		Object: "list",
		Results: []notion.Page{
			*testPage,
		},
	}

	testSearchDBBodyJSON = `{"query":"Jamboree","sort":{"direction":"descending","timestamp":"last_edited_time"},"filter":{"value":"database","property":"object"}}` + "\n"
	testSearchDBResJSON  = `{"object":"list","results":[{"object":"database","id":"123","created_time":"2021-05-23T07:41:16.751Z","last_edited_time":"2021-05-23T07:41:00.000Z","title":[{"type":"text","text":{"content":"Jamboree"}}],"properties":{"Name":{"id":"title","type":"title","title":{}}}}]}` + "\n"
	testSearchDBRes      = &notion.SearchDatabaseResponse{
		Object: "list",
		Results: []notion.Database{
			{
				Object:         "database",
				ID:             "123",
				CreatedTime:    "2021-05-23T07:41:16.751Z",
				LastEditedTime: "2021-05-23T07:41:00.000Z",
				Title: []notion.RichText{
					{
						Type: "text",
						Text: &notion.Text{
							Content: "Jamboree",
						},
					},
				},
				Properties: map[string]notion.PropertyObj{
					"Name": {
						ID:    "title",
						Type:  "title",
						Title: map[string]interface{}{},
					},
				},
			},
		},
	}
)

func TestSearchService_SearchPage(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/search/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, testSearchPageBodyJSON, r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testSearchPageResJSON)
	})

	client := notion.NewClient(httpClient, "0000")
	resp, _, err := client.Search.SearchPage(testSearchBody)
	assert.Nil(t, err)
	assert.Equal(t, testSearchPageRes, resp)
}

func TestSearchService_SearchDatabase(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/search/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, testSearchDBBodyJSON, r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testSearchDBResJSON)
	})

	client := notion.NewClient(httpClient, "0000")
	resp, _, err := client.Search.SearchDatabase(testSearchBody)
	assert.Nil(t, err)
	assert.Equal(t, testSearchDBRes, resp)
}
