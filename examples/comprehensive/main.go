package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	discord "github.com/andresperezl/discordctl"
	core "github.com/andresperezl/discordctl/core"
)

func getClientID() int64 {
	if len(os.Args) > 1 {
		if id, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
			return id
		}
	}
	if env := os.Getenv("DISCORD_CLIENT_ID"); env != "" {
		if id, err := strconv.ParseInt(env, 10, 64); err == nil {
			return id
		}
	}
	// Default for testing
	return 1311711649018941501
}

func main() {
	clientID := getClientID()
	fmt.Printf("Starting Discord SDK with client ID: %d\n", clientID)

	config := discord.DefaultClientConfig(clientID)
	client, err := discord.NewClient(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize Discord SDK: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Println("Discord SDK initialized successfully!")

	// Fetch current user
	user, err := client.GetCurrentUser(5 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fetch current user: %v\n", err)
	} else {
		fmt.Printf("Current user: %s#%s (ID: %d)\n", user.Username, user.Discriminator, user.ID)
	}

	// Set a simple activity
	activity := &core.Activity{
		State:      "In the Go example!",
		Details:    "Comprehensive Example",
		Timestamps: core.ActivityTimestamps{Start: time.Now().Unix()},
	}
	err = client.Activity().SetActivity(activity)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set activity: %v\n", err)
	} else {
		fmt.Println("Activity set successfully!")
	}

	fmt.Println("Waiting 5 seconds to show activity...")
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting.")
}
