package discord

// User represents a Discord user
type User struct {
	ID            int64
	Username      string
	Discriminator string
	Avatar        string
	Bot           bool
}

// OAuth2Token represents an OAuth2 token
type OAuth2Token struct {
	AccessToken string
	Scopes      string
	Expires     int64
}

// ImageHandle represents an image handle
type ImageHandle struct {
	Type ImageType
	ID   int64
	Size uint32
}

// ImageDimensions represents image dimensions
type ImageDimensions struct {
	Width  uint32
	Height uint32
}

// ActivityTimestamps represents activity timestamps
type ActivityTimestamps struct {
	Start int64
	End   int64
}

// ActivityAssets represents activity assets
type ActivityAssets struct {
	LargeImage string
	LargeText  string
	SmallImage string
	SmallText  string
}

// PartySize represents party size
type PartySize struct {
	CurrentSize int32
	MaxSize     int32
}

// ActivityParty represents an activity party
type ActivityParty struct {
	ID      string
	Size    PartySize
	Privacy ActivityPartyPrivacy
}

// ActivitySecrets represents activity secrets
type ActivitySecrets struct {
	Match    string
	Join     string
	Spectate string
}

// Activity represents a Discord activity
type Activity struct {
	Type               ActivityType
	ApplicationID      int64
	Name               string
	State              string
	Details            string
	Timestamps         ActivityTimestamps
	Assets             ActivityAssets
	Party              ActivityParty
	Secrets            ActivitySecrets
	Instance           bool
	SupportedPlatforms uint32
}

// Presence represents a user's presence
type Presence struct {
	Status   Status
	Activity Activity
}

// Relationship represents a relationship between users
type Relationship struct {
	Type     RelationshipType
	User     User
	Presence Presence
}

// Lobby represents a Discord lobby
type Lobby struct {
	ID       int64
	Type     LobbyType
	OwnerID  int64
	Secret   string
	Capacity uint32
	Locked   bool
}

// FileStat represents file statistics
type FileStat struct {
	Filename     string
	Size         uint64
	LastModified uint64
}

// Entitlement represents an entitlement
type Entitlement struct {
	ID    int64
	Type  EntitlementType
	SkuID int64
}

// SkuPrice represents a SKU price
type SkuPrice struct {
	Amount   uint32
	Currency string
}

// Sku represents a SKU
type Sku struct {
	ID    int64
	Type  SkuType
	Name  string
	Price SkuPrice
}

// InputMode represents input mode settings
type InputMode struct {
	Type     InputModeType
	Shortcut string
}

// UserAchievement represents a user achievement
type UserAchievement struct {
	UserID          int64
	AchievementID   int64
	PercentComplete uint8
	UnlockedAt      string
}

// Enums
type ImageType int32

const (
	ImageTypeUser ImageType = 0
)

type ActivityPartyPrivacy int32

const (
	ActivityPartyPrivacyPrivate ActivityPartyPrivacy = 0
	ActivityPartyPrivacyPublic  ActivityPartyPrivacy = 1
)

type ActivityType int32

const (
	ActivityTypePlaying   ActivityType = 0
	ActivityTypeStreaming ActivityType = 1
	ActivityTypeListening ActivityType = 2
	ActivityTypeWatching  ActivityType = 3
)

type ActivityActionType int32

const (
	ActivityActionTypeJoin     ActivityActionType = 1
	ActivityActionTypeSpectate ActivityActionType = 2
)

type ActivityJoinRequestReply int32

const (
	ActivityJoinRequestReplyNo     ActivityJoinRequestReply = 0
	ActivityJoinRequestReplyYes    ActivityJoinRequestReply = 1
	ActivityJoinRequestReplyIgnore ActivityJoinRequestReply = 2
)

type Status int32

const (
	StatusOffline      Status = 0
	StatusOnline       Status = 1
	StatusIdle         Status = 2
	StatusDoNotDisturb Status = 3
)

type RelationshipType int32

const (
	RelationshipTypeNone            RelationshipType = 0
	RelationshipTypeFriend          RelationshipType = 1
	RelationshipTypeBlocked         RelationshipType = 2
	RelationshipTypePendingIncoming RelationshipType = 3
	RelationshipTypePendingOutgoing RelationshipType = 4
	RelationshipTypeImplicit        RelationshipType = 5
)

type LobbyType int32

const (
	LobbyTypePrivate LobbyType = 1
	LobbyTypePublic  LobbyType = 2
)

type LobbySearchComparison int32

const (
	LobbySearchComparisonLessThanOrEqual    LobbySearchComparison = -2
	LobbySearchComparisonLessThan           LobbySearchComparison = -1
	LobbySearchComparisonEqual              LobbySearchComparison = 0
	LobbySearchComparisonGreaterThan        LobbySearchComparison = 1
	LobbySearchComparisonGreaterThanOrEqual LobbySearchComparison = 2
	LobbySearchComparisonNotEqual           LobbySearchComparison = 3
)

type LobbySearchCast int32

const (
	LobbySearchCastString LobbySearchCast = 1
	LobbySearchCastNumber LobbySearchCast = 2
)

type LobbySearchDistance int32

const (
	LobbySearchDistanceLocal    LobbySearchDistance = 0
	LobbySearchDistanceDefault  LobbySearchDistance = 1
	LobbySearchDistanceExtended LobbySearchDistance = 2
	LobbySearchDistanceGlobal   LobbySearchDistance = 3
)

type KeyVariant int32

const (
	KeyVariantNormal KeyVariant = 0
	KeyVariantRight  KeyVariant = 1
	KeyVariantLeft   KeyVariant = 2
)

type MouseButton int32

const (
	MouseButtonLeft   MouseButton = 0
	MouseButtonMiddle MouseButton = 1
	MouseButtonRight  MouseButton = 2
)

type EntitlementType int32

const (
	EntitlementTypePurchase            EntitlementType = 1
	EntitlementTypePremiumSubscription EntitlementType = 2
	EntitlementTypeDeveloperGift       EntitlementType = 3
	EntitlementTypeTestModePurchase    EntitlementType = 4
	EntitlementTypeFreePurchase        EntitlementType = 5
	EntitlementTypeUserGift            EntitlementType = 6
	EntitlementTypePremiumPurchase     EntitlementType = 7
)

type SkuType int32

const (
	SkuTypeApplication SkuType = 1
	SkuTypeDLC         SkuType = 2
	SkuTypeConsumable  SkuType = 3
	SkuTypeBundle      SkuType = 4
)

type InputModeType int32

const (
	InputModeTypeVoiceActivity InputModeType = 0
	InputModeTypePushToTalk    InputModeType = 1
)

type UserFlag uint32

const (
	UserFlagPartner         UserFlag = 2
	UserFlagHypeSquadEvents UserFlag = 4
	UserFlagHypeSquadHouse1 UserFlag = 64
	UserFlagHypeSquadHouse2 UserFlag = 128
	UserFlagHypeSquadHouse3 UserFlag = 256
)

type PremiumType int32

const (
	PremiumTypeNone  PremiumType = 0
	PremiumTypeTier1 PremiumType = 1
	PremiumTypeTier2 PremiumType = 2
)
