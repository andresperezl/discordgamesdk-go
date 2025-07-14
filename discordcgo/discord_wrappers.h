#ifndef DISCORD_WRAPPERS_H
#define DISCORD_WRAPPERS_H

#include "discord_game_sdk.h"

// Core wrapper functions
enum EDiscordResult discord_core_create(DiscordVersion version, struct DiscordCreateParams* params, struct IDiscordCore** result);
void discord_core_destroy(void* core);
enum EDiscordResult discord_core_run_callbacks(void* core);
void discord_core_set_log_hook(void* core, enum EDiscordLogLevel min_level, void* hook_data, void (*hook)(void* hook_data, enum EDiscordLogLevel level, const char* message));
struct IDiscordApplicationManager* discord_core_get_application_manager(void* core);
struct IDiscordUserManager* discord_core_get_user_manager(void* core);
struct IDiscordActivityManager* discord_core_get_activity_manager(void* core);
struct IDiscordLobbyManager* discord_core_get_lobby_manager(void* core);
struct IDiscordNetworkManager* discord_core_get_network_manager(void* core);
struct IDiscordOverlayManager* discord_core_get_overlay_manager(void* core);
struct IDiscordStorageManager* discord_core_get_storage_manager(void* core);
struct IDiscordStoreManager* discord_core_get_store_manager(void* core);
struct IDiscordVoiceManager* discord_core_get_voice_manager(void* core);
struct IDiscordAchievementManager* discord_core_get_achievement_manager(void* core);

// Application manager wrappers
void discord_application_manager_validate_or_exit(struct IDiscordApplicationManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_application_manager_get_current_locale(struct IDiscordApplicationManager* manager, DiscordLocale* locale);
void discord_application_manager_get_current_branch(struct IDiscordApplicationManager* manager, DiscordBranch* branch);
void discord_application_manager_get_oauth2_token(struct IDiscordApplicationManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordOAuth2Token* oauth2_token));
void discord_application_manager_get_ticket(struct IDiscordApplicationManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, const char* data));

// User manager wrappers
enum EDiscordResult discord_user_manager_get_current_user(struct IDiscordUserManager* manager, struct DiscordUser* current_user);
void discord_user_manager_get_user(struct IDiscordUserManager* manager, DiscordUserId user_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordUser* user));
enum EDiscordResult discord_user_manager_get_current_user_premium_type(struct IDiscordUserManager* manager, enum EDiscordPremiumType* premium_type);
enum EDiscordResult discord_user_manager_current_user_has_flag(struct IDiscordUserManager* manager, enum EDiscordUserFlag flag, bool* has_flag);

