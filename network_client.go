package discord

import (
	"fmt"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// NOTE: The Discord Game SDK does not provide APIs to query open peers, channel counts, or message counts.
// Methods such as IsPeerOpen, GetOpenPeerCount, GetOpenPeerId, GetChannelCount, GetChannelId, GetChannelMessageCount, GetChannelMessage, and ClearChannelMessages
// have been removed because they cannot be implemented with the current SDK.

// NetworkClient provides Go-like interfaces for network management
type NetworkClient struct {
	manager *core.NetworkManager
	core    *core.Core
}

// GetPeerID gets the local peer ID
func (nc *NetworkClient) GetPeerID() (uint64, error) {
	if nc.manager == nil {
		return 0, fmt.Errorf("network manager not available")
	}
	return nc.manager.GetPeerID(), nil
}

// Flush flushes the network
func (nc *NetworkClient) Flush() error {
	if nc.manager == nil {
		return fmt.Errorf("network manager not available")
	}
	result := nc.manager.Flush()
	if result != core.ResultOk {
		return fmt.Errorf("failed to flush: %v", result)
	}
	return nil
}

// OpenPeer opens a peer connection
func (nc *NetworkClient) OpenPeer(peerID uint64, route string) error {
	if nc.manager == nil {
		return fmt.Errorf("network manager not available")
	}
	result := nc.manager.OpenPeer(peerID, route)
	if result != core.ResultOk {
		return fmt.Errorf("failed to open peer: %v", result)
	}
	return nil
}

// UpdatePeer updates a peer connection
func (nc *NetworkClient) UpdatePeer(peerID uint64, route string) error {
	if nc.manager == nil {
		return fmt.Errorf("network manager not available")
	}
	result := nc.manager.UpdatePeer(peerID, route)
	if result != core.ResultOk {
		return fmt.Errorf("failed to update peer: %v", result)
	}
	return nil
}

// ClosePeer closes a peer connection
func (nc *NetworkClient) ClosePeer(peerID uint64) error {
	if nc.manager == nil {
		return fmt.Errorf("network manager not available")
	}
	result := nc.manager.ClosePeer(peerID)
	if result != core.ResultOk {
		return fmt.Errorf("failed to close peer: %v", result)
	}
	return nil
}

// OpenChannel opens a channel to a peer
func (nc *NetworkClient) OpenChannel(peerID uint64, channelID uint8, reliable bool) error {
	if nc.manager == nil {
		return fmt.Errorf("network manager not available")
	}
	result := nc.manager.OpenChannel(peerID, channelID, reliable)
	if result != core.ResultOk {
		return fmt.Errorf("failed to open channel: %v", result)
	}
	return nil
}

// CloseChannel closes a channel to a peer
func (nc *NetworkClient) CloseChannel(peerID uint64, channelID uint8) error {
	if nc.manager == nil {
		return fmt.Errorf("network manager not available")
	}
	result := nc.manager.CloseChannel(peerID, channelID)
	if result != core.ResultOk {
		return fmt.Errorf("failed to close channel: %v", result)
	}
	return nil
}

// SendMessage sends a message to a peer
func (nc *NetworkClient) SendMessage(peerID uint64, channelID uint8, data []byte) error {
	if nc.manager == nil {
		return fmt.Errorf("network manager not available")
	}
	result := nc.manager.SendMessage(peerID, channelID, data)
	if result != core.ResultOk {
		return fmt.Errorf("failed to send message: %v", result)
	}
	return nil
}
