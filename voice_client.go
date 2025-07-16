package discord

import (
	"fmt"

	core "github.com/andresperezl/discordgamesdk-go/core"
)

// VoiceClient provides Go-like interfaces for voice management
type VoiceClient struct {
	manager *core.VoiceManager
	core    *core.Core
}

// SetInputMode sets the input mode
func (vc *VoiceClient) SetInputMode(mode core.InputMode) error {
	if vc.manager == nil {
		return fmt.Errorf("voice manager not available")
	}
	res := vc.manager.SetInputMode(mode)
	if res != core.ResultOk {
		return fmt.Errorf("failed to set input mode: %v", res)
	}
	return nil
}

// GetInputMode gets the input mode
func (vc *VoiceClient) GetInputMode() (core.InputMode, error) {
	if vc.manager == nil {
		return core.InputMode{}, fmt.Errorf("voice manager not available")
	}
	mode, res := vc.manager.GetInputMode()
	if res != core.ResultOk {
		return core.InputMode{}, fmt.Errorf("failed to get input mode: %v", res)
	}
	return mode, nil
}

// IsSelfMute checks if self-mute is enabled
func (vc *VoiceClient) IsSelfMute() (bool, error) {
	if vc.manager == nil {
		return false, fmt.Errorf("voice manager not available")
	}
	mute, res := vc.manager.IsSelfMute()
	if res != core.ResultOk {
		return false, fmt.Errorf("failed to get self mute: %v", res)
	}
	return mute, nil
}

// SetSelfMute sets self-mute
func (vc *VoiceClient) SetSelfMute(mute bool) error {
	if vc.manager == nil {
		return fmt.Errorf("voice manager not available")
	}
	res := vc.manager.SetSelfMute(mute)
	if res != core.ResultOk {
		return fmt.Errorf("failed to set self mute: %v", res)
	}
	return nil
}

// IsSelfDeaf checks if self-deaf is enabled
func (vc *VoiceClient) IsSelfDeaf() (bool, error) {
	if vc.manager == nil {
		return false, fmt.Errorf("voice manager not available")
	}
	deaf, res := vc.manager.IsSelfDeaf()
	if res != core.ResultOk {
		return false, fmt.Errorf("failed to get self deaf: %v", res)
	}
	return deaf, nil
}

// SetSelfDeaf sets self-deaf
func (vc *VoiceClient) SetSelfDeaf(deaf bool) error {
	if vc.manager == nil {
		return fmt.Errorf("voice manager not available")
	}
	res := vc.manager.SetSelfDeaf(deaf)
	if res != core.ResultOk {
		return fmt.Errorf("failed to set self deaf: %v", res)
	}
	return nil
}

// IsLocalMute checks if a user is locally muted
func (vc *VoiceClient) IsLocalMute(userID int64) (bool, error) {
	if vc.manager == nil {
		return false, fmt.Errorf("voice manager not available")
	}
	mute, res := vc.manager.IsLocalMute(userID)
	if res != core.ResultOk {
		return false, fmt.Errorf("failed to get local mute: %v", res)
	}
	return mute, nil
}

// SetLocalMute sets local mute for a user
func (vc *VoiceClient) SetLocalMute(userID int64, mute bool) error {
	if vc.manager == nil {
		return fmt.Errorf("voice manager not available")
	}
	res := vc.manager.SetLocalMute(userID, mute)
	if res != core.ResultOk {
		return fmt.Errorf("failed to set local mute: %v", res)
	}
	return nil
}

// GetLocalVolume gets the local volume for a user
func (vc *VoiceClient) GetLocalVolume(userID int64) (uint8, error) {
	if vc.manager == nil {
		return 0, fmt.Errorf("voice manager not available")
	}
	volume, res := vc.manager.GetLocalVolume(userID)
	if res != core.ResultOk {
		return 0, fmt.Errorf("failed to get local volume: %v", res)
	}
	return volume, nil
}

// SetLocalVolume sets the local volume for a user
func (vc *VoiceClient) SetLocalVolume(userID int64, volume uint8) error {
	if vc.manager == nil {
		return fmt.Errorf("voice manager not available")
	}
	res := vc.manager.SetLocalVolume(userID, volume)
	if res != core.ResultOk {
		return fmt.Errorf("failed to set local volume: %v", res)
	}
	return nil
}
