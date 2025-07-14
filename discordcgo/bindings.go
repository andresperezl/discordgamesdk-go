package discordcgo

/*
#cgo CFLAGS: -I${SRCDIR}/../lib
#cgo LDFLAGS: -L${SRCDIR}/../lib -ldiscord_game_sdk
#include "discord_game_sdk.h"
#include "discord_wrappers.h"
*/
import "C"
import "unsafe"

// Core wrappers
func CoreCreate(version int32, params unsafe.Pointer, result unsafe.Pointer) int32 {
	return int32(C.discord_core_create(C.DiscordVersion(version), (*C.struct_DiscordCreateParams)(params), (**C.struct_IDiscordCore)(result)))
}

func CoreCreateHelper(clientID int64, flags uint64) (unsafe.Pointer, int32) {
	// Create the DiscordCreateParams structure
	var params C.struct_DiscordCreateParams
	C.DiscordCreateParamsSetDefault(&params)

	// Set the client ID
	params.client_id = C.DiscordClientId(clientID)
	params.flags = C.uint64_t(flags)

	// Set default versions for all managers
	params.application_version = C.DISCORD_APPLICATION_MANAGER_VERSION
	params.user_version = C.DISCORD_USER_MANAGER_VERSION
	params.image_version = C.DISCORD_IMAGE_MANAGER_VERSION
	params.activity_version = C.DISCORD_ACTIVITY_MANAGER_VERSION
	params.relationship_version = C.DISCORD_RELATIONSHIP_MANAGER_VERSION
	params.lobby_version = C.DISCORD_LOBBY_MANAGER_VERSION
	params.network_version = C.DISCORD_NETWORK_MANAGER_VERSION
	params.overlay_version = C.DISCORD_OVERLAY_MANAGER_VERSION
	params.storage_version = C.DISCORD_STORAGE_MANAGER_VERSION
	params.store_version = C.DISCORD_STORE_MANAGER_VERSION
	params.voice_version = C.DISCORD_VOICE_MANAGER_VERSION
	params.achievement_version = C.DISCORD_ACHIEVEMENT_MANAGER_VERSION

	// Create the core
	var core *C.struct_IDiscordCore
	result := C.discord_core_create(3, &params, &core)

	return unsafe.Pointer(core), int32(result)
}

func CoreDestroy(core unsafe.Pointer) {
	C.discord_core_destroy(core)
}

func CoreRunCallbacks(core unsafe.Pointer) int32 {
	return int32(C.discord_core_run_callbacks(core))
}

func CoreSetLogHook(core unsafe.Pointer, minLevel int32, hookData unsafe.Pointer, hook unsafe.Pointer) {
	C.discord_core_set_log_hook(core, C.enum_EDiscordLogLevel(minLevel), hookData, (*[0]byte)(hook))
}

func CoreGetApplicationManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_application_manager(core))
}

func CoreGetUserManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_user_manager(core))
}

func CoreGetActivityManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_activity_manager(core))
}

func CoreGetLobbyManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_lobby_manager(core))
}

func CoreGetNetworkManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_network_manager(core))
}

func CoreGetOverlayManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_overlay_manager(core))
}

func CoreGetStorageManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_storage_manager(core))
}

func CoreGetStoreManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_store_manager(core))
}

func CoreGetVoiceManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_voice_manager(core))
}

func CoreGetAchievementManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_achievement_manager(core))
}

// Application manager wrappers
func ApplicationManagerValidateOrExit(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_application_manager_validate_or_exit((*C.struct_IDiscordApplicationManager)(manager), callbackData, (*[0]byte)(callback))
}

func ApplicationManagerGetCurrentLocale(manager unsafe.Pointer, locale unsafe.Pointer) {
	C.discord_application_manager_get_current_locale((*C.struct_IDiscordApplicationManager)(manager), (*C.DiscordLocale)(locale))
}

func ApplicationManagerGetCurrentBranch(manager unsafe.Pointer, branch unsafe.Pointer) {
	C.discord_application_manager_get_current_branch((*C.struct_IDiscordApplicationManager)(manager), (*C.DiscordBranch)(branch))
}

func ApplicationManagerGetOAuth2Token(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_application_manager_get_oauth2_token((*C.struct_IDiscordApplicationManager)(manager), callbackData, (*[0]byte)(callback))
}

func ApplicationManagerGetTicket(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_application_manager_get_ticket((*C.struct_IDiscordApplicationManager)(manager), callbackData, (*[0]byte)(callback))
}

// User manager wrappers
func UserManagerGetCurrentUser(manager unsafe.Pointer, user unsafe.Pointer) int32 {
	return int32(C.discord_user_manager_get_current_user((*C.struct_IDiscordUserManager)(manager), (*C.struct_DiscordUser)(user)))
}

