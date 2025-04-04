package main

import (
	"context"
	"log"

	"github.com/xataio/xata-go/xata"
)

func main() {
	// Create a client for your workspace
	client, err := xata.NewDatabasesClient("https://your-workspace-name.xata.sh", nil)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: Fetching a list of databases
	databases, err := client.Databases.List(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch databases: %v", err)
	}

	log.Printf("Databases: %+v\n", databases)
}
