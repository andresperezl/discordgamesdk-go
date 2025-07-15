package discordcgo

/*
#cgo CFLAGS: -I${SRCDIR}/../lib
#cgo LDFLAGS: -L${SRCDIR}/../lib -ldiscord_game_sdk
#include "discord_game_sdk.h"
#include "discord_wrappers.h"
#include <stdint.h>
#include <string.h>
#include <stdlib.h>

// Typedef for the Go callback trampoline
typedef void (*go_storage_read_async_callback_t)(void* go_callback_data, enum EDiscordResult result, uint8_t* data, uint32_t data_length);
extern void go_storage_read_async_callback_trampoline(void* go_callback_data, enum EDiscordResult result, uint8_t* data, uint32_t data_length);

int64_t get_discord_sku_id(struct DiscordSku* sku) { return sku->id; }
int32_t get_discord_sku_type(struct DiscordSku* sku) { return sku->type; }
void get_discord_sku_name(struct DiscordSku* sku, char* out, int outlen) { strncpy(out, sku->name, outlen-1); out[outlen-1] = '\0'; }
uint32_t get_discord_sku_price_amount(struct DiscordSku* sku) { return sku->price.amount; }
void get_discord_sku_price_currency(struct DiscordSku* sku, char* out, int outlen) { strncpy(out, sku->price.currency, outlen-1); out[outlen-1] = '\0'; }

int64_t get_discord_entitlement_id(struct DiscordEntitlement* ent) { return ent->id; }
int32_t get_discord_entitlement_type(struct DiscordEntitlement* ent) { return ent->type; }
int64_t get_discord_entitlement_sku_id(struct DiscordEntitlement* ent) { return ent->sku_id; }
uint64_t get_discord_file_stat_size(struct DiscordFileStat* stat) { return stat->size; }
uint64_t get_discord_file_stat_last_modified(struct DiscordFileStat* stat) { return stat->last_modified; }
void get_discord_file_stat_filename(struct DiscordFileStat* stat, char* out, int outlen) { strncpy(out, stat->filename, outlen-1); out[outlen-1] = '\0'; }
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// Dispatcher for serializing all SDK calls on a single OS thread
type sdkCall func()

type sdkDispatcher struct {
	calls chan sdkCall
	once  sync.Once
}

var dispatcher = &sdkDispatcher{
	calls: make(chan sdkCall, 128),
}

func (d *sdkDispatcher) start() {
	d.once.Do(func() {
		go func() {
			runtime.LockOSThread()
			for call := range d.calls {
				call()
			}
		}()
	})
}

func runOnDispatcher(call sdkCall) {
	dispatcher.start()
	dispatcher.calls <- call
}

// Run a function on the dispatcher and wait for its result
func RunOnDispatcherSync[T any](fn func() T) T {
	ch := make(chan T, 1)
	runOnDispatcher(func() {
		ch <- fn()
	})
	return <-ch
}

// Callback registry for Go async storage read
var storageReadAsyncCallbacks sync.Map // map[uintptr]func(result int32, data []byte)

//export go_storage_read_async_callback_trampoline
func go_storage_read_async_callback_trampoline(go_callback_data unsafe.Pointer, result C.enum_EDiscordResult, data *C.uint8_t, data_length C.uint32_t) {
	cbID := uintptr(go_callback_data)
	runOnDispatcher(func() {
		if cb, ok := storageReadAsyncCallbacks.Load(cbID); ok {
			callback := cb.(func(result int32, data []byte))
			var goData []byte
			if data != nil && data_length > 0 {
				goData = C.GoBytes(unsafe.Pointer(data), C.int(data_length))
			}
			callback(int32(result), goData)
			storageReadAsyncCallbacks.Delete(cbID)
		}
	})
}

// StorageManagerReadAsync with Go callback trampoline
func StorageManagerReadAsync(manager unsafe.Pointer, name *C.char, callback func(result int32, data []byte)) {
	cbID := uintptr(unsafe.Pointer(&callback))
	storageReadAsyncCallbacks.Store(cbID, callback)
	runOnDispatcher(func() {
		C.discord_storage_manager_read_async_trampoline((*C.struct_IDiscordStorageManager)(manager), name, unsafe.Pointer(cbID))
	})
}

// Callback registry for Go async storage write
var storageWriteAsyncCallbacks sync.Map // map[uintptr]func(result int32)

//export go_storage_write_async_callback_trampoline
func go_storage_write_async_callback_trampoline(go_callback_data unsafe.Pointer, result C.enum_EDiscordResult) {
	cbID := uintptr(go_callback_data)
	runOnDispatcher(func() {
		if cb, ok := storageWriteAsyncCallbacks.Load(cbID); ok {
			callback := cb.(func(result int32))
			callback(int32(result))
			storageWriteAsyncCallbacks.Delete(cbID)
		}
	})
}

// StorageManagerWriteAsync with Go callback trampoline
func StorageManagerWriteAsync(manager unsafe.Pointer, name *C.char, data unsafe.Pointer, dataLength uint32, callback func(result int32)) {
	cbID := uintptr(unsafe.Pointer(&callback))
	storageWriteAsyncCallbacks.Store(cbID, callback)
	runOnDispatcher(func() {
		C.discord_storage_manager_write_async_trampoline((*C.struct_IDiscordStorageManager)(manager), name, (*C.uint8_t)(data), C.uint32_t(dataLength), unsafe.Pointer(cbID))
	})
}

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

func CoreGetImageManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_image_manager(core))
}

func CoreGetRelationshipManager(core unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.discord_core_get_relationship_manager(core))
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

// LobbyTransactionSetMetadata sets metadata on a lobby transaction
func LobbyTransactionSetMetadata(transaction unsafe.Pointer, key *C.char, value *C.char) int32 {
	return int32(C.discord_lobby_transaction_set_metadata((*C.struct_IDiscordLobbyTransaction)(transaction), key, value))
}

// LobbyTransactionDeleteMetadata deletes metadata from a lobby transaction
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

// GoStringToCChar converts a Go string to a *C.char (null-terminated C string)
func GoStringToCChar(s string) *C.char {
	if s == "" {
		return nil
	}
	return C.CString(s)
}

// FreeCChar frees a *C.char allocated by GoStringToCChar
func FreeCChar(cstr *C.char) {
	if cstr != nil {
		C.free(unsafe.Pointer(cstr))
	}
}

// Additional storage manager wrappers
func StorageManagerReadAsyncPartial(manager unsafe.Pointer, name *C.char, offset uint64, length uint64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_storage_manager_read_async_partial((*C.struct_IDiscordStorageManager)(manager), name, C.uint64_t(offset), C.uint64_t(length), callbackData, (*[0]byte)(callback))
}

func StorageManagerStat(manager unsafe.Pointer, name *C.char, stat unsafe.Pointer) int32 {
	return int32(C.discord_storage_manager_stat((*C.struct_IDiscordStorageManager)(manager), name, (*C.struct_DiscordFileStat)(stat)))
}

func StorageManagerStatAt(manager unsafe.Pointer, index int32, stat unsafe.Pointer) int32 {
	return int32(C.discord_storage_manager_stat_at((*C.struct_IDiscordStorageManager)(manager), C.int32_t(index), (*C.struct_DiscordFileStat)(stat)))
}

func StorageManagerGetPath(manager unsafe.Pointer, path unsafe.Pointer) int32 {
	return int32(C.discord_storage_manager_get_path((*C.struct_IDiscordStorageManager)(manager), (*C.DiscordPath)(path)))
}

// Additional overlay manager wrappers
func OverlayManagerInitDrawingDXGI(manager unsafe.Pointer, swapchain unsafe.Pointer, useMessageForwarding bool) int32 {
	return int32(C.discord_overlay_manager_init_drawing_dxgi((*C.struct_IDiscordOverlayManager)(manager), swapchain, C.bool(useMessageForwarding)))
}

func OverlayManagerOnPresent(manager unsafe.Pointer) {
	C.discord_overlay_manager_on_present((*C.struct_IDiscordOverlayManager)(manager))
}

func OverlayManagerForwardMessage(manager unsafe.Pointer, message unsafe.Pointer) {
	C.discord_overlay_manager_forward_message((*C.struct_IDiscordOverlayManager)(manager), message)
}

func OverlayManagerKeyEvent(manager unsafe.Pointer, down bool, keyCode *C.char, variant int32) {
	C.discord_overlay_manager_key_event((*C.struct_IDiscordOverlayManager)(manager), C.bool(down), keyCode, C.enum_EDiscordKeyVariant(variant))
}

func OverlayManagerCharEvent(manager unsafe.Pointer, character *C.char) {
	C.discord_overlay_manager_char_event((*C.struct_IDiscordOverlayManager)(manager), character)
}

func OverlayManagerMouseButtonEvent(manager unsafe.Pointer, down uint8, clickCount int32, which int32, x int32, y int32) {
	C.discord_overlay_manager_mouse_button_event((*C.struct_IDiscordOverlayManager)(manager), C.uint8_t(down), C.int32_t(clickCount), C.enum_EDiscordMouseButton(which), C.int32_t(x), C.int32_t(y))
}

func OverlayManagerMouseMotionEvent(manager unsafe.Pointer, x int32, y int32) {
	C.discord_overlay_manager_mouse_motion_event((*C.struct_IDiscordOverlayManager)(manager), C.int32_t(x), C.int32_t(y))
}

func OverlayManagerImeCommitText(manager unsafe.Pointer, text *C.char) {
	C.discord_overlay_manager_ime_commit_text((*C.struct_IDiscordOverlayManager)(manager), text)
}

func OverlayManagerImeSetComposition(manager unsafe.Pointer, text *C.char, underlines unsafe.Pointer, underlinesLength uint32, from int32, to int32) {
	C.discord_overlay_manager_ime_set_composition((*C.struct_IDiscordOverlayManager)(manager), text, (*C.struct_DiscordImeUnderline)(underlines), C.uint32_t(underlinesLength), C.int32_t(from), C.int32_t(to))
}

func OverlayManagerImeCancelComposition(manager unsafe.Pointer) {
	C.discord_overlay_manager_ime_cancel_composition((*C.struct_IDiscordOverlayManager)(manager))
}

func OverlayManagerSetImeCompositionRangeCallback(manager unsafe.Pointer, onImeCompositionRangeChangedData unsafe.Pointer, onImeCompositionRangeChanged unsafe.Pointer) {
	C.discord_overlay_manager_set_ime_composition_range_callback((*C.struct_IDiscordOverlayManager)(manager), onImeCompositionRangeChangedData, (*[0]byte)(onImeCompositionRangeChanged))
}

func OverlayManagerSetImeSelectionBoundsCallback(manager unsafe.Pointer, onImeSelectionBoundsChangedData unsafe.Pointer, onImeSelectionBoundsChanged unsafe.Pointer) {
	C.discord_overlay_manager_set_ime_selection_bounds_callback((*C.struct_IDiscordOverlayManager)(manager), onImeSelectionBoundsChangedData, (*[0]byte)(onImeSelectionBoundsChanged))
}

func OverlayManagerIsPointInsideClickZone(manager unsafe.Pointer, x int32, y int32) bool {
	return bool(C.discord_overlay_manager_is_point_inside_click_zone((*C.struct_IDiscordOverlayManager)(manager), C.int32_t(x), C.int32_t(y)))
}

// Store manager wrappers
func StoreManagerFetchSkus(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_store_manager_fetch_skus((*C.struct_IDiscordStoreManager)(manager), callbackData, (*[0]byte)(callback))
}

func StoreManagerCountSkus(manager unsafe.Pointer, count unsafe.Pointer) {
	C.discord_store_manager_count_skus((*C.struct_IDiscordStoreManager)(manager), (*C.int32_t)(count))
}

func StoreManagerGetSku(manager unsafe.Pointer, skuID int64, sku unsafe.Pointer) int32 {
	return int32(C.discord_store_manager_get_sku((*C.struct_IDiscordStoreManager)(manager), C.DiscordSnowflake(skuID), (*C.struct_DiscordSku)(sku)))
}

func StoreManagerGetSkuAt(manager unsafe.Pointer, index int32, sku unsafe.Pointer) int32 {
	return int32(C.discord_store_manager_get_sku_at((*C.struct_IDiscordStoreManager)(manager), C.int32_t(index), (*C.struct_DiscordSku)(sku)))
}

func StoreManagerFetchEntitlements(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_store_manager_fetch_entitlements((*C.struct_IDiscordStoreManager)(manager), callbackData, (*[0]byte)(callback))
}

func StoreManagerCountEntitlements(manager unsafe.Pointer, count unsafe.Pointer) {
	C.discord_store_manager_count_entitlements((*C.struct_IDiscordStoreManager)(manager), (*C.int32_t)(count))
}

func StoreManagerGetEntitlement(manager unsafe.Pointer, entitlementID int64, entitlement unsafe.Pointer) int32 {
	return int32(C.discord_store_manager_get_entitlement((*C.struct_IDiscordStoreManager)(manager), C.DiscordSnowflake(entitlementID), (*C.struct_DiscordEntitlement)(entitlement)))
}

func StoreManagerGetEntitlementAt(manager unsafe.Pointer, index int32, entitlement unsafe.Pointer) int32 {
	return int32(C.discord_store_manager_get_entitlement_at((*C.struct_IDiscordStoreManager)(manager), C.int32_t(index), (*C.struct_DiscordEntitlement)(entitlement)))
}

func StoreManagerHasSkuEntitlement(manager unsafe.Pointer, skuID int64, hasEntitlement unsafe.Pointer) int32 {
	return int32(C.discord_store_manager_has_sku_entitlement((*C.struct_IDiscordStoreManager)(manager), C.DiscordSnowflake(skuID), (*C.bool)(hasEntitlement)))
}

func StoreManagerStartPurchase(manager unsafe.Pointer, skuID int64, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_store_manager_start_purchase((*C.struct_IDiscordStoreManager)(manager), C.DiscordSnowflake(skuID), callbackData, (*[0]byte)(callback))
}

// Voice manager wrappers
func VoiceManagerGetInputMode(manager unsafe.Pointer, inputMode unsafe.Pointer) int32 {
	return int32(C.discord_voice_manager_get_input_mode((*C.struct_IDiscordVoiceManager)(manager), (*C.struct_DiscordInputMode)(inputMode)))
}

func VoiceManagerSetInputMode(manager unsafe.Pointer, inputMode unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_voice_manager_set_input_mode((*C.struct_IDiscordVoiceManager)(manager), *(*C.struct_DiscordInputMode)(inputMode), callbackData, (*[0]byte)(callback))
}

func VoiceManagerIsSelfMute(manager unsafe.Pointer, mute unsafe.Pointer) int32 {
	return int32(C.discord_voice_manager_is_self_mute((*C.struct_IDiscordVoiceManager)(manager), (*C.bool)(mute)))
}

func VoiceManagerSetSelfMute(manager unsafe.Pointer, mute bool) int32 {
	return int32(C.discord_voice_manager_set_self_mute((*C.struct_IDiscordVoiceManager)(manager), C.bool(mute)))
}

func VoiceManagerIsSelfDeaf(manager unsafe.Pointer, deaf unsafe.Pointer) int32 {
	return int32(C.discord_voice_manager_is_self_deaf((*C.struct_IDiscordVoiceManager)(manager), (*C.bool)(deaf)))
}

func VoiceManagerSetSelfDeaf(manager unsafe.Pointer, deaf bool) int32 {
	return int32(C.discord_voice_manager_set_self_deaf((*C.struct_IDiscordVoiceManager)(manager), C.bool(deaf)))
}

func VoiceManagerIsLocalMute(manager unsafe.Pointer, userID int64, mute unsafe.Pointer) int32 {
	return int32(C.discord_voice_manager_is_local_mute((*C.struct_IDiscordVoiceManager)(manager), C.DiscordSnowflake(userID), (*C.bool)(mute)))
}

func VoiceManagerSetLocalMute(manager unsafe.Pointer, userID int64, mute bool) int32 {
	return int32(C.discord_voice_manager_set_local_mute((*C.struct_IDiscordVoiceManager)(manager), C.DiscordSnowflake(userID), C.bool(mute)))
}

func VoiceManagerGetLocalVolume(manager unsafe.Pointer, userID int64, volume unsafe.Pointer) int32 {
	return int32(C.discord_voice_manager_get_local_volume((*C.struct_IDiscordVoiceManager)(manager), C.DiscordSnowflake(userID), (*C.uint8_t)(volume)))
}

func VoiceManagerSetLocalVolume(manager unsafe.Pointer, userID int64, volume uint8) int32 {
	return int32(C.discord_voice_manager_set_local_volume((*C.struct_IDiscordVoiceManager)(manager), C.DiscordSnowflake(userID), C.uint8_t(volume)))
}

// Achievement manager wrappers
func AchievementManagerSetUserAchievement(manager unsafe.Pointer, achievementID int64, percentComplete uint8, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_achievement_manager_set_user_achievement((*C.struct_IDiscordAchievementManager)(manager), C.DiscordSnowflake(achievementID), C.uint8_t(percentComplete), callbackData, (*[0]byte)(callback))
}

func AchievementManagerFetchUserAchievements(manager unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_achievement_manager_fetch_user_achievements((*C.struct_IDiscordAchievementManager)(manager), callbackData, (*[0]byte)(callback))
}

func AchievementManagerCountUserAchievements(manager unsafe.Pointer, count unsafe.Pointer) {
	C.discord_achievement_manager_count_user_achievements((*C.struct_IDiscordAchievementManager)(manager), (*C.int32_t)(count))
}

func AchievementManagerGetUserAchievement(manager unsafe.Pointer, userAchievementID int64, userAchievement unsafe.Pointer) int32 {
	return int32(C.discord_achievement_manager_get_user_achievement((*C.struct_IDiscordAchievementManager)(manager), C.DiscordSnowflake(userAchievementID), (*C.struct_DiscordUserAchievement)(userAchievement)))
}

func AchievementManagerGetUserAchievementAt(manager unsafe.Pointer, index int32, userAchievement unsafe.Pointer) int32 {
	return int32(C.discord_achievement_manager_get_user_achievement_at((*C.struct_IDiscordAchievementManager)(manager), C.int32_t(index), (*C.struct_DiscordUserAchievement)(userAchievement)))
}

// Go-friendly storage manager wrappers
func StorageManagerReadGo(manager unsafe.Pointer, name string, data []byte, read *uint32) int32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var datPtr unsafe.Pointer
	if len(data) > 0 {
		datPtr = unsafe.Pointer(&data[0])
	}
	return StorageManagerRead(manager, cname, datPtr, uint32(len(data)), unsafe.Pointer(read))
}

func StorageManagerWriteGo(manager unsafe.Pointer, name string, data []byte) int32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var datPtr unsafe.Pointer
	if len(data) > 0 {
		datPtr = unsafe.Pointer(&data[0])
	}
	return StorageManagerWrite(manager, cname, datPtr, uint32(len(data)))
}

func StorageManagerDeleteGo(manager unsafe.Pointer, name string) int32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return StorageManagerDelete_(manager, cname)
}

func StorageManagerExistsGo(manager unsafe.Pointer, name string, exists *bool) int32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var cExists C.bool
	res := StorageManagerExists(manager, cname, unsafe.Pointer(&cExists))
	*exists = bool(cExists)
	return res
}

func StorageManagerStatGo(manager unsafe.Pointer, name string, stat unsafe.Pointer) int32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return StorageManagerStat(manager, cname, stat)
}

func StorageManagerStatAtGo(manager unsafe.Pointer, index int32, stat unsafe.Pointer) int32 {
	return StorageManagerStatAt(manager, index, stat)
}

func StorageManagerGetPathGo(manager unsafe.Pointer, path unsafe.Pointer) int32 {
	return StorageManagerGetPath(manager, path)
}

// NetworkManagerOpenPeerHelper is a Go-friendly wrapper for NetworkManagerOpenPeer
func NetworkManagerOpenPeerHelper(manager unsafe.Pointer, peerID uint64, routeData string) int32 {
	cRoute := GoStringToCChar(routeData)
	defer FreeCChar(cRoute)
	return NetworkManagerOpenPeer(manager, peerID, cRoute)
}

// NetworkManagerUpdatePeerHelper is a Go-friendly wrapper for NetworkManagerUpdatePeer
func NetworkManagerUpdatePeerHelper(manager unsafe.Pointer, peerID uint64, routeData string) int32 {
	cRoute := GoStringToCChar(routeData)
	defer FreeCChar(cRoute)
	return NetworkManagerUpdatePeer(manager, peerID, cRoute)
}

// DiscordSku field accessors
func GetDiscordSkuID(ptr unsafe.Pointer) int64 {
	return int64(C.get_discord_sku_id((*C.struct_DiscordSku)(ptr)))
}
func GetDiscordSkuType(ptr unsafe.Pointer) int32 {
	return int32(C.get_discord_sku_type((*C.struct_DiscordSku)(ptr)))
}
func GetDiscordSkuName(ptr unsafe.Pointer) string {
	var buf [256]C.char
	C.get_discord_sku_name((*C.struct_DiscordSku)(ptr), &buf[0], 256)
	return C.GoString(&buf[0])
}
func GetDiscordSkuPriceAmount(ptr unsafe.Pointer) uint32 {
	return uint32(C.get_discord_sku_price_amount((*C.struct_DiscordSku)(ptr)))
}
func GetDiscordSkuPriceCurrency(ptr unsafe.Pointer) string {
	var buf [16]C.char
	C.get_discord_sku_price_currency((*C.struct_DiscordSku)(ptr), &buf[0], 16)
	return C.GoString(&buf[0])
}

// DiscordEntitlement field accessors
func GetDiscordEntitlementID(ptr unsafe.Pointer) int64 {
	return int64(C.get_discord_entitlement_id((*C.struct_DiscordEntitlement)(ptr)))
}
func GetDiscordEntitlementType(ptr unsafe.Pointer) int32 {
	return int32(C.get_discord_entitlement_type((*C.struct_DiscordEntitlement)(ptr)))
}
func GetDiscordEntitlementSkuID(ptr unsafe.Pointer) int64 {
	return int64(C.get_discord_entitlement_sku_id((*C.struct_DiscordEntitlement)(ptr)))
}

// FileStat field accessors
func GetDiscordFileStatFilename(stat *DiscordFileStat) string {
	var buf [260]C.char
	C.get_discord_file_stat_filename((*C.struct_DiscordFileStat)(unsafe.Pointer(stat)), &buf[0], 260)
	return C.GoString(&buf[0])
}
func GetDiscordFileStatSize(stat *DiscordFileStat) uint64 {
	return uint64(C.get_discord_file_stat_size((*C.struct_DiscordFileStat)(unsafe.Pointer(stat))))
}
func GetDiscordFileStatLastModified(stat *DiscordFileStat) uint64 {
	return uint64(C.get_discord_file_stat_last_modified((*C.struct_DiscordFileStat)(unsafe.Pointer(stat))))
}

// Type aliases for C structs

type DiscordSku C.struct_DiscordSku

type DiscordEntitlement C.struct_DiscordEntitlement

// Malloc and Free helpers for DiscordSku and DiscordEntitlement
func MallocDiscordSku() unsafe.Pointer {
	return C.malloc(C.sizeof_struct_DiscordSku)
}

func MallocDiscordEntitlement() unsafe.Pointer {
	return C.malloc(C.sizeof_struct_DiscordEntitlement)
}

func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

// Get typed pointer from unsafe.Pointer
func GetDiscordSku(ptr unsafe.Pointer) *DiscordSku {
	return (*DiscordSku)(ptr)
}

func GetDiscordEntitlement(ptr unsafe.Pointer) *DiscordEntitlement {
	return (*DiscordEntitlement)(ptr)
}

// Type aliases for C structs used in storage

type DiscordFileStat C.struct_DiscordFileStat

type DiscordPath C.DiscordPath

// GoString helper for C strings
func GoString(cstr *C.char) string {
	return C.GoString(cstr)
}

// GoStringFromBytes helper for byte buffers
func GoStringFromBytes(b *byte) string {
	return C.GoString((*C.char)(unsafe.Pointer(b)))
}

// Go-friendly StoreManager SKU helpers
func StoreManagerGetSkuGo(manager unsafe.Pointer, skuID int64) *DiscordSku {
	ptr := MallocDiscordSku()
	defer Free(ptr)
	res := StoreManagerGetSku(manager, skuID, ptr)
	if res != 0 {
		return nil
	}
	return GetDiscordSku(ptr)
}

func StoreManagerGetSkuAtGo(manager unsafe.Pointer, index int32) *DiscordSku {
	ptr := MallocDiscordSku()
	defer Free(ptr)
	res := StoreManagerGetSkuAt(manager, index, ptr)
	if res != 0 {
		return nil
	}
	return GetDiscordSku(ptr)
}

// Go-friendly StoreManager Entitlement helpers
func StoreManagerGetEntitlementGo(manager unsafe.Pointer, entitlementID int64) *DiscordEntitlement {
	ptr := MallocDiscordEntitlement()
	defer Free(ptr)
	res := StoreManagerGetEntitlement(manager, entitlementID, ptr)
	if res != 0 {
		return nil
	}
	return GetDiscordEntitlement(ptr)
}

func StoreManagerGetEntitlementAtGo(manager unsafe.Pointer, index int32) *DiscordEntitlement {
	ptr := MallocDiscordEntitlement()
	defer Free(ptr)
	res := StoreManagerGetEntitlementAt(manager, index, ptr)
	if res != 0 {
		return nil
	}
	return GetDiscordEntitlement(ptr)
}

func LobbyManagerGetLobbyGo(manager unsafe.Pointer, lobbyID int64) (id int64, typ int32, ownerID int64, secret string, capacity uint32, locked bool, res int32) {
	var cLobby C.struct_DiscordLobby
	res = LobbyManagerGetLobby(manager, lobbyID, unsafe.Pointer(&cLobby))
	id = int64(cLobby.id)
	typ = int32(cLobby._type)
	ownerID = int64(cLobby.owner_id)
	secret = C.GoString(&cLobby.secret[0])
	capacity = uint32(cLobby.capacity)
	locked = bool(cLobby.locked)
	return
}

func LobbyManagerGetLobbyActivitySecret(manager unsafe.Pointer, lobbyID int64, secret unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_lobby_activity_secret((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.DiscordLobbySecret)(secret)))
}

// Image manager wrappers
func ImageManagerFetch(manager unsafe.Pointer, handle unsafe.Pointer, refresh bool, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_image_manager_fetch((*C.struct_IDiscordImageManager)(manager), *(*C.struct_DiscordImageHandle)(handle), C.bool(refresh), callbackData, (*[0]byte)(callback))
}

func ImageManagerGetDimensions(manager unsafe.Pointer, handle unsafe.Pointer, dimensions unsafe.Pointer) int32 {
	return int32(C.discord_image_manager_get_dimensions((*C.struct_IDiscordImageManager)(manager), *(*C.struct_DiscordImageHandle)(handle), (*C.struct_DiscordImageDimensions)(dimensions)))
}

func ImageManagerGetData(manager unsafe.Pointer, handle unsafe.Pointer, data unsafe.Pointer, dataLength uint32) int32 {
	return int32(C.discord_image_manager_get_data((*C.struct_IDiscordImageManager)(manager), *(*C.struct_DiscordImageHandle)(handle), (*C.uint8_t)(data), C.uint32_t(dataLength)))
}

// Relationship manager wrappers
func RelationshipManagerFilter(manager unsafe.Pointer, filterData unsafe.Pointer, filter unsafe.Pointer) {
	C.discord_relationship_manager_filter((*C.struct_IDiscordRelationshipManager)(manager), filterData, (*[0]byte)(filter))
}

func RelationshipManagerCount(manager unsafe.Pointer, count unsafe.Pointer) int32 {
	return int32(C.discord_relationship_manager_count((*C.struct_IDiscordRelationshipManager)(manager), (*C.int32_t)(count)))
}

func RelationshipManagerGet(manager unsafe.Pointer, userID int64, relationship unsafe.Pointer) int32 {
	return int32(C.discord_relationship_manager_get((*C.struct_IDiscordRelationshipManager)(manager), C.DiscordUserId(userID), (*C.struct_DiscordRelationship)(relationship)))
}

func RelationshipManagerGetAt(manager unsafe.Pointer, index uint32, relationship unsafe.Pointer) int32 {
	return int32(C.discord_relationship_manager_get_at((*C.struct_IDiscordRelationshipManager)(manager), C.uint32_t(index), (*C.struct_DiscordRelationship)(relationship)))
}

// Additional missing lobby manager wrappers
func LobbyManagerConnectLobbyWithActivitySecret(manager unsafe.Pointer, activitySecret unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_connect_lobby_with_activity_secret((*C.struct_IDiscordLobbyManager)(manager), (*C.char)(activitySecret), callbackData, (*[0]byte)(callback))
}

func LobbyManagerGetMemberUpdateTransaction(manager unsafe.Pointer, lobbyID int64, userID int64, transaction unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_member_update_transaction((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.DiscordUserId(userID), (**C.struct_IDiscordLobbyMemberTransaction)(transaction)))
}

// NOTE: key must be a pointer to [256]C.char, value to [4096]C.char, cast as *C.char
func LobbyManagerGetLobbyMetadataValue(manager unsafe.Pointer, lobbyID int64, key unsafe.Pointer, value unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_lobby_metadata_value((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.char)(key), (*C.char)(value)))
}

func LobbyManagerGetLobbyMetadataKey(manager unsafe.Pointer, lobbyID int64, index int32, key unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_lobby_metadata_key((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.int32_t(index), (*C.char)(key)))
}

func LobbyManagerLobbyMetadataCount(manager unsafe.Pointer, lobbyID int64, count unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_lobby_metadata_count((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.int32_t)(count)))
}

func LobbyManagerMemberCount(manager unsafe.Pointer, lobbyID int64, count unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_member_count((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), (*C.int32_t)(count)))
}

func LobbyManagerGetMemberUserID(manager unsafe.Pointer, lobbyID int64, index int32, userID unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_member_user_id((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.int32_t(index), (*C.DiscordUserId)(userID)))
}

func LobbyManagerGetMemberUser(manager unsafe.Pointer, lobbyID int64, userID int64, user unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_member_user((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.DiscordUserId(userID), (*C.struct_DiscordUser)(user)))
}

// NOTE: key must be a pointer to [256]C.char, value to [4096]C.char, cast as *C.char
func LobbyManagerGetMemberMetadataValue(manager unsafe.Pointer, lobbyID int64, userID int64, key unsafe.Pointer, value unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_member_metadata_value((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.DiscordUserId(userID), (*C.char)(key), (*C.char)(value)))
}

func LobbyManagerGetMemberMetadataKey(manager unsafe.Pointer, lobbyID int64, userID int64, index int32, key unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_member_metadata_key((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.DiscordUserId(userID), C.int32_t(index), (*C.char)(key)))
}

func LobbyManagerMemberMetadataCount(manager unsafe.Pointer, lobbyID int64, userID int64, count unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_member_metadata_count((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.DiscordUserId(userID), (*C.int32_t)(count)))
}

func LobbyManagerUpdateMember(manager unsafe.Pointer, lobbyID int64, userID int64, transaction unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_update_member((*C.struct_IDiscordLobbyManager)(manager), C.DiscordLobbyId(lobbyID), C.DiscordUserId(userID), (*C.struct_IDiscordLobbyMemberTransaction)(transaction), callbackData, (*[0]byte)(callback))
}

func LobbyManagerGetSearchQuery(manager unsafe.Pointer, query unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_search_query((*C.struct_IDiscordLobbyManager)(manager), (**C.struct_IDiscordLobbySearchQuery)(query)))
}

func LobbyManagerSearch(manager unsafe.Pointer, query unsafe.Pointer, callbackData unsafe.Pointer, callback unsafe.Pointer) {
	C.discord_lobby_manager_search((*C.struct_IDiscordLobbyManager)(manager), (*C.struct_IDiscordLobbySearchQuery)(query), callbackData, (*[0]byte)(callback))
}

func LobbyManagerLobbyCount(manager unsafe.Pointer, count unsafe.Pointer) {
	C.discord_lobby_manager_lobby_count((*C.struct_IDiscordLobbyManager)(manager), (*C.int32_t)(count))
}

func LobbyManagerGetLobbyID(manager unsafe.Pointer, index int32, lobbyID unsafe.Pointer) int32 {
	return int32(C.discord_lobby_manager_get_lobby_id((*C.struct_IDiscordLobbyManager)(manager), C.int32_t(index), (*C.DiscordLobbyId)(lobbyID)))
}
