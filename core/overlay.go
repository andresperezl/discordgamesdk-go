package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// OverlayManager provides access to overlay-related functionality
type OverlayManager struct {
	manager unsafe.Pointer
}

// IsEnabled checks if the overlay is enabled
func (o *OverlayManager) IsEnabled() bool {
	if o.manager == nil {
		return false
	}

	var enabled bool
	dcgo.OverlayManagerIsEnabled(o.manager, unsafe.Pointer(&enabled))
	return enabled
}

// IsLocked checks if the overlay is locked
func (o *OverlayManager) IsLocked() bool {
	if o.manager == nil {
		return false
	}

	var locked bool
	dcgo.OverlayManagerIsLocked(o.manager, unsafe.Pointer(&locked))
	return locked
}

// SetLocked sets the overlay locked state
func (o *OverlayManager) SetLocked(locked bool, callback func(result Result)) {
	if o.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Call the C wrapper function
	dcgo.OverlayManagerSetLocked(o.manager, locked, nil, nil)

	// For now, call the callback immediately since we don't have proper callback support
	if callback != nil {
		callback(ResultOk)
	}
}

// OpenActivityInvite opens an activity invite
func (o *OverlayManager) OpenActivityInvite(actionType ActivityActionType, callback func(result Result)) {
	if o.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Call the C wrapper function
	dcgo.OverlayManagerOpenActivityInvite(o.manager, int32(actionType), nil, nil)

	// For now, call the callback immediately since we don't have proper callback support
	if callback != nil {
		callback(ResultOk)
	}
}

// OpenGuildInvite opens a guild invite
func (o *OverlayManager) OpenGuildInvite(code string, callback func(result Result)) {
	if o.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}
	cCode := dcgo.GoStringToCChar(code)
	defer dcgo.FreeCChar(cCode)
	dcgo.OverlayManagerOpenGuildInvite(o.manager, cCode, nil, nil)
	if callback != nil {
		callback(ResultOk)
	}
}

// OpenVoiceSettings opens voice settings
func (o *OverlayManager) OpenVoiceSettings(callback func(result Result)) {
	if o.manager == nil {
		if callback != nil {
			callback(ResultInternalError)
		}
		return
	}

	// Call the C wrapper function
	dcgo.OverlayManagerOpenVoiceSettings(o.manager, nil, nil)

	// For now, call the callback immediately since we don't have proper callback support
	if callback != nil {
		callback(ResultOk)
	}
}

// InitDrawingDXGI initializes drawing with DXGI
func (o *OverlayManager) InitDrawingDXGI(swapchain unsafe.Pointer, useMessageForwarding bool) Result {
	if o.manager == nil {
		return ResultInternalError
	}
	return Result(dcgo.OverlayManagerInitDrawingDXGI(o.manager, swapchain, useMessageForwarding))
}

// OnPresent handles present events
func (o *OverlayManager) OnPresent() {
	if o.manager == nil {
		return
	}
	dcgo.OverlayManagerOnPresent(o.manager)
}

// ForwardMessage forwards a message
func (o *OverlayManager) ForwardMessage(message unsafe.Pointer) {
	if o.manager == nil {
		return
	}
	dcgo.OverlayManagerForwardMessage(o.manager, message)
}

// KeyEvent handles key events
func (o *OverlayManager) KeyEvent(down bool, keyCode string, variant KeyVariant) {
	if o.manager == nil {
		return
	}
	cKey := dcgo.GoStringToCChar(keyCode)
	defer dcgo.FreeCChar(cKey)
	dcgo.OverlayManagerKeyEvent(o.manager, down, cKey, int32(variant))
}

// CharEvent handles character events
func (o *OverlayManager) CharEvent(character string) {
	if o.manager == nil {
		return
	}
	cChar := dcgo.GoStringToCChar(character)
	defer dcgo.FreeCChar(cChar)
	dcgo.OverlayManagerCharEvent(o.manager, cChar)
}

// MouseButtonEvent handles mouse button events
func (o *OverlayManager) MouseButtonEvent(down uint8, clickCount int32, which MouseButton, x, y int32) {
	if o.manager == nil {
		return
	}
	dcgo.OverlayManagerMouseButtonEvent(o.manager, down, clickCount, int32(which), x, y)
}

// MouseMotionEvent handles mouse motion events
func (o *OverlayManager) MouseMotionEvent(x, y int32) {
	if o.manager == nil {
		return
	}
	dcgo.OverlayManagerMouseMotionEvent(o.manager, x, y)
}

// IsPointInsideClickZone checks if a point is inside the click zone
func (o *OverlayManager) IsPointInsideClickZone(x, y int32) bool {
	if o.manager == nil {
		return false
	}
	return dcgo.OverlayManagerIsPointInsideClickZone(o.manager, x, y)
}
