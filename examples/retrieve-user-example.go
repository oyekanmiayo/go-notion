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

func main() {
	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	accessToken := flags.String("access-token", "", "Notion API Key / Notion Access Token")
	userID := flags.String("user-id", "", "User ID")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	client := notion.NewClient(http.DefaultClient, *accessToken)

	// List all users in workspace
	// Sample command: go run retrieve-user-example.go --access-token=<token> --user-id=<user-id>
	db, _, err := client.Users.RetrieveUser(*userID)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(db)
	fmt.Println(string(jsonBody))
}
