package version1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	testSearchPageBodyJSON = `{"query":"Jamboree","sort":{"direction":"descending","timestamp":"last_edited_time"},"filter":{"value":"page","property":"object"}}` + "\n"
	testSearchBody         = &SearchBodyParams{
		Query: "Jamboree",
		Sort: &Sort{
			Direction: "descending",
			Timestamp: "last_edited_time",
		},
	}
	testSearchPageResJSON = `{"object":"list","results":[{"object":"page","id":"5678","created_time":"2021-05-14T01:06:32.845Z","last_edited_time":"2021-05-23T08:02:00.000Z","parent":{"database_id":"38923","type":"database_id"},"properties":{"Name":{"id":"title","type":"title","title":[{"plain_text":"Jamboree"}]},"Recommended":{"id":"EZMA","type":"checkbox","checkbox":true},"Tags":{"id":"VSvn","type":"multi_select","multi_select":[{"id":"44645","name":"TagTest","color":"purple"}]}}}]}` + "\n"
	testSearchPageRes     = &SearchPageResponse{
		Object: "list",
		Results: []Page{
			*testPage,
		},
	}

	testSearchDBBodyJSON = `{"query":"Jamboree","sort":{"direction":"descending","timestamp":"last_edited_time"},"filter":{"value":"database","property":"object"}}` + "\n"
	testSearchDBResJSON  = `{"object":"list","results":[{"object":"database","id":"123","created_time":"2021-05-23T07:41:16.751Z","last_edited_time":"2021-05-23T07:41:00.000Z","title":[{"type":"text","text":{"content":"Jamboree"}}],"properties":{"Name":{"id":"title","type":"title","title":{}}}}]}` + "\n"
	testSearchDBRes      = &SearchDatabaseResponse{
		Object: "list",
		Results: []Database{
			{
				Object:         "database",
				ID:             "123",
				CreatedTime:    "2021-05-23T07:41:16.751Z",
				LastEditedTime: "2021-05-23T07:41:00.000Z",
				Title: []RichText{
					{
						Type: "text",
						Text: &Text{
							Content: "Jamboree",
						},
					},
				},
				Properties: map[string]PropertyObj{
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

	client := NewClient(httpClient, "0000")
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

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Search.SearchDatabase(testSearchBody)
	assert.Nil(t, err)
	assert.Equal(t, testSearchDBRes, resp)
}
