package discord

import "unsafe"

// CoreEvents contains all event callbacks for the Discord SDK
type CoreEvents struct {
	events              unsafe.Pointer
	eventData           unsafe.Pointer
	applicationEvents   unsafe.Pointer
	applicationVersion  int32
	userEvents          unsafe.Pointer
	userVersion         int32
	imageEvents         unsafe.Pointer
	imageVersion        int32
	activityEvents      unsafe.Pointer
	activityVersion     int32
	relationshipEvents  unsafe.Pointer
	relationshipVersion int32
	lobbyEvents         unsafe.Pointer
	lobbyVersion        int32
	networkEvents       unsafe.Pointer
	networkVersion      int32
	overlayEvents       unsafe.Pointer
	overlayVersion      int32
	storageEvents       unsafe.Pointer
	storageVersion      int32
	storeEvents         unsafe.Pointer
	storeVersion        int32
	voiceEvents         unsafe.Pointer
	voiceVersion        int32
	achievementEvents   unsafe.Pointer
	achievementVersion  int32
}

// ApplicationEvents defines callbacks for application-related events
type ApplicationEvents struct {
	OnValidateOrExit func(result Result)
	OnOAuth2Token    func(result Result, token *OAuth2Token)
	OnTicket         func(result Result, data string)
}

// UserEvents defines callbacks for user-related events
type UserEvents struct {
	OnCurrentUserUpdate func()
}

// ImageEvents defines callbacks for image-related events
type ImageEvents struct {
	OnFetch func(result Result, handle *ImageHandle)
}

// ActivityEvents defines callbacks for activity-related events
type ActivityEvents struct {
	OnActivityJoin        func(secret string)
	OnActivitySpectate    func(secret string)
	OnActivityJoinRequest func(user *User)
	OnActivityInvite      func(actionType ActivityActionType, user *User, activity *Activity)
}

// RelationshipEvents defines callbacks for relationship-related events
type RelationshipEvents struct {
	OnRefresh            func()
	OnRelationshipUpdate func(relationship *Relationship)
}

// LobbyEvents defines callbacks for lobby-related events
type LobbyEvents struct {
	OnLobbyUpdate      func(lobbyID int64)
	OnLobbyDelete      func(lobbyID int64, reason uint32)
	OnMemberConnect    func(lobbyID int64, userID int64)
	OnMemberUpdate     func(lobbyID int64, userID int64)
	OnMemberDisconnect func(lobbyID int64, userID int64)
	OnLobbyMessage     func(lobbyID int64, userID int64, data []byte)
	OnSpeaking         func(lobbyID int64, userID int64, speaking bool)
	OnNetworkMessage   func(lobbyID int64, userID int64, channelID uint8, data []byte)
}

// NetworkEvents defines callbacks for network-related events
type NetworkEvents struct {
	OnMessage     func(peerID uint64, channelID uint8, data []byte)
	OnRouteUpdate func(routeData string)
}

// OverlayEvents defines callbacks for overlay-related events
type OverlayEvents struct {
	OnToggle func(locked bool)
}

// StorageEvents defines callbacks for storage-related events
type StorageEvents struct {
	OnRead  func(result Result, data []byte)
	OnWrite func(result Result)
	OnCount func(count int32)
	OnStat  func(result Result, stat *FileStat)
}

// StoreEvents defines callbacks for store-related events
type StoreEvents struct {
	OnEntitlementCreate func(entitlement *Entitlement)
	OnEntitlementDelete func(entitlement *Entitlement)
}

// VoiceEvents defines callbacks for voice-related events
type VoiceEvents struct {
	OnSettingsUpdate func()
}

// AchievementEvents defines callbacks for achievement-related events
type AchievementEvents struct {
	OnUserAchievementUpdate func(userAchievement *UserAchievement)
}

