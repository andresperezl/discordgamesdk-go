#include "discord_game_sdk.h"

// Core wrapper functions
enum EDiscordResult discord_core_create(DiscordVersion version, struct DiscordCreateParams* params, struct IDiscordCore** result) {
    return DiscordCreate(version, params, result);
}

void discord_core_destroy(void* core) {
    ((struct IDiscordCore*)core)->destroy((struct IDiscordCore*)core);
}

enum EDiscordResult discord_core_run_callbacks(void* core) {
    return ((struct IDiscordCore*)core)->run_callbacks((struct IDiscordCore*)core);
}

void discord_core_set_log_hook(void* core, enum EDiscordLogLevel min_level, void* hook_data, void (*hook)(void* hook_data, enum EDiscordLogLevel level, const char* message)) {
    ((struct IDiscordCore*)core)->set_log_hook((struct IDiscordCore*)core, min_level, hook_data, hook);
}

struct IDiscordApplicationManager* discord_core_get_application_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_application_manager((struct IDiscordCore*)core);
}

struct IDiscordUserManager* discord_core_get_user_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_user_manager((struct IDiscordCore*)core);
}

struct IDiscordActivityManager* discord_core_get_activity_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_activity_manager((struct IDiscordCore*)core);
}

struct IDiscordLobbyManager* discord_core_get_lobby_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_lobby_manager((struct IDiscordCore*)core);
}

struct IDiscordNetworkManager* discord_core_get_network_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_network_manager((struct IDiscordCore*)core);
}

struct IDiscordOverlayManager* discord_core_get_overlay_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_overlay_manager((struct IDiscordCore*)core);
}

struct IDiscordStorageManager* discord_core_get_storage_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_storage_manager((struct IDiscordCore*)core);
}

struct IDiscordStoreManager* discord_core_get_store_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_store_manager((struct IDiscordCore*)core);
}

struct IDiscordVoiceManager* discord_core_get_voice_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_voice_manager((struct IDiscordCore*)core);
}

struct IDiscordAchievementManager* discord_core_get_achievement_manager(void* core) {
    return ((struct IDiscordCore*)core)->get_achievement_manager((struct IDiscordCore*)core);
}





// Application manager wrappers
void discord_application_manager_validate_or_exit(struct IDiscordApplicationManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->validate_or_exit(manager, callback_data, callback);
}

void discord_application_manager_get_current_locale(struct IDiscordApplicationManager* manager, DiscordLocale* locale) {
    manager->get_current_locale(manager, locale);
}

void discord_application_manager_get_current_branch(struct IDiscordApplicationManager* manager, DiscordBranch* branch) {
    manager->get_current_branch(manager, branch);
}

void discord_application_manager_get_oauth2_token(struct IDiscordApplicationManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordOAuth2Token* oauth2_token)) {
    manager->get_oauth2_token(manager, callback_data, callback);
}

void discord_application_manager_get_ticket(struct IDiscordApplicationManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, const char* data)) {
    manager->get_ticket(manager, callback_data, callback);
}

// User manager wrappers
enum EDiscordResult discord_user_manager_get_current_user(struct IDiscordUserManager* manager, struct DiscordUser* current_user) {
    return manager->get_current_user(manager, current_user);
}

void discord_user_manager_get_user(struct IDiscordUserManager* manager, DiscordUserId user_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordUser* user)) {
    manager->get_user(manager, user_id, callback_data, callback);
}

enum EDiscordResult discord_user_manager_get_current_user_premium_type(struct IDiscordUserManager* manager, enum EDiscordPremiumType* premium_type) {
    return manager->get_current_user_premium_type(manager, premium_type);
}

enum EDiscordResult discord_user_manager_current_user_has_flag(struct IDiscordUserManager* manager, enum EDiscordUserFlag flag, bool* has_flag) {
    return manager->current_user_has_flag(manager, flag, has_flag);
}

// Activity manager wrappers
enum EDiscordResult discord_activity_manager_register_command(struct IDiscordActivityManager* manager, const char* command) {
    return manager->register_command(manager, command);
}

enum EDiscordResult discord_activity_manager_register_steam(struct IDiscordActivityManager* manager, uint32_t steam_id) {
    return manager->register_steam(manager, steam_id);
}

void discord_activity_manager_update_activity(struct IDiscordActivityManager* manager, struct DiscordActivity* activity, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->update_activity(manager, activity, callback_data, callback);
}

