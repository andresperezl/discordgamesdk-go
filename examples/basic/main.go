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
	discordClient, err := discord.NewClient(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize Discord SDK: %v\n", err)
		os.Exit(1)
	}
	defer discordClient.Close()

	fmt.Println("Discord SDK initialized successfully! Waiting 3 seconds...")
	time.Sleep(3 * time.Second)
	fmt.Println("Exiting.")
}
