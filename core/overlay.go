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

	// For now, call the callback immediately since we don't have proper callback support
	// TODO: Implement proper string conversion and C wrapper call
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

	// NOTE: DXGI initialization not implemented in wrapper approach yet
	return ResultInternalError
}

// OnPresent handles present events
func (o *OverlayManager) OnPresent() {
	if o.manager == nil {
		return
	}

	// NOTE: Present handling not implemented in wrapper approach yet
}

// ForwardMessage forwards a message
func (o *OverlayManager) ForwardMessage(message unsafe.Pointer) {
	if o.manager == nil {
		return
	}

	// NOTE: Message forwarding not implemented in wrapper approach yet
}

// KeyEvent handles key events
func (o *OverlayManager) KeyEvent(down bool, keyCode string, variant KeyVariant) {
	if o.manager == nil {
		return
	}

	// NOTE: C.CString/C.free not needed, use Go string or []byte if needed by wrapper
	// NOTE: Key event handling not implemented in wrapper approach yet
}

// CharEvent handles character events
func (o *OverlayManager) CharEvent(character string) {
	if o.manager == nil {
		return
	}

	// NOTE: C.CString/C.free not needed, use Go string or []byte if needed by wrapper
	// NOTE: Char event handling not implemented in wrapper approach yet
}

// MouseButtonEvent handles mouse button events
func (o *OverlayManager) MouseButtonEvent(down uint8, clickCount int32, which MouseButton, x, y int32) {
	if o.manager == nil {
		return
	}

	// NOTE: Mouse button event handling not implemented in wrapper approach yet
}

// MouseMotionEvent handles mouse motion events
func (o *OverlayManager) MouseMotionEvent(x, y int32) {
	if o.manager == nil {
		return
	}

	// NOTE: Mouse motion event handling not implemented in wrapper approach yet
}

// IsPointInsideClickZone checks if a point is inside the click zone
func (o *OverlayManager) IsPointInsideClickZone(x, y int32) bool {
	if o.manager == nil {
		return false
	}

	// NOTE: Click zone checking not implemented in wrapper approach yet
	return false
}
