package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flags := flag.NewFlagSet("notion-databases-example", flag.ExitOnError)
	accessToken := flags.String("access-token", "", "Notion API Key / Notion Access Token")

	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Print(accessToken)
}
