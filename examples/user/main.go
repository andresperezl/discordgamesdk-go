package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK User Example ===")
	fmt.Println("This example demonstrates user information retrieval.")

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
	fmt.Println("✓ Current user retrieved")
	fmt.Printf("  User ID: %d\n", user.ID)
	fmt.Printf("  Username: %s\n", user.Username)
	fmt.Printf("  Discriminator: %s\n", user.Discriminator)
	fmt.Printf("  Avatar: %s\n", user.Avatar)
	fmt.Printf("  Bot: %t\n", user.Bot)

	userManager := core.GetUserManager()
	if userManager == nil {
		log.Fatal("Failed to get user manager")
	}

	// Get current user's premium type
	premiumType, result := userManager.GetCurrentUserPremiumType()
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to get premium type: %v\n", result)
	} else {
		fmt.Println("✓ Premium type retrieved")
		switch premiumType {
		case discord.PremiumTypeNone:
			fmt.Println("  Premium Type: None")
		case discord.PremiumTypeTier1:
			fmt.Println("  Premium Type: Nitro Classic")
		case discord.PremiumTypeTier2:
			fmt.Println("  Premium Type: Nitro")
		default:
			fmt.Printf("  Premium Type: Unknown (%d)\n", premiumType)
		}
	}

	// Check user flags
	flags := []struct {
		name string
		flag discord.UserFlag
	}{
		{"Partner", discord.UserFlagPartner},
		{"HypeSquad Events", discord.UserFlagHypeSquadEvents},
		{"HypeSquad House 1", discord.UserFlagHypeSquadHouse1},
		{"HypeSquad House 2", discord.UserFlagHypeSquadHouse2},
		{"HypeSquad House 3", discord.UserFlagHypeSquadHouse3},
	}

	fmt.Println("Checking user flags...")
	for _, flagInfo := range flags {
		hasFlag, result := userManager.CurrentUserHasFlag(flagInfo.flag)
		if result != discord.ResultOk {
			fmt.Printf("  Warning: Failed to check %s flag: %v\n", flagInfo.name, result)
		} else {
			status := "No"
			if hasFlag {
				status = "Yes"
			}
			fmt.Printf("  %s: %s\n", flagInfo.name, status)
		}
	}

	// Let the callback loop process any remaining events
	fmt.Println("Processing user operations...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 User example completed!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- Current user information retrieval")
	fmt.Println("- Premium type checking")
	fmt.Println("- User flag verification")
}
