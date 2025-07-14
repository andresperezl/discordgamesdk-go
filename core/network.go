package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// NetworkManager provides access to network-related functionality
type NetworkManager struct {
	manager unsafe.Pointer
}

// GetPeerID gets the local peer ID
func (n *NetworkManager) GetPeerID() uint64 {
	if n.manager == nil {
		return 0
	}

	var peerID uint64
	dcgo.NetworkManagerGetPeerID(n.manager, unsafe.Pointer(&peerID))
	return peerID
}

// Flush flushes pending network messages
func (n *NetworkManager) Flush() Result {
	if n.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.NetworkManagerFlush(n.manager))
}

// OpenPeer opens a connection to a remote peer
func (n *NetworkManager) OpenPeer(peerID uint64, routeData string) Result {
	if n.manager == nil {
		return ResultInternalError
	}
	return Result(dcgo.NetworkManagerOpenPeerHelper(n.manager, peerID, routeData))
}

// UpdatePeer updates the route data for a connected peer
func (n *NetworkManager) UpdatePeer(peerID uint64, routeData string) Result {
	if n.manager == nil {
		return ResultInternalError
	}
	return Result(dcgo.NetworkManagerUpdatePeerHelper(n.manager, peerID, routeData))
}

// ClosePeer closes the connection to a remote peer
func (n *NetworkManager) ClosePeer(peerID uint64) Result {
	if n.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.NetworkManagerClosePeer(n.manager, peerID))
}

// OpenChannel opens a message channel to a connected peer
func (n *NetworkManager) OpenChannel(peerID uint64, channelID uint8, reliable bool) Result {
	if n.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.NetworkManagerOpenChannel(n.manager, peerID, channelID, reliable))
}

// CloseChannel closes a message channel to a connected peer
func (n *NetworkManager) CloseChannel(peerID uint64, channelID uint8) Result {
	if n.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.NetworkManagerCloseChannel(n.manager, peerID, channelID))
}

// SendMessage sends a message to a connected peer
func (n *NetworkManager) SendMessage(peerID uint64, channelID uint8, data []byte) Result {
	if n.manager == nil {
		return ResultInternalError
	}

	var cData unsafe.Pointer
	if len(data) > 0 {
		cData = unsafe.Pointer(&data[0])
	}

	return Result(dcgo.NetworkManagerSendMessage(n.manager, peerID, channelID, cData, uint32(len(data))))
}
