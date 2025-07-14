#ifndef _DISCORD_GAME_SDK_SIZEOF_H_
#define _DISCORD_GAME_SDK_SIZEOF_H_

#include "discord_game_sdk.h"

// Helper functions to get sizeof values for array fields
static inline size_t sizeof_DiscordUser_username() { return sizeof(((struct DiscordUser*)0)->username); }
static inline size_t sizeof_DiscordUser_discriminator() { return sizeof(((struct DiscordUser*)0)->discriminator); }
static inline size_t sizeof_DiscordUser_avatar() { return sizeof(((struct DiscordUser*)0)->avatar); }
static inline size_t sizeof_DiscordActivity_name() { return sizeof(((struct DiscordActivity*)0)->name); }
static inline size_t sizeof_DiscordActivity_state() { return sizeof(((struct DiscordActivity*)0)->state); }
static inline size_t sizeof_DiscordActivity_details() { return sizeof(((struct DiscordActivity*)0)->details); }
static inline size_t sizeof_DiscordActivityAssets_large_image() { return sizeof(((struct DiscordActivityAssets*)0)->large_image); }
static inline size_t sizeof_DiscordActivityAssets_large_text() { return sizeof(((struct DiscordActivityAssets*)0)->large_text); }
static inline size_t sizeof_DiscordActivityAssets_small_image() { return sizeof(((struct DiscordActivityAssets*)0)->small_image); }
static inline size_t sizeof_DiscordActivityAssets_small_text() { return sizeof(((struct DiscordActivityAssets*)0)->small_text); }
static inline size_t sizeof_DiscordActivityParty_id() { return sizeof(((struct DiscordActivityParty*)0)->id); }
static inline size_t sizeof_DiscordActivitySecrets_match() { return sizeof(((struct DiscordActivitySecrets*)0)->match); }
static inline size_t sizeof_DiscordActivitySecrets_join() { return sizeof(((struct DiscordActivitySecrets*)0)->join); }
static inline size_t sizeof_DiscordActivitySecrets_spectate() { return sizeof(((struct DiscordActivitySecrets*)0)->spectate); }
static inline size_t sizeof_DiscordLobby_secret() { return sizeof(((struct DiscordLobby*)0)->secret); }
static inline size_t sizeof_DiscordInputMode_shortcut() { return sizeof(((struct DiscordInputMode*)0)->shortcut); }

#endif // _DISCORD_GAME_SDK_SIZEOF_H_ 