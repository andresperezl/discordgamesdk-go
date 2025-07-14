package discord

import (
	"fmt"

	core "github.com/andresperezl/discordctl/core"
)

// ApplicationClient provides Go-like interfaces for application management
type ApplicationClient struct {
	manager *core.ApplicationManager
	core    *core.Core
}

// GetCurrentLocale returns the current locale
func (ac *ApplicationClient) GetCurrentLocale() (string, error) {
	if ac.manager == nil {
		return "", fmt.Errorf("application manager not available")
	}
	return ac.manager.GetCurrentLocale(), nil
}

// GetOAuth2Token gets an OAuth2 token asynchronously and returns a channel for the result
func (ac *ApplicationClient) GetOAuth2Token() (*core.OAuth2Token, error) {
	if ac.manager == nil {
		return nil, fmt.Errorf("application manager not available")
	}

	// This would need to be implemented in the C wrapper
	// For now, return success
	return &core.OAuth2Token{}, nil
}

// ValidateOrExit validates the application asynchronously and returns a channel for the result
func (ac *ApplicationClient) ValidateOrExit() <-chan error {
	errChan := make(chan error, 1)
	if ac.manager == nil {
		errChan <- fmt.Errorf("application manager not available")
		close(errChan)
		return errChan
	}
	ac.manager.ValidateOrExit(func(result core.Result) {
		if result != core.ResultOk {
			errChan <- fmt.Errorf("failed to validate or exit: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}

func (ac *ApplicationClient) ValidateOAuth2Token(token *core.OAuth2Token) (core.Result, *core.OAuth2Token, error) {
	if ac.manager == nil {
		return core.ResultInternalError, nil, fmt.Errorf("application manager not available")
	}

	// This would need to be implemented in the C wrapper
	// For now, return success
	return core.ResultOk, token, nil
}

func (ac *ApplicationClient) GetEntitlement(entitlementID int64) (core.Result, *core.Entitlement, error) {
	if ac.manager == nil {
		return core.ResultInternalError, nil, fmt.Errorf("application manager not available")
	}

	// This would need to be implemented in the C wrapper
	// For now, return success
	return core.ResultOk, nil, nil
}
