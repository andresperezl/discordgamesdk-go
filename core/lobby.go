package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
	discordlog "github.com/andresperezl/discordctl/discordlog"
)

// LobbyManager provides access to lobby-related functionality
type LobbyManager struct {
	manager unsafe.Pointer
}

// CreateLobby creates a new lobby
func (l *LobbyManager) CreateLobby(transaction *LobbyTransaction, callback func(result Result, lobby *Lobby)) {
	discordlog.GetLogger().Info("LobbyManager.CreateLobby called", "transaction", transaction)
	if l.manager == nil {
		if callback != nil {
			discordlog.GetLogger().Warn("LobbyManager.CreateLobby: manager is nil")
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerCreateLobby(l.manager, cTransaction, callbackData, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerUpdateLobby(l.manager, lobbyID, cTransaction, callbackData, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerDeleteLobby(l.manager, lobbyID, callbackData, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerConnectLobby(l.manager, lobbyID, unsafe.Pointer(&cSecret[0]), callbackData, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerDisconnectLobby(l.manager, lobbyID, callbackData, nil)
		return nil
	})
}

// GetLobby gets a lobby by ID
func (l *LobbyManager) GetLobby(lobbyID int64) (*Lobby, Result) {
	discordlog.GetLogger().Info("LobbyManager.GetLobby called", "lobbyID", lobbyID)
	if l.manager == nil {
		discordlog.GetLogger().Warn("LobbyManager.GetLobby: manager is nil")
		return nil, ResultInternalError
	}

	var id int64
	var typ int32
	var ownerID int64
	var secret string
	var capacity uint32
	var locked bool
	var res int32
	dcgo.RunOnDispatcherSync(func() any {
		id, typ, ownerID, secret, capacity, locked, res = dcgo.LobbyManagerGetLobbyGo(l.manager, lobbyID)
		return nil
	})
	if res != 0 {
		return nil, Result(res)
	}
	return &Lobby{
		ID:       id,
		Type:     LobbyType(typ),
		OwnerID:  ownerID,
		Secret:   secret,
		Capacity: capacity,
		Locked:   locked,
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerSendLobbyMessage(l.manager, lobbyID, cData, uint32(len(data)), callbackData, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerConnectVoice(l.manager, lobbyID, callbackData, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerDisconnectVoice(l.manager, lobbyID, callbackData, nil)
		return nil
	})
}

// ConnectNetwork connects network to a lobby
func (l *LobbyManager) ConnectNetwork(lobbyID int64) Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerConnectNetwork(l.manager, lobbyID)
	}))
}

// DisconnectNetwork disconnects network from a lobby
func (l *LobbyManager) DisconnectNetwork(lobbyID int64) Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerDisconnectNetwork(l.manager, lobbyID)
	}))
}

// FlushNetwork flushes network messages
func (l *LobbyManager) FlushNetwork() Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerFlushNetwork(l.manager)
	}))
}

// OpenNetworkChannel opens a network channel
func (l *LobbyManager) OpenNetworkChannel(lobbyID int64, channelID uint8, reliable bool) Result {
	if l.manager == nil {
		return ResultInternalError
	}

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerOpenNetworkChannel(l.manager, lobbyID, channelID, reliable)
	}))
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

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerSendNetworkMessage(l.manager, lobbyID, userID, channelID, cData, uint32(len(data)))
	}))
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
	result := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetLobbyCreateTransaction(l.manager, unsafe.Pointer(&transaction))
	})

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
	result := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetLobbyUpdateTransaction(l.manager, lobbyID, unsafe.Pointer(&transaction))
	})

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

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyTransactionSetType(t.transaction, int32(lobbyType))
	}))
}

// SetOwner sets the lobby owner
func (t *LobbyTransaction) SetOwner(ownerID int64) Result {
	if t.transaction == nil {
		return ResultInternalError
	}

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyTransactionSetOwner(t.transaction, ownerID)
	}))
}

// SetCapacity sets the lobby capacity
func (t *LobbyTransaction) SetCapacity(capacity uint32) Result {
	if t.transaction == nil {
		return ResultInternalError
	}

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyTransactionSetCapacity(t.transaction, capacity)
	}))
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
	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyTransactionSetMetadata(t.transaction, cKey, cValue)
	}))
}

// DeleteMetadata deletes lobby metadata
func (t *LobbyTransaction) DeleteMetadata(key string) Result {
	if t.transaction == nil {
		return ResultInternalError
	}
	cKey := dcgo.GoStringToCChar(key)
	defer dcgo.FreeCChar(cKey)
	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyTransactionDeleteMetadata(t.transaction, cKey)
	}))
}

// SetLocked sets the lobby locked state
func (t *LobbyTransaction) SetLocked(locked bool) Result {
	if t.transaction == nil {
		return ResultInternalError
	}

	return Result(dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyTransactionSetLocked(t.transaction, locked)
	}))
}

// Core-level Go functions for new IDiscordLobbyManager methods
// (import block to be removed)
// import (
// 	"unsafe"
// 	dcgo "github.com/andresperezl/discordctl/discordcgo"
// )

// ConnectLobbyWithActivitySecret connects to a lobby using an activity secret
func (lm *LobbyManager) ConnectLobbyWithActivitySecret(activitySecret string, callbackData, callback unsafe.Pointer) {
	cSecret := dcgo.GoStringToCChar(activitySecret)
	defer dcgo.FreeCChar(cSecret)
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerConnectLobbyWithActivitySecret(lm.manager, unsafe.Pointer(cSecret), callbackData, callback)
		return nil
	})
}

