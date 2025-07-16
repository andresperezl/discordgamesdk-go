package discord

import (
	"context"
	"log"
	"time"
)

// ExampleStorageClient_WriteWithContext demonstrates how to use WriteWithContext with a timeout.
// This example is for documentation only and requires a real, initialized StorageClient.
func ExampleStorageClient_WriteWithContext() {
	var storageClient *StorageClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := storageClient.WriteWithContext(ctx, "file.txt", []byte("hello"))
	if err != nil {
		log.Fatalf("failed to write: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleStorageClient_ReadWithContext demonstrates how to use ReadWithContext with a timeout.
// This example is for documentation only and requires a real, initialized StorageClient.
func ExampleStorageClient_ReadWithContext() {
	var storageClient *StorageClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data, err := storageClient.ReadWithContext(ctx, "file.txt")
	if err != nil {
		log.Fatalf("failed to read: %v", err)
	}
	log.Printf("Read %d bytes", len(data))
	// No Output: (documentation only)
}

// ExampleStoreClient_FetchSkusWithContext demonstrates how to use FetchSkusWithContext with a timeout.
// This example is for documentation only and requires a real, initialized StoreClient.
func ExampleStoreClient_FetchSkusWithContext() {
	var storeClient *StoreClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	skus, err := storeClient.FetchSkusWithContext(ctx)
	if err != nil {
		log.Fatalf("failed to fetch SKUs: %v", err)
	}
	log.Printf("Fetched %d SKUs", len(skus))
	// No Output: (documentation only)
}

// ExampleStoreClient_FetchEntitlementsWithContext demonstrates how to use FetchEntitlementsWithContext with a timeout.
// This example is for documentation only and requires a real, initialized StoreClient.
func ExampleStoreClient_FetchEntitlementsWithContext() {
	var storeClient *StoreClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ents, err := storeClient.FetchEntitlementsWithContext(ctx)
	if err != nil {
		log.Fatalf("failed to fetch entitlements: %v", err)
	}
	log.Printf("Fetched %d entitlements", len(ents))
	// No Output: (documentation only)
}
