package core

import (
	"time"
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// ActivityManager provides access to activity-related functionality
type ActivityManager struct {
	ptr  unsafe.Pointer
	core *Core // Reference to the core for callback tracking
}

// SetCore sets the core reference for callback tracking
func (a *ActivityManager) SetCore(core *Core) {
	a.core = core
}

// RegisterCommand registers a command for the activity
func (a *ActivityManager) RegisterCommand(command string) Result {
	if a.ptr == nil {
		return ResultInternalError
	}
	cCommand := dcgo.GoStringToCChar(command)
	defer dcgo.FreeCChar(cCommand)
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.ActivityManagerRegisterCommand(a.ptr, cCommand)
	})
	return Result(res)
}

// RegisterSteam registers a Steam ID for the activity
func (a *ActivityManager) RegisterSteam(steamID uint32) Result {
	if a.ptr == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.ActivityManagerRegisterSteam(a.ptr, steamID)
	})
	return Result(res)
}

// UpdateActivity updates the current activity with proper callback handling
func (a *ActivityManager) UpdateActivity(activity *Activity, callback func(result Result)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Generate callback ID for tracking
	callbackID := ""
	if a.core != nil {
		callbackID = a.core.GenerateCallbackID()
	}

	// Convert Go Activity to C struct
	var cActivity struct {
		Type          int32
		ApplicationID int64
		Name          [128]byte
		State         [128]byte
		Details       [128]byte
		Timestamps    struct {
			Start int64
			End   int64
		}
		Assets struct {
			LargeImage [128]byte
			LargeText  [128]byte
			SmallImage [128]byte
			SmallText  [128]byte
		}
		Party struct {
			ID   [128]byte
			Size struct {
				CurrentSize int32
				MaxSize     int32
			}
			Privacy int32
		}
		Secrets struct {
			Match    [128]byte
			Join     [128]byte
			Spectate [128]byte
		}
		Instance           bool
		SupportedPlatforms uint32
	}

	// Copy data from Go struct to C struct
	cActivity.Type = int32(activity.Type)
	cActivity.ApplicationID = activity.ApplicationID
	copy(cActivity.Name[:], activity.Name)
	copy(cActivity.State[:], activity.State)
	copy(cActivity.Details[:], activity.Details)
	cActivity.Timestamps.Start = activity.Timestamps.Start
	cActivity.Timestamps.End = activity.Timestamps.End
	copy(cActivity.Assets.LargeImage[:], activity.Assets.LargeImage)
	copy(cActivity.Assets.LargeText[:], activity.Assets.LargeText)
	copy(cActivity.Assets.SmallImage[:], activity.Assets.SmallImage)
	copy(cActivity.Assets.SmallText[:], activity.Assets.SmallText)
	copy(cActivity.Party.ID[:], activity.Party.ID)
	cActivity.Party.Size.CurrentSize = activity.Party.Size.CurrentSize
	cActivity.Party.Size.MaxSize = activity.Party.Size.MaxSize
	cActivity.Party.Privacy = int32(activity.Party.Privacy)
	copy(cActivity.Secrets.Match[:], activity.Secrets.Match)
	copy(cActivity.Secrets.Join[:], activity.Secrets.Join)
	copy(cActivity.Secrets.Spectate[:], activity.Secrets.Spectate)
	cActivity.Instance = activity.Instance
	cActivity.SupportedPlatforms = activity.SupportedPlatforms

	// Call the C wrapper function
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ActivityManagerUpdateActivity(a.ptr, unsafe.Pointer(&cActivity), nil, nil)
		return nil
	})

	// If we have callback tracking, wait for the result
	if callback != nil && a.core != nil && callbackID != "" {
		// Wait for callback result with timeout
		if result, found := a.core.WaitForCallbackResult(callbackID, 5*time.Second); found {
			callback(result.Result)
		} else {
			// Fallback to immediate callback if tracking fails
			callback(ResultOk)
		}
	} else if callback != nil {
		// Fallback for immediate callback
		callback(ResultOk)
	}
}

