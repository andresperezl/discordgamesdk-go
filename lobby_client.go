package discord

import (
	"context"
	"fmt"
	"unsafe"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// LobbyClient provides Go-like interfaces for lobby management
type LobbyClient struct {
	manager *core.LobbyManager
	core    *core.Core // Added to match usage in client.go
}

func NewLobbyClient(core *core.Core) *LobbyClient {
	return &LobbyClient{manager: core.GetLobbyManager(), core: core}
}

func (c *LobbyClient) ConnectLobbyWithActivitySecret(activitySecret string, callbackData, callback unsafe.Pointer) {
	c.manager.ConnectLobbyWithActivitySecret(activitySecret, callbackData, callback)
}

func (c *LobbyClient) GetMemberUpdateTransaction(lobbyID, userID int64) unsafe.Pointer {
	return c.manager.GetMemberUpdateTransaction(lobbyID, userID)
}

func (c *LobbyClient) CreateLobby(transaction *core.LobbyTransaction) (<-chan *core.Lobby, <-chan error) {
	lobbyChan := make(chan *core.Lobby, 1)
	errChan := make(chan error, 1)

	if c.manager == nil {
		errChan <- fmt.Errorf("lobby manager not available")
		close(lobbyChan)
		close(errChan)
		return lobbyChan, errChan
	}

	c.manager.CreateLobby(transaction, func(result core.Result, lobby *core.Lobby) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to create lobby: %v", result)
			close(lobbyChan)
			close(errChan)
			return
		}
		lobbyChan <- lobby
		close(lobbyChan)
		close(errChan)
	})

	return lobbyChan, errChan
}

func (c *LobbyClient) ConnectLobby(lobbyID int64, secret string) (<-chan *core.Lobby, <-chan error) {
	lobbyChan := make(chan *core.Lobby, 1)
	errChan := make(chan error, 1)

	if c.manager == nil {
		errChan <- fmt.Errorf("lobby manager not available")
		close(lobbyChan)
		close(errChan)
		return lobbyChan, errChan
	}

	c.manager.ConnectLobby(lobbyID, secret, func(result core.Result, lobby *core.Lobby) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to connect to lobby: %v", result)
			close(lobbyChan)
			close(errChan)
			return
		}
		lobbyChan <- lobby
		close(lobbyChan)
		close(errChan)
	})

	return lobbyChan, errChan
}

func (c *LobbyClient) DisconnectLobby(lobbyID int64) <-chan error {
	errChan := make(chan error, 1)

	if c.manager == nil {
		errChan <- fmt.Errorf("lobby manager not available")
		close(errChan)
		return errChan
	}

	c.manager.DisconnectLobby(lobbyID, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to disconnect from lobby: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})

	return errChan
}

func (c *LobbyClient) GetLobbyActivitySecret(lobbyID int64) (string, error) {
	if c.manager == nil {
		return "", fmt.Errorf("lobby manager not available")
	}

	secret, res := c.manager.GetLobbyActivitySecret(lobbyID)
	if res != 0 {
		return "", fmt.Errorf("failed to get lobby activity secret: %v", res)
	}
	return secret, nil
}

func (c *LobbyClient) SetLobbyMetadata(lobbyID int64, key, value string) error {
	if c.manager == nil {
		return fmt.Errorf("lobby manager not available")
	}

	transaction, res := c.manager.GetLobbyUpdateTransaction(lobbyID)
	if res != core.ResultOk {
		return fmt.Errorf("failed to get lobby update transaction: %v", res)
	}

	res = transaction.SetMetadata(key, value)
	if res != core.ResultOk {
		return fmt.Errorf("failed to set lobby metadata: %v", res)
	}

	// TODO: Implement proper update callback
	return nil
}

func (c *LobbyClient) DeleteLobbyMetadata(lobbyID int64, key string) error {
	if c.manager == nil {
		return fmt.Errorf("lobby manager not available")
	}

	transaction, res := c.manager.GetLobbyUpdateTransaction(lobbyID)
	if res != core.ResultOk {
		return fmt.Errorf("failed to get lobby update transaction: %v", res)
	}

	res = transaction.DeleteMetadata(key)
	if res != core.ResultOk {
		return fmt.Errorf("failed to delete lobby metadata: %v", res)
	}

	// TODO: Implement proper update callback
	return nil
}