void discord_activity_manager_clear_activity(struct IDiscordActivityManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->clear_activity(manager, callback_data, callback);
}

void discord_activity_manager_send_request_reply(struct IDiscordActivityManager* manager, DiscordUserId user_id, enum EDiscordActivityJoinRequestReply reply, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->send_request_reply(manager, user_id, reply, callback_data, callback);
}

void discord_activity_manager_send_invite(struct IDiscordActivityManager* manager, DiscordUserId user_id, enum EDiscordActivityActionType type, const char* content, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->send_invite(manager, user_id, type, content, callback_data, callback);
}

void discord_activity_manager_accept_invite(struct IDiscordActivityManager* manager, DiscordUserId user_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->accept_invite(manager, user_id, callback_data, callback);
}

// Lobby manager wrappers
enum EDiscordResult discord_lobby_manager_get_lobby_create_transaction(struct IDiscordLobbyManager* manager, struct IDiscordLobbyTransaction** transaction) {
    return manager->get_lobby_create_transaction(manager, transaction);
}

void discord_lobby_manager_create_lobby(struct IDiscordLobbyManager* manager, struct IDiscordLobbyTransaction* transaction, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordLobby* lobby)) {
    manager->create_lobby(manager, transaction, callback_data, callback);
}

void discord_lobby_manager_connect_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordLobbySecret secret, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, struct DiscordLobby* lobby)) {
    manager->connect_lobby(manager, lobby_id, secret, callback_data, callback);
}

enum EDiscordResult discord_lobby_manager_get_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, struct DiscordLobby* lobby) {
    return manager->get_lobby(manager, lobby_id, lobby);
}

enum EDiscordResult discord_lobby_manager_get_lobby_activity_secret(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordLobbySecret* secret) {
    return manager->get_lobby_activity_secret(manager, lobby_id, secret);
}

// Network manager wrappers
void discord_network_manager_get_peer_id(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId* peer_id) {
    manager->get_peer_id(manager, peer_id);
}

enum EDiscordResult discord_network_manager_flush(struct IDiscordNetworkManager* manager) {
    return manager->flush(manager);
}

enum EDiscordResult discord_network_manager_open_peer(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, const char* route_data) {
    return manager->open_peer(manager, peer_id, route_data);
}

enum EDiscordResult discord_network_manager_update_peer(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, const char* route_data) {
    return manager->update_peer(manager, peer_id, route_data);
}

enum EDiscordResult discord_network_manager_close_peer(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id) {
    return manager->close_peer(manager, peer_id);
}

enum EDiscordResult discord_network_manager_open_channel(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, DiscordNetworkChannelId channel_id, bool reliable) {
    return manager->open_channel(manager, peer_id, channel_id, reliable);
}

enum EDiscordResult discord_network_manager_close_channel(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, DiscordNetworkChannelId channel_id) {
    return manager->close_channel(manager, peer_id, channel_id);
}

enum EDiscordResult discord_network_manager_send_message(struct IDiscordNetworkManager* manager, DiscordNetworkPeerId peer_id, DiscordNetworkChannelId channel_id, uint8_t* data, uint32_t data_length) {
    return manager->send_message(manager, peer_id, channel_id, data, data_length);
}

// Storage manager wrappers
enum EDiscordResult discord_storage_manager_read(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length, uint32_t* read) {
    return manager->read(manager, name, data, data_length, read);
}

void discord_storage_manager_read_async(struct IDiscordStorageManager* manager, const char* name, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, uint8_t* data, uint32_t data_length)) {
    manager->read_async(manager, name, callback_data, callback);
}

enum EDiscordResult discord_storage_manager_write(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length) {
    return manager->write(manager, name, data, data_length);
}

void discord_storage_manager_write_async(struct IDiscordStorageManager* manager, const char* name, uint8_t* data, uint32_t data_length, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->write_async(manager, name, data, data_length, callback_data, callback);
}

enum EDiscordResult discord_storage_manager_delete_(struct IDiscordStorageManager* manager, const char* name) {
    return manager->delete_(manager, name);
}

enum EDiscordResult discord_storage_manager_exists(struct IDiscordStorageManager* manager, const char* name, bool* exists) {
    return manager->exists(manager, name, exists);
}

void discord_storage_manager_count(struct IDiscordStorageManager* manager, int32_t* count) {
    manager->count(manager, count);
}