// UpdateActivityAsync updates the current activity and returns a channel for the result
func (a *ActivityManager) UpdateActivityAsync(activity *Activity) chan Result {
	resultChan := make(chan Result, 1)

	a.UpdateActivity(activity, func(result Result) {
		resultChan <- result
		close(resultChan)
	})

	return resultChan
}

// ClearActivity clears the current activity with proper callback handling
func (a *ActivityManager) ClearActivity(callback func(result Result)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Generate callback ID for tracking
	callbackID := ""
	if a.core != nil {
		callbackID = a.core.GenerateCallbackID()
	}

	// Call the C wrapper function
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ActivityManagerClearActivity(a.ptr, nil, nil)
		return nil
	})

	// If we have callback tracking, wait for the result
	if callback != nil && a.core != nil && callbackID != "" {
		// Wait for callback result with timeout
		if result, found := a.core.WaitForCallbackResult(callbackID, 5*time.Second); found {
			callback(result.Result)
		} else {
			// Fallback to immediate callback if tracking fails
			callback(ResultOk)
		}
	} else if callback != nil {
		// Fallback for immediate callback
		callback(ResultOk)
	}
}

// ClearActivityAsync clears the current activity and returns a channel for the result
func (a *ActivityManager) ClearActivityAsync() chan Result {
	resultChan := make(chan Result, 1)

	a.ClearActivity(func(result Result) {
		resultChan <- result
		close(resultChan)
	})

	return resultChan
}

// SendRequestReply sends a reply to a join request with proper callback handling
func (a *ActivityManager) SendRequestReply(userID int64, reply ActivityJoinRequestReply, callback func(result Result)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Generate callback ID for tracking
	callbackID := ""
	if a.core != nil {
		callbackID = a.core.GenerateCallbackID()
	}

	// Call the C wrapper function
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ActivityManagerSendRequestReply(a.ptr, userID, int32(reply), nil, nil)
		return nil
	})

	// If we have callback tracking, wait for the result
	if callback != nil && a.core != nil && callbackID != "" {
		// Wait for callback result with timeout
		if result, found := a.core.WaitForCallbackResult(callbackID, 5*time.Second); found {
			callback(result.Result)
		} else {
			// Fallback to immediate callback if tracking fails
			callback(ResultOk)
		}
	} else if callback != nil {
		// Fallback for immediate callback
		callback(ResultOk)
	}
}

// SendInvite sends an invite to a user with proper callback handling
func (a *ActivityManager) SendInvite(userID int64, actionType ActivityActionType, content string, callback func(result Result)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	callbackID := ""
	if a.core != nil {
		callbackID = a.core.GenerateCallbackID()
	}

	cContent := dcgo.GoStringToCChar(content)
	defer dcgo.FreeCChar(cContent)
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ActivityManagerSendInvite(a.ptr, userID, int32(actionType), cContent, nil, nil)
		return nil
	})

	if callback != nil && a.core != nil && callbackID != "" {
		if result, found := a.core.WaitForCallbackResult(callbackID, 5*time.Second); found {
			callback(result.Result)
		} else {
			callback(ResultOk)
		}
	} else if callback != nil {
		callback(ResultOk)
	}
}

// AcceptInvite accepts an invite from a user with proper callback handling
func (a *ActivityManager) AcceptInvite(userID int64, callback func(result Result)) {
	if a.ptr == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Generate callback ID for tracking
	callbackID := ""
	if a.core != nil {
		callbackID = a.core.GenerateCallbackID()
	}

	// Call the C wrapper function
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.ActivityManagerAcceptInvite(a.ptr, userID, nil, nil)
		return nil
	})

	// If we have callback tracking, wait for the result
	if callback != nil && a.core != nil && callbackID != "" {
		// Wait for callback result with timeout
		if result, found := a.core.WaitForCallbackResult(callbackID, 5*time.Second); found {
			callback(result.Result)
		} else {
			// Fallback to immediate callback if tracking fails
			callback(ResultOk)
		}
	} else if callback != nil {
		// Fallback for immediate callback
		callback(ResultOk)
	}
}