// Activity manager wrappers
enum EDiscordResult discord_activity_manager_register_command(struct IDiscordActivityManager* manager, const char* command);
enum EDiscordResult discord_activity_manager_register_steam(struct IDiscordActivityManager* manager, uint32_t steam_id);
void discord_activity_manager_update_activity(struct IDiscordActivityManager* manager, struct DiscordActivity* activity, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_activity_manager_clear_activity(struct IDiscordActivityManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_activity_manager_send_request_reply(struct IDiscordActivityManager* manager, DiscordUserId user_id, enum EDiscordActivityJoinRequestReply reply, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_activity_manager_send_invite(struct IDiscordActivityManager* manager, DiscordUserId user_id, enum EDiscordActivityActionType type, const char* content, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_activity_manager_accept_invite(struct IDiscordActivityManager* manager, DiscordUserId user_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));

// Lobby manager wrappers
enum EDiscordResult discord_lobby_manager_get_lobby_create_transaction(struct IDiscordLobbyManager* manager, struct IDiscordLobbyTransaction** transaction);
void discord_lobby_manager_create_lobby(struct IDiscordLobbyManager* manager, struct IDiscordLobbyTransaction* transaction, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordLobby* lobby));
void discord_lobby_manager_connect_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordLobbySecret secret, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordLobby* lobby));
enum EDiscordResult discord_lobby_manager_get_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, struct DiscordLobby* lobby);
void discord_lobby_manager_update_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, struct IDiscordLobbyTransaction* transaction, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_lobby_manager_delete_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_lobby_manager_disconnect_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_lobby_manager_send_lobby_message(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, uint8_t* data, uint32_t data_length, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_lobby_manager_connect_voice(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_lobby_manager_disconnect_voice(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
enum EDiscordResult discord_lobby_manager_connect_network(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id);
enum EDiscordResult discord_lobby_manager_disconnect_network(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id);
enum EDiscordResult discord_lobby_manager_flush_network(struct IDiscordLobbyManager* manager);
enum EDiscordResult discord_lobby_manager_open_network_channel(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, uint8_t channel_id, bool reliable);
enum EDiscordResult discord_lobby_manager_send_network_message(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, uint8_t channel_id, uint8_t* data, uint32_t data_length);
enum EDiscordResult discord_lobby_manager_get_lobby_update_transaction(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, struct IDiscordLobbyTransaction** transaction);
enum EDiscordResult discord_lobby_transaction_set_type(struct IDiscordLobbyTransaction* transaction, enum EDiscordLobbyType type);
enum EDiscordResult discord_lobby_transaction_set_owner(struct IDiscordLobbyTransaction* transaction, DiscordUserId owner_id);
enum EDiscordResult discord_lobby_transaction_set_capacity(struct IDiscordLobbyTransaction* transaction, uint32_t capacity);
enum EDiscordResult discord_lobby_transaction_set_metadata(struct IDiscordLobbyTransaction* transaction, DiscordMetadataKey key, DiscordMetadataValue value);
enum EDiscordResult discord_lobby_transaction_delete_metadata(struct IDiscordLobbyTransaction* transaction, DiscordMetadataKey key);
enum EDiscordResult discord_lobby_transaction_set_locked(struct IDiscordLobbyTransaction* transaction, bool locked);

// Network manager wrappers
void discord_network_manager_get_peer_id(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId* peer_id);
enum EDiscordResult discord_network_manager_flush(struct IDiscordNetworkManager* manager);
enum EDiscordResult discord_network_manager_open_peer(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, const char* route_data);
enum EDiscordResult discord_network_manager_update_peer(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, const char* route_data);
enum EDiscordResult discord_network_manager_close_peer(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id);
enum EDiscordResult discord_network_manager_open_channel(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, DiscordNetworkChannelId channel_id, bool reliable);
enum EDiscordResult discord_network_manager_close_channel(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, DiscordNetworkChannelId channel_id);
enum EDiscordResult discord_network_manager_send_message(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, DiscordNetworkChannelId channel_id, uint8_t* data, uint32_t data_length);

// Storage manager wrappers
enum EDiscordResult discord_storage_manager_read(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length, uint32_t* read);
void discord_storage_manager_read_async(struct IDiscordStorageManager* manager, const char* name, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, uint8_t* data, uint32_t data_length));
enum EDiscordResult discord_storage_manager_write(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length);
void discord_storage_manager_write_async(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
enum EDiscordResult discord_storage_manager_delete_(struct IDiscordStorageManager* manager, const char* name);
enum EDiscordResult discord_storage_manager_exists(struct IDiscordStorageManager* manager, const char* name, bool* exists);
void discord_storage_manager_count(struct IDiscordStorageManager* manager, int32_t* count);

// Overlay manager wrappers
void discord_overlay_manager_is_enabled(struct IDiscordOverlayManager* manager, bool* enabled);
void discord_overlay_manager_is_locked(struct IDiscordOverlayManager* manager, bool* locked);
void discord_overlay_manager_set_locked(struct IDiscordOverlayManager* manager, bool locked, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_overlay_manager_open_activity_invite(struct IDiscordOverlayManager* manager, enum EDiscordActivityActionType type, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_overlay_manager_open_guild_invite(struct IDiscordOverlayManager* manager, const char* code, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_overlay_manager_open_voice_settings(struct IDiscordOverlayManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));

#endif // DISCORD_WRAPPERS_H 