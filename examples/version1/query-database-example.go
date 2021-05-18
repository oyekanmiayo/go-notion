package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	notion "github.com/oyekanmiayo/go-notion/notion/version1"
)

func main() {
	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	accessToken := flags.String("access-token", "", "Notion API Key / Notion Access Token")
	databaseID := flags.String("db-id", "", "Notion Database ID")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// Query DB with SingleFilter
	// Sample command: go run query-database-example.go --access-token=<token> --db-id=<database-id>
	params := &notion.QueryDatabaseBodyParams{
		Filter: &notion.SingleFilter{
			Property: "Tags",
			MultiSelect: &notion.MultiSelectCondition{
				Contains: "Tag1",
			},
		},
	}
	resp, _, err := client.Databases.QueryDatabase(*databaseID, params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(resp)
	fmt.Println(string(jsonBody))

	fmt.Println()

	// Query DB with CompoundFilter
	params = &notion.QueryDatabaseBodyParams{
		Filter: &notion.CompoundFilter{
			AND: []notion.SingleFilter{
				{
					Property: "Tags",
					MultiSelect: &notion.MultiSelectCondition{
						Contains: "Tag1",
					},
				},
				{
					Property: "Recommended",
					Checkbox: &notion.CheckboxCondition{
						Equals: true,
					},
				},
			},
		},
	}
	resp, _, err = client.Databases.QueryDatabase(*databaseID, params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ = json.Marshal(resp)
	fmt.Println(string(jsonBody))
}
