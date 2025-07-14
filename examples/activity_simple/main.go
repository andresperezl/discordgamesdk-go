package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Simple Activity Test ===")

	// Initialize Discord SDK
	clientID := int64(1311711649018941501)
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

	// Get activity manager
	activityManager := core.GetActivityManager()
	if activityManager == nil {
		log.Fatal("Failed to get activity manager")
	}
	fmt.Println("✓ Activity manager retrieved")

	// Create a simple activity without complex assets
	activity := discord.Activity{
		Type:          discord.ActivityTypePlaying,
		ApplicationID: clientID,
		Name:          "Discord Go SDK Test",
		State:         "Testing",
		Details:       "Simple activity test",
		Timestamps: discord.ActivityTimestamps{
			Start: time.Now().Unix(),
		},
		Instance: true,
	}

	// Update the activity
	fmt.Println("Updating activity...")
	activityManager.UpdateActivity(&activity, func(result discord.Result) {
		fmt.Printf("Activity update callback result: %v\n", result)
		if result == discord.ResultOk {
			fmt.Println("✓ Activity updated successfully")
		} else {
			fmt.Printf("✗ Activity update failed: %v\n", result)
		}
	})

	// Let the callback loop process the update
	fmt.Println("Processing activity update...")
	time.Sleep(2 * time.Second)

	fmt.Println("\nPress Enter to clear activity and exit...")
	fmt.Scanln()

	// Clear the activity
	activityManager.ClearActivity(func(result discord.Result) {
		fmt.Printf("Activity clear callback result: %v\n", result)
		if result == discord.ResultOk {
			fmt.Println("✓ Activity cleared successfully")
		} else {
			fmt.Printf("✗ Activity clear failed: %v\n", result)
		}
	})

	// Let the callback loop process the clear
	time.Sleep(1 * time.Second)

	fmt.Println("Test completed!")
}
