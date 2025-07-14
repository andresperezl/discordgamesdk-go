package discord

import (
	"fmt"
)

// ApplicationClient provides Go-like interfaces for application management
type ApplicationClient struct {
	manager *ApplicationManager
	core    *Core
}

// GetCurrentLocale returns the current locale
func (ac *ApplicationClient) GetCurrentLocale() (string, error) {
	if ac.manager == nil {
		return "", fmt.Errorf("application manager not available")
	}
	return ac.manager.GetCurrentLocale(), nil
}

// GetOAuth2Token gets an OAuth2 token asynchronously and returns a channel for the result
func (ac *ApplicationClient) GetOAuth2Token() (<-chan *OAuth2Token, <-chan error) {
	tokenChan := make(chan *OAuth2Token, 1)
	errChan := make(chan error, 1)
	if ac.manager == nil {
		errChan <- fmt.Errorf("application manager not available")
		close(tokenChan)
		close(errChan)
		return tokenChan, errChan
	}
	ac.manager.GetOAuth2Token(func(result Result, token *OAuth2Token) {
		if result != ResultOk {
			errChan <- fmt.Errorf("failed to get OAuth2 token: %v", result)
			close(tokenChan)
			close(errChan)
			return
		}
		tokenChan <- token
		close(tokenChan)
		close(errChan)
	})
	return tokenChan, errChan
}

// ValidateOrExit validates the application asynchronously and returns a channel for the result
func (ac *ApplicationClient) ValidateOrExit() <-chan error {
	errChan := make(chan error, 1)
	if ac.manager == nil {
		errChan <- fmt.Errorf("application manager not available")
		close(errChan)
		return errChan
	}
	ac.manager.ValidateOrExit(func(result Result) {
		if result != ResultOk {
			errChan <- fmt.Errorf("failed to validate or exit: %v", result)
		} else {
			errChan <- nil
		}
		close(errChan)
	})
	return errChan
}
