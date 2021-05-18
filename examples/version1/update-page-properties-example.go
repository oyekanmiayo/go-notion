package main

import (
	"encoding/json"
	"flag"
	"fmt"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"log"
	"net/http"
	"os"
)

func main() {
	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	accessToken := flags.String("access-token", "", "Notion API Key / Notion Access Token")
	pageID := flags.String("page-id", "", "Page ID")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// Sample command: go run update-page-properties-example.go --access-token=<token> --page-id=<page-id>
	params := &notion.UpdatePagePropertiesBodyParams{
		// first keys are the names or ids of the properties :)
		// id for title is "title"
		// See examples here: https://developers.notion.com/reference/page#page-property-value
		Properties: map[string]map[string]interface{}{
			"Name": {
				"title": []notion.RichText{
					{
						Type: "text",
						Text: &notion.Text{
							Content: "Jamaican Cuisines II",
						},
					},
				},
			},
			"Recommended": {
				"checkbox": false,
			},
		},
	}
	resp, _, err := client.Pages.UpdatePageProperties(*pageID, params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(resp)
	fmt.Println(string(jsonBody))
}
