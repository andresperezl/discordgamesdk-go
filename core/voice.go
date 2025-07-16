package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordgamesdk-go/discordcgo"
)

// SetInputMode sets the input mode
func (v *VoiceManager) SetInputMode(mode InputMode) Result {
	if v.manager == nil {
		return ResultInternalError
	}
	var cMode struct {
		Type     int32
		Shortcut [256]byte
	}
	cMode.Type = int32(mode.Type)
	copy(cMode.Shortcut[:], mode.Shortcut)
	dcgo.VoiceManagerSetInputMode(v.manager, unsafe.Pointer(&cMode), nil, nil)
	return ResultOk // TODO: callback support
}

// GetInputMode gets the input mode
func (v *VoiceManager) GetInputMode() (InputMode, Result) {
	if v.manager == nil {
		return InputMode{}, ResultInternalError
	}
	var cMode struct {
		Type     int32
		Shortcut [256]byte
	}
	res := dcgo.VoiceManagerGetInputMode(v.manager, unsafe.Pointer(&cMode))
	if res != 0 {
		return InputMode{}, Result(res)
	}
	return InputMode{
		Type:     InputModeType(cMode.Type),
		Shortcut: string(cMode.Shortcut[:]),
	}, ResultOk
}

// IsSelfMute checks if self-mute is enabled
func (v *VoiceManager) IsSelfMute() (bool, Result) {
	if v.manager == nil {
		return false, ResultInternalError
	}
	var mute bool
	res := dcgo.VoiceManagerIsSelfMute(v.manager, unsafe.Pointer(&mute))
	return mute, Result(res)
}

// SetSelfMute sets self-mute
func (v *VoiceManager) SetSelfMute(mute bool) Result {
	if v.manager == nil {
		return ResultInternalError
	}
	res := dcgo.VoiceManagerSetSelfMute(v.manager, mute)
	return Result(res)
}

// IsSelfDeaf checks if self-deaf is enabled
func (v *VoiceManager) IsSelfDeaf() (bool, Result) {
	if v.manager == nil {
		return false, ResultInternalError
	}
	var deaf bool
	res := dcgo.VoiceManagerIsSelfDeaf(v.manager, unsafe.Pointer(&deaf))
	return deaf, Result(res)
}

// SetSelfDeaf sets self-deaf
func (v *VoiceManager) SetSelfDeaf(deaf bool) Result {
	if v.manager == nil {
		return ResultInternalError
	}
	res := dcgo.VoiceManagerSetSelfDeaf(v.manager, deaf)
	return Result(res)
}

// IsLocalMute checks if a user is locally muted
func (v *VoiceManager) IsLocalMute(userID int64) (bool, Result) {
	if v.manager == nil {
		return false, ResultInternalError
	}
	var mute bool
	res := dcgo.VoiceManagerIsLocalMute(v.manager, userID, unsafe.Pointer(&mute))
	return mute, Result(res)
}

// SetLocalMute sets local mute for a user
func (v *VoiceManager) SetLocalMute(userID int64, mute bool) Result {
	if v.manager == nil {
		return ResultInternalError
	}
	res := dcgo.VoiceManagerSetLocalMute(v.manager, userID, mute)
	return Result(res)
}

// GetLocalVolume gets the local volume for a user
func (v *VoiceManager) GetLocalVolume(userID int64) (uint8, Result) {
	if v.manager == nil {
		return 0, ResultInternalError
	}
	var volume uint8
	res := dcgo.VoiceManagerGetLocalVolume(v.manager, userID, unsafe.Pointer(&volume))
	return volume, Result(res)
}

// SetLocalVolume sets the local volume for a user
func (v *VoiceManager) SetLocalVolume(userID int64, volume uint8) Result {
	if v.manager == nil {
		return ResultInternalError
	}
	res := dcgo.VoiceManagerSetLocalVolume(v.manager, userID, volume)
	return Result(res)
}
