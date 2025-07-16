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
struct IDiscordImageManager* discord_core_get_image_manager(void* core);
struct IDiscordRelationshipManager* discord_core_get_relationship_manager(void* core);

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
// Callback trampoline for Go
typedef void (*go_storage_read_async_callback_t)(void* go_callback_data, enum EDiscordResult result, uint8_t* data, uint32_t data_length);

// Update the read_async wrapper to accept only go_callback_data
void discord_storage_manager_read_async_trampoline(struct IDiscordStorageManager* manager, const char* name, void* go_callback_data);
enum EDiscordResult discord_storage_manager_write(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length);
// Callback trampoline for Go write async
typedef void (*go_storage_write_async_callback_t)(void* go_callback_data, enum EDiscordResult result);

// Update the write_async wrapper to accept only go_callback_data
void discord_storage_manager_write_async_trampoline(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length, void* go_callback_data);
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

// Store manager wrappers
void discord_store_manager_fetch_skus(struct IDiscordStoreManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_store_manager_count_skus(struct IDiscordStoreManager* manager, int32_t* count);
enum EDiscordResult discord_store_manager_get_sku(struct IDiscordStoreManager* manager, DiscordSnowflake sku_id, struct DiscordSku* sku);
enum EDiscordResult discord_store_manager_get_sku_at(struct IDiscordStoreManager* manager, int32_t index, struct DiscordSku* sku);
void discord_store_manager_fetch_entitlements(struct IDiscordStoreManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_store_manager_count_entitlements(struct IDiscordStoreManager* manager, int32_t* count);
enum EDiscordResult discord_store_manager_get_entitlement(struct IDiscordStoreManager* manager, DiscordSnowflake entitlement_id, struct DiscordEntitlement* entitlement);
enum EDiscordResult discord_store_manager_get_entitlement_at(struct IDiscordStoreManager* manager, int32_t index, struct DiscordEntitlement* entitlement);
enum EDiscordResult discord_store_manager_has_sku_entitlement(struct IDiscordStoreManager* manager, DiscordSnowflake sku_id, bool* has_entitlement);
void discord_store_manager_start_purchase(struct IDiscordStoreManager* manager, DiscordSnowflake sku_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));

// Voice manager wrappers
enum EDiscordResult discord_voice_manager_get_input_mode(struct IDiscordVoiceManager* manager, struct DiscordInputMode* input_mode);
void discord_voice_manager_set_input_mode(struct IDiscordVoiceManager* manager, struct DiscordInputMode input_mode, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
enum EDiscordResult discord_voice_manager_is_self_mute(struct IDiscordVoiceManager* manager, bool* mute);
enum EDiscordResult discord_voice_manager_set_self_mute(struct IDiscordVoiceManager* manager, bool mute);
enum EDiscordResult discord_voice_manager_is_self_deaf(struct IDiscordVoiceManager* manager, bool* deaf);
enum EDiscordResult discord_voice_manager_set_self_deaf(struct IDiscordVoiceManager* manager, bool deaf);
enum EDiscordResult discord_voice_manager_is_local_mute(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, bool* mute);
enum EDiscordResult discord_voice_manager_set_local_mute(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, bool mute);
enum EDiscordResult discord_voice_manager_get_local_volume(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, uint8_t* volume);
enum EDiscordResult discord_voice_manager_set_local_volume(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, uint8_t volume);

// Achievement manager wrappers
void discord_achievement_manager_set_user_achievement(struct IDiscordAchievementManager* manager, DiscordSnowflake achievement_id, uint8_t percent_complete, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_achievement_manager_fetch_user_achievements(struct IDiscordAchievementManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_achievement_manager_count_user_achievements(struct IDiscordAchievementManager* manager, int32_t* count);
enum EDiscordResult discord_achievement_manager_get_user_achievement(struct IDiscordAchievementManager* manager, DiscordSnowflake user_achievement_id, struct DiscordUserAchievement* user_achievement);
enum EDiscordResult discord_achievement_manager_get_user_achievement_at(struct IDiscordAchievementManager* manager, int32_t index, struct DiscordUserAchievement* user_achievement);

// Additional overlay manager wrappers
enum EDiscordResult discord_overlay_manager_init_drawing_dxgi(struct IDiscordOverlayManager* manager, void* swapchain, bool use_message_forwarding);
void discord_overlay_manager_on_present(struct IDiscordOverlayManager* manager);
void discord_overlay_manager_forward_message(struct IDiscordOverlayManager* manager, void* message);
void discord_overlay_manager_key_event(struct IDiscordOverlayManager* manager, bool down, const char* key_code, enum EDiscordKeyVariant variant);
void discord_overlay_manager_char_event(struct IDiscordOverlayManager* manager, const char* character);
void discord_overlay_manager_mouse_button_event(struct IDiscordOverlayManager* manager, uint8_t down, int32_t click_count, enum EDiscordMouseButton which, int32_t x, int32_t y);
void discord_overlay_manager_mouse_motion_event(struct IDiscordOverlayManager* manager, int32_t x, int32_t y);
void discord_overlay_manager_ime_commit_text(struct IDiscordOverlayManager* manager, const char* text);
void discord_overlay_manager_ime_set_composition(struct IDiscordOverlayManager* manager, const char* text, struct DiscordImeUnderline* underlines, uint32_t underlines_length, int32_t from, int32_t to);
void discord_overlay_manager_ime_cancel_composition(struct IDiscordOverlayManager* manager);
void discord_overlay_manager_set_ime_composition_range_callback(struct IDiscordOverlayManager* manager, void* on_ime_composition_range_changed_data, void (*on_ime_composition_range_changed)(void* on_ime_composition_range_changed_data, int32_t from, int32_t to, struct DiscordRect* bounds, uint32_t bounds_length));
void discord_overlay_manager_set_ime_selection_bounds_callback(struct IDiscordOverlayManager* manager, void* on_ime_selection_bounds_changed_data, void (*on_ime_selection_bounds_changed)(void* on_ime_selection_bounds_changed_data, struct DiscordRect anchor, struct DiscordRect focus, bool is_anchor_first));
bool discord_overlay_manager_is_point_inside_click_zone(struct IDiscordOverlayManager* manager, int32_t x, int32_t y);

// Additional missing lobby manager wrappers
void discord_lobby_manager_connect_lobby_with_activity_secret(struct IDiscordLobbyManager* manager, DiscordLobbySecret activity_secret, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordLobby* lobby));
enum EDiscordResult discord_lobby_manager_get_member_update_transaction(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, struct IDiscordLobbyMemberTransaction** transaction);
enum EDiscordResult discord_lobby_manager_get_lobby_metadata_value(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, const char* key, char* value);
enum EDiscordResult discord_lobby_manager_get_lobby_metadata_key(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, int32_t index, char* key);
enum EDiscordResult discord_lobby_manager_lobby_metadata_count(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, int32_t* count);
enum EDiscordResult discord_lobby_manager_member_count(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, int32_t* count);
enum EDiscordResult discord_lobby_manager_get_member_user_id(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, int32_t index, DiscordUserId* user_id);
enum EDiscordResult discord_lobby_manager_get_member_user(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, struct DiscordUser* user);
enum EDiscordResult discord_lobby_manager_get_member_metadata_value(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, const char* key, char* value);
enum EDiscordResult discord_lobby_manager_get_member_metadata_key(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, int32_t index, char* key);
enum EDiscordResult discord_lobby_manager_member_metadata_count(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, int32_t* count);
void discord_lobby_manager_update_member(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, struct IDiscordLobbyMemberTransaction* transaction, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
enum EDiscordResult discord_lobby_manager_get_search_query(struct IDiscordLobbyManager* manager, struct IDiscordLobbySearchQuery** query);
void discord_lobby_manager_search(struct IDiscordLobbyManager* manager, struct IDiscordLobbySearchQuery* query, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
void discord_lobby_manager_lobby_count(struct IDiscordLobbyManager* manager, int32_t* count);
enum EDiscordResult discord_lobby_manager_get_lobby_id(struct IDiscordLobbyManager* manager, int32_t index, DiscordLobbyId* lobby_id);

// Additional storage manager wrappers
void discord_storage_manager_read_async_partial(struct IDiscordStorageManager* manager, const char* name, uint64_t offset, uint64_t length, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, uint8_t* data, uint32_t data_length));
void discord_storage_manager_write_async(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result));
enum EDiscordResult discord_storage_manager_stat(struct IDiscordStorageManager* manager, const char* name, struct DiscordFileStat* stat);
enum EDiscordResult discord_storage_manager_stat_at(struct IDiscordStorageManager* manager, int32_t index, struct DiscordFileStat* stat);
enum EDiscordResult discord_storage_manager_get_path(struct IDiscordStorageManager* manager, DiscordPath* path);

// Store manager wrappers 
enum EDiscordResult discord_lobby_manager_get_lobby_activity_secret(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordLobbySecret* secret);
// Image manager wrappers
void discord_image_manager_fetch(struct IDiscordImageManager* manager, struct DiscordImageHandle handle, bool refresh, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordImageHandle handle_result));
enum EDiscordResult discord_image_manager_get_dimensions(struct IDiscordImageManager* manager, struct DiscordImageHandle handle, struct DiscordImageDimensions* dimensions);
enum EDiscordResult discord_image_manager_get_data(struct IDiscordImageManager* manager, struct DiscordImageHandle handle, uint8_t* data, uint32_t data_length);
// Relationship manager wrappers
void discord_relationship_manager_filter(struct IDiscordRelationshipManager* manager, void* filter_data, bool (*filter)(void* filter_data, struct DiscordRelationship* relationship));
enum EDiscordResult discord_relationship_manager_count(struct IDiscordRelationshipManager* manager, int32_t* count);
enum EDiscordResult discord_relationship_manager_get(struct IDiscordRelationshipManager* manager, DiscordUserId user_id, struct DiscordRelationship* relationship);
enum EDiscordResult discord_relationship_manager_get_at(struct IDiscordRelationshipManager* manager, uint32_t index, struct DiscordRelationship* relationship);
// Field accessors for DiscordSku
int64_t get_discord_sku_id(struct DiscordSku* sku);
int32_t get_discord_sku_type(struct DiscordSku* sku);
void get_discord_sku_name(struct DiscordSku* sku, char* buf, int bufsize);
uint32_t get_discord_sku_price_amount(struct DiscordSku* sku);
void get_discord_sku_price_currency(struct DiscordSku* sku, char* buf, int bufsize);
// Field accessors for DiscordEntitlement
int64_t get_discord_entitlement_id(struct DiscordEntitlement* ent);
int32_t get_discord_entitlement_type(struct DiscordEntitlement* ent);
int64_t get_discord_entitlement_sku_id(struct DiscordEntitlement* ent);
// Field accessors for DiscordFileStat
void get_discord_file_stat_filename(struct DiscordFileStat* stat, char* buf, int bufsize);
uint64_t get_discord_file_stat_size(struct DiscordFileStat* stat);
uint64_t get_discord_file_stat_last_modified(struct DiscordFileStat* stat);
#endif 