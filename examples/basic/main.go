package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
	core "github.com/andresperezl/discordctl/core"
)

func main() {
	fmt.Println("=== Discord Game SDK Basic Example ===")
	fmt.Println("This example demonstrates basic SDK initialization and manager access using the new Go-like Client wrapper.")

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

	// Get current user using the new Go-like interface
	fmt.Println("Getting current user...")
	user, err := client.GetCurrentUser(5 * time.Second)
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
	fmt.Printf("✓ Connected as user: %s\n", user.Username)

	// Test getting managers using the new client methods
	fmt.Println("\n=== Testing Manager Access ===")

	// Test activity manager
	fmt.Println("Testing activity manager...")
	activityClient := client.Activity()

	// Create an activity using the builder pattern
	activity := discord.NewActivity().
		SetType(core.ActivityTypePlaying).
		SetApplicationID(clientID).
		SetName("Discord Go SDK").
		SetState("Testing New Client").
		SetDetails("Enhanced Go-like interface").
		SetTimestamps(time.Now().Unix(), 0).
		Build()

	// Set activity using the new Go-like interface
	err = activityClient.SetActivity(activity)
	if err != nil {
		fmt.Printf("⚠ Activity update failed: %v\n", err)
	} else {
		fmt.Println("✓ Activity updated successfully")
	}

	// Test user manager
	fmt.Println("Testing user manager...")
	userClient := client.User()

	// Get current user premium type
	premiumType, err := userClient.GetCurrentUserPremiumType()
	if err != nil {
		fmt.Printf("⚠ Failed to get premium type: %v\n", err)
	} else {
		fmt.Printf("✓ Premium type: %v\n", premiumType)
	}

	// Test storage manager
	fmt.Println("Testing storage manager...")
	storageClient := client.Storage()

	// Get storage count
	count, err := storageClient.Count()
	if err != nil {
		fmt.Printf("⚠ Failed to get storage count: %v\n", err)
	} else {
		fmt.Printf("✓ Storage count: %d\n", count)
	}

	// Test application manager
	fmt.Println("Testing application manager...")
	appClient := client.Application()

	// Get current locale
	locale, err := appClient.GetCurrentLocale()
	if err != nil {
		fmt.Printf("⚠ Failed to get locale: %v\n", err)
	} else {
		fmt.Printf("✓ Current locale: %s\n", locale)
	}

	// Test overlay manager
	fmt.Println("Testing overlay manager...")
	overlayClient := client.Overlay()

	// Check if overlay is enabled
	enabled, err := overlayClient.IsEnabled()
	if err != nil {
		fmt.Printf("⚠ Failed to check overlay status: %v\n", err)
	} else {
		fmt.Printf("✓ Overlay enabled: %t\n", enabled)
	}

	// Test network manager
	fmt.Println("Testing network manager...")
	networkClient := client.Network()

	// Get peer ID
	peerID, err := networkClient.GetPeerID()
	if err != nil {
		fmt.Printf("⚠ Failed to get peer ID: %v\n", err)
	} else {
		fmt.Printf("✓ Peer ID: %d\n", peerID)
	}

	// Test lobby manager
	fmt.Println("Testing lobby manager...")
	_ = client.Lobby() // Lobby operations are async and return channels

	// Note: Lobby operations are async and return channels
	// This is just a demonstration of the interface
	fmt.Println("✓ Lobby manager available (operations are async)")

	// Clear activity at the end
	fmt.Println("\nClearing activity...")
	err = activityClient.ClearActivity()
	if err != nil {
		fmt.Printf("⚠ Failed to clear activity: %v\n", err)
	} else {
		fmt.Println("✓ Activity cleared successfully")
	}

	// Let the client run for a moment to process any remaining events
	fmt.Println("Processing SDK events...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 Basic example completed successfully!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- SDK initialization with the new Client wrapper")
	fmt.Println("- Go-like error handling")
	fmt.Println("- Manager access through client methods")
	fmt.Println("- Activity builder pattern")
	fmt.Println("- Async operations with channels")
	fmt.Println("- Enhanced error messages with Result.String()")
}
