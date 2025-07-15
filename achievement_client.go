package discord

import (
	"fmt"
	"unsafe"

	core "github.com/andresperezl/discordctl/core"
)

// AchievementClient provides Go-like interfaces for achievement management
type AchievementClient struct {
	manager *core.AchievementManager
	core    *core.Core
}

// SetUserAchievement sets a user achievement
func (ac *AchievementClient) SetUserAchievement(achievementID int64, percentComplete uint8) error {
	if ac.manager == nil {
		return fmt.Errorf("achievement manager not available")
	}
	res := ac.manager.SetUserAchievement(achievementID, percentComplete)
	if res != core.ResultOk {
		return fmt.Errorf("failed to set user achievement: %v", res)
	}
	return nil
}

// GetUserAchievement gets a user achievement
func (ac *AchievementClient) GetUserAchievement(achievementID int64) (*core.UserAchievement, error) {
	if ac.manager == nil {
		return nil, fmt.Errorf("achievement manager not available")
	}
	ach, res := ac.manager.GetUserAchievement(achievementID)
	if res != core.ResultOk {
		return nil, fmt.Errorf("failed to get user achievement: %v", res)
	}
	return ach, nil
}

// GetUserAchievementAt gets a user achievement at index
func (ac *AchievementClient) GetUserAchievementAt(index int32) (*core.UserAchievement, error) {
	if ac.manager == nil {
		return nil, fmt.Errorf("achievement manager not available")
	}
	ach, res := ac.manager.GetUserAchievementAt(index)
	if res != core.ResultOk {
		return nil, fmt.Errorf("failed to get user achievement at index: %v", res)
	}
	return ach, nil
}

// GetUserAchievementCount gets the user achievement count
func (ac *AchievementClient) GetUserAchievementCount() (int32, error) {
	if ac.manager == nil {
		return 0, fmt.Errorf("achievement manager not available")
	}
	count, res := ac.manager.GetUserAchievementCount()
	if res != core.ResultOk {
		return 0, fmt.Errorf("failed to get user achievement count: %v", res)
	}
	return count, nil
}

// FetchUserAchievements fetches user achievements asynchronously
func (ac *AchievementClient) FetchUserAchievements(callbackData, callback unsafe.Pointer) {
	if ac.manager == nil {
		return
	}
	ac.manager.FetchUserAchievements(callbackData, callback)
}