func UserManagerGetUser(manager unsafe.Pointer, userID int64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_user_manager_get_user((*C.struct_IDiscordUserManager)(manager), C.DiscordUserId(userID), callbackData, (*[0]byte)(callback))
}

func UserManagerGetCurrentUserPremiumType(manager unsafe.Pointer, premiumType unsafe.Pointer) int32 {
	return int32(C.discord_user_manager_get_current_user_premium_type((*C.struct_IDiscordUserManager)(manager), (*C.enum_EDiscordPremiumType)(premiumType)))
}

func UserManagerCurrentUserHasFlag(manager unsafe.Pointer, flag int32, hasFlag unsafe.Pointer) int32 {
	return int32(C.discord_user_manager_current_user_has_flag((*C.struct_IDiscordUserManager)(manager), C.enum_EDiscordUserFlag(flag), (*C.bool)(hasFlag)))
}

// Activity manager wrappers
func ActivityManagerRegisterCommand(manager unsafe.Pointer, command *C.char) int32 {
	return int32(C.discord_activity_manager_register_command((*C.struct_IDiscordActivityManager)(manager), command))
}

func ActivityManagerRegisterSteam(manager unsafe.Pointer, steamID uint32) int32 {
	return int32(C.discord_activity_manager_register_steam((*C.struct_IDiscordActivityManager)(manager), C.uint32_t(steamID)))
}

func ActivityManagerUpdateActivity(manager unsafe.Pointer, activity unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_activity_manager_update_activity((*C.struct_IDiscordActivityManager)(manager), (*C.struct_DiscordActivity)(activity), callbackData, (*[0]byte)(callback))
}

func ActivityManagerClearActivity(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_activity_manager_clear_activity((*C.struct_IDiscordActivityManager)(manager), callbackData, (*[0]byte)(callback))
}

func ActivityManagerSendRequestReply(manager unsafe.Pointer, userID int64, reply int32, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_activity_manager_send_request_reply((*C.struct_IDiscordActivityManager)(manager), C.DiscordUserId(userID), C.enum_EDiscordActivityJoinRequestReply(reply), callbackData, (*[0]byte)(callback))
}

func ActivityManagerSendInvite(manager unsafe.Pointer, userID int64, actionType int32, content *C.char, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_activity_manager_send_invite((*C.struct_IDiscordActivityManager)(manager), C.DiscordUserId(userID), C.enum_EDiscordActivityActionType(actionType), content, callbackData, (*[0]byte)(callback))
}

func ActivityManagerAcceptInvite(manager unsafe.Pointer, userID int64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_activity_manager_accept_invite((*C.struct_IDiscordActivityManager)(manager), C.DiscordUserId(userID), callbackData, (*[0]byte)(callback))
}

// Lobby manager wrappers
func LobbyManagerGetLobbyCreateTransaction(manager unsafe.Pointer, transaction unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_lobby_create_transaction((*C.struct_IDiscordLobbyManager)(manager), (**C.struct_IDiscordLobbyTransaction)(transaction)))
}

func LobbyManagerCreateLobby(manager unsafe.Pointer, transaction unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_create_lobby((*C.struct_IDiscordLobbyManager)(manager), (*C.struct_IDiscordLobbyTransaction)(transaction), callbackData, (*[0]byte)(callback))
}

func LobbyManagerConnectLobby(manager unsafe.Pointer, lobbyID int64, secret unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_connect_lobby((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.char)(secret), callbackData, (*[0]byte)(callback))
}

func LobbyManagerGetLobby(manager unsafe.Pointer, lobbyID int64, lobby unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_lobby((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.struct_DiscordLobby)(lobby)))
}

func LobbyManagerUpdateLobby(manager unsafe.Pointer, lobbyID int64, transaction unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_update_lobby((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.struct_IDiscordLobbyTransaction)(transaction), callbackData, (*[0]byte)(callback))
}

func LobbyManagerDeleteLobby(manager unsafe.Pointer, lobbyID int64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_delete_lobby((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), callbackData, (*[0]byte)(callback))
}

func LobbyManagerDisconnectLobby(manager unsafe.Pointer, lobbyID int64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_disconnect_lobby((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), callbackData, (*[0]byte)(callback))
}

func LobbyManagerSendLobbyMessage(manager unsafe.Pointer, lobbyID int64, data unsafe.Pointer, dataLength uint32, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_send_lobby_message((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.uint8_t)(data), C.uint32_t(dataLength), callbackData, (*[0]byte)(callback))
}

func LobbyManagerConnectVoice(manager unsafe.Pointer, lobbyID int64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_connect_voice((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), callbackData, (*[0]byte)(callback))
}

