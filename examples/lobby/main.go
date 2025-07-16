package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"log/slog"

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
	h := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(h)

	clientID := getClientID()
	slog.Info("Starting Discord SDK", "clientID", clientID)

	config := discord.DefaultClientConfig(clientID)
	client, err := discord.NewClient(config)
	if err != nil {
		slog.Error("Failed to initialize Discord SDK", "error", err)
		os.Exit(1)
	}
	defer client.Close()

	slog.Info("Discord SDK initialized successfully!")

	// Create a public lobby
	lobbyManager := client.Lobby()
	createTxn, err := lobbyManager.GetLobbyCreateTransaction()
	if err != nil {
		slog.Error("Failed to get lobby create transaction", "error", err)
		return
	}
	slog.Info("Got lobby create transaction", "txn", createTxn)

	// Set up timeout from env
	timeoutSecs := 5
	if val := os.Getenv("LOBBY_TIMEOUT_SECS"); val != "" {
		if n, err := strconv.Atoi(val); err == nil && n > 0 {
			timeoutSecs = n
		}
	}
	timeout := time.Duration(timeoutSecs) * time.Second
	slog.Info("Lobby creation timeout set", "timeout", timeout)

	lobbyCh, errCh := lobbyManager.CreateLobby(createTxn)

	slog.Info("Called CreateLobby, waiting for result", "timeout", timeout)
	start := time.Now()
	var lobby *core.Lobby
	var lobbyErr error
	callbackDone := make(chan struct{})
	go func() {
		slog.Info("[goroutine] Waiting for lobby creation callback...")
		lobby = <-lobbyCh
		slog.Info("[goroutine] Received lobby from lobbyCh", "lobby", lobby)
		lobbyErr = <-errCh
		slog.Info("[goroutine] Received error from errCh", "error", lobbyErr)
		close(callbackDone)
		if lobbyErr != nil {
			slog.Error("[goroutine] Lobby creation error", "error", lobbyErr)
		} else if lobby != nil {
			slog.Info("[goroutine] Lobby created", "lobbyID", lobby.ID)
		} else {
			slog.Warn("[goroutine] Lobby is nil after creation")
		}
	}()

	select {
	case <-callbackDone:
		elapsed := time.Since(start)
		slog.Info("Lobby creation callback completed", "elapsed", elapsed, "lobby", lobby, "error", lobbyErr)
		if lobbyErr != nil {
			slog.Error("Lobby creation failed", "error", lobbyErr)
			fmt.Println("Lobby creation failed:", lobbyErr)
			return
		}
		if lobby == nil {
			slog.Error("Lobby is nil after creation")
			fmt.Println("Lobby is nil after creation")
			return
		}
		fmt.Println("Lobby created! ID:", lobby.ID)
	case <-time.After(timeout):
		elapsed := time.Since(start)
		slog.Error("Timed out waiting for lobby creation", "timeout", timeout, "elapsed", elapsed)
		fmt.Println("Timed out waiting for lobby creation.")
	}
	// Final log
	slog.Info("Exiting main")
}
