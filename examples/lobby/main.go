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

	// Create a public lobby
	lobbyManager := client.Lobby()
	createTxn, err := lobbyManager.GetLobbyCreateTransaction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get lobby create transaction: %v\n", err)
		return
	}
	createTxn.SetType(core.LobbyTypePublic)
	createTxn.SetCapacity(4)

	lobbyChan, errChan := lobbyManager.CreateLobby(createTxn)

	select {
	case lobby := <-lobbyChan:
		if lobby != nil {
			fmt.Printf("Lobby created! ID: %d, Secret: %s\n", lobby.ID, lobby.Secret)
		} else {
			fmt.Println("Lobby creation returned nil lobby.")
		}
	case err := <-errChan:
		fmt.Fprintf(os.Stderr, "Failed to create lobby: %v\n", err)
		return
	case <-time.After(5 * time.Second):
		fmt.Fprintln(os.Stderr, "Timed out waiting for lobby creation.")
		return
	}

	fmt.Println("Waiting 5 seconds before exit...")
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting.")
}
