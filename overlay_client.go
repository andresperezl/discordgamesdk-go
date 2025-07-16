package discord

import (
	"context"
	"fmt"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// OverlayClient provides Go-like interfaces for overlay management
type OverlayClient struct {
	manager *core.OverlayManager
	core    *core.Core
}

// IsEnabled checks if the overlay is enabled
func (oc *OverlayClient) IsEnabled() (bool, error) {
	if oc.manager == nil {
		return false, fmt.Errorf("overlay manager not available")
	}
	return oc.manager.IsEnabled(), nil
}

// IsLocked checks if the overlay is locked
func (oc *OverlayClient) IsLocked() (bool, error) {
	if oc.manager == nil {
		return false, fmt.Errorf("overlay manager not available")
	}
	return oc.manager.IsLocked(), nil
}

// SetLocked sets the overlay locked state asynchronously and returns a channel for the result
func (oc *OverlayClient) SetLocked(locked bool) <-chan error {
	errChan := make(chan error, 1)
	if oc.manager == nil {
		errChan <- fmt.Errorf("overlay manager not available")
		close(errChan)
		return errChan
	}
	oc.manager.SetLocked(locked, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to set locked: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

// OpenActivityInvite opens an activity invite asynchronously and returns a channel for the result
func (oc *OverlayClient) OpenActivityInvite(actionType core.ActivityActionType) <-chan error {
	errChan := make(chan error, 1)
	if oc.manager == nil {
		errChan <- fmt.Errorf("overlay manager not available")
		close(errChan)
		return errChan
	}
	oc.manager.OpenActivityInvite(actionType, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to open activity invite: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

// OpenGuildInvite opens a guild invite asynchronously and returns a channel for the result
func (oc *OverlayClient) OpenGuildInvite(code string) <-chan error {
	errChan := make(chan error, 1)
	if oc.manager == nil {
		errChan <- fmt.Errorf("overlay manager not available")
		close(errChan)
		return errChan
	}
	oc.manager.OpenGuildInvite(code, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to open guild invite: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

// OpenVoiceSettings opens the voice settings asynchronously and returns a channel for the result
func (oc *OverlayClient) OpenVoiceSettings() <-chan error {
	errChan := make(chan error, 1)
	if oc.manager == nil {
		errChan <- fmt.Errorf("overlay manager not available")
		close(errChan)
		return errChan
	}
	oc.manager.OpenVoiceSettings(func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to open voice settings: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

// OpenVoiceSettingsWithContext opens the voice settings, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Overlay().OpenVoiceSettingsWithContext(ctx)
//	if err != nil {
//	    log.Fatalf("failed to open voice settings: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the operation fails.
func (oc *OverlayClient) OpenVoiceSettingsWithContext(ctx context.Context) error {
	if oc.manager == nil {
		return fmt.Errorf("overlay manager not available")
	}
	errChan := make(chan error, 1)
	oc.manager.OpenVoiceSettings(func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to open voice settings: %v", result)
		} else {
			errChan <- nil
		}
	})
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetLockedWithContext sets the overlay locked state, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Overlay().SetLockedWithContext(ctx, true)
//	if err != nil {
//	    log.Fatalf("failed to set locked: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the operation fails.
func (oc *OverlayClient) SetLockedWithContext(ctx context.Context, locked bool) error {
	if oc.manager == nil {
		return fmt.Errorf("overlay manager not available")
	}
	errChan := make(chan error, 1)
	oc.manager.SetLocked(locked, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to set locked: %v", result)
		} else {
			errChan <- nil
		}
	})
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// OpenActivityInviteWithContext opens an activity invite, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Overlay().OpenActivityInviteWithContext(ctx, core.ActivityActionTypeJoin)
//	if err != nil {
//	    log.Fatalf("failed to open activity invite: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the operation fails.
func (oc *OverlayClient) OpenActivityInviteWithContext(ctx context.Context, actionType core.ActivityActionType) error {
	if oc.manager == nil {
		return fmt.Errorf("overlay manager not available")
	}
	errChan := make(chan error, 1)
	oc.manager.OpenActivityInvite(actionType, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to open activity invite: %v", result)
		} else {
			errChan <- nil
		}
	})
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// OpenGuildInviteWithContext opens a guild invite, respecting context cancellation and timeout.
//
// Example usage:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	err := client.Overlay().OpenGuildInviteWithContext(ctx, "inviteCode")
//	if err != nil {
//	    log.Fatalf("failed to open guild invite: %v", err)
//	}
//
// Returns an error if the context is cancelled, deadline exceeded, or the operation fails.
func (oc *OverlayClient) OpenGuildInviteWithContext(ctx context.Context, code string) error {
	if oc.manager == nil {
		return fmt.Errorf("overlay manager not available")
	}
	errChan := make(chan error, 1)
	oc.manager.OpenGuildInvite(code, func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to open guild invite: %v", result)
		} else {
			errChan <- nil
		}
	})
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