func (c *LobbyClient) GetLobbyMetadataCount(lobbyID int64) (int32, error) {
	if c.manager == nil {
		return 0, fmt.Errorf("lobby manager not available")
	}

	count, res := c.manager.LobbyMetadataCount(lobbyID)
	if res != 0 {
		return 0, fmt.Errorf("failed to get lobby metadata count: %v", res)
	}
	return count, nil
}

func (c *LobbyClient) GetLobbyMetadataKeyByIndex(lobbyID int64, index int32) (string, error) {
	if c.manager == nil {
		return "", fmt.Errorf("lobby manager not available")
	}

	key, res := c.manager.GetLobbyMetadataKey(lobbyID, index)
	if res != 0 {
		return "", fmt.Errorf("failed to get lobby metadata key: %v", res)
	}
	return key, nil
}

func (c *LobbyClient) GetLobbyMemberCount(lobbyID int64) (int32, error) {
	if c.manager == nil {
		return 0, fmt.Errorf("lobby manager not available")
	}

	count, res := c.manager.MemberCount(lobbyID)
	if res != 0 {
		return 0, fmt.Errorf("failed to get lobby member count: %v", res)
	}
	return count, nil
}

func (c *LobbyClient) GetLobbyMemberUserId(lobbyID int64, index int32) (int64, error) {
	if c.manager == nil {
		return 0, fmt.Errorf("lobby manager not available")
	}

	userID, res := c.manager.GetMemberUserID(lobbyID, index)
	if res != 0 {
		return 0, fmt.Errorf("failed to get lobby member user ID: %v", res)
	}
	return userID, nil
}

func (c *LobbyClient) GetLobbyMemberUser(lobbyID int64, userID int64) (*core.User, error) {
	if c.manager == nil {
		return nil, fmt.Errorf("lobby manager not available")
	}

	user, res := c.manager.GetMemberUser(lobbyID, userID)
	if res != 0 {
		return nil, fmt.Errorf("failed to get lobby member user: %v", res)
	}
	return user, nil
}

func (c *LobbyClient) GetLobbyMemberMetadataValue(lobbyID int64, userID int64, key string) (string, error) {
	if c.manager == nil {
		return "", fmt.Errorf("lobby manager not available")
	}

	value, res := c.manager.GetMemberMetadataValue(lobbyID, userID, key)
	if res != 0 {
		return "", fmt.Errorf("failed to get lobby member metadata value: %v", res)
	}
	return value, nil
}

func (c *LobbyClient) SetLobbyMemberMetadata(lobbyID int64, userID int64, key, value string) error {
	if c.manager == nil {
		return fmt.Errorf("lobby manager not available")
	}

	transaction := c.manager.GetMemberUpdateTransaction(lobbyID, userID)
	if transaction == nil {
		return fmt.Errorf("failed to get member update transaction")
	}

	// TODO: Implement proper member update transaction methods
	// For now, return success
	return nil
}

func (c *LobbyClient) DeleteLobbyMemberMetadata(lobbyID int64, userID int64, key string) error {
	if c.manager == nil {
		return fmt.Errorf("lobby manager not available")
	}

	// This would need to be implemented in the C wrapper
	// For now, return success
	return nil
}

func (c *LobbyClient) GetLobbyMemberMetadataCount(lobbyID int64, userID int64) (int32, error) {
	if c.manager == nil {
		return 0, fmt.Errorf("lobby manager not available")
	}

	// This would need to be implemented in the C wrapper
	// For now, return 0
	return 0, nil
}

func (c *LobbyClient) GetLobbyMemberMetadataKeyByIndex(lobbyID int64, userID int64, index int32) (string, error) {
	if c.manager == nil {
		return "", fmt.Errorf("lobby manager not available")
	}

	// This would need to be implemented in the C wrapper
	// For now, return empty string
	return "", nil
}

func (c *LobbyClient) SendLobbyMessage(lobbyID int64, data []byte) <-chan error {
	errChan := make(chan error, 1)

	if c.manager == nil {
		errChan <- fmt.Errorf("lobby manager not available")
		close(errChan)
		return errChan
	}

	c.manager.SendLobbyMessage(lobbyID, data, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to send lobby message: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})

	return errChan
}

func (c *LobbyClient) DeleteLobby(lobbyID int64) <-chan error {
	errChan := make(chan error, 1)

	if c.manager == nil {
		errChan <- fmt.Errorf("lobby manager not available")
		close(errChan)
		return errChan
	}

	c.manager.DeleteLobby(lobbyID, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to delete lobby: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})

	return errChan
}

