package main

import (
	"flag"
	"fmt"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"golang.org/x/oauth2"
	"log"
	"os"
)

// go run oauth2-flow-part1-example.go --client-id= --client-secret= --code=
func main() {

	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	clientID := flags.String("client-id", "", "Client ID")
	clientSecret := flags.String("client-secret", "", "Client ID")
	code := flags.String("code", "", "Auth Code from Notion")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c := oauth2.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.notion.com/v1/oauth/authorize",
			TokenURL: "https://api.notion.com/v1/oauth/token",
		},
		RedirectURL: "http://a8bc9f95180e.ngrok.io",
	}

	authURL := c.AuthCodeURL("sample-state")
	fmt.Println(authURL)

	resp, _ := notion.AccessToken(&c, *code)
	fmt.Println(resp)
}
