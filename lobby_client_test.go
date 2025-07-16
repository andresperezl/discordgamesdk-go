package discord

import (
	"context"
	"log"
	"time"

	core "github.com/andresperezl/discordctl/core"
)

// ExampleLobbyClient_CreateLobbyWithContext demonstrates how to use CreateLobbyWithContext with a timeout.
// This example is for documentation only and requires a real, initialized LobbyClient and transaction.
func ExampleLobbyClient_CreateLobbyWithContext() {
	var lobbyClient *LobbyClient           // Assume this is properly initialized
	var transaction *core.LobbyTransaction // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lobby, err := lobbyClient.CreateLobbyWithContext(ctx, transaction)
	if err != nil {
		log.Fatalf("failed to create lobby: %v", err)
	}
	log.Printf("Created lobby with ID: %d", lobby.ID)
	// No Output: (documentation only)
}

// ExampleLobbyClient_UpdateLobbyWithContext demonstrates how to use UpdateLobbyWithContext with a timeout.
// This example is for documentation only and requires a real, initialized LobbyClient and transaction.
func ExampleLobbyClient_UpdateLobbyWithContext() {
	var lobbyClient *LobbyClient           // Assume this is properly initialized
	var transaction *core.LobbyTransaction // Assume this is properly initialized
	var lobbyID int64                      // Assume this is properly set

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := lobbyClient.UpdateLobbyWithContext(ctx, lobbyID, transaction)
	if err != nil {
		log.Fatalf("failed to update lobby: %v", err)
	}
	// No Output: (documentation only)
}
