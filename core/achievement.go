package core

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// SetUserAchievement sets a user achievement
func (a *AchievementManager) SetUserAchievement(achievementID int64, percentComplete uint8) Result {
	if a.manager == nil {
		return ResultInternalError
	}
	dcgo.AchievementManagerSetUserAchievement(a.manager, achievementID, percentComplete, nil, nil)
	return ResultOk // TODO: callback support
}

// GetUserAchievement gets a user achievement
func (a *AchievementManager) GetUserAchievement(userAchievementID int64) (*UserAchievement, Result) {
	if a.manager == nil {
		return nil, ResultInternalError
	}
	var cAch struct {
		UserID          int64
		AchievementID   int64
		PercentComplete uint8
		UnlockedAt      [64]byte
	}
	res := dcgo.AchievementManagerGetUserAchievement(a.manager, userAchievementID, unsafe.Pointer(&cAch))
	if res != 0 {
		return nil, Result(res)
	}
	return &UserAchievement{
		UserID:          cAch.UserID,
		AchievementID:   cAch.AchievementID,
		PercentComplete: cAch.PercentComplete,
		UnlockedAt:      string(cAch.UnlockedAt[:]),
	}, ResultOk
}

// GetUserAchievementAt gets a user achievement at index
func (a *AchievementManager) GetUserAchievementAt(index int32) (*UserAchievement, Result) {
	if a.manager == nil {
		return nil, ResultInternalError
	}
	var cAch struct {
		UserID          int64
		AchievementID   int64
		PercentComplete uint8
		UnlockedAt      [64]byte
	}
	res := dcgo.AchievementManagerGetUserAchievementAt(a.manager, index, unsafe.Pointer(&cAch))
	if res != 0 {
		return nil, Result(res)
	}
	return &UserAchievement{
		UserID:          cAch.UserID,
		AchievementID:   cAch.AchievementID,
		PercentComplete: cAch.PercentComplete,
		UnlockedAt:      string(cAch.UnlockedAt[:]),
	}, ResultOk
}

// GetUserAchievementCount gets the number of user achievements
func (a *AchievementManager) GetUserAchievementCount() (int32, Result) {
	if a.manager == nil {
		return 0, ResultInternalError
	}
	var count int32
	dcgo.AchievementManagerCountUserAchievements(a.manager, unsafe.Pointer(&count))
	return count, ResultOk
}

// FetchUserAchievements fetches user achievements asynchronously
func (a *AchievementManager) FetchUserAchievements(callbackData unsafe.Pointer, callback unsafe.Pointer) {
	if a.manager == nil {
		return
	}
	dcgo.AchievementManagerFetchUserAchievements(a.manager, callbackData, callback)
}
