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

	// Search the workspace for pages with titles that contain this
	// Sample command: go run search-example.go --access-token=<token>
	params := &notion.SearchBodyParams{
		Query: "Yurts",
		Sort: &notion.Sort{
			Direction: "descending",
			Timestamp: "last_edited_time",
		},
	}
	res, _, err := client.Search.SearchPage(params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(res)
	fmt.Println(string(jsonBody))

	fmt.Println("<==============================>")

	// Search the workspace for databases with titles that contain this
	paramsII := &notion.SearchBodyParams{
		Query: "Yurts",
		Sort: &notion.Sort{
			Direction: "descending",
			Timestamp: "last_edited_time",
		},
	}
	resII, _, err := client.Search.SearchDatabase(paramsII)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBodyII, _ := json.Marshal(resII)
	fmt.Println(string(jsonBodyII))
}
