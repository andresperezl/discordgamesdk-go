package core

import (
	"unsafe"

	"github.com/andresperezl/discordctl/discordcgo"
	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// LobbyManager provides access to lobby-related functionality
type LobbyManager struct {
	manager unsafe.Pointer
}

// CreateLobby creates a new lobby
func (l *LobbyManager) CreateLobby(transaction *LobbyTransaction, callback func(result Result, lobby *Lobby)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError, nil)
		}
		return
	}

	var cTransaction unsafe.Pointer
	if transaction != nil {
		cTransaction = transaction.transaction
	}

	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerCreateLobby(l.manager, cTransaction, callbackData, nil)
}

// UpdateLobby updates a lobby
func (l *LobbyManager) UpdateLobby(lobbyID int64, transaction *LobbyTransaction, callback func(result Result)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	var cTransaction unsafe.Pointer
	if transaction != nil {
		cTransaction = transaction.transaction
	}

	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerUpdateLobby(l.manager, lobbyID, cTransaction, callbackData, nil)
}

// DeleteLobby deletes a lobby
func (l *LobbyManager) DeleteLobby(lobbyID int64, callback func(result Result)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerDeleteLobby(l.manager, lobbyID, callbackData, nil)
}

// ConnectLobby connects to a lobby
func (l *LobbyManager) ConnectLobby(lobbyID int64, secret string, callback func(result Result, lobby *Lobby)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError, nil)
		}
		return
	}

	cSecret := []byte(secret)
	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerConnectLobby(l.manager, lobbyID, unsafe.Pointer(&cSecret[0]), callbackData, nil)
}

// DisconnectLobby disconnects from a lobby
func (l *LobbyManager) DisconnectLobby(lobbyID int64, callback func(result Result)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerDisconnectLobby(l.manager, lobbyID, callbackData, nil)
}

// GetLobby gets a lobby by ID
func (l *LobbyManager) GetLobby(lobbyID int64) (*Lobby, Result) {
	if l.manager == nil {
		return nil, ResultInternalError
	}

	var cLobby struct {
		ID       int64
		Type     int32
		OwnerID  int64
		Secret   [128]byte
		Capacity uint32
		Locked   bool
	}

	result := discordcgo.LobbyManagerGetLobby(l.manager, lobbyID, unsafe.Pointer(&cLobby))

	if result != int32(ResultOk) {
		return nil, Result(result)
	}

	return &Lobby{
		ID:       cLobby.ID,
		Type:     LobbyType(cLobby.Type),
		OwnerID:  cLobby.OwnerID,
		Secret:   string(cLobby.Secret[:]),
		Capacity: cLobby.Capacity,
		Locked:   cLobby.Locked,
	}, ResultOk
}

// SendLobbyMessage sends a message to a lobby
func (l *LobbyManager) SendLobbyMessage(lobbyID int64, data []byte, callback func(result Result)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	var cData unsafe.Pointer
	if len(data) > 0 {
		cData = unsafe.Pointer(&data[0])
	}

	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerSendLobbyMessage(l.manager, lobbyID, cData, uint32(len(data)), callbackData, nil)
}

// ConnectVoice connects voice to a lobby
func (l *LobbyManager) ConnectVoice(lobbyID int64, callback func(result Result)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerConnectVoice(l.manager, lobbyID, callbackData, nil)
}

// DisconnectVoice disconnects voice from a lobby
func (l *LobbyManager) DisconnectVoice(lobbyID int64, callback func(result Result)) {
	if l.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	var callbackData unsafe.Pointer
	if callback != nil {
		callbackData = unsafe.Pointer(&callback)
	}

	discordcgo.LobbyManagerDisconnectVoice(l.manager, lobbyID, callbackData, nil)
}

// ConnectNetwork connects network to a lobby
func (l *LobbyManager) ConnectNetwork(lobbyID int64) Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyManagerConnectNetwork(l.manager, lobbyID))
}

// DisconnectNetwork disconnects network from a lobby
func (l *LobbyManager) DisconnectNetwork(lobbyID int64) Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyManagerDisconnectNetwork(l.manager, lobbyID))
}

