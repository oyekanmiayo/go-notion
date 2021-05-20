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

	// Retrieve DB
	// Sample command: go run retrieve-database-example.go --access-token=<token> --db-id=<database-id>
	db, _, err := client.Databases.RetrieveDatabase(*databaseID)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	// Print column names and type for DB table
	for k, v := range db.Properties {
		fmt.Printf("%v, %v", k, v.Type)
		fmt.Println()
	}

	// View response in json
	jsonBody, _ := json.Marshal(db)
	fmt.Println(string(jsonBody))
}
