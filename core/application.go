package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// ApplicationManager provides access to application-related functionality
type ApplicationManager struct {
	ptr unsafe.Pointer
}

// ValidateOrExit validates the application or exits if validation fails
func (a *ApplicationManager) ValidateOrExit(callback func(result Result)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Call the C wrapper function
	dcgo.ApplicationManagerValidateOrExit(a.ptr, nil, nil)

	// For now, call the callback immediately since we don't have proper callback support
	if callback != nil {
		callback(ResultOk)
	}
}

// GetCurrentLocale gets the current locale
func (a *ApplicationManager) GetCurrentLocale() string {
	if a.ptr == nil {
		return ""
	}

	var locale [128]byte // DiscordLocale is 128 bytes
	dcgo.ApplicationManagerGetCurrentLocale(a.ptr, unsafe.Pointer(&locale[0]))
	return string(locale[:])
}

// GetCurrentBranch gets the current branch
func (a *ApplicationManager) GetCurrentBranch() string {
	if a.ptr == nil {
		return ""
	}

	var branch [4096]byte // DiscordBranch is 4096 bytes
	dcgo.ApplicationManagerGetCurrentBranch(a.ptr, unsafe.Pointer(&branch[0]))
	return string(branch[:])
}

// GetOAuth2Token gets an OAuth2 token
func (a *ApplicationManager) GetOAuth2Token(callback func(result Result, token *OAuth2Token)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError, nil)
		}
		return
	}

	// Call the C wrapper function
	dcgo.ApplicationManagerGetOAuth2Token(a.ptr, nil, nil)

	// For now, call the callback immediately since we don't have proper callback support
	if callback != nil {
		callback(ResultOk, nil)
	}
}

// GetTicket gets a ticket
func (a *ApplicationManager) GetTicket(callback func(result Result, data string)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError, "")
		}
		return
	}

	// Call the C wrapper function
	dcgo.ApplicationManagerGetTicket(a.ptr, nil, nil)

	// For now, call the callback immediately since we don't have proper callback support
	if callback != nil {
		callback(ResultOk, "")
	}
}