// GetMemberUpdateTransaction gets a member update transaction
func (lm *LobbyManager) GetMemberUpdateTransaction(lobbyID, userID int64) unsafe.Pointer {
	var transaction unsafe.Pointer
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerGetMemberUpdateTransaction(lm.manager, lobbyID, userID, unsafe.Pointer(&transaction))
		return nil
	})
	return transaction
}

// GetLobbyMetadataValue retrieves a metadata value for a lobby
func (lm *LobbyManager) GetLobbyMetadataValue(lobbyID int64, key string) (string, int32) {
	var value [4096]byte
	cKey := dcgo.GoStringToCChar(key)
	defer dcgo.FreeCChar(cKey)
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetLobbyMetadataValue(lm.manager, lobbyID, unsafe.Pointer(cKey), unsafe.Pointer(&value[0]))
	})
	return dcgo.GoString(cKey), res
}

// GetLobbyMetadataKey retrieves a metadata key for a lobby by index
func (lm *LobbyManager) GetLobbyMetadataKey(lobbyID int64, index int32) (string, int32) {
	var key [256]byte
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetLobbyMetadataKey(lm.manager, lobbyID, index, unsafe.Pointer(&key[0]))
	})
	return dcgo.GoStringFromBytes(&key[0]), res
}

// LobbyMetadataCount returns the number of metadata entries for a lobby
func (lm *LobbyManager) LobbyMetadataCount(lobbyID int64) (int32, int32) {
	var count int32
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerLobbyMetadataCount(lm.manager, lobbyID, unsafe.Pointer(&count))
	})
	return count, res
}

// MemberCount returns the number of members in a lobby
func (lm *LobbyManager) MemberCount(lobbyID int64) (int32, int32) {
	var count int32
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerMemberCount(lm.manager, lobbyID, unsafe.Pointer(&count))
	})
	return count, res
}

// GetMemberUserID retrieves a user ID for a member by index
func (lm *LobbyManager) GetMemberUserID(lobbyID int64, index int32) (int64, int32) {
	var userID int64
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetMemberUserID(lm.manager, lobbyID, index, unsafe.Pointer(&userID))
	})
	return userID, res
}

// GetMemberUser retrieves a user struct for a member
func (lm *LobbyManager) GetMemberUser(lobbyID, userID int64) (*User, int32) {
	var user User
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetMemberUser(lm.manager, lobbyID, userID, unsafe.Pointer(&user))
	})
	return &user, res
}

// GetMemberMetadataValue retrieves a metadata value for a member
func (lm *LobbyManager) GetMemberMetadataValue(lobbyID, userID int64, key string) (string, int32) {
	var value [4096]byte
	cKey := dcgo.GoStringToCChar(key)
	defer dcgo.FreeCChar(cKey)
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetMemberMetadataValue(lm.manager, lobbyID, userID, unsafe.Pointer(cKey), unsafe.Pointer(&value[0]))
	})
	return dcgo.GoString(cKey), res
}

// GetMemberMetadataKey retrieves a metadata key for a member by index
func (lm *LobbyManager) GetMemberMetadataKey(lobbyID, userID int64, index int32) (string, int32) {
	var key [256]byte
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetMemberMetadataKey(lm.manager, lobbyID, userID, index, unsafe.Pointer(&key[0]))
	})
	return dcgo.GoStringFromBytes(&key[0]), res
}

// MemberMetadataCount returns the number of metadata entries for a member
func (lm *LobbyManager) MemberMetadataCount(lobbyID, userID int64) (int32, int32) {
	var count int32
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerMemberMetadataCount(lm.manager, lobbyID, userID, unsafe.Pointer(&count))
	})
	return count, res
}

// UpdateMember updates a member using a transaction
func (lm *LobbyManager) UpdateMember(lobbyID, userID int64, transaction unsafe.Pointer, callbackData, callback unsafe.Pointer) {
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerUpdateMember(lm.manager, lobbyID, userID, transaction, callbackData, callback)
		return nil
	})
}

// GetSearchQuery gets a search query object
func (lm *LobbyManager) GetSearchQuery() unsafe.Pointer {
	var query unsafe.Pointer
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerGetSearchQuery(lm.manager, unsafe.Pointer(&query))
		return nil
	})
	return query
}

// Search performs a lobby search
func (lm *LobbyManager) Search(query unsafe.Pointer, callbackData, callback unsafe.Pointer) {
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerSearch(lm.manager, query, callbackData, callback)
		return nil
	})
}

// LobbyCount returns the number of lobbies
func (lm *LobbyManager) LobbyCount() int32 {
	var count int32
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.LobbyManagerLobbyCount(lm.manager, unsafe.Pointer(&count))
		return nil
	})
	return count
}

// GetLobbyID retrieves a lobby ID by index
func (lm *LobbyManager) GetLobbyID(index int32) (int64, int32) {
	var lobbyID int64
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetLobbyID(lm.manager, index, unsafe.Pointer(&lobbyID))
	})
	return lobbyID, res
}

// GetLobbyActivitySecret retrieves the activity secret for a lobby
func (lm *LobbyManager) GetLobbyActivitySecret(lobbyID int64) (string, int32) {
	var secret [4096]byte
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.LobbyManagerGetLobbyActivitySecret(lm.manager, lobbyID, unsafe.Pointer(&secret[0]))
	})
	return dcgo.GoStringFromBytes(&secret[0]), res
}
