package discord

import (
	"context"
	"log"
	"time"

	core "github.com/andresperezl/discordgamesdk-go/core"
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

// ExampleLobbyClient_LobbyEventsChannel demonstrates how to use LobbyEventsChannel to receive all lobby event streams.
// This example is for documentation only and requires a real, initialized LobbyClient.
func ExampleLobbyClient_LobbyEventsChannel() {
	var lobbyClient *LobbyClient // Assume this is properly initialized

	events := lobbyClient.LobbyEventsChannel()

	go func() {
		for userID := range events.MemberJoin {
			log.Printf("User joined: %d", userID)
		}
	}()
	go func() {
		for userID := range events.MemberLeave {
			log.Printf("User left: %d", userID)
		}
	}()
	go func() {
		for msg := range events.LobbyMessage {
			log.Printf("Message in lobby %d from %d: %s", msg.LobbyID, msg.UserID, string(msg.Data))
		}
	}()
	go func() {
		for lobbyID := range events.LobbyUpdate {
			log.Printf("Lobby updated: %d", lobbyID)
		}
	}()
	go func() {
		for del := range events.LobbyDelete {
			log.Printf("Lobby deleted: %d (reason: %d)", del.LobbyID, del.Reason)
		}
	}()
	go func() {
		for upd := range events.MemberUpdate {
			log.Printf("Member updated: lobby %d user %d", upd.LobbyID, upd.UserID)
		}
	}()
	go func() {
		for spk := range events.Speaking {
			log.Printf("Speaking: lobby %d user %d speaking=%v", spk.LobbyID, spk.UserID, spk.Speaking)
		}
	}()
	go func() {
		for net := range events.NetworkMessage {
			log.Printf("Network message: lobby %d user %d channel %d data %v", net.LobbyID, net.UserID, net.ChannelID, net.Data)
		}
	}()
	// No Output: (documentation only)
}
