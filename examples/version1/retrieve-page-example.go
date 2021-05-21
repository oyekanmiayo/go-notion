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

	// Retrieve a page using its pageID
	// Sample command: go run retrieve-page-example.go --access-token=<token> --page-id=<page-id>
	page, _, err := client.Pages.RetrievePage(*pageID)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	// Print page response as json
	jsonBody, _ := json.Marshal(page)
	fmt.Println(string(jsonBody))
}
