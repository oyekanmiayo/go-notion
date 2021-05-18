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
	blockID := flags.String("block-id", "", "Page ID")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// List all users in workspace
	// Sample command: go run retrieve-page-example.go --access-token=<token> --block-id=<block-id>
	params := &notion.RetrieveBlockChildrenParams{}
	db, _, err := client.Blocks.RetrieveBlockChildren(*blockID, params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(db)
	fmt.Println(string(jsonBody))
}
