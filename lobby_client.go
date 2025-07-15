package discord

import (
	"fmt"
	"unsafe"

	core "github.com/andresperezl/discordctl/core"
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

// NOTE: The Discord Game SDK does not provide APIs for lobby metadata value by index, lobby message history, or direct search result access.
// Methods such as GetLobbyMetadataValueByIndex, GetLobbyMemberMetadataValueByIndex, GetLobbyMessageCount, GetLobbyMessageUserId, GetLobbyMessageData,
// Search, SearchWithFilter, GetSearchResultCount, and GetSearchResult have been removed because they cannot be implemented with the current SDK.
//
// TODO: Methods like SetLobbyMetadata, DeleteLobbyMetadata, SetLobbyMemberMetadata, and DeleteLobbyMemberMetadata should be implemented using the transaction pattern.
