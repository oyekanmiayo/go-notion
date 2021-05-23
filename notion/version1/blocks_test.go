package version1_test

import (
	"fmt"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	testRetrieveBlockChildrenResJSON = `{"object":"list","results":[{"object":"block","type":"to_do","to_do":{"text":[{"plain_text":"Sample Task","type":""}]}}]}`
	testRetrieveBlockChildrenRes     = &notion.RetrieveBlockChildrenResponse{
		Object: "list",
		Results: []notion.Block{
			{
				Object: "block",
				Type:   "to_do",
				ToDo: &notion.ToDoBlock{
					Text: []notion.RichText{
						{
							PlainText: "Sample Task",
						},
					},
				},
			},
		},
	}

	testAppendBlockChildrenBodyParamsJSON = `{"children":[{"object":"block","type":"heading_2","heading_2":{"text":[{"type":"text","text":{"content":"Header Two Test"}}]}}]}` + "\n"
	testAppendBlockChildrenBodyParams     = &notion.AppendBlockChildrenBodyParams{
		Children: []notion.Block{
			{
				Object: "block",
				Type:   "heading_2",
				HeadingTwo: &notion.HeadingTwoBlock{
					Text: []notion.RichText{
						{
							Type: "text",
							Text: &notion.Text{
								Content: "Header Two Test",
							},
						},
					},
				},
			},
		},
	}
	testAppendBlockChildrenResJSON = `
	{
	  "object": "block",
	  "id": "123",
	  "type": "child_page",
	  "created_time": "2021-05-16T14:35:46.713Z",
	  "last_edited_time": "2021-05-23T07:09:07.935Z",
	  "has_children": true,
	  "child_page": {
		"title": "Yurts in Big Surr, California 2 "
	  }
	}`
	testAppendBlockChildrenRes = &notion.Block{
		Object:         "block",
		ID:             "123",
		Type:           "child_page",
		CreatedTime:    "2021-05-16T14:35:46.713Z",
		LastEditedTime: "2021-05-23T07:09:07.935Z",
		HasChildren:    true,
		ChildPage: &notion.ChildPageBlock{
			Title: "Yurts in Big Surr, California 2 ",
		},
	}
)

func TestBlockService_RetrieveBlockChildren(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/blocks/123/children", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testRetrieveBlockChildrenResJSON)
	})

	client := notion.NewClient(httpClient, "0000")
	params := notion.RetrieveBlockChildrenParams{}
	resp, _, err := client.Blocks.RetrieveBlockChildren("123", &params)
	assert.Nil(t, err)
	assert.Equal(t, testRetrieveBlockChildrenRes, resp)
}

func TestBlockService_AppendBlockChildren(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/blocks/123/children", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PATCH", r)
		assertPostJSON(t, testAppendBlockChildrenBodyParamsJSON, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testAppendBlockChildrenResJSON)
	})

	client := notion.NewClient(httpClient, "0000")
	resp, _, err := client.Blocks.AppendBlockChildren("123", testAppendBlockChildrenBodyParams)
	assert.Nil(t, err)
	assert.Equal(t, testAppendBlockChildrenRes, resp)
}