// NewCoreEvents creates a new CoreEvents structure with default values
func NewCoreEvents() *CoreEvents {
	return &CoreEvents{
		applicationVersion:  1, // DISCORD_APPLICATION_MANAGER_VERSION
		userVersion:         1, // DISCORD_USER_MANAGER_VERSION
		imageVersion:        1, // DISCORD_IMAGE_MANAGER_VERSION
		activityVersion:     1, // DISCORD_ACTIVITY_MANAGER_VERSION
		relationshipVersion: 1, // DISCORD_RELATIONSHIP_MANAGER_VERSION
		lobbyVersion:        1, // DISCORD_LOBBY_MANAGER_VERSION
		networkVersion:      1, // DISCORD_NETWORK_MANAGER_VERSION
		overlayVersion:      2, // DISCORD_OVERLAY_MANAGER_VERSION
		storageVersion:      1, // DISCORD_STORAGE_MANAGER_VERSION
		storeVersion:        1, // DISCORD_STORE_MANAGER_VERSION
		voiceVersion:        1, // DISCORD_VOICE_MANAGER_VERSION
		achievementVersion:  1, // DISCORD_ACHIEVEMENT_MANAGER_VERSION
	}
}

// SetApplicationEvents sets the application events
func (e *CoreEvents) SetApplicationEvents(events *ApplicationEvents) {
	if events == nil {
		e.applicationEvents = nil
		return
	}

	// ApplicationEvents is a void* typedef, so we just pass the events pointer
	e.applicationEvents = unsafe.Pointer(events)
}

// SetUserEvents sets the user events
func (e *CoreEvents) SetUserEvents(events *UserEvents) {
	if events == nil {
		e.userEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.userEvents = unsafe.Pointer(events)
	e.eventData = unsafe.Pointer(events)
}

// SetImageEvents sets the image events
func (e *CoreEvents) SetImageEvents(events *ImageEvents) {
	if events == nil {
		e.imageEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.imageEvents = unsafe.Pointer(events)
}

// SetActivityEvents sets the activity events
func (e *CoreEvents) SetActivityEvents(events *ActivityEvents) {
	if events == nil {
		e.activityEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.activityEvents = unsafe.Pointer(events)
}

// SetRelationshipEvents sets the relationship events
func (e *CoreEvents) SetRelationshipEvents(events *RelationshipEvents) {
	if events == nil {
		e.relationshipEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.relationshipEvents = unsafe.Pointer(events)
}

// SetLobbyEvents sets the lobby events
func (e *CoreEvents) SetLobbyEvents(events *LobbyEvents) {
	if events == nil {
		e.lobbyEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.lobbyEvents = unsafe.Pointer(events)
}

// SetNetworkEvents sets the network events
func (e *CoreEvents) SetNetworkEvents(events *NetworkEvents) {
	if events == nil {
		e.networkEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.networkEvents = unsafe.Pointer(events)
}

// SetOverlayEvents sets the overlay events
func (e *CoreEvents) SetOverlayEvents(events *OverlayEvents) {
	if events == nil {
		e.overlayEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.overlayEvents = unsafe.Pointer(events)
}

// SetStorageEvents sets the storage events
func (e *CoreEvents) SetStorageEvents(events *StorageEvents) {
	if events == nil {
		e.storageEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.storageEvents = unsafe.Pointer(events)
}

// SetStoreEvents sets the store events
func (e *CoreEvents) SetStoreEvents(events *StoreEvents) {
	if events == nil {
		e.storeEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.storeEvents = unsafe.Pointer(events)
}

// SetVoiceEvents sets the voice events
func (e *CoreEvents) SetVoiceEvents(events *VoiceEvents) {
	if events == nil {
		e.voiceEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.voiceEvents = unsafe.Pointer(events)
}

// SetAchievementEvents sets the achievement events
func (e *CoreEvents) SetAchievementEvents(events *AchievementEvents) {
	if events == nil {
		e.achievementEvents = nil
		return
	}

	// NOTE: C structure creation not supported in wrapper approach
	// Just store the events pointer for now
	e.achievementEvents = unsafe.Pointer(events)
}
