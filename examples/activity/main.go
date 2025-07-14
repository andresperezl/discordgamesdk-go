package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Activity Example ===")
	fmt.Println("This example demonstrates activity management using the new Go-like Client wrapper.")

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

	// Create a rich activity using the builder pattern
	fmt.Println("\n=== Creating Rich Activity ===")

	activity := discord.NewActivity().
		SetType(discord.ActivityTypePlaying).
		SetApplicationID(clientID).
		SetName("Discord Go SDK").
		SetState("Testing Activity Builder").
		SetDetails("Building activities with Go-like interfaces").
		SetTimestamps(time.Now().Unix(), time.Now().Add(3600).Unix()).
		SetAssets("large_image_key", "Playing Discord Go SDK", "small_image_key", "Go Language").
		SetParty("party_id_123", 1, 4, discord.ActivityPartyPrivacyPublic).
		SetSecrets("match_secret", "join_secret", "spectate_secret").
		SetInstance(true).
		Build()

	// Set the activity
	fmt.Println("Setting rich activity...")
	err = activityClient.SetActivity(activity)
	if err != nil {
		log.Fatalf("Failed to set activity: %v", err)
	}
	fmt.Println("✓ Rich activity set successfully")

	// Wait a moment to see the activity
	time.Sleep(2 * time.Second)

	// Update the activity with new details
	fmt.Println("\n=== Updating Activity ===")

	updatedActivity := discord.NewActivity().
		SetType(discord.ActivityTypePlaying).
		SetApplicationID(clientID).
		SetName("Discord Go SDK").
		SetState("Activity Updated!").
		SetDetails("Updated with new information").
		SetTimestamps(time.Now().Unix(), time.Now().Add(7200).Unix()).
		Build()

	fmt.Println("Updating activity...")
	err = activityClient.SetActivity(updatedActivity)
	if err != nil {
		log.Fatalf("Failed to update activity: %v", err)
	}
	fmt.Println("✓ Activity updated successfully")

	// Wait a moment to see the updated activity
	time.Sleep(2 * time.Second)

	// Test async activity operations
	fmt.Println("\n=== Testing Async Operations ===")

	// Test async activity update
	fmt.Println("Testing async activity update...")
	activityClient.SetActivityWithCallback(updatedActivity, func(err error) {
		if err != nil {
			fmt.Printf("⚠ Async activity update failed: %v\n", err)
		} else {
			fmt.Println("✓ Async activity update completed")
		}
	})

	// Wait a moment for the callback
	time.Sleep(1 * time.Second)

	// Test async activity clear
	fmt.Println("Testing async activity clear...")
	activityClient.ClearActivityWithCallback(func(err error) {
		if err != nil {
			fmt.Printf("⚠ Async activity clear failed: %v\n", err)
		} else {
			fmt.Println("✓ Async activity clear completed")
		}
	})

	// Wait a moment for the callback
	time.Sleep(1 * time.Second)

	// Test activity with different types
	fmt.Println("\n=== Testing Different Activity Types ===")

	// Listening activity
	listeningActivity := discord.NewActivity().
		SetType(discord.ActivityTypeListening).
		SetApplicationID(clientID).
		SetName("Spotify").
		SetState("Listening to music").
		SetDetails("Song Title - Artist").
		Build()

	fmt.Println("Setting listening activity...")
	err = activityClient.SetActivity(listeningActivity)
	if err != nil {
		fmt.Printf("⚠ Failed to set listening activity: %v\n", err)
	} else {
		fmt.Println("✓ Listening activity set successfully")
	}

	time.Sleep(2 * time.Second)

	// Streaming activity
	streamingActivity := discord.NewActivity().
		SetType(discord.ActivityTypeStreaming).
		SetApplicationID(clientID).
		SetName("Twitch").
		SetState("Live on Twitch").
		SetDetails("Playing some games").
		Build()

	fmt.Println("Setting streaming activity...")
	err = activityClient.SetActivity(streamingActivity)
	if err != nil {
		fmt.Printf("⚠ Failed to set streaming activity: %v\n", err)
	} else {
		fmt.Println("✓ Streaming activity set successfully")
	}

	time.Sleep(2 * time.Second)

	// Watching activity
	watchingActivity := discord.NewActivity().
		SetType(discord.ActivityTypeWatching).
		SetApplicationID(clientID).
		SetName("YouTube").
		SetState("Watching videos").
		SetDetails("Learning Go programming").
		Build()

	fmt.Println("Setting watching activity...")
	err = activityClient.SetActivity(watchingActivity)
	if err != nil {
		fmt.Printf("⚠ Failed to set watching activity: %v\n", err)
	} else {
		fmt.Println("✓ Watching activity set successfully")
	}

	time.Sleep(2 * time.Second)

	// Clear all activities
	fmt.Println("\n=== Clearing Activity ===")
	err = activityClient.ClearActivity()
	if err != nil {
		log.Fatalf("Failed to clear activity: %v", err)
	}
	fmt.Println("✓ Activity cleared successfully")

	// Let the client run for a moment to process any remaining events
	fmt.Println("Processing SDK events...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n🎉 Activity example completed successfully!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- SDK initialization with the new Client wrapper")
	fmt.Println("- Activity builder pattern for easy activity creation")
	fmt.Println("- Different activity types (Playing, Listening, Streaming, Watching)")
	fmt.Println("- Async operations with callback support")
	fmt.Println("- Go-like error handling")
	fmt.Println("- Enhanced activity management")
}
