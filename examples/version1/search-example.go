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
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// Search the workspace
	// Sample command: go run search-example.go --access-token=<token>
	params := &notion.SearchBodyParams{
		Query: "Yurts",
		Sort: &notion.Sort{
			Direction: "descending",
			Timestamp: "last_edited_time",
		},
		Filter: &notion.SearchFilter{
			Value:    "page",
			Property: "object",
		},
	}
	db, _, err := client.Search.Search(params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(db)
	fmt.Println(string(jsonBody))
}
