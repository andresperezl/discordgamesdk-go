package discord

import (
	"fmt"
)

// OverlayClient provides Go-like interfaces for overlay management
type OverlayClient struct {
	manager *OverlayManager
	core    *Core
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
	oc.manager.SetLocked(locked, func(result Result) {
		if result != ResultOk {
			errChan <- fmt.Errorf("failed to set locked: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

// OpenActivityInvite opens an activity invite asynchronously and returns a channel for the result
func (oc *OverlayClient) OpenActivityInvite(actionType ActivityActionType) <-chan error {
	errChan := make(chan error, 1)
	if oc.manager == nil {
		errChan <- fmt.Errorf("overlay manager not available")
		close(errChan)
		return errChan
	}
	oc.manager.OpenActivityInvite(actionType, func(result Result) {
		if result != ResultOk {
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
	oc.manager.OpenGuildInvite(code, func(result Result) {
		if result != ResultOk {
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
	oc.manager.OpenVoiceSettings(func(result Result) {
		if result != ResultOk {
			errChan <- fmt.Errorf("failed to open voice settings: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}
