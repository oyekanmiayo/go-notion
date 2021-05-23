package main

import (
	"encoding/json"
	"fmt"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
)

func main() {

	params := &notion.SearchDatabaseResponse{
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

	res, _ := json.Marshal(params)
	fmt.Println(string(res))
}
