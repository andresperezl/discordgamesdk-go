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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerIsEnabled(o.manager, unsafe.Pointer(&enabled))
		return nil
	})
	return enabled
}

// IsLocked checks if the overlay is locked
func (o *OverlayManager) IsLocked() bool {
	if o.manager == nil {
		return false
	}

	var locked bool
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerIsLocked(o.manager, unsafe.Pointer(&locked))
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerSetLocked(o.manager, locked, nil, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerOpenActivityInvite(o.manager, int32(actionType), nil, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerOpenGuildInvite(o.manager, cCode, nil, nil)
		return nil
	})
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
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerOpenVoiceSettings(o.manager, nil, nil)
		return nil
	})
	if callback != nil {
		callback(ResultOk)
	}
}

// InitDrawingDXGI initializes drawing with DXGI
func (o *OverlayManager) InitDrawingDXGI(swapchain unsafe.Pointer, useMessageForwarding bool) Result {
	if o.manager == nil {
		return ResultInternalError
	}
	res := dcgo.RunOnDispatcherSync(func() int32 {
		return dcgo.OverlayManagerInitDrawingDXGI(o.manager, swapchain, useMessageForwarding)
	})
	return Result(res)
}

// OnPresent handles present events
func (o *OverlayManager) OnPresent() {
	if o.manager == nil {
		return
	}
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerOnPresent(o.manager)
		return nil
	})
}

// ForwardMessage forwards a message
func (o *OverlayManager) ForwardMessage(message unsafe.Pointer) {
	if o.manager == nil {
		return
	}
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerForwardMessage(o.manager, message)
		return nil
	})
}

// KeyEvent handles key events
func (o *OverlayManager) KeyEvent(down bool, keyCode string, variant KeyVariant) {
	if o.manager == nil {
		return
	}
	cKey := dcgo.GoStringToCChar(keyCode)
	defer dcgo.FreeCChar(cKey)
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerKeyEvent(o.manager, down, cKey, int32(variant))
		return nil
	})
}

// CharEvent handles character events
func (o *OverlayManager) CharEvent(character string) {
	if o.manager == nil {
		return
	}
	cChar := dcgo.GoStringToCChar(character)
	defer dcgo.FreeCChar(cChar)
	dcgo.RunOnDispatcherSync(func() any {
		dcgo.OverlayManagerCharEvent(o.manager, cChar)
		return nil
	})
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
