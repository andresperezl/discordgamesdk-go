package discord

import (
	"context"
	"fmt"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// UserClient provides Go-like interfaces for user management
type UserClient struct {
	manager *core.UserManager
	core    *core.Core
}

// GetCurrentUser returns the current user with Go-like error handling
func (uc *UserClient) GetCurrentUser() (*core.User, error) {
	if uc.manager == nil {
		return nil, fmt.Errorf("user manager not available")
	}

	user, result := uc.manager.GetCurrentUser()
	if result != core.ResultOk {
		return nil, fmt.Errorf("failed to get current user: %v", result)
	}

	return user, nil
}

// GetUser gets a user by ID with Go-like error handling
func (uc *UserClient) GetUser(userID int64) (*core.User, error) {
	if uc.manager == nil {
		return nil, fmt.Errorf("user manager not available")
	}

	// Create a channel to receive the result
	resultChan := make(chan *core.User, 1)
	errChan := make(chan error, 1)

	uc.manager.GetUser(userID, func(result core.Result, user *core.User) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to get user: %v", result)
			return
		}
		resultChan <- user
	})

	select {
	case user := <-resultChan:
		return user, nil
	case err := <-errChan:
		return nil, err
	}
}

// GetCurrentUserPremiumType returns the current user's premium type
func (uc *UserClient) GetCurrentUserPremiumType() (core.PremiumType, error) {
	if uc.manager == nil {
		return core.PremiumTypeNone, fmt.Errorf("user manager not available")
	}

	premiumType, result := uc.manager.GetCurrentUserPremiumType()
	if result != core.ResultOk {
		return core.PremiumTypeNone, fmt.Errorf("failed to get premium type: %v", result)
	}

	return premiumType, nil
}

// CurrentUserHasFlag checks if the current user has a specific flag
func (uc *UserClient) CurrentUserHasFlag(flag core.UserFlag) (bool, error) {
	if uc.manager == nil {
		return false, fmt.Errorf("user manager not available")
	}

	hasFlag, result := uc.manager.CurrentUserHasFlag(flag)
	if result != core.ResultOk {
		return false, fmt.Errorf("failed to check user flag: %v", result)
	}

	return hasFlag, nil
}

// GetUserWithContext gets a user by ID, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	user, err := client.User().GetUserWithContext(ctx, userID)
//	if err != nil {
//	    log.Fatalf("failed to get user: %v", err)
//	}
//	fmt.Printf("Fetched user: %s\n", user.Username)
//
// Returns the user or error if the context is cancelled, deadline exceeded, or the fetch fails.
func (uc *UserClient) GetUserWithContext(ctx context.Context, userID int64) (*core.User, error) {
	if uc.manager == nil {
		return nil, fmt.Errorf("user manager not available")
	}
	resultChan := make(chan *core.User, 1)
	errChan := make(chan error, 1)

	uc.manager.GetUser(userID, func(result core.Result, user *core.User) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to get user: %v", result)
			return
		}
		resultChan <- user
	})

	select {
	case user := <-resultChan:
		return user, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// UserBuilder helps build user-related queries with a fluent interface
type UserBuilder struct {
	userID int64
}

// NewUser creates a new user builder
func NewUser(userID int64) *UserBuilder {
	return &UserBuilder{
		userID: userID,
	}
}

// Get retrieves the user
func (ub *UserBuilder) Get(client *Client) (*core.User, error) {
	return client.User().GetUser(ub.userID)
}
