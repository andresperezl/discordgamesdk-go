package discord

import (
	"context"
	"fmt"
	"time"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// ActivityClient provides Go-like interfaces for activity management
type ActivityClient struct {
	manager *core.ActivityManager
	core    *core.Core
}

// SetActivity sets the current activity with Go-like error handling
func (ac *ActivityClient) SetActivity(activity *core.Activity) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	resultChan := ac.manager.UpdateActivityAsync(activity)

	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to set activity: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("activity update timed out")
	}
}

// SetActivityWithCallback sets the current activity with a callback
func (ac *ActivityClient) SetActivityWithCallback(activity *core.Activity, callback func(error)) {
	if ac.manager == nil {
		if callback != nil {
			callback(fmt.Errorf("activity manager not available"))
		}
		return
	}

	ac.manager.UpdateActivity(activity, func(result core.Result) {
		if callback != nil {
			if result != core.ResultOk {
				callback(fmt.Errorf("failed to set activity: %v", result))
			} else {
				callback(nil)
			}
		}
	})
}

// SetActivityWithContext sets the current activity, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.SetActivityWithContext(ctx, activity)
//	if err != nil {
//	    log.Fatalf("failed to set activity: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the update fails.
func (ac *ActivityClient) SetActivityWithContext(ctx context.Context, activity *core.Activity) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	resultChan := ac.manager.UpdateActivityAsync(activity)

	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to set activity: %v", result)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ClearActivity clears the current activity with Go-like error handling
func (ac *ActivityClient) ClearActivity() error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	resultChan := ac.manager.ClearActivityAsync()

	select {
	case result := <-resultChan:
		if result != core.ResultOk {
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

	ac.manager.ClearActivity(func(result core.Result) {
		if callback != nil {
			if result != core.ResultOk {
				callback(fmt.Errorf("failed to clear activity: %v", result))
			} else {
				callback(nil)
			}
		}
	})
}

// ClearActivityWithContext clears the current activity, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Activity().ClearActivityWithContext(ctx)
//	if err != nil {
//	    log.Fatalf("failed to clear activity: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the clear fails.
func (ac *ActivityClient) ClearActivityWithContext(ctx context.Context) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}
	resultChan := ac.manager.ClearActivityAsync()
	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to clear activity: %v", result)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SendRequestReply sends a reply to a join request with Go-like error handling
func (ac *ActivityClient) SendRequestReply(userID int64, reply core.ActivityJoinRequestReply) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	// Create a channel to receive the result
	resultChan := make(chan core.Result, 1)

	ac.manager.SendRequestReply(userID, reply, func(result core.Result) {
		resultChan <- result
		close(resultChan)
	})

	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to send request reply: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("send request reply timed out")
	}
}

// SendRequestReplyWithContext sends a reply to a join request, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Activity().SendRequestReplyWithContext(ctx, userID, core.ActivityJoinRequestReplyYes)
//	if err != nil {
//	    log.Fatalf("failed to send request reply: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the operation fails.
func (ac *ActivityClient) SendRequestReplyWithContext(ctx context.Context, userID int64, reply core.ActivityJoinRequestReply) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}
	resultChan := make(chan core.Result, 1)
	ac.manager.SendRequestReply(userID, reply, func(result core.Result) {
		resultChan <- result
		close(resultChan)
	})
	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to send request reply: %v", result)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SendInvite sends an invite to a user with Go-like error handling
func (ac *ActivityClient) SendInvite(userID int64, actionType core.ActivityActionType, content string) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	// Create a channel to receive the result
	resultChan := make(chan core.Result, 1)

	ac.manager.SendInvite(userID, actionType, content, func(result core.Result) {
		resultChan <- result
		close(resultChan)
	})

	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to send invite: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("send invite timed out")
	}
}

