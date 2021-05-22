package main

import (
	"flag"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"os"
)

// go run oauth2-flow-part1-example.go --client-id= --url=
func main() {

	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	clientID := flags.String("client-id", "", "Client ID")
	redirectURL := flags.String("url", "", "Redirect url")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c := oauth2.Config{
		ClientID:     *clientID,
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.notion.com/v1/oauth/authorize",
			TokenURL: "https://api.notion.com/v1/oauth/token",
		},
		RedirectURL: *redirectURL,
	}

	authURL := c.AuthCodeURL("")
	fmt.Println(authURL)
}
