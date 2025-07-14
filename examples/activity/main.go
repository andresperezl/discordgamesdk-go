package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Activity Example ===")
	fmt.Println("This example demonstrates rich presence functionality.")

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

	// Get activity manager
	activityManager := core.GetActivityManager()
	if activityManager == nil {
		log.Fatal("Failed to get activity manager")
	}
	fmt.Println("✓ Activity manager retrieved")

	// Register the application
	result = activityManager.RegisterCommand("discordctl-activity-example")
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to register command: %v\n", result)
	} else {
		fmt.Println("✓ Application registered")
	}

	// Create a rich presence activity
	activity := discord.Activity{
		Type:          discord.ActivityTypePlaying,
		ApplicationID: clientID,
		Name:          "Discord Go SDK",
		State:         "Testing Rich Presence",
		Details:       "Running activity example",
		Timestamps: discord.ActivityTimestamps{
			Start: time.Now().Unix(),
		},
		Assets: discord.ActivityAssets{
			LargeImage: "logo",
			LargeText:  "Discord Go SDK",
			SmallImage: "go",
			SmallText:  "Go Language",
		},
		Party: discord.ActivityParty{
			Size: discord.PartySize{
				CurrentSize: 1,
				MaxSize:     4,
			},
			Privacy: discord.ActivityPartyPrivacyPublic,
		},
		Secrets: discord.ActivitySecrets{
			Match:    "secret-match",
			Join:     "secret-join",
			Spectate: "secret-spectate",
		},
		Instance: true,
	}

	// Update the activity
	fmt.Println("Updating rich presence...")
	activityManager.UpdateActivity(&activity, func(result discord.Result) {
		if result == discord.ResultOk {
			fmt.Println("✓ Rich presence updated successfully")
		} else {
			fmt.Printf("Warning: Failed to update activity: %v\n", result)
		}
	})

	// Let the callback loop process the update
	fmt.Println("Processing activity update...")
	time.Sleep(2 * time.Second)

	fmt.Println("\n🎉 Activity example completed!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- Rich presence creation")
	fmt.Println("- Activity updates")
	fmt.Println("- Party information")
	fmt.Println("- Activity secrets")

	fmt.Println("\nPress Enter to clear activity and exit...")
	fmt.Scanln()

	// Clear the activity
	activityManager.ClearActivity(func(result discord.Result) {
		if result == discord.ResultOk {
			fmt.Println("✓ Activity cleared successfully")
		} else {
			fmt.Printf("Warning: Failed to clear activity: %v\n", result)
		}
	})

	// Let the callback loop process the clear
	time.Sleep(1 * time.Second)
}