// Overlay manager wrappers
void discord_overlay_manager_is_enabled(struct IDiscordOverlayManager* manager, bool* enabled) {
    manager->is_enabled(manager, enabled);
}

void discord_overlay_manager_is_locked(struct IDiscordOverlayManager* manager, bool* locked) {
    manager->is_locked(manager, locked);
}

void discord_overlay_manager_set_locked(struct IDiscordOverlayManager* manager, bool locked, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->set_locked(manager, locked, callback_data, callback);
}

void discord_overlay_manager_open_activity_invite(struct IDiscordOverlayManager* manager, enum EDiscordActivityActionType type, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->open_activity_invite(manager, type, callback_data, callback);
}

void discord_overlay_manager_open_guild_invite(struct IDiscordOverlayManager* manager, const char* code, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->open_guild_invite(manager, code, callback_data, callback);
}

void discord_overlay_manager_open_voice_settings(struct IDiscordOverlayManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->open_voice_settings(manager, callback_data, callback);
}

// Additional lobby manager wrappers
void discord_lobby_manager_update_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, struct IDiscordLobbyTransaction* transaction, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->update_lobby(manager, lobby_id, transaction, callback_data, callback);
}

void discord_lobby_manager_delete_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->delete_lobby(manager, lobby_id, callback_data, callback);
}

void discord_lobby_manager_disconnect_lobby(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->disconnect_lobby(manager, lobby_id, callback_data, callback);
}

void discord_lobby_manager_send_lobby_message(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, uint8_t* data, uint32_t data_length, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->send_lobby_message(manager, lobby_id, data, data_length, callback_data, callback);
}

void discord_lobby_manager_connect_voice(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->connect_voice(manager, lobby_id, callback_data, callback);
}

void discord_lobby_manager_disconnect_voice(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->disconnect_voice(manager, lobby_id, callback_data, callback);
}

enum EDiscordResult discord_lobby_manager_connect_network(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id) {
    return manager->connect_network(manager, lobby_id);
}

enum EDiscordResult discord_lobby_manager_disconnect_network(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id) {
    return manager->disconnect_network(manager, lobby_id);
}

enum EDiscordResult discord_lobby_manager_flush_network(struct IDiscordLobbyManager* manager) {
    return manager->flush_network(manager);
}

enum EDiscordResult discord_lobby_manager_open_network_channel(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, uint8_t channel_id, bool reliable) {
    return manager->open_network_channel(manager, lobby_id, channel_id, reliable);
}

enum EDiscordResult discord_lobby_manager_send_network_message(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, DiscordUserId user_id, uint8_t channel_id, uint8_t* data, uint32_t data_length) {
    return manager->send_network_message(manager, lobby_id, user_id, channel_id, data, data_length);
}

enum EDiscordResult discord_lobby_manager_get_lobby_update_transaction(struct IDiscordLobbyManager* manager, DiscordLobbyId lobby_id, struct IDiscordLobbyTransaction** transaction) {
    return manager->get_lobby_update_transaction(manager, lobby_id, transaction);
}

enum EDiscordResult discord_lobby_transaction_set_type(struct IDiscordLobbyTransaction* transaction, enum EDiscordLobbyType type) {
    return transaction->set_type(transaction, type);
}

enum EDiscordResult discord_lobby_transaction_set_owner(struct IDiscordLobbyTransaction* transaction, DiscordUserId owner_id) {
    return transaction->set_owner(transaction, owner_id);
}

enum EDiscordResult discord_lobby_transaction_set_capacity(struct IDiscordLobbyTransaction* transaction, uint32_t capacity) {
    return transaction->set_capacity(transaction, capacity);
}

enum EDiscordResult discord_lobby_transaction_set_metadata(struct IDiscordLobbyTransaction* transaction, DiscordMetadataKey key, DiscordMetadataValue value) {
    return transaction->set_metadata(transaction, key, value);
}

enum EDiscordResult discord_lobby_transaction_delete_metadata(struct IDiscordLobbyTransaction* transaction, DiscordMetadataKey key) {
    return transaction->delete_metadata(transaction, key);
}

enum EDiscordResult discord_lobby_transaction_set_locked(struct IDiscordLobbyTransaction* transaction, bool locked) {
    return transaction->set_locked(transaction, locked);
}

// Additional storage manager wrappers
void discord_storage_manager_read_async_partial(struct IDiscordStorageManager* manager, const char* name, uint64_t offset, uint64_t length, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result, uint8_t* data, uint32_t data_length)) {
    manager->read_async_partial(manager, name, offset, length, callback_data, callback);
}