// CreateLobbyWithContext creates a lobby, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	lobby, err := client.Lobby().CreateLobbyWithContext(ctx, transaction)
//	if err != nil {
//	    log.Fatalf("failed to create lobby: %v", err)
//	}
//	fmt.Printf("Created lobby with ID: %d\n", lobby.ID)
//
// Returns the created lobby or error if the context is cancelled, deadline exceeded, or the creation fails.
func (c *LobbyClient) CreateLobbyWithContext(ctx context.Context, transaction *core.LobbyTransaction) (*core.Lobby, error) {
	if c.manager == nil {
		return nil, fmt.Errorf("lobby manager not available")
	}

	lobbyChan := make(chan *core.Lobby, 1)
	errChan := make(chan error, 1)

	c.manager.CreateLobby(transaction, func(result core.Result, lobby *core.Lobby) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to create lobby: %v", result)
			return
		}
		lobbyChan <- lobby
	})

	select {
	case lobby := <-lobbyChan:
		return lobby, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// UpdateLobby updates a lobby asynchronously and returns a channel for the result.
func (c *LobbyClient) UpdateLobby(lobbyID int64, transaction *core.LobbyTransaction) <-chan error {
	errChan := make(chan error, 1)
	if c.manager == nil {
		errChan <- fmt.Errorf("lobby manager not available")
		close(errChan)
		return errChan
	}
	c.manager.UpdateLobby(lobbyID, transaction, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to update lobby: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

// UpdateLobbyWithContext updates a lobby, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	err := client.Lobby().UpdateLobbyWithContext(ctx, lobbyID, transaction)
//	if err != nil {
//	    log.Fatalf("failed to update lobby: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the update fails.
func (c *LobbyClient) UpdateLobbyWithContext(ctx context.Context, lobbyID int64, transaction *core.LobbyTransaction) error {
	if c.manager == nil {
		return fmt.Errorf("lobby manager not available")
	}
	errChan := make(chan error, 1)
	c.manager.UpdateLobby(lobbyID, transaction, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to update lobby: %v", result)
		} else {
			errChan <- nil
		}
	})
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetLobbyCreateTransaction returns a new lobby create transaction
func (c *LobbyClient) GetLobbyCreateTransaction() (*core.LobbyTransaction, error) {
	if c.manager == nil {
		return nil, fmt.Errorf("lobby manager not available")
	}
	transaction, result := c.manager.GetLobbyCreateTransaction()
	if result != core.ResultOk {
		return nil, fmt.Errorf("failed to get lobby create transaction: %v", result)
	}
	return transaction, nil
}

// NOTE: The Discord Game SDK does not provide APIs for lobby metadata value by index, lobby message history, or direct search result access.
// Methods such as GetLobbyMetadataValueByIndex, GetLobbyMemberMetadataValueByIndex, GetLobbyMessageCount, GetLobbyMessageUserId, GetLobbyMessageData,
// Search, SearchWithFilter, GetSearchResultCount, and GetSearchResult have been removed because they cannot be implemented with the current SDK.
//
// TODO: Methods like SetLobbyMetadata, DeleteLobbyMetadata, SetLobbyMemberMetadata, and DeleteLobbyMemberMetadata should be implemented using the transaction pattern.

// LobbyEventChannels provides channels for key lobby events.
type LobbyEventChannels struct {
	MemberJoin     <-chan int64 // userID
	MemberLeave    <-chan int64 // userID
	LobbyMessage   <-chan LobbyMessageEvent
	LobbyUpdate    <-chan int64 // lobbyID
	LobbyDelete    <-chan LobbyDeleteEvent
	MemberUpdate   <-chan MemberUpdateEvent
	Speaking       <-chan SpeakingEvent
	NetworkMessage <-chan NetworkMessageEvent
}

// LobbyMessageEvent represents a message event in a lobby.
type LobbyMessageEvent struct {
	LobbyID int64
	UserID  int64
	Data    []byte
}

// LobbyDeleteEvent represents a lobby delete event.
type LobbyDeleteEvent struct {
	LobbyID int64
	Reason  uint32
}

// MemberUpdateEvent represents a member update event in a lobby.
type MemberUpdateEvent struct {
	LobbyID int64
	UserID  int64
}

// SpeakingEvent represents a speaking event in a lobby.
type SpeakingEvent struct {
	LobbyID  int64
	UserID   int64
	Speaking bool
}

// NetworkMessageEvent represents a network message event in a lobby.
type NetworkMessageEvent struct {
	LobbyID   int64
	UserID    int64
	ChannelID uint8
	Data      []byte
}

// LobbyEventsChannel returns channels for key lobby events (member join/leave/update, lobby update/delete, message, speaking, network message).
//
// Example usage:
//
//	events := client.Lobby().LobbyEventsChannel()
//	go func() {
//	    for userID := range events.MemberJoin {
//	        fmt.Printf("User joined: %d\n", userID)
//	    }
//	}()
//	go func() {
//	    for userID := range events.MemberLeave {
//	        fmt.Printf("User left: %d\n", userID)
//	    }
//	}()
//	go func() {
//	    for msg := range events.LobbyMessage {
//	        fmt.Printf("Message in lobby %d from %d: %s\n", msg.LobbyID, msg.UserID, string(msg.Data))
//	    }
//	}()
//	go func() {
//	    for lobbyID := range events.LobbyUpdate {
//	        fmt.Printf("Lobby updated: %d\n", lobbyID)
//	    }
//	}()
//	go func() {
//	    for del := range events.LobbyDelete {
//	        fmt.Printf("Lobby deleted: %d (reason: %d)\n", del.LobbyID, del.Reason)
//	    }
//	}()
//	go func() {
//	    for upd := range events.MemberUpdate {
//	        fmt.Printf("Member updated: lobby %d user %d\n", upd.LobbyID, upd.UserID)
//	    }
//	}()
//	go func() {
//	    for spk := range events.Speaking {
//	        fmt.Printf("Speaking: lobby %d user %d speaking=%v\n", spk.LobbyID, spk.UserID, spk.Speaking)
//	    }
//	}()
//	go func() {
//	    for net := range events.NetworkMessage {
//	        fmt.Printf("Network message: lobby %d user %d channel %d data %v\n", net.LobbyID, net.UserID, net.ChannelID, net.Data)
//	    }
//	}()
func (c *LobbyClient) LobbyEventsChannel() *LobbyEventChannels {
	memberJoin := make(chan int64, 8)
	memberLeave := make(chan int64, 8)
	lobbyMessage := make(chan LobbyMessageEvent, 8)
	lobbyUpdate := make(chan int64, 8)
	lobbyDelete := make(chan LobbyDeleteEvent, 8)
	memberUpdate := make(chan MemberUpdateEvent, 8)
	speaking := make(chan SpeakingEvent, 8)
	networkMessage := make(chan NetworkMessageEvent, 8)

	if c.core != nil {
		events := &core.LobbyEvents{
			OnMemberConnect: func(lobbyID, userID int64) {
				memberJoin <- userID
			},
			OnMemberDisconnect: func(lobbyID, userID int64) {
				memberLeave <- userID
			},
			OnLobbyMessage: func(lobbyID, userID int64, data []byte) {
				lobbyMessage <- LobbyMessageEvent{LobbyID: lobbyID, UserID: userID, Data: data}
			},
			OnLobbyUpdate: func(lobbyID int64) {
				lobbyUpdate <- lobbyID
			},
			OnLobbyDelete: func(lobbyID int64, reason uint32) {
				lobbyDelete <- LobbyDeleteEvent{LobbyID: lobbyID, Reason: reason}
			},
			OnMemberUpdate: func(lobbyID, userID int64) {
				memberUpdate <- MemberUpdateEvent{LobbyID: lobbyID, UserID: userID}
			},
			OnSpeaking: func(lobbyID, userID int64, speakingVal bool) {
				speaking <- SpeakingEvent{LobbyID: lobbyID, UserID: userID, Speaking: speakingVal}
			},
			OnNetworkMessage: func(lobbyID, userID int64, channelID uint8, data []byte) {
				networkMessage <- NetworkMessageEvent{LobbyID: lobbyID, UserID: userID, ChannelID: channelID, Data: data}
			},
		}
		c.core.SetLobbyEvents(events)
	}

	return &LobbyEventChannels{
		MemberJoin:     memberJoin,
		MemberLeave:    memberLeave,
		LobbyMessage:   lobbyMessage,
		LobbyUpdate:    lobbyUpdate,
		LobbyDelete:    lobbyDelete,
		MemberUpdate:   memberUpdate,
		Speaking:       speaking,
		NetworkMessage: networkMessage,
	}
}
