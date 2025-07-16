package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	discord "github.com/andresperezl/discordgamesdk-go"
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
	overlay := client.Overlay()
	if overlay == nil {
		fmt.Println("Overlay manager not available.")
		return
	}

	enabled, err := overlay.IsEnabled()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to check overlay enabled: %v\n", err)
		return
	}
	fmt.Printf("Overlay enabled: %v\n", enabled)

	locked, err := overlay.IsLocked()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to check overlay locked: %v\n", err)
		return
	}
	fmt.Printf("Overlay locked: %v\n", locked)

	// Toggle lock state
	newLock := !locked
	errChan := overlay.SetLocked(newLock)
	err = <-errChan
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set overlay lock: %v\n", err)
	} else {
		fmt.Printf("Overlay lock state set to: %v\n", newLock)
	}

	fmt.Println("Waiting 3 seconds before exit...")
	time.Sleep(3 * time.Second)
	fmt.Println("Exiting.")
}
