package main

import (
	"flag"
	"fmt"
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

// go run oauth2-flow-part2-example.go --client-id= --client-secret= --code= --url=
func main() {

	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	clientID := flags.String("client-id", "", "Client ID")
	clientSecret := flags.String("client-secret", "", "Client ID")
	code := flags.String("code", "", "Auth Code from Notion")
	redirectURL := flags.String("url", "", "Redirect url")
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
		RedirectURL: *redirectURL,
	}

	fmt.Println("Hello")
	client := notion.AuthClient(http.DefaultClient)

	// This code is sent in as a query param to the redirect_uri after the user has authorized notion
	//<redirect_uri>?code=<code>&state=<state>
	resp, _, _ := client.Auth.AccessToken(&c, *code)
	fmt.Println(resp)
}
