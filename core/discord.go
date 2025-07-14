package core

import (
	"fmt"
	"sync"
	"time"
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// StoreManager provides access to store-related functionality
type StoreManager struct {
	manager unsafe.Pointer
}

// VoiceManager provides access to voice-related functionality
type VoiceManager struct {
	manager unsafe.Pointer
}

// AchievementManager provides access to achievement-related functionality
type AchievementManager struct {
	manager unsafe.Pointer
}

// Version constants
const (
	DiscordVersion = 3 // Should match DISCORD_VERSION in the SDK
)

// Result codes
type Result int32

const (
	ResultOk                              Result = 0
	ResultServiceUnavailable              Result = 1
	ResultInvalidVersion                  Result = 2
	ResultLockFailed                      Result = 3
	ResultInternalError                   Result = 4
	ResultInvalidPayload                  Result = 5
	ResultInvalidCommand                  Result = 6
	ResultInvalidPermissions              Result = 7
	ResultNotFetched                      Result = 8
	ResultNotFound                        Result = 9
	ResultConflict                        Result = 10
	ResultInvalidSecret                   Result = 11
	ResultInvalidJoinSecret               Result = 12
	ResultNoEligibleActivity              Result = 13
	ResultInvalidInvite                   Result = 14
	ResultNotAuthenticated                Result = 15
	ResultInvalidAccessToken              Result = 16
	ResultApplicationMismatch             Result = 17
	ResultInvalidDataUrl                  Result = 18
	ResultInvalidBase64                   Result = 19
	ResultNotFiltered                     Result = 20
	ResultLobbyFull                       Result = 21
	ResultInvalidLobbySecret              Result = 22
	ResultInvalidFilename                 Result = 23
	ResultInvalidFileSize                 Result = 24
	ResultInvalidEntitlement              Result = 25
	ResultNotInstalled                    Result = 26
	ResultNotRunning                      Result = 27
	ResultInsufficientBuffer              Result = 28
	ResultPurchaseCanceled                Result = 29
	ResultInvalidGuild                    Result = 30
	ResultInvalidEvent                    Result = 31
	ResultInvalidChannel                  Result = 32
	ResultInvalidOrigin                   Result = 33
	ResultRateLimited                     Result = 34
	ResultOAuth2Error                     Result = 35
	ResultSelectChannelTimeout            Result = 36
	ResultGetGuildTimeout                 Result = 37
	ResultSelectVoiceForceRequired        Result = 38
	ResultCaptureShortcutAlreadyListening Result = 39
	ResultUnauthorizedForAchievement      Result = 40
	ResultInvalidGiftCode                 Result = 41
	ResultPurchaseError                   Result = 42
	ResultTransactionAborted              Result = 43
	ResultDrawingInitFailed               Result = 44
)

// String returns a string representation of the Result
func (r Result) String() string {
	switch r {
	case ResultOk:
		return "ResultOk(0)"
	case ResultServiceUnavailable:
		return "ResultServiceUnavailable(1)"
	case ResultInvalidVersion:
		return "ResultInvalidVersion(2)"
	case ResultLockFailed:
		return "ResultLockFailed(3)"
	case ResultInternalError:
		return "ResultInternalError(4)"
	case ResultInvalidPayload:
		return "ResultInvalidPayload(5)"
	case ResultInvalidCommand:
		return "ResultInvalidCommand(6)"
	case ResultInvalidPermissions:
		return "ResultInvalidPermissions(7)"
	case ResultNotFetched:
		return "ResultNotFetched(8)"
	case ResultNotFound:
		return "ResultNotFound(9)"
	case ResultConflict:
		return "ResultConflict(10)"
	case ResultInvalidSecret:
		return "ResultInvalidSecret(11)"
	case ResultInvalidJoinSecret:
		return "ResultInvalidJoinSecret(12)"
	case ResultNoEligibleActivity:
		return "ResultNoEligibleActivity(13)"
	case ResultInvalidInvite:
		return "ResultInvalidInvite(14)"
	case ResultNotAuthenticated:
		return "ResultNotAuthenticated(15)"
	case ResultInvalidAccessToken:
		return "ResultInvalidAccessToken(16)"
	case ResultApplicationMismatch:
		return "ResultApplicationMismatch(17)"
	case ResultInvalidDataUrl:
		return "ResultInvalidDataUrl(18)"
	case ResultInvalidBase64:
		return "ResultInvalidBase64(19)"
	case ResultNotFiltered:
		return "ResultNotFiltered(20)"
	case ResultLobbyFull:
		return "ResultLobbyFull(21)"
	case ResultInvalidLobbySecret:
		return "ResultInvalidLobbySecret(22)"
	case ResultInvalidFilename:
		return "ResultInvalidFilename(23)"
	case ResultInvalidFileSize:
		return "ResultInvalidFileSize(24)"
	case ResultInvalidEntitlement:
		return "ResultInvalidEntitlement(25)"
	case ResultNotInstalled:
		return "ResultNotInstalled(26)"
	case ResultNotRunning:
		return "ResultNotRunning(27)"
	case ResultInsufficientBuffer:
		return "ResultInsufficientBuffer(28)"
	case ResultPurchaseCanceled:
		return "ResultPurchaseCanceled(29)"
	case ResultInvalidGuild:
		return "ResultInvalidGuild(30)"
	case ResultInvalidEvent:
		return "ResultInvalidEvent(31)"
	case ResultInvalidChannel:
		return "ResultInvalidChannel(32)"
	case ResultInvalidOrigin:
		return "ResultInvalidOrigin(33)"
	case ResultRateLimited:
		return "ResultRateLimited(34)"
	case ResultOAuth2Error:
		return "ResultOAuth2Error(35)"
	case ResultSelectChannelTimeout:
		return "ResultSelectChannelTimeout(36)"
	case ResultGetGuildTimeout:
		return "ResultGetGuildTimeout(37)"
	case ResultSelectVoiceForceRequired:
		return "ResultSelectVoiceForceRequired(38)"
	case ResultCaptureShortcutAlreadyListening:
		return "ResultCaptureShortcutAlreadyListening(39)"
	case ResultUnauthorizedForAchievement:
		return "ResultUnauthorizedForAchievement(40)"
	case ResultInvalidGiftCode:
		return "ResultInvalidGiftCode(41)"
	case ResultPurchaseError:
		return "ResultPurchaseError(42)"
	case ResultTransactionAborted:
		return "ResultTransactionAborted(43)"
	case ResultDrawingInitFailed:
		return "ResultDrawingInitFailed(44)"
	default:
		return fmt.Sprintf("ResultUnknown(%d)", int32(r))
	}
}

// Create flags
type CreateFlags uint64

const (
	CreateFlagsDefault          CreateFlags = 0
	CreateFlagsNoRequireDiscord CreateFlags = 1
)

// Log levels
type LogLevel int32

const (
	LogLevelError LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelInfo  LogLevel = 3
	LogLevelDebug LogLevel = 4
)

// CallbackResult represents a callback that has been executed with its result
type CallbackResult struct {
	CallbackID string
	Result     Result
	Data       interface{}
	Timestamp  time.Time
}

// Core represents the main Discord SDK instance
// Enhanced with robust callback handling and initialization tracking
type Core struct {
	ptr          unsafe.Pointer
	callbackStop chan struct{} // Used to signal the callback goroutine to stop
	callbackDone chan struct{} // Used to signal when the callback goroutine has stopped

	// Enhanced callback handling
	callbackQueue   []CallbackResult
	callbackMutex   sync.RWMutex
	initialized     bool
	initMutex       sync.RWMutex
	callbackID      int64
	callbackIDMutex sync.Mutex
}

// Start begins a background goroutine that continuously calls RunCallbacks.
// This ensures the SDK processes all events and state changes.
func (c *Core) Start() {
	if c.callbackStop != nil {
		return // Already started
	}
	c.callbackStop = make(chan struct{})
	c.callbackDone = make(chan struct{})
	go func() {
		defer close(c.callbackDone)
		for {
			select {
			case <-c.callbackStop:
				return
			default:
				result := c.RunCallbacks()
				if result == ResultOk {
					// Mark as initialized after first successful callback run
					c.initMutex.Lock()
					if !c.initialized {
						c.initialized = true
					}
					c.initMutex.Unlock()
				}
				// 50ms is a good balance for responsiveness and CPU usage
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
}

// Shutdown stops the callback goroutine and cleans up resources.
func (c *Core) Shutdown() {
	if c.callbackStop != nil {
		close(c.callbackStop)
		<-c.callbackDone
		c.callbackStop = nil
		c.callbackDone = nil
	}
	c.Destroy()
}

// WaitForInitialization blocks until the SDK is fully initialized
// Returns true if initialized within timeout, false otherwise
func (c *Core) WaitForInitialization(timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		c.initMutex.RLock()
		if c.initialized {
			c.initMutex.RUnlock()
			return true
		}
		c.initMutex.RUnlock()
		time.Sleep(50 * time.Millisecond)
	}
	return false
}

// WaitForUser blocks until GetCurrentUser returns a valid user or timeout.
// Returns the user and result code. Use this after Start().
func (c *Core) WaitForUser(timeout time.Duration) (*User, Result) {
	// First wait for initialization
	if !c.WaitForInitialization(timeout) {
		return nil, ResultInternalError
	}

	userManager := c.GetUserManager()
	if userManager == nil {
		return nil, ResultInternalError
	}
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		user, result := userManager.GetCurrentUser()
		if result == ResultOk && user != nil && user.ID != 0 {
			return user, ResultOk
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil, ResultNotFound
}

// AddCallbackResult adds a callback result to the queue for tracking
func (c *Core) AddCallbackResult(callbackID string, result Result, data interface{}) {
	c.callbackMutex.Lock()
	defer c.callbackMutex.Unlock()

	c.callbackQueue = append(c.callbackQueue, CallbackResult{
		CallbackID: callbackID,
		Result:     result,
		Data:       data,
		Timestamp:  time.Now(),
	})
}

// GetCallbackResult retrieves a specific callback result by ID
func (c *Core) GetCallbackResult(callbackID string) (CallbackResult, bool) {
	c.callbackMutex.RLock()
	defer c.callbackMutex.RUnlock()

	for _, result := range c.callbackQueue {
		if result.CallbackID == callbackID {
			return result, true
		}
	}
	return CallbackResult{}, false
}

// WaitForCallbackResult waits for a specific callback result
func (c *Core) WaitForCallbackResult(callbackID string, timeout time.Duration) (CallbackResult, bool) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if result, found := c.GetCallbackResult(callbackID); found {
			return result, true
		}
		time.Sleep(50 * time.Millisecond)
	}
	return CallbackResult{}, false
}

// GenerateCallbackID generates a unique callback ID
func (c *Core) GenerateCallbackID() string {
	c.callbackIDMutex.Lock()
	defer c.callbackIDMutex.Unlock()
	c.callbackID++
	return fmt.Sprintf("callback_%d", c.callbackID)
}

// Create creates a new Discord SDK instance
func Create(clientID int64, flags CreateFlags, events *CoreEvents) (*Core, Result) {
	// Use the helper function from discordcgo package
	core, result := dcgo.CoreCreateHelper(clientID, uint64(flags))

	if result != 0 {
		return nil, Result(result)
	}

	return &Core{ptr: core}, ResultOk
}

// Destroy destroys the Discord SDK instance
func (c *Core) Destroy() {
	if c.ptr != nil {
		dcgo.CoreDestroy(c.ptr)
		c.ptr = nil
	}
}

// RunCallbacks runs the Discord SDK callbacks
func (c *Core) RunCallbacks() Result {
	if c.ptr == nil {
		return ResultInternalError
	}
	return Result(dcgo.CoreRunCallbacks(c.ptr))
}

// SetLogHook sets a log hook for the Discord SDK
func (c *Core) SetLogHook(minLevel LogLevel, hook LogHook) {
	if c.ptr == nil {
		return
	}

	// Call the C wrapper function
	dcgo.CoreSetLogHook(c.ptr, int32(minLevel), nil, nil)

	// TODO: Implement proper callback support for log hooks
}

// GetApplicationManager returns the application manager
func (c *Core) GetApplicationManager() *ApplicationManager {
	if c.ptr == nil {
		return nil
	}
	appManager := dcgo.CoreGetApplicationManager(c.ptr)
	if appManager == nil {
		return nil
	}
	return &ApplicationManager{ptr: appManager}
}

// GetUserManager returns the user manager
func (c *Core) GetUserManager() *UserManager {
	if c.ptr == nil {
		return nil
	}
	userManager := dcgo.CoreGetUserManager(c.ptr)
	if userManager == nil {
		return nil
	}
	return &UserManager{ptr: userManager}
}

// GetActivityManager returns the activity manager
func (c *Core) GetActivityManager() *ActivityManager {
	if c.ptr == nil {
		return nil
	}
	activityManager := dcgo.CoreGetActivityManager(c.ptr)
	if activityManager == nil {
		return nil
	}
	manager := &ActivityManager{ptr: activityManager}
	manager.SetCore(c)
	return manager
}

// GetLobbyManager returns the lobby manager
func (c *Core) GetLobbyManager() *LobbyManager {
	if c.ptr == nil {
		return nil
	}
	lobbyManager := dcgo.CoreGetLobbyManager(c.ptr)
	if lobbyManager == nil {
		return nil
	}
	return &LobbyManager{manager: lobbyManager}
}

// GetNetworkManager returns the network manager
func (c *Core) GetNetworkManager() *NetworkManager {
	if c.ptr == nil {
		return nil
	}
	networkManager := dcgo.CoreGetNetworkManager(c.ptr)
	if networkManager == nil {
		return nil
	}
	return &NetworkManager{manager: networkManager}
}

// GetOverlayManager returns the overlay manager
func (c *Core) GetOverlayManager() *OverlayManager {
	if c.ptr == nil {
		return nil
	}
	overlayManager := dcgo.CoreGetOverlayManager(c.ptr)
	if overlayManager == nil {
		return nil
	}
	return &OverlayManager{manager: overlayManager}
}

// GetStorageManager returns the storage manager
func (c *Core) GetStorageManager() *StorageManager {
	if c.ptr == nil {
		return nil
	}
	storageManager := dcgo.CoreGetStorageManager(c.ptr)
	if storageManager == nil {
		return nil
	}
	return &StorageManager{manager: storageManager}
}

// GetStoreManager returns the store manager
func (c *Core) GetStoreManager() *StoreManager {
	if c.ptr == nil {
		return nil
	}
	storeManager := dcgo.CoreGetStoreManager(c.ptr)
	if storeManager == nil {
		return nil
	}
	return &StoreManager{manager: storeManager}
}

// GetVoiceManager returns the voice manager
func (c *Core) GetVoiceManager() *VoiceManager {
	if c.ptr == nil {
		return nil
	}
	voiceManager := dcgo.CoreGetVoiceManager(c.ptr)
	if voiceManager == nil {
		return nil
	}
	return &VoiceManager{manager: voiceManager}
}

// GetAchievementManager returns the achievement manager
func (c *Core) GetAchievementManager() *AchievementManager {
	if c.ptr == nil {
		return nil
	}
	achievementManager := dcgo.CoreGetAchievementManager(c.ptr)
	if achievementManager == nil {
		return nil
	}
	return &AchievementManager{manager: achievementManager}
}

// LogHook represents a log hook function
type LogHook func(level LogLevel, message string)

// StoreManager methods

// CountSkus returns the number of SKUs
func (s *StoreManager) CountSkus() (int32, Result) {
	if s.manager == nil {
		return 0, ResultInternalError
	}
	var count int32
	dcgo.StoreManagerCountSkus(s.manager, unsafe.Pointer(&count))
	return count, ResultOk
}

// GetSku retrieves a SKU by its ID
func (s *StoreManager) GetSku(skuID int64) (*Sku, Result) {
	if s.manager == nil {
		return nil, ResultInternalError
	}
	sku := dcgo.StoreManagerGetSkuGo(s.manager, skuID)
	if sku == nil {
		return nil, ResultInternalError
	}
	return convertDiscordSku(sku), ResultOk
}

// GetSkuAt retrieves a SKU by index
func (s *StoreManager) GetSkuAt(index int32) (*Sku, Result) {
	if s.manager == nil {
		return nil, ResultInternalError
	}
	sku := dcgo.StoreManagerGetSkuAtGo(s.manager, index)
	if sku == nil {
		return nil, ResultInternalError
	}
	return convertDiscordSku(sku), ResultOk
}

// GetEntitlement gets a single entitlement by ID
func (s *StoreManager) GetEntitlement(entitlementID int64) (*Entitlement, Result) {
	if s.manager == nil {
		return nil, ResultInternalError
	}
	ptr := dcgo.MallocDiscordEntitlement()
	defer dcgo.Free(ptr)
	res := dcgo.StoreManagerGetEntitlement(s.manager, entitlementID, ptr)
	if res != 0 {
		return nil, Result(res)
	}
	return &Entitlement{
		ID:    dcgo.GetDiscordEntitlementID(ptr),
		Type:  EntitlementType(dcgo.GetDiscordEntitlementType(ptr)),
		SkuID: dcgo.GetDiscordEntitlementSkuID(ptr),
	}, ResultOk
}

// GetEntitlementAt gets an entitlement at index
func (s *StoreManager) GetEntitlementAt(index int32) (*Entitlement, Result) {
	if s.manager == nil {
		return nil, ResultInternalError
	}
	ptr := dcgo.MallocDiscordEntitlement()
	defer dcgo.Free(ptr)
	res := dcgo.StoreManagerGetEntitlementAt(s.manager, index, ptr)
	if res != 0 {
		return nil, Result(res)
	}
	return &Entitlement{
		ID:    dcgo.GetDiscordEntitlementID(ptr),
		Type:  EntitlementType(dcgo.GetDiscordEntitlementType(ptr)),
		SkuID: dcgo.GetDiscordEntitlementSkuID(ptr),
	}, ResultOk
}

// CountEntitlements gets the count of entitlements
func (s *StoreManager) CountEntitlements() (int32, Result) {
	if s.manager == nil {
		return 0, ResultInternalError
	}
	var count int32
	dcgo.StoreManagerCountEntitlements(s.manager, unsafe.Pointer(&count))
	return count, ResultOk
}

// HasSkuEntitlement checks if a SKU has an entitlement
func (s *StoreManager) HasSkuEntitlement(skuID int64) (bool, Result) {
	if s.manager == nil {
		return false, ResultInternalError
	}
	var has bool
	res := dcgo.StoreManagerHasSkuEntitlement(s.manager, skuID, unsafe.Pointer(&has))
	return has, Result(res)
}

// Helper conversion functions
func convertDiscordSku(sku *dcgo.DiscordSku) *Sku {
	if sku == nil {
		return nil
	}
	return &Sku{
		ID:   int64(dcgo.GetDiscordSkuID(unsafe.Pointer(sku))),
		Type: SkuType(dcgo.GetDiscordSkuType(unsafe.Pointer(sku))),
		Name: dcgo.GetDiscordSkuName(unsafe.Pointer(sku)),
		Price: SkuPrice{
			Amount:   uint32(dcgo.GetDiscordSkuPriceAmount(unsafe.Pointer(sku))),
			Currency: dcgo.GetDiscordSkuPriceCurrency(unsafe.Pointer(sku)),
		},
	}
}

func convertDiscordEntitlement(ent *dcgo.DiscordEntitlement) *Entitlement {
	if ent == nil {
		return nil
	}
	return &Entitlement{
		ID:    int64(dcgo.GetDiscordEntitlementID(unsafe.Pointer(ent))),
		Type:  EntitlementType(dcgo.GetDiscordEntitlementType(unsafe.Pointer(ent))),
		SkuID: int64(dcgo.GetDiscordEntitlementSkuID(unsafe.Pointer(ent))),
	}
}