enum EDiscordResult discord_storage_manager_stat(struct IDiscordStorageManager* manager, const char* name, struct DiscordFileStat* stat) {
    return manager->stat(manager, name, stat);
}

enum EDiscordResult discord_storage_manager_stat_at(struct IDiscordStorageManager* manager, int32_t index, struct DiscordFileStat* stat) {
    return manager->stat_at(manager, index, stat);
}

enum EDiscordResult discord_storage_manager_get_path(struct IDiscordStorageManager* manager, DiscordPath* path) {
    return manager->get_path(manager, path);
}

// Additional overlay manager wrappers
enum EDiscordResult discord_overlay_manager_init_drawing_dxgi(struct IDiscordOverlayManager* manager, void* swapchain, bool use_message_forwarding) {
    return manager->init_drawing_dxgi(manager, swapchain, use_message_forwarding);
}

void discord_overlay_manager_on_present(struct IDiscordOverlayManager* manager) {
    manager->on_present(manager);
}

void discord_overlay_manager_forward_message(struct IDiscordOverlayManager* manager, void* message) {
    manager->forward_message(manager, message);
}

void discord_overlay_manager_key_event(struct IDiscordOverlayManager* manager, bool down, const char* key_code, enum EDiscordKeyVariant variant) {
    manager->key_event(manager, down, key_code, variant);
}

void discord_overlay_manager_char_event(struct IDiscordOverlayManager* manager, const char* character) {
    manager->char_event(manager, character);
}

void discord_overlay_manager_mouse_button_event(struct IDiscordOverlayManager* manager, uint8_t down, int32_t click_count, enum EDiscordMouseButton which, int32_t x, int32_t y) {
    manager->mouse_button_event(manager, down, click_count, which, x, y);
}

void discord_overlay_manager_mouse_motion_event(struct IDiscordOverlayManager* manager, int32_t x, int32_t y) {
    manager->mouse_motion_event(manager, x, y);
}

void discord_overlay_manager_ime_commit_text(struct IDiscordOverlayManager* manager, const char* text) {
    manager->ime_commit_text(manager, text);
}

void discord_overlay_manager_ime_set_composition(struct IDiscordOverlayManager* manager, const char* text, struct DiscordImeUnderline* underlines, uint32_t underlines_length, int32_t from, int32_t to) {
    manager->ime_set_composition(manager, text, underlines, underlines_length, from, to);
}

void discord_overlay_manager_ime_cancel_composition(struct IDiscordOverlayManager* manager) {
    manager->ime_cancel_composition(manager);
}

void discord_overlay_manager_set_ime_composition_range_callback(struct IDiscordOverlayManager* manager, void* on_ime_composition_range_changed_data, void (*on_ime_composition_range_changed)(void* on_ime_composition_range_changed_data, int32_t from, int32_t to, struct DiscordRect* bounds, uint32_t bounds_length)) {
    manager->set_ime_composition_range_callback(manager, on_ime_composition_range_changed_data, on_ime_composition_range_changed);
}

void discord_overlay_manager_set_ime_selection_bounds_callback(struct IDiscordOverlayManager* manager, void* on_ime_selection_bounds_changed_data, void (*on_ime_selection_bounds_changed)(void* on_ime_selection_bounds_changed_data, struct DiscordRect anchor, struct DiscordRect focus, bool is_anchor_first)) {
    manager->set_ime_selection_bounds_callback(manager, on_ime_selection_bounds_changed_data, on_ime_selection_bounds_changed);
}

bool discord_overlay_manager_is_point_inside_click_zone(struct IDiscordOverlayManager* manager, int32_t x, int32_t y) {
    return manager->is_point_inside_click_zone(manager, x, y);
}

// Store manager wrappers
void discord_store_manager_fetch_skus(struct IDiscordStoreManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->fetch_skus(manager, callback_data, callback);
}

void discord_store_manager_count_skus(struct IDiscordStoreManager* manager, int32_t* count) {
    manager->count_skus(manager, count);
}

enum EDiscordResult discord_store_manager_get_sku(struct IDiscordStoreManager* manager, DiscordSnowflake sku_id, struct DiscordSku* sku) {
    return manager->get_sku(manager, sku_id, sku);
}

enum EDiscordResult discord_store_manager_get_sku_at(struct IDiscordStoreManager* manager, int32_t index, struct DiscordSku* sku) {
    return manager->get_sku_at(manager, index, sku);
}

