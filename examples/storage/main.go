package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Storage Example ===")
	fmt.Println("This example demonstrates local storage functionality.")

	// Initialize Discord SDK
	clientID := int64(1311711649018941501) // Replace with your actual client ID
	fmt.Printf("Initializing Discord SDK with client ID: %d\n", clientID)

	core, err := discord.Create(clientID, discord.CreateFlagsDefault, nil)
	if err != discord.ResultOk {
		log.Fatalf("Failed to create Discord core: %v", err)
	}
	// Start the callback loop for robust event processing
	core.Start()
	defer core.Shutdown()

	fmt.Println("✓ Discord SDK initialized successfully")

	// Wait for user info to become available (robust pattern)
	fmt.Println("Waiting for user info...")
	user, result := core.WaitForUser(5 * time.Second)
	if result != discord.ResultOk {
		log.Fatalf("Failed to get current user: %v", result)
	}
	fmt.Printf("✓ Connected as user: %s\n", user.Username)

	// Get storage manager
	storageManager := core.GetStorageManager()
	if storageManager == nil {
		log.Fatal("Failed to get storage manager")
	}
	fmt.Println("✓ Storage manager retrieved")

	// Test data to store
	testData := []byte("Hello from Discord Go SDK!")
	testKey := "test_key"

	// Write data to storage
	fmt.Printf("Writing data to storage with key: %s\n", testKey)
	result = storageManager.Write(testKey, testData)
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to write data: %v\n", result)
	} else {
		fmt.Println("✓ Data written successfully")
	}

	// Check if the data exists
	exists, result := storageManager.Exists(testKey)
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to check if data exists: %v\n", result)
	} else {
		if exists {
			fmt.Println("✓ Data exists in storage")
		} else {
			fmt.Println("✗ Data does not exist in storage")
		}
	}

	// Read data from storage
	fmt.Printf("Reading data from storage with key: %s\n", testKey)
	readData := make([]byte, len(testData))
	bytesRead, result := storageManager.Read(testKey, readData)
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to read data: %v\n", result)
	} else {
		fmt.Println("✓ Data read successfully")
		fmt.Printf("  Original data: %s\n", string(testData))
		fmt.Printf("  Read data: %s\n", string(readData[:bytesRead]))
		if string(testData) == string(readData[:bytesRead]) {
			fmt.Println("✓ Data matches!")
		} else {
			fmt.Println("✗ Data does not match!")
		}
	}

	// Get storage statistics
	count, result := storageManager.Count()
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to get storage count: %v\n", result)
	} else {
		fmt.Printf("Total files in storage: %d\n", count)
	}

	// List storage files (if any)
	if count > 0 {
		fmt.Println("Storage files:")
		for i := int32(0); i < count; i++ {
			stat, result := storageManager.StatAt(i)
			if result == discord.ResultOk {
				fmt.Printf("  %d: %s (%d bytes, modified: %d)\n",
					i, stat.Filename, stat.Size, stat.LastModified)
			}
		}
	}

	// Get storage path
	path, result := storageManager.GetPath()
	if result == discord.ResultOk {
		fmt.Printf("Storage path: %s\n", path)
	} else {
		fmt.Printf("Warning: Failed to get storage path: %v\n", result)
	}

	// Let the callback loop process any pending operations
	fmt.Println("Processing storage operations...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 Storage example completed!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- Writing data to local storage")
	fmt.Println("- Reading data from local storage")
	fmt.Println("- Checking if data exists")
	fmt.Println("- Getting storage statistics")
	fmt.Println("- Listing storage files")
	fmt.Println("- Getting storage path")

	fmt.Println("\nPress Enter to delete test data and exit...")
	fmt.Scanln()

	// Delete the test data
	result = storageManager.Delete(testKey)
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to delete data: %v\n", result)
	} else {
		fmt.Println("✓ Test data deleted successfully")
	}

	// Let the callback loop process the deletion
	time.Sleep(500 * time.Millisecond)
}
