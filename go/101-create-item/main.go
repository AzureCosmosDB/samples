package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type Product struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Sale     bool   `json:"sale"`
}

func main() {
	endpoint := os.Getenv("COSMOS_ENDPOINT")
	if endpoint == "" {
		log.Fatalf("COSMOS_ENDPOINT environment variable is not set")
	}
	key := os.Getenv("COSMOS_KEY")
	if key == "" {
		log.Fatalf("COSMOS_KEY environment variable is not set")
	}

	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		log.Fatalf("failed to create key credential: %v", err)
	}

	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		log.Fatalf("failed to create cosmos client: %v", err)
	}

	ctx := context.Background()

	// Create database (ignore conflict if already exists)
	_, err = client.CreateDatabase(ctx, azcosmos.DatabaseProperties{ID: "cosmicworks"}, nil)
	if err != nil && !isConflict(err) {
		log.Fatalf("failed to create database: %v", err)
	}
	fmt.Println("Created/retrieved database: cosmicworks")

	dbClient, err := client.NewDatabase("cosmicworks")
	if err != nil {
		log.Fatalf("failed to get database client: %v", err)
	}

	// Create container (ignore conflict if already exists)
	_, err = dbClient.CreateContainer(ctx, azcosmos.ContainerProperties{
		ID: "products",
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{"/category"},
		},
	}, nil)
	if err != nil && !isConflict(err) {
		log.Fatalf("failed to create container: %v", err)
	}
	fmt.Println("Created/retrieved container: products")

	containerClient, err := client.NewContainer("cosmicworks", "products")
	if err != nil {
		log.Fatalf("failed to get container client: %v", err)
	}

	product := Product{
		ID:       "70b63682-b93a-4c77-aad2-65501347265f",
		Category: "gear-surf-surfboards",
		Name:     "Yamba Surfboard",
		Quantity: 12,
		Sale:     false,
	}
	pk := azcosmos.NewPartitionKeyString(product.Category)

	// CREATE
	itemBytes, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("failed to marshal product: %v", err)
	}
	_, err = containerClient.UpsertItem(ctx, pk, itemBytes, nil)
	if err != nil {
		log.Fatalf("failed to upsert item: %v", err)
	}
	fmt.Printf("Upserted item: %s\n", product.ID)

	// READ
	readResp, err := containerClient.ReadItem(ctx, pk, product.ID, nil)
	if err != nil {
		log.Fatalf("failed to read item: %v", err)
	}
	var readProduct Product
	if err = json.Unmarshal(readResp.Value, &readProduct); err != nil {
		log.Fatalf("failed to unmarshal read response: %v", err)
	}
	fmt.Printf("Read item: %s | Quantity: %d\n", readProduct.Name, readProduct.Quantity)

	// UPDATE
	readProduct.Quantity = 20
	updatedBytes, err := json.Marshal(readProduct)
	if err != nil {
		log.Fatalf("failed to marshal updated product: %v", err)
	}
	_, err = containerClient.ReplaceItem(ctx, pk, readProduct.ID, updatedBytes, nil)
	if err != nil {
		log.Fatalf("failed to replace item: %v", err)
	}
	fmt.Printf("Updated item: %s | New quantity: %d\n", readProduct.ID, readProduct.Quantity)

	// DELETE
	_, err = containerClient.DeleteItem(ctx, pk, product.ID, nil)
	if err != nil {
		log.Fatalf("failed to delete item: %v", err)
	}
	fmt.Printf("Deleted item: %s\n", product.ID)
}

func isConflict(err error) bool {
	var responseErr *azcore.ResponseError
	return errors.As(err, &responseErr) && responseErr.StatusCode == http.StatusConflict
}