// SendInviteWithContext sends an invite to a user, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Activity().SendInviteWithContext(ctx, userID, core.ActivityActionTypeJoin, "Let's play!")
//	if err != nil {
//	    log.Fatalf("failed to send invite: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the operation fails.
func (ac *ActivityClient) SendInviteWithContext(ctx context.Context, userID int64, actionType core.ActivityActionType, content string) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}
	resultChan := make(chan core.Result, 1)
	ac.manager.SendInvite(userID, actionType, content, func(result core.Result) {
		resultChan <- result
		close(resultChan)
	})
	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to send invite: %v", result)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// AcceptInvite accepts an invite with Go-like error handling
func (ac *ActivityClient) AcceptInvite(userID int64) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	// Create a channel to receive the result
	resultChan := make(chan core.Result, 1)

	ac.manager.AcceptInvite(userID, func(result core.Result) {
		resultChan <- result
		close(resultChan)
	})

	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to accept invite: %v", result)
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("accept invite timed out")
	}
}

// AcceptInviteWithContext accepts an invite, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Activity().AcceptInviteWithContext(ctx, userID)
//	if err != nil {
//	    log.Fatalf("failed to accept invite: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the operation fails.
func (ac *ActivityClient) AcceptInviteWithContext(ctx context.Context, userID int64) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}
	resultChan := make(chan core.Result, 1)
	ac.manager.AcceptInvite(userID, func(result core.Result) {
		resultChan <- result
		close(resultChan)
	})
	select {
	case result := <-resultChan:
		if result != core.ResultOk {
			return fmt.Errorf("failed to accept invite: %v", result)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RegisterCommand registers a command for the activity
func (ac *ActivityClient) RegisterCommand(command string) error {
	if ac.manager == nil {
		return fmt.Errorf("activity manager not available")
	}

	result := ac.manager.RegisterCommand(command)
	if result != core.ResultOk {
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
	if result != core.ResultOk {
		return fmt.Errorf("failed to register Steam ID: %v", result)
	}

	return nil
}

// ActivityJoinRequests returns a channel that receives activity join requests as *core.User.
// The channel is closed when the ActivityClient is no longer in use.
//
// Example usage:
//
//	joinRequests := client.Activity().ActivityJoinRequests()
//	go func() {
//	    for user := range joinRequests {
//	        fmt.Printf("Received join request from: %s\n", user.Username)
//	    }
//	}()
func (ac *ActivityClient) ActivityJoinRequests() <-chan *core.User {
	ch := make(chan *core.User, 8)
	if ac.core == nil {
		close(ch)
		return ch
	}

	events := &core.ActivityEvents{
		OnActivityJoinRequest: func(user *core.User) {
			ch <- user
		},
	}
	ac.core.SetActivityEvents(events)
	return ch
}

// ActivityBuilder helps build activities with a fluent interface
type ActivityBuilder struct {
	activity *core.Activity
}

// NewActivity creates a new activity builder
func NewActivity() *ActivityBuilder {
	return &ActivityBuilder{
		activity: &core.Activity{},
	}
}

// SetType sets the activity type
func (ab *ActivityBuilder) SetType(activityType core.ActivityType) *ActivityBuilder {
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
	ab.activity.Timestamps = core.ActivityTimestamps{Start: start, End: end}
	return ab
}

// SetAssets sets the activity assets
func (ab *ActivityBuilder) SetAssets(assets core.ActivityAssets) *ActivityBuilder {
	ab.activity.Assets = assets
	return ab
}

// SetParty sets the activity party
func (ab *ActivityBuilder) SetParty(id string, currentSize, maxSize int32, privacy core.ActivityPartyPrivacy) *ActivityBuilder {
	ab.activity.Party = core.ActivityParty{
		ID: id,
		Size: core.PartySize{
			CurrentSize: currentSize,
			MaxSize:     maxSize,
		},
		Privacy: privacy,
	}
	return ab
}

// SetSecrets sets the activity secrets
func (ab *ActivityBuilder) SetSecrets(match, join, spectate string) *ActivityBuilder {
	ab.activity.Secrets = core.ActivitySecrets{
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
func (ab *ActivityBuilder) Build() *core.Activity {
	return ab.activity
}
