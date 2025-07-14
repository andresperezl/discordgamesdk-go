package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Basic Example ===")
	fmt.Println("This example demonstrates basic SDK initialization and manager access.")

	// Initialize Discord SDK
	clientID := int64(1311711649018941501) // Replace with your actual client ID
	fmt.Printf("Initializing Discord SDK with client ID: %d\n", clientID)

	core, err := discord.Create(clientID, discord.CreateFlagsDefault, nil)
	if err != discord.ResultOk {
		log.Fatalf("Failed to create Discord core: %v", err)
	}

	// Start the callback loop for robust event processing
	fmt.Println("Starting callback loop...")
	core.Start()
	defer core.Shutdown()

	// Wait for SDK initialization
	fmt.Println("Waiting for SDK initialization...")
	if !core.WaitForInitialization(5 * time.Second) {
		log.Fatal("Failed to initialize Discord SDK within timeout")
	}
	fmt.Println("✓ Discord SDK initialized successfully")

	// Wait for user info to become available (robust pattern)
	fmt.Println("Waiting for user info...")
	user, result := core.WaitForUser(5 * time.Second)
	if result != discord.ResultOk {
		log.Fatalf("Failed to get current user: %v", result)
	}
	fmt.Printf("✓ Connected as user: %s\n", user.Username)

	// Test getting managers
	appManager := core.GetApplicationManager()
	if appManager == nil {
		log.Fatal("Failed to get application manager")
	}
	fmt.Println("✓ Application manager retrieved")

	userManager := core.GetUserManager()
	if userManager == nil {
		log.Fatal("Failed to get user manager")
	}
	fmt.Println("✓ User manager retrieved")

	activityManager := core.GetActivityManager()
	if activityManager == nil {
		log.Fatal("Failed to get activity manager")
	}
	fmt.Println("✓ Activity manager retrieved")

	storageManager := core.GetStorageManager()
	if storageManager == nil {
		log.Fatal("Failed to get storage manager")
	}
	fmt.Println("✓ Storage manager retrieved")

	// Test callback result tracking
	fmt.Println("\n=== Testing Callback Result Tracking ===")

	// Test async activity update
	fmt.Println("Testing async activity update...")
	resultChan := activityManager.UpdateActivityAsync(&discord.Activity{
		Type:          discord.ActivityTypePlaying,
		ApplicationID: clientID,
		Name:          "Discord Go SDK",
		State:         "Testing Callbacks",
		Details:       "Enhanced callback handling",
		Timestamps: discord.ActivityTimestamps{
			Start: time.Now().Unix(),
		},
	})

	// Wait for result with timeout
	select {
	case result := <-resultChan:
		if result == discord.ResultOk {
			fmt.Println("✓ Async activity update completed successfully")
		} else {
			fmt.Printf("⚠ Async activity update failed: %v\n", result)
		}
	case <-time.After(3 * time.Second):
		fmt.Println("⚠ Async activity update timed out")
	}

	// Test async activity clear
	fmt.Println("Testing async activity clear...")
	clearChan := activityManager.ClearActivityAsync()

	select {
	case result := <-clearChan:
		if result == discord.ResultOk {
			fmt.Println("✓ Async activity clear completed successfully")
		} else {
			fmt.Printf("⚠ Async activity clear failed: %v\n", result)
		}
	case <-time.After(3 * time.Second):
		fmt.Println("⚠ Async activity clear timed out")
	}

	// Let the callback loop process any remaining events
	fmt.Println("Processing SDK events...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 Basic example completed successfully!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- SDK initialization with proper waiting")
	fmt.Println("- Manager access")
	fmt.Println("- Robust callback processing")
	fmt.Println("- User connection verification")
	fmt.Println("- Async callback result tracking")
	fmt.Println("- Enhanced error handling")
}
