package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordgamesdk-go/discordcgo"
	discordlog "github.com/andresperezl/discordgamesdk-go/discordlog"
)

// NetworkManager provides access to network-related functionality
type NetworkManager struct {
	manager unsafe.Pointer
}

// GetPeerID gets the local peer ID
func (n *NetworkManager) GetPeerID() uint64 {
	discordlog.GetLogger().Info("NetworkManager.GetPeerID called")
	if n.manager == nil {
		discordlog.GetLogger().Warn("NetworkManager.GetPeerID: manager is nil")
		return 0
	}

	var peerID uint64
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.NetworkManagerGetPeerID(n.manager, unsafe.Pointer(&peerID))
		return nil
	})
	return peerID
}

// Flush flushes pending network messages
func (n *NetworkManager) Flush() Result {
	if n.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.NetworkManagerFlush(n.manager)
	})
	return Result(res)
}

// OpenPeer opens a connection to a remote peer
func (n *NetworkManager) OpenPeer(peerID uint64, routeData string) Result {
	discordlog.GetLogger().Info("NetworkManager.OpenPeer called", "peerID", peerID, "routeData", routeData)
	if n.manager == nil {
		discordlog.GetLogger().Warn("NetworkManager.OpenPeer: manager is nil")
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.NetworkManagerOpenPeerHelper(n.manager, peerID, routeData)
	})
	return Result(res)
}

// UpdatePeer updates the route data for a connected peer
func (n *NetworkManager) UpdatePeer(peerID uint64, routeData string) Result {
	if n.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.NetworkManagerUpdatePeerHelper(n.manager, peerID, routeData)
	})
	return Result(res)
}

// ClosePeer closes the connection to a remote peer
func (n *NetworkManager) ClosePeer(peerID uint64) Result {
	if n.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.NetworkManagerClosePeer(n.manager, peerID)
	})
	return Result(res)
}

// OpenChannel opens a message channel to a connected peer
func (n *NetworkManager) OpenChannel(peerID uint64, channelID uint8, reliable bool) Result {
	if n.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.NetworkManagerOpenChannel(n.manager, peerID, channelID, reliable)
	})
	return Result(res)
}

// CloseChannel closes a message channel to a connected peer
func (n *NetworkManager) CloseChannel(peerID uint64, channelID uint8) Result {
	if n.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.NetworkManagerCloseChannel(n.manager, peerID, channelID)
	})
	return Result(res)
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
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.NetworkManagerSendMessage(n.manager, peerID, channelID, cData, uint32(len(data)))
	})
	return Result(res)
}
