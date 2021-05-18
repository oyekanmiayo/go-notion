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
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// List all DBs in a workspace
	// Sample command: go run list-databases-example.go --access-token=<token>
	params := &notion.ListDatabasesQueryParams{
		PageSize: 20,
	}
	resp, _, err := client.Databases.ListDatabases(params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(resp)
	fmt.Println(string(jsonBody))
}
