package main

import (
	"log"
	"os"

	"github.com/mas9612/slackbase"
)

func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	if apiToken == "" {
		log.Fatalln("SLACK_API_TOKEN environment variable is not set")
	}

	slackbase.Run(apiToken)
}
