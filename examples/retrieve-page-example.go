package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/oyekanmiayo/go-notion/notion"
	"log"
	"net/http"
	"os"
)

// 90b357cd32ed478b8e3dbb62b7116225

func main() {
	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	accessToken := flags.String("access-token", "", "Notion API Key / Notion Access Token")
	pageID := flags.String("page-id", "", "Page ID")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// List all users in workspace
	// Sample command: go run retrieve-page-example.go --access-token=<token> --page-id=<page-id>
	db, _, err := client.Pages.RetrievePage(*pageID)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(db)
	fmt.Println(string(jsonBody))
}
