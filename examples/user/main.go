package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
	core "github.com/andresperezl/discordctl/core"
)

func main() {
	fmt.Println("=== Discord Game SDK User Example ===")
	fmt.Println("This example demonstrates user management using the new Go-like Client wrapper.")

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

	// Get user manager
	userClient := client.User()

	// Get current user
	fmt.Println("\n=== Getting Current User ===")
	user, err := userClient.GetCurrentUser()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	fmt.Printf("✓ Current user: %s#%s\n", user.Username, user.Discriminator)
	fmt.Printf("  User ID: %d\n", user.ID)
	fmt.Printf("  Avatar: %s\n", user.Avatar)
	fmt.Printf("  Bot: %t\n", user.Bot)

	// Get current user's premium type
	fmt.Println("\n=== Getting Premium Type ===")
	premiumType, err := userClient.GetCurrentUserPremiumType()
	if err != nil {
		fmt.Printf("⚠ Failed to get premium type: %v\n", err)
	} else {
		fmt.Printf("✓ Premium type: %v\n", premiumType)
	}

	// Check if current user has specific flags
	fmt.Println("\n=== Checking User Flags ===")

	// Check for Partner flag
	hasPartner, err := userClient.CurrentUserHasFlag(core.UserFlagPartner)
	if err != nil {
		fmt.Printf("⚠ Failed to check Partner flag: %v\n", err)
	} else {
		fmt.Printf("✓ Has Partner: %t\n", hasPartner)
	}

	// Check for HypeSquad Events flag
	hasHypeSquadEvents, err := userClient.CurrentUserHasFlag(core.UserFlagHypeSquadEvents)
	if err != nil {
		fmt.Printf("⚠ Failed to check HypeSquad Events flag: %v\n", err)
	} else {
		fmt.Printf("✓ Has HypeSquad Events: %t\n", hasHypeSquadEvents)
	}

	// Check for HypeSquad House 1 flag
	hasHypeSquadHouse1, err := userClient.CurrentUserHasFlag(core.UserFlagHypeSquadHouse1)
	if err != nil {
		fmt.Printf("⚠ Failed to check HypeSquad House 1 flag: %v\n", err)
	} else {
		fmt.Printf("✓ Has HypeSquad House 1: %t\n", hasHypeSquadHouse1)
	}

	// Check for HypeSquad House 2 flag
	hasHypeSquadHouse2, err := userClient.CurrentUserHasFlag(core.UserFlagHypeSquadHouse2)
	if err != nil {
		fmt.Printf("⚠ Failed to check HypeSquad House 2 flag: %v\n", err)
	} else {
		fmt.Printf("✓ Has HypeSquad House 2: %t\n", hasHypeSquadHouse2)
	}

	// Check for HypeSquad House 3 flag
	hasHypeSquadHouse3, err := userClient.CurrentUserHasFlag(core.UserFlagHypeSquadHouse3)
	if err != nil {
		fmt.Printf("⚠ Failed to check HypeSquad House 3 flag: %v\n", err)
	} else {
		fmt.Printf("✓ Has HypeSquad House 3: %t\n", hasHypeSquadHouse3)
	}

	// Test getting a specific user (this would require a valid user ID)
	fmt.Println("\n=== Testing Get User ===")
	fmt.Println("Note: GetUser requires a valid user ID and is async")
	fmt.Println("This is just a demonstration of the interface")

	// Let the client run for a moment to process any remaining events
	fmt.Println("Processing SDK events...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 User example completed successfully!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- SDK initialization with the new Client wrapper")
	fmt.Println("- Getting current user information")
	fmt.Println("- Checking user premium status")
	fmt.Println("- Checking user flags")
	fmt.Println("- Go-like error handling")
	fmt.Println("- Enhanced user management")
}
