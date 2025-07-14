package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Storage Example ===")
	fmt.Println("This example demonstrates storage management using the new Go-like Client wrapper.")

	// Initialize Discord SDK with the new Client wrapper
	clientID := int64(1311711649018941501) // Replace with your actual client ID
	fmt.Printf("Initializing Discord SDK with client ID: %d\n", clientID)

	config := discord.DefaultClientConfig(clientID)
	client, err := discord.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Discord client: %v", err)
	}
	defer client.Close()

	fmt.Println("✓ Discord SDK initialized successfully")

	// Get current user
	user, err := client.GetCurrentUser(5 * time.Second)
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
	fmt.Printf("✓ Connected as user: %s\n", user.Username)

	// Get storage manager
	storageClient := client.Storage()

	// Test storage operations
	fmt.Println("\n=== Testing Storage Operations ===")

	// Get storage count
	fmt.Println("Getting storage count...")
	count, err := storageClient.Count()
	if err != nil {
		fmt.Printf("⚠ Failed to get storage count: %v\n", err)
	} else {
		fmt.Printf("✓ Storage count: %d\n", count)
	}

	// Test writing data
	fmt.Println("\n=== Testing Write Operations ===")

	testData := []byte("Hello from Discord Go SDK!")
	testFileName := "test_file.txt"

	fmt.Printf("Writing data to '%s'...\n", testFileName)
	err = storageClient.Write(testFileName, testData)
	if err != nil {
		fmt.Printf("⚠ Failed to write data: %v\n", err)
	} else {
		fmt.Println("✓ Data written successfully")
	}

	// Test reading data
	fmt.Println("\n=== Testing Read Operations ===")

	fmt.Printf("Reading data from '%s'...\n", testFileName)
	readData, err := storageClient.Read(testFileName)
	if err != nil {
		fmt.Printf("⚠ Failed to read data: %v\n", err)
	} else {
		fmt.Printf("✓ Data read successfully: %s\n", string(readData))
	}

	// Test checking if file exists
	fmt.Println("\n=== Testing Exists Check ===")

	fmt.Printf("Checking if '%s' exists...\n", testFileName)
	exists, err := storageClient.Exists(testFileName)
	if err != nil {
		fmt.Printf("⚠ Failed to check existence: %v\n", err)
	} else {
		fmt.Printf("✓ File exists: %t\n", exists)
	}

	// Test getting file statistics
	fmt.Println("\n=== Testing File Statistics ===")

	fmt.Printf("Getting statistics for '%s'...\n", testFileName)
	stat, err := storageClient.Stat(testFileName)
	if err != nil {
		fmt.Printf("⚠ Failed to get file statistics: %v\n", err)
	} else {
		fmt.Printf("✓ File statistics:\n")
		fmt.Printf("  Filename: %s\n", stat.Filename)
		fmt.Printf("  Size: %d bytes\n", stat.Size)
		fmt.Printf("  Last Modified: %d\n", stat.LastModified)
	}

	// Test async operations
	fmt.Println("\n=== Testing Async Operations ===")

	// Test async read
	fmt.Println("Testing async read...")
	dataChan, errChan := storageClient.ReadAsync(testFileName)

	select {
	case data := <-dataChan:
		fmt.Printf("✓ Async read successful: %s\n", string(data))
	case err := <-errChan:
		fmt.Printf("⚠ Async read failed: %v\n", err)
	case <-time.After(3 * time.Second):
		fmt.Println("⚠ Async read timed out")
	}

	// Test async write
	fmt.Println("Testing async write...")
	writeData := []byte("Async write test from Discord Go SDK!")
	writeErrChan := storageClient.WriteAsync("async_test.txt", writeData)

	select {
	case err := <-writeErrChan:
		if err != nil {
			fmt.Printf("⚠ Async write failed: %v\n", err)
		} else {
			fmt.Println("✓ Async write successful")
		}
	case <-time.After(3 * time.Second):
		fmt.Println("⚠ Async write timed out")
	}

	// Test getting storage path
	fmt.Println("\n=== Testing Storage Path ===")

	path, err := storageClient.GetPath()
	if err != nil {
		fmt.Printf("⚠ Failed to get storage path: %v\n", err)
	} else {
		fmt.Printf("✓ Storage path: %s\n", path)
	}

	// Test deleting file
	fmt.Println("\n=== Testing Delete Operation ===")

	fmt.Printf("Deleting '%s'...\n", testFileName)
	err = storageClient.Delete(testFileName)
	if err != nil {
		fmt.Printf("⚠ Failed to delete file: %v\n", err)
	} else {
		fmt.Println("✓ File deleted successfully")
	}

	// Verify deletion
	fmt.Printf("Checking if '%s' still exists...\n", testFileName)
	exists, err = storageClient.Exists(testFileName)
	if err != nil {
		fmt.Printf("⚠ Failed to check existence after deletion: %v\n", err)
	} else {
		fmt.Printf("✓ File exists after deletion: %t\n", exists)
	}

	// Let the client run for a moment to process any remaining events
	fmt.Println("Processing SDK events...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 Storage example completed successfully!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- SDK initialization with the new Client wrapper")
	fmt.Println("- Storage read/write operations")
	fmt.Println("- File existence checking")
	fmt.Println("- File statistics")
	fmt.Println("- Async operations with channels")
	fmt.Println("- File deletion")
	fmt.Println("- Go-like error handling")
	fmt.Println("- Enhanced storage management")
}
