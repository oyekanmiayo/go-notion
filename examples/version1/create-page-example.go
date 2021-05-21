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
	databaseID := flags.String("db-id", "", "DB ID")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// Create a page in a database
	// Sample command: go run create-page-example.go --access-token=<token> --page-id=<page-id> --db-id=<db-id>
	params := &notion.CreatePageBodyParams{
		Parent: &notion.DatabaseParent{
			DatabaseID: *databaseID,
		},
		Properties: map[string]notion.PageProperty{
			"Name": {
				Title: []notion.RichText{
					{
						Text: &notion.Text{
							Content: "Creating Page Sample",
						},
					},
				},
			},
			"Tags": {
				MultiSelect: []notion.MultiSelectPropertyOpts{
					{
						Name: "Tag1",
					},
					{
						Name: "Tag3",
					},
				},
			},
			"Recommended": {
				Checkbox: true,
			},
		},
	}
	resp, _, err := client.Pages.CreatePage(params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(resp)
	fmt.Println(string(jsonBody))

	// Create a page within a page
	params = &notion.CreatePageBodyParams{
		Parent: &notion.PageParent{
			PageID: *pageID,
		},
		Properties: &notion.PageProperty{
			Title: []notion.RichText{
				{
					Text: &notion.Text{
						Content: "Creating Page Within A Page Sample",
					},
				},
			},
		},
	}

	resp, _, err = client.Pages.CreatePage(params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ = json.Marshal(resp)
	fmt.Println(string(jsonBody))
}
