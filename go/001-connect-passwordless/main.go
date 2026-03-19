package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

func main() {
	endpoint := os.Getenv("COSMOS_ENDPOINT")
	if endpoint == "" {
		log.Fatalf("COSMOS_ENDPOINT environment variable is not set")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	client, err := azcosmos.NewClient(endpoint, cred, nil)
	if err != nil {
		log.Fatalf("failed to create cosmos client: %v", err)
	}

	ctx := context.Background()

	pager := client.NewQueryDatabasesPager("SELECT * FROM c", nil)
	var databases []string
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to list databases: %v", err)
		}
		for _, db := range page.Databases {
			databases = append(databases, db.ID)
		}
	}
	fmt.Printf("Connected to: %s | Databases: %d\n", client.Endpoint(), len(databases))
}
