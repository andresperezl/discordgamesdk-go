package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
	core "github.com/andresperezl/discordctl/core"
)

func main() {
	fmt.Println("=== Discord Game SDK Simple Activity Example ===")
	fmt.Println("This example demonstrates simple activity management using the new Go-like Client wrapper.")

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

	// Get activity manager
	activityClient := client.Activity()

	// Create a simple activity using the builder pattern
	fmt.Println("\n=== Creating Simple Activity ===")

	activity := discord.NewActivity().
		SetType(core.ActivityTypePlaying).
		SetApplicationID(clientID).
		SetName("Discord Go SDK").
		SetState("Simple Example").
		SetDetails("Testing simple activity").
		Build()

	// Set the activity
	fmt.Println("Setting simple activity...")
	err = activityClient.SetActivity(activity)
	if err != nil {
		log.Fatalf("Failed to set activity: %v", err)
	}
	fmt.Println("✓ Simple activity set successfully")

	// Wait a moment to see the activity
	time.Sleep(3 * time.Second)

	// Clear the activity
	fmt.Println("\n=== Clearing Activity ===")
	err = activityClient.ClearActivity()
	if err != nil {
		log.Fatalf("Failed to clear activity: %v", err)
	}
	fmt.Println("✓ Activity cleared successfully")

	// Let the client run for a moment to process any remaining events
	fmt.Println("Processing SDK events...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 Simple activity example completed successfully!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- SDK initialization with the new Client wrapper")
	fmt.Println("- Simple activity creation with the builder pattern")
	fmt.Println("- Go-like error handling")
	fmt.Println("- Clean activity management")
}
