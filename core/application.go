package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordgamesdk-go/discordcgo"
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ApplicationManagerValidateOrExit(a.ptr, nil, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ApplicationManagerGetCurrentLocale(a.ptr, unsafe.Pointer(&locale[0]))
		return nil
	})
	return string(locale[:])
}

// GetCurrentBranch gets the current branch
func (a *ApplicationManager) GetCurrentBranch() string {
	if a.ptr == nil {
		return ""
	}

	var branch [4096]byte // DiscordBranch is 4096 bytes
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ApplicationManagerGetCurrentBranch(a.ptr, unsafe.Pointer(&branch[0]))
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ApplicationManagerGetOAuth2Token(a.ptr, nil, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ApplicationManagerGetTicket(a.ptr, nil, nil)
		return nil
	})
	if callback != nil {
		callback(ResultOk, "")
	}
}
