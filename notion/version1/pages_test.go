package version1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//TODO: Clean up tests. Extract reused names.
var (
	testPageJSON = `{"object":"page","id":"5678","created_time":"2021-05-14T01:06:32.845Z","last_edited_time":"2021-05-23T08:02:00.000Z","parent":{"database_id":"38923","type":"database_id"},"properties":{"Name":{"id":"title","type":"title","title":[{"plain_text":"Jamboree"}]},"Recommended":{"id":"EZMA","type":"checkbox","checkbox":true},"Tags":{"id":"VSvn","type":"multi_select","multi_select":[{"id":"44645","name":"TagTest","color":"purple"}]}}}` + "\n"
	testPage     = &Page{
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
	}
	testRetrievePageResJSON = testPageJSON
	testRetrievePageRes     = testPage

	testPageProperty = map[string]PageProperty{
		"Name": {
			Title: []RichText{
				{
					Text: &Text{
						Content: "Jamboree",
					},
				},
			},
		},
		"Tags": {
			MultiSelect: []MultiSelectPropertyOpts{
				{
					Name: "TagTest",
				},
			},
		},
		"Recommended": {
			Checkbox: true,
		},
	}

	testCreatePageBodyJSON = `{"parent":{"database_id":"7d6410f1-0c2d-4c75-8199-3fd7d90ff4ff"},"properties":{"Name":{"title":[{"text":{"content":"Jamboree"}}]},"Recommended":{"checkbox":true},"Tags":{"multi_select":[{"name":"TagTest"}]}}}` + "\n"
	testCreatePageBody     = &CreatePageBodyParams{
		Parent: &DatabaseParent{
			DatabaseID: "7d6410f1-0c2d-4c75-8199-3fd7d90ff4ff",
		},
		Properties: testPageProperty,
	}
	testCreatePageResJSON = testPageJSON
	testCreatePageRes     = testPage

	testUpdatePropertiesBodyJSON = `{"properties":{"Name":{"title":[{"text":{"content":"Jamboree"}}]},"Recommended":{"checkbox":true},"Tags":{"multi_select":[{"name":"TagTest"}]}}}` + "\n"
	testUpdatePropertiesBody     = &UpdatePagePropertiesBodyParams{
		Properties: testPageProperty,
	}
	testUpdatePropertiesResJSON = testPageJSON
	testUpdatePropertiesRes     = testPage
)

func TestPageService_RetrievePage(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/pages/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testRetrievePageResJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Pages.RetrievePage("123")
	assert.Nil(t, err)
	assert.Equal(t, testRetrievePageRes, resp)
}

func TestPageService_CreatePage(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/pages/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, testCreatePageBodyJSON, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testCreatePageResJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Pages.CreatePage(testCreatePageBody)
	assert.Nil(t, err)
	assert.Equal(t, testCreatePageRes, resp)
}

func TestPageService_UpdatePageProperties(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/pages/123", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		assertPostJSON(t, testUpdatePropertiesBodyJSON, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testUpdatePropertiesResJSON)
	})

	client := NewClient(httpClient, "0000")
	resp, _, err := client.Pages.UpdatePageProperties("123", testUpdatePropertiesBody)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatePropertiesRes, resp)
}