// FlushNetwork flushes network messages
func (l *LobbyManager) FlushNetwork() Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyManagerFlushNetwork(l.manager))
}

// OpenNetworkChannel opens a network channel
func (l *LobbyManager) OpenNetworkChannel(lobbyID int64, channelID uint8, reliable bool) Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyManagerOpenNetworkChannel(l.manager, lobbyID, channelID, reliable))
}

// SendNetworkMessage sends a network message
func (l *LobbyManager) SendNetworkMessage(lobbyID int64, userID int64, channelID uint8, data []byte) Result {
	if l.manager == nil {
		return ResultInternalError
	}

	var cData unsafe.Pointer
	if len(data) > 0 {
		cData = unsafe.Pointer(&data[0])
	}

	return Result(discordcgo.LobbyManagerSendNetworkMessage(l.manager, lobbyID, userID, channelID, cData, uint32(len(data))))
}

// LobbyTransaction represents a lobby transaction
type LobbyTransaction struct {
	transaction unsafe.Pointer
}

// GetLobbyCreateTransaction gets a lobby create transaction
func (l *LobbyManager) GetLobbyCreateTransaction() (*LobbyTransaction, Result) {
	if l.manager == nil {
		return nil, ResultInternalError
	}

	var transaction unsafe.Pointer
	result := discordcgo.LobbyManagerGetLobbyCreateTransaction(l.manager, unsafe.Pointer(&transaction))

	if result != int32(ResultOk) {
		return nil, Result(result)
	}

	return &LobbyTransaction{transaction: transaction}, ResultOk
}

// GetLobbyUpdateTransaction gets a lobby update transaction
func (l *LobbyManager) GetLobbyUpdateTransaction(lobbyID int64) (*LobbyTransaction, Result) {
	if l.manager == nil {
		return nil, ResultInternalError
	}

	var transaction unsafe.Pointer
	result := discordcgo.LobbyManagerGetLobbyUpdateTransaction(l.manager, lobbyID, unsafe.Pointer(&transaction))

	if result != int32(ResultOk) {
		return nil, Result(result)
	}

	return &LobbyTransaction{transaction: transaction}, ResultOk
}

// SetType sets the lobby type
func (t *LobbyTransaction) SetType(lobbyType LobbyType) Result {
	if t.transaction == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyTransactionSetType(t.transaction, int32(lobbyType)))
}

// SetOwner sets the lobby owner
func (t *LobbyTransaction) SetOwner(ownerID int64) Result {
	if t.transaction == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyTransactionSetOwner(t.transaction, ownerID))
}

// SetCapacity sets the lobby capacity
func (t *LobbyTransaction) SetCapacity(capacity uint32) Result {
	if t.transaction == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyTransactionSetCapacity(t.transaction, capacity))
}

// SetMetadata sets lobby metadata
func (t *LobbyTransaction) SetMetadata(key, value string) Result {
	if t.transaction == nil {
		return ResultInternalError
	}
	cKey := dcgo.GoStringToCChar(key)
	defer dcgo.FreeCChar(cKey)
	cValue := dcgo.GoStringToCChar(value)
	defer dcgo.FreeCChar(cValue)
	return Result(discordcgo.LobbyTransactionSetMetadata(t.transaction, cKey, cValue))
}

// DeleteMetadata deletes lobby metadata
func (t *LobbyTransaction) DeleteMetadata(key string) Result {
	if t.transaction == nil {
		return ResultInternalError
	}
	cKey := dcgo.GoStringToCChar(key)
	defer dcgo.FreeCChar(cKey)
	return Result(discordcgo.LobbyTransactionDeleteMetadata(t.transaction, cKey))
}

// SetLocked sets the lobby locked state
func (t *LobbyTransaction) SetLocked(locked bool) Result {
	if t.transaction == nil {
		return ResultInternalError
	}

	return Result(discordcgo.LobbyTransactionSetLocked(t.transaction, locked))
}