void discord_store_manager_fetch_entitlements(struct IDiscordStoreManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->fetch_entitlements(manager, callback_data, callback);
}

void discord_store_manager_count_entitlements(struct IDiscordStoreManager* manager, int32_t* count) {
    manager->count_entitlements(manager, count);
}

enum EDiscordResult discord_store_manager_get_entitlement(struct IDiscordStoreManager* manager, DiscordSnowflake entitlement_id, struct DiscordEntitlement* entitlement) {
    return manager->get_entitlement(manager, entitlement_id, entitlement);
}

enum EDiscordResult discord_store_manager_get_entitlement_at(struct IDiscordStoreManager* manager, int32_t index, struct DiscordEntitlement* entitlement) {
    return manager->get_entitlement_at(manager, index, entitlement);
}

enum EDiscordResult discord_store_manager_has_sku_entitlement(struct IDiscordStoreManager* manager, DiscordSnowflake sku_id, bool* has_entitlement) {
    return manager->has_sku_entitlement(manager, sku_id, has_entitlement);
}

void discord_store_manager_start_purchase(struct IDiscordStoreManager* manager, DiscordSnowflake sku_id, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->start_purchase(manager, sku_id, callback_data, callback);
}

// Voice manager wrappers
enum EDiscordResult discord_voice_manager_get_input_mode(struct IDiscordVoiceManager* manager, struct DiscordInputMode* input_mode) {
    return manager->get_input_mode(manager, input_mode);
}

void discord_voice_manager_set_input_mode(struct IDiscordVoiceManager* manager, struct DiscordInputMode input_mode, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->set_input_mode(manager, input_mode, callback_data, callback);
}

enum EDiscordResult discord_voice_manager_is_self_mute(struct IDiscordVoiceManager* manager, bool* mute) {
    return manager->is_self_mute(manager, mute);
}

enum EDiscordResult discord_voice_manager_set_self_mute(struct IDiscordVoiceManager* manager, bool mute) {
    return manager->set_self_mute(manager, mute);
}

enum EDiscordResult discord_voice_manager_is_self_deaf(struct IDiscordVoiceManager* manager, bool* deaf) {
    return manager->is_self_deaf(manager, deaf);
}

enum EDiscordResult discord_voice_manager_set_self_deaf(struct IDiscordVoiceManager* manager, bool deaf) {
    return manager->set_self_deaf(manager, deaf);
}

enum EDiscordResult discord_voice_manager_is_local_mute(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, bool* mute) {
    return manager->is_local_mute(manager, user_id, mute);
}

enum EDiscordResult discord_voice_manager_set_local_mute(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, bool mute) {
    return manager->set_local_mute(manager, user_id, mute);
}

enum EDiscordResult discord_voice_manager_get_local_volume(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, uint8_t* volume) {
    return manager->get_local_volume(manager, user_id, volume);
}

enum EDiscordResult discord_voice_manager_set_local_volume(struct IDiscordVoiceManager* manager, DiscordSnowflake user_id, uint8_t volume) {
    return manager->set_local_volume(manager, user_id, volume);
}

// Achievement manager wrappers
void discord_achievement_manager_set_user_achievement(struct IDiscordAchievementManager* manager, DiscordSnowflake achievement_id, uint8_t percent_complete, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->set_user_achievement(manager, achievement_id, percent_complete, callback_data, callback);
}

void discord_achievement_manager_fetch_user_achievements(struct IDiscordAchievementManager* manager, void* callback_data, void (*callback)(void* callback_data, enum EDiscordResult result)) {
    manager->fetch_user_achievements(manager, callback_data, callback);
}

void discord_achievement_manager_count_user_achievements(struct IDiscordAchievementManager* manager, int32_t* count) {
    manager->count_user_achievements(manager, count);
}

enum EDiscordResult discord_achievement_manager_get_user_achievement(struct IDiscordAchievementManager* manager, DiscordSnowflake user_achievement_id, struct DiscordUserAchievement* user_achievement) {
    return manager->get_user_achievement(manager, user_achievement_id, user_achievement);
}

enum EDiscordResult discord_achievement_manager_get_user_achievement_at(struct IDiscordAchievementManager* manager, int32_t index, struct DiscordUserAchievement* user_achievement) {
    return manager->get_user_achievement_at(manager, index, user_achievement);
} 