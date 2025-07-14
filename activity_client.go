package discord

import (
	"fmt"
	"time"
)

// ActivityClient provides Go-like interfaces for activity management
type ActivityClient struct {
	manager *ActivityManager
	core    *Core
}

// SetActivity sets the current activity with Go-like error handling
func (ac *ActivityClient) SetActivity(activity *Activity) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	resultChan := ac.manager.UpdateActivityAsync(activity)

	select {
	case result := <-resultChan:
		if result != ResultOk {
			return fmt.Errorf("failed to set activity: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("activity update timed out")
	}
}

// SetActivityWithCallback sets the current activity with a callback
func (ac *ActivityClient) SetActivityWithCallback(activity *Activity, callback func(error)) {
	if ac.manager == nil {
		if callback != nil {
			callback(fmt.Errorf("activity manager not available"))
		}
		return
	}

	ac.manager.UpdateActivity(activity, func(result Result) {
		if callback != nil {
			if result != ResultOk {
				callback(fmt.Errorf("failed to set activity: %v", result))
			} else {
				callback(nil)
			}
		}
	})
}

// ClearActivity clears the current activity with Go-like error handling
func (ac *ActivityClient) ClearActivity() error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	resultChan := ac.manager.ClearActivityAsync()

	select {
	case result := <-resultChan:
		if result != ResultOk {
			return fmt.Errorf("failed to clear activity: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("activity clear timed out")
	}
}

// ClearActivityWithCallback clears the current activity with a callback
func (ac *ActivityClient) ClearActivityWithCallback(callback func(error)) {
	if ac.manager == nil {
		if callback != nil {
			callback(fmt.Errorf("activity manager not available"))
		}
		return
	}

	ac.manager.ClearActivity(func(result Result) {
		if callback != nil {
			if result != ResultOk {
				callback(fmt.Errorf("failed to clear activity: %v", result))
			} else {
				callback(nil)
			}
		}
	})
}

// SendRequestReply sends a reply to a join request with Go-like error handling
func (ac *ActivityClient) SendRequestReply(userID int64, reply ActivityJoinRequestReply) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	// Create a channel to receive the result
	resultChan := make(chan Result, 1)

	ac.manager.SendRequestReply(userID, reply, func(result Result) {
		resultChan <- result
		close(resultChan)
	})

	select {
	case result := <-resultChan:
		if result != ResultOk {
			return fmt.Errorf("failed to send request reply: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("send request reply timed out")
	}
}

// SendInvite sends an invite to a user with Go-like error handling
func (ac *ActivityClient) SendInvite(userID int64, actionType ActivityActionType, content string) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	// Create a channel to receive the result
	resultChan := make(chan Result, 1)

	ac.manager.SendInvite(userID, actionType, content, func(result Result) {
		resultChan <- result
		close(resultChan)
	})

	select {
	case result := <-resultChan:
		if result != ResultOk {
			return fmt.Errorf("failed to send invite: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("send invite timed out")
	}
}

// AcceptInvite accepts an invite with Go-like error handling
func (ac *ActivityClient) AcceptInvite(userID int64) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	// Create a channel to receive the result
	resultChan := make(chan Result, 1)

	ac.manager.AcceptInvite(userID, func(result Result) {
		resultChan <- result
		close(resultChan)
	})

	select {
	case result := <-resultChan:
		if result != ResultOk {
			return fmt.Errorf("failed to accept invite: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("accept invite timed out")
	}
}

// RegisterCommand registers a command for the activity
func (ac *ActivityClient) RegisterCommand(command string) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	result := ac.manager.RegisterCommand(command)
	if result != ResultOk {
		return fmt.Errorf("failed to register command: %v", result)
	}

	return nil
}

// RegisterSteam registers a Steam ID for the activity
func (ac *ActivityClient) RegisterSteam(steamID uint32) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	result := ac.manager.RegisterSteam(steamID)
	if result != ResultOk {
		return fmt.Errorf("failed to register Steam ID: %v", result)
	}

	return nil
}

// ActivityBuilder helps build activities with a fluent interface
type ActivityBuilder struct {
	activity *Activity
}

// NewActivity creates a new activity builder
func NewActivity() *ActivityBuilder {
	return &ActivityBuilder{
		activity: &Activity{},
	}
}

// SetType sets the activity type
func (ab *ActivityBuilder) SetType(activityType ActivityType) *ActivityBuilder {
	ab.activity.Type = activityType
	return ab
}

// SetApplicationID sets the application ID
func (ab *ActivityBuilder) SetApplicationID(appID int64) *ActivityBuilder {
	ab.activity.ApplicationID = appID
	return ab
}

// SetName sets the activity name
func (ab *ActivityBuilder) SetName(name string) *ActivityBuilder {
	ab.activity.Name = name
	return ab
}

// SetState sets the activity state
func (ab *ActivityBuilder) SetState(state string) *ActivityBuilder {
	ab.activity.State = state
	return ab
}

// SetDetails sets the activity details
func (ab *ActivityBuilder) SetDetails(details string) *ActivityBuilder {
	ab.activity.Details = details
	return ab
}

// SetTimestamps sets the activity timestamps
func (ab *ActivityBuilder) SetTimestamps(start, end int64) *ActivityBuilder {
	ab.activity.Timestamps = ActivityTimestamps{
		Start: start,
		End:   end,
	}
	return ab
}

// SetAssets sets the activity assets
func (ab *ActivityBuilder) SetAssets(largeImage, largeText, smallImage, smallText string) *ActivityBuilder {
	ab.activity.Assets = ActivityAssets{
		LargeImage: largeImage,
		LargeText:  largeText,
		SmallImage: smallImage,
		SmallText:  smallText,
	}
	return ab
}

// SetParty sets the activity party
func (ab *ActivityBuilder) SetParty(id string, currentSize, maxSize int32, privacy ActivityPartyPrivacy) *ActivityBuilder {
	ab.activity.Party = ActivityParty{
		ID: id,
		Size: PartySize{
			CurrentSize: currentSize,
			MaxSize:     maxSize,
		},
		Privacy: privacy,
	}
	return ab
}

// SetSecrets sets the activity secrets
func (ab *ActivityBuilder) SetSecrets(match, join, spectate string) *ActivityBuilder {
	ab.activity.Secrets = ActivitySecrets{
		Match:    match,
		Join:     join,
		Spectate: spectate,
	}
	return ab
}

// SetInstance sets whether this is an instance
func (ab *ActivityBuilder) SetInstance(instance bool) *ActivityBuilder {
	ab.activity.Instance = instance
	return ab
}

// SetSupportedPlatforms sets the supported platforms
func (ab *ActivityBuilder) SetSupportedPlatforms(platforms uint32) *ActivityBuilder {
	ab.activity.SupportedPlatforms = platforms
	return ab
}

// Build returns the built activity
func (ab *ActivityBuilder) Build() *Activity {
	return ab.activity
}