func LobbyManagerDisconnectVoice(manager unsafe.Pointer, lobbyID int64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_disconnect_voice((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), callbackData, (*[0]byte)(callback))
}

func LobbyManagerConnectNetwork(manager unsafe.Pointer, lobbyID int64) int32 {
	return int32(C.discord_lobby_manager_connect_network((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID)))
}

func LobbyManagerDisconnectNetwork(manager unsafe.Pointer, lobbyID int64) int32 {
	return int32(C.discord_lobby_manager_disconnect_network((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID)))
}

func LobbyManagerFlushNetwork(manager unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_flush_network((*C.struct_IDiscordLobbyManager)(manager)))
}

func LobbyManagerOpenNetworkChannel(manager unsafe.Pointer, lobbyID int64, channelID uint8, reliable bool) int32 {
	return int32(C.discord_lobby_manager_open_network_channel((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.uint8_t(channelID), C.bool(reliable)))
}

func LobbyManagerSendNetworkMessage(manager unsafe.Pointer, lobbyID int64, userID int64, channelID uint8, data unsafe.Pointer, dataLength uint32) int32 {
	return int32(C.discord_lobby_manager_send_network_message((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.DiscordUserId(userID), C.uint8_t(channelID), (*C.uint8_t)(data), C.uint32_t(dataLength)))
}

func LobbyManagerGetLobbyUpdateTransaction(manager unsafe.Pointer, lobbyID int64, transaction unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_lobby_update_transaction((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (**C.struct_IDiscordLobbyTransaction)(transaction)))
}

func LobbyTransactionSetType(transaction unsafe.Pointer, lobbyType int32) int32 {
	return int32(C.discord_lobby_transaction_set_type((*C.struct_IDiscordLobbyTransaction)(transaction), C.enum_EDiscordLobbyType(lobbyType)))
}

func LobbyTransactionSetOwner(transaction unsafe.Pointer, ownerID int64) int32 {
	return int32(C.discord_lobby_transaction_set_owner((*C.struct_IDiscordLobbyTransaction)(transaction), C.DiscordUserId(ownerID)))
}

func LobbyTransactionSetCapacity(transaction unsafe.Pointer, capacity uint32) int32 {
	return int32(C.discord_lobby_transaction_set_capacity((*C.struct_IDiscordLobbyTransaction)(transaction), C.uint32_t(capacity)))
}

func LobbyTransactionSetMetadata(transaction unsafe.Pointer, key *C.char, value *C.char) int32 {
	return int32(C.discord_lobby_transaction_set_metadata((*C.struct_IDiscordLobbyTransaction)(transaction), key, value))
}

func LobbyTransactionDeleteMetadata(transaction unsafe.Pointer, key *C.char) int32 {
	return int32(C.discord_lobby_transaction_delete_metadata((*C.struct_IDiscordLobbyTransaction)(transaction), key))
}

func LobbyTransactionSetLocked(transaction unsafe.Pointer, locked bool) int32 {
	return int32(C.discord_lobby_transaction_set_locked((*C.struct_IDiscordLobbyTransaction)(transaction), C.bool(locked)))
}

// Network manager wrappers
func NetworkManagerGetPeerID(manager unsafe.Pointer, peerID unsafe.Pointer) {
	C.discord_network_manager_get_peer_id((*C.struct_IDiscordNetworkManager)(manager), (*C.DiscordNetworkPeerId)(peerID))
}

func NetworkManagerFlush(manager unsafe.Pointer) int32 {
	return int32(C.discord_network_manager_flush((*C.struct_IDiscordNetworkManager)(manager)))
}

func NetworkManagerOpenPeer(manager unsafe.Pointer, peerID uint64, routeData *C.char) int32 {
	return int32(C.discord_network_manager_open_peer((*C.struct_IDiscordNetworkManager)(manager), C.DiscordNetworkPeerId(peerID), routeData))
}

func NetworkManagerUpdatePeer(manager unsafe.Pointer, peerID uint64, routeData *C.char) int32 {
	return int32(C.discord_network_manager_update_peer((*C.struct_IDiscordNetworkManager)(manager), C.DiscordNetworkPeerId(peerID), routeData))
}

func NetworkManagerClosePeer(manager unsafe.Pointer, peerID uint64) int32 {
	return int32(C.discord_network_manager_close_peer((*C.struct_IDiscordNetworkManager)(manager), C.DiscordNetworkPeerId(peerID)))
}

func NetworkManagerOpenChannel(manager unsafe.Pointer, peerID uint64, channelID uint8, reliable bool) int32 {
	return int32(C.discord_network_manager_open_channel((*C.struct_IDiscordNetworkManager)(manager), C.DiscordNetworkPeerId(peerID), C.DiscordNetworkChannelId(channelID), C.bool(reliable)))
}

func NetworkManagerCloseChannel(manager unsafe.Pointer, peerID uint64, channelID uint8) int32 {
	return int32(C.discord_network_manager_close_channel((*C.struct_IDiscordNetworkManager)(manager), C.DiscordNetworkPeerId(peerID), C.DiscordNetworkChannelId(channelID)))
}

func NetworkManagerSendMessage(manager unsafe.Pointer, peerID uint64, channelID uint8, data unsafe.Pointer, dataLength uint32) int32 {
	return int32(C.discord_network_manager_send_message((*C.struct_IDiscordNetworkManager)(manager), C.DiscordNetworkPeerId(peerID), C.DiscordNetworkChannelId(channelID), (*C.uint8_t)(data), C.uint32_t(dataLength)))
}

// Overlay manager wrappers
func OverlayManagerIsEnabled(manager unsafe.Pointer, enabled unsafe.Pointer) {
	C.discord_overlay_manager_is_enabled((*C.struct_IDiscordOverlayManager)(manager), (*C.bool)(enabled))
}

func OverlayManagerIsLocked(manager unsafe.Pointer, locked unsafe.Pointer) {
	C.discord_overlay_manager_is_locked((*C.struct_IDiscordOverlayManager)(manager), (*C.bool)(locked))
}

func OverlayManagerSetLocked(manager unsafe.Pointer, locked bool, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_overlay_manager_set_locked((*C.struct_IDiscordOverlayManager)(manager), C.bool(locked), callbackData, (*[0]byte)(callback))
}

func OverlayManagerOpenActivityInvite(manager unsafe.Pointer, actionType int32, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_overlay_manager_open_activity_invite((*C.struct_IDiscordOverlayManager)(manager), C.enum_EDiscordActivityActionType(actionType), callbackData, (*[0]byte)(callback))
}

func OverlayManagerOpenGuildInvite(manager unsafe.Pointer, code *C.char, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_overlay_manager_open_guild_invite((*C.struct_IDiscordOverlayManager)(manager), code, callbackData, (*[0]byte)(callback))
}

func OverlayManagerOpenVoiceSettings(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_overlay_manager_open_voice_settings((*C.struct_IDiscordOverlayManager)(manager), callbackData, (*[0]byte)(callback))
}

func OverlayManagerOpenGuildInviteHelper(manager unsafe.Pointer, code string, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	codeBytes := []byte(code)
	var codePtr *C.char
	if len(codeBytes) > 0 {
		codePtr = (*C.char)(unsafe.Pointer(&codeBytes[0]))
	} else {
		codePtr = nil
	}
	C.discord_overlay_manager_open_guild_invite((*C.struct_IDiscordOverlayManager)(manager), codePtr, callbackData, (*[0]byte)(callback))
}

// Storage manager wrappers
func StorageManagerRead(manager unsafe.Pointer, name *C.char, data unsafe.Pointer, dataLength uint32, read unsafe.Pointer) int32 {
	return int32(C.discord_storage_manager_read((*C.struct_IDiscordStorageManager)(manager), name, (*C.uint8_t)(data), C.uint32_t(dataLength), (*C.uint32_t)(read)))
}

func StorageManagerWrite(manager unsafe.Pointer, name *C.char, data unsafe.Pointer, dataLength uint32) int32 {
	return int32(C.discord_storage_manager_write((*C.struct_IDiscordStorageManager)(manager), name, (*C.uint8_t)(data), C.uint32_t(dataLength)))
}

func StorageManagerDelete_(manager unsafe.Pointer, name *C.char) int32 {
	return int32(C.discord_storage_manager_delete_((*C.struct_IDiscordStorageManager)(manager), name))
}

func StorageManagerExists(manager unsafe.Pointer, name *C.char, exists unsafe.Pointer) int32 {
	return int32(C.discord_storage_manager_exists((*C.struct_IDiscordStorageManager)(manager), name, (*C.bool)(exists)))
}

func StorageManagerCount(manager unsafe.Pointer, count unsafe.Pointer) {
	C.discord_storage_manager_count((*C.struct_IDiscordStorageManager)(manager), (*C.int32_t)(count))
}

// String conversion helper functions
func StringToCChar(s string) unsafe.Pointer {
	if s == "" {
		return unsafe.Pointer(nil)
	}
	bytes := []byte(s)
	return unsafe.Pointer(&bytes[0])
}

func StringToCCharPtr(s string) unsafe.Pointer {
	if s == "" {
		return unsafe.Pointer(nil)
	}
	bytes := []byte(s)
	return unsafe.Pointer(&bytes[0])
}
