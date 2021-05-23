package version1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	testDB = &Database{
		Object:         "database",
		ID:             "123",
		CreatedTime:    "2021-05-23T07:41:16.751Z",
		LastEditedTime: "2021-05-23T07:41:00.000Z",
		Title: []RichText{
			{
				PlainText: "TestDB",
				Annotations: &Annotations{
					Color: "default",
				},
				Type: "text",
				Text: &Text{
					Content: "TestDB",
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
	}

	testDatabaseResJSON = `{"object":"database","id":"123","created_time":"2021-05-23T07:41:16.751Z","last_edited_time":"2021-05-23T07:41:00.000Z","title":[{"plain_text":"TestDB","annotations":{"color":"default"},"type":"text","text":{"content":"TestDB"}}],"properties":{"Name":{"id":"title","type":"title","title":{}}}}` + "\n"
	testDatabaseRes     = testDB

	testSingleFilterParamsJSON = `{"filter":{"property":"Tags","multi_select":{"contains":"TagTest"}}}` + "\n"
	testSingleFilterParams     = &QueryDatabaseBodyParams{
		Filter: &SingleFilter{
			Property: "Tags",
			MultiSelect: &MultiSelectCondition{
				Contains: "TagTest",
			},
		},
	}
	testFilterResJSON = `{"object":"list","results":[{"object":"page","id":"5678","created_time":"2021-05-14T01:06:32.845Z","last_edited_time":"2021-05-23T08:02:00.000Z","parent":{"database_id":"38923","type":"database_id"},"properties":{"Name":{"id":"title","type":"title","title":[{"plain_text":"Jamboree"}]},"Recommended":{"id":"EZMA","type":"checkbox","checkbox":true},"Tags":{"id":"VSvn","type":"multi_select","multi_select":[{"id":"44645","name":"TagTest","color":"purple"}]}}}]}` + "\n"
	testFilterRes     = &QueryDatabaseResponse{
		Object: "list",
		Results: []Page{
			{
				Object:         "page",
				ID:             "5678",
				CreatedTime:    "2021-05-14T01:06:32.845Z",
				LastEditedTime: "2021-05-23T08:02:00.000Z",
				// Technically this is &DatabaseParent but Go will see type as map[string]interface{}
				// Haven't found a workaround yet :)
				Parent: map[string]interface{}{
					"database_id": "38923",
					"type":        "database_id",
				},
				Properties: map[string]PageProperty{
					"Name": {
						ID:   "title",
						Type: "title",
						Title: []RichText{
							{
								PlainText: "Jamboree",
							},
						},
					},
					"Recommended": {
						ID:       "EZMA",
						Type:     "checkbox",
						Checkbox: true,
					},
					"Tags": {
						ID:   "VSvn",
						Type: "multi_select",
						MultiSelect: []MultiSelectPropertyOpts{
							{
								ID:    "44645",
								Name:  "TagTest",
								Color: "purple",
							},
						},
					},
				},
			},
		},
	}

	testCompFilterParamsJSON = `{"filter":{"and":[{"property":"Tags","multi_select":{"contains":"TagTest"}},{"property":"Name","title":{"contains":"Jamboree"}}]}}` + "\n"
	testCompFilterParams     = &QueryDatabaseBodyParams{
		Filter: &CompoundFilter{
			AND: []SingleFilter{
				{
					Property: "Tags",
					MultiSelect: &MultiSelectCondition{
						Contains: "TagTest",
					},
				},
				{
					Property: "Name",
					Title: &TextCondition{
						Contains: "Jamboree",
					},
				},
			},
		},
	}

	testListDBResJSON = `{"results":[{"object":"database","id":"f170fb7c-af53-4173-a0b2-c4cfc8da7789","created_time":"2021-05-23T07:41:16.751Z","last_edited_time":"2021-05-23T08:04:00.000Z","title":[{"plain_text":"DB for Testing"}],"properties":{"Name":{"id":"title","type":"title","title":{}},"Tags":{"id":"M}hZ","type":"multi_select","multi_select":{"options":[{"name":"TagTest","id":"f3bbc99e-f70b-4585-a1a4-f3f895e01104","color":"gray"}]}}}}]}` + "\n"
	testListDBRes     = &ListDatabasesResponse{
		Results: []Database{
			{
				Object:         "database",
				ID:             "f170fb7c-af53-4173-a0b2-c4cfc8da7789",
				CreatedTime:    "2021-05-23T07:41:16.751Z",
				LastEditedTime: "2021-05-23T08:04:00.000Z",
				Title: []RichText{
					{
						PlainText: "DB for Testing",
					},
				},
				Properties: map[string]PropertyObj{
					"Name": {
						ID:    "title",
						Type:  "title",
						Title: map[string]interface{}{},
					},
					"Tags": {
						ID:   "M}hZ",
						Type: "multi_select",
						MultiSelect: &MultiSelectConfig{
							Options: []MultiSelectOption{
								{
									Name:  "TagTest",
									ID:    "f3bbc99e-f70b-4585-a1a4-f3f895e01104",
									Color: "gray",
								},
							},
						},
					},
				},
			},
		},
	}
)

func TestDatabaseService_RetrieveDatabase(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/databases/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDatabaseResJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Databases.RetrieveDatabase("123")
	assert.Nil(t, err)
	assert.Equal(t, testDatabaseRes, resp)
}

func TestDatabaseService_QueryDatabase_SingleFilter(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/databases/123/query", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, testSingleFilterParamsJSON, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testFilterResJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Databases.QueryDatabase("123", testSingleFilterParams)
	assert.Nil(t, err)
	assert.Equal(t, testFilterRes, resp)

}

func TestDatabaseService_QueryDatabase_CompoundFilter(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/databases/123/query", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, testCompFilterParamsJSON, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testFilterResJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Databases.QueryDatabase("123", testCompFilterParams)
	assert.Nil(t, err)
	assert.Equal(t, testFilterRes, resp)

}

func TestDatabaseService_ListDatabases(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/databases/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testListDBResJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Databases.ListDatabases(&ListDatabasesQueryParams{})
	assert.Nil(t, err)
	assert.Equal(t, testListDBRes, resp)
}
