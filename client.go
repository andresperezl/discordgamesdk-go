package discord

import (
	"context"
	"fmt"
	"time"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// Client represents a Discord client with Go-like interfaces
type Client struct {
	core        *core.Core
	clientID    int64
	initialized bool
	ctx         context.Context
	cancel      context.CancelFunc
}

// ClientConfig holds configuration for creating a Discord client
type ClientConfig struct {
	ClientID int64
	Flags    core.CreateFlags
	Timeout  time.Duration
}

// DefaultClientConfig returns a default configuration
func DefaultClientConfig(clientID int64) *ClientConfig {
	return &ClientConfig{
		ClientID: clientID,
		Flags:    core.CreateFlagsDefault,
		Timeout:  10 * time.Second,
	}
}

// NewClient creates a new Discord client with Go-like interfaces
func NewClient(config *ClientConfig) (*Client, error) {
	if config == nil {
		config = DefaultClientConfig(0)
	}

	ctx, cancel := context.WithCancel(context.Background())

	coreObj, result := core.Create(config.ClientID, config.Flags, nil)
	if result != core.ResultOk {
		cancel()
		return nil, fmt.Errorf("failed to create Discord core: %v", result)
	}

	client := &Client{
		core:     coreObj,
		clientID: config.ClientID,
		ctx:      ctx,
		cancel:   cancel,
	}

	// Start the callback loop
	client.core.Start()

	// Wait for initialization
	if !client.core.WaitForInitialization(config.Timeout) {
		client.Close()
		return nil, fmt.Errorf("failed to initialize Discord SDK within %v", config.Timeout)
	}

	client.initialized = true
	return client, nil
}

// Close shuts down the client and cleans up resources
func (c *Client) Close() error {
	if c.cancel != nil {
		c.cancel()
	}
	if c.core != nil {
		c.core.Shutdown()
	}
	return nil
}

// IsInitialized returns whether the client is fully initialized
func (c *Client) IsInitialized() bool {
	return c.initialized
}

// GetCurrentUser returns the current user, waiting if necessary
func (c *Client) GetCurrentUser(timeout time.Duration) (*core.User, error) {
	if !c.initialized {
		return nil, fmt.Errorf("client not initialized")
	}

	user, result := c.core.WaitForUser(timeout)
	if result != core.ResultOk {
		return nil, fmt.Errorf("failed to get current user: %v", result)
	}

	return user, nil
}

// GetCurrentUserAsync returns a channel that will receive the current user
func (c *Client) GetCurrentUserAsync() <-chan *core.User {
	ch := make(chan *core.User, 1)

	go func() {
		defer close(ch)
		user, err := c.GetCurrentUser(5 * time.Second)
		if err != nil {
			// Send nil to indicate error
			ch <- nil
			return
		}
		ch <- user
	}()

	return ch
}

// Activity returns an activity manager with Go-like methods
func (c *Client) Activity() *ActivityClient {
	return &ActivityClient{
		manager: c.core.GetActivityManager(),
		core:    c.core,
	}
}

// User returns a user manager with Go-like methods
func (c *Client) User() *UserClient {
	return &UserClient{
		manager: c.core.GetUserManager(),
		core:    c.core,
	}
}

// Application returns an application manager with Go-like methods
func (c *Client) Application() *ApplicationClient {
	return &ApplicationClient{
		manager: c.core.GetApplicationManager(),
		core:    c.core,
	}
}

// Storage returns a storage manager with Go-like methods
func (c *Client) Storage() *StorageClient {
	return &StorageClient{
		manager: c.core.GetStorageManager(),
		core:    c.core,
	}
}

// Lobby returns a lobby manager with Go-like methods
func (c *Client) Lobby() *LobbyClient {
	return &LobbyClient{
		manager: c.core.GetLobbyManager(),
		core:    c.core,
	}
}

// Network returns a network manager with Go-like methods
func (c *Client) Network() *NetworkClient {
	return &NetworkClient{
		manager: c.core.GetNetworkManager(),
		core:    c.core,
	}
}

// Overlay returns an overlay manager with Go-like methods
func (c *Client) Overlay() *OverlayClient {
	return &OverlayClient{
		manager: c.core.GetOverlayManager(),
		core:    c.core,
	}
}

// Store returns a store manager with Go-like methods
func (c *Client) Store() *StoreClient {
	return &StoreClient{
		manager: c.core.GetStoreManager(),
		core:    c.core,
	}
}

// Voice returns a voice manager with Go-like methods
func (c *Client) Voice() *VoiceClient {
	return &VoiceClient{
		manager: c.core.GetVoiceManager(),
		core:    c.core,
	}
}

// Achievement returns an achievement manager with Go-like methods
func (c *Client) Achievement() *AchievementClient {
	return &AchievementClient{
		manager: c.core.GetAchievementManager(),
		core:    c.core,
	}
}

// Run starts the client's event loop
func (c *Client) Run() {
	// Run until context is cancelled
	<-c.ctx.Done()
}

// RunWithTimeout runs the client for a specified duration
func (c *Client) RunWithTimeout(timeout time.Duration) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-c.ctx.Done():
	case <-timer.C:
		c.Close()
	}
}
