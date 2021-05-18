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

	// Append block children to a block
	// Sample command: go run append-block-children-example.go --access-token=<token> --block-id=<block-id>
	params := &notion.AppendBlockChildrenBodyParams{
		Children: []notion.Block{
			{
				Object: "block",
				Type:   "heading_2",
				HeadingTwo: notion.HeadingTwo{
					Text: []notion.RichText{
						{
							Type: "text",
							Text: &notion.Text{
								Content: "Header Two Test",
							},
						},
					},
				},
			},
			{
				Object: "block",
				Type:   "paragraph",
				Paragraph: notion.Paragraph{
					Text: []notion.RichText{
						{
							Type: "text",
							Text: &notion.Text{
								Content: "Paragraph Test",
							},
						},
					},
				},
			},
		},
	}

	db, _, err := client.Blocks.AppendBlockChildren(*blockID, params)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	jsonBody, _ := json.Marshal(db)
	fmt.Println()
	fmt.Println(string(jsonBody))
}
