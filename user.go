package discord

import (
	"unsafe"

	dcgo "github.com/andresperezl/discordctl/discordcgo"
)

// UserManager provides access to user-related functionality
type UserManager struct {
	ptr unsafe.Pointer
}

// GetCurrentUser gets the current user
func (u *UserManager) GetCurrentUser() (*User, Result) {
	if u.ptr == nil {
		return nil, ResultInternalError
	}

	// Use the proper C DiscordUser struct with explicit alignment
	var cUser struct {
		id            int64
		username      [256]byte
		discriminator [8]byte
		avatar        [128]byte
		bot           bool
		_padding      [7]byte // Ensure proper alignment
	}

	// Zero out the struct to ensure clean state
	for i := range cUser.username {
		cUser.username[i] = 0
	}
	for i := range cUser.discriminator {
		cUser.discriminator[i] = 0
	}
	for i := range cUser.avatar {
		cUser.avatar[i] = 0
	}

	result := dcgo.UserManagerGetCurrentUser(u.ptr, unsafe.Pointer(&cUser))

	if result != int32(ResultOk) {
		return nil, Result(result)
	}

	return userFromStruct(&cUser), ResultOk
}

// GetUser gets a user by ID
func (u *UserManager) GetUser(userID int64, callback func(result Result, user *User)) {
	if u.ptr == nil {
		if callback != nil {
			callback(ResultInternalError, nil)
		}
		return
	}

	// Call the C wrapper function
	dcgo.UserManagerGetUser(u.ptr, userID, nil, nil)

	// For now, call the callback immediately since we don't have proper callback support
	if callback != nil {
		callback(ResultOk, nil)
	}
}

// GetCurrentUserPremiumType gets the current user's premium type
func (u *UserManager) GetCurrentUserPremiumType() (PremiumType, Result) {
	if u.ptr == nil {
		return PremiumTypeNone, ResultInternalError
	}

	var premiumType int32
	result := dcgo.UserManagerGetCurrentUserPremiumType(u.ptr, unsafe.Pointer(&premiumType))

	if result != int32(ResultOk) {
		return PremiumTypeNone, Result(result)
	}

	return PremiumType(premiumType), ResultOk
}

// CurrentUserHasFlag checks if the current user has a specific flag
func (u *UserManager) CurrentUserHasFlag(flag UserFlag) (bool, Result) {
	if u.ptr == nil {
		return false, ResultInternalError
	}

	var hasFlag bool
	result := dcgo.UserManagerCurrentUserHasFlag(u.ptr, int32(flag), unsafe.Pointer(&hasFlag))

	if result != int32(ResultOk) {
		return false, Result(result)
	}

	return hasFlag, ResultOk
}

// Helper function to convert struct to User
func userFromStruct(cUser *struct {
	id            int64
	username      [256]byte
	discriminator [8]byte
	avatar        [128]byte
	bot           bool
	_padding      [7]byte
}) *User {
	// Helper function to convert C string to Go string
	cStringToGoString := func(bytes []byte) string {
		// Find null terminator
		for i, b := range bytes {
			if b == 0 {
				return string(bytes[:i])
			}
		}
		return string(bytes)
	}

	return &User{
		ID:            cUser.id,
		Username:      cStringToGoString(cUser.username[:]),
		Discriminator: cStringToGoString(cUser.discriminator[:]),
		Avatar:        cStringToGoString(cUser.avatar[:]),
		Bot:           cUser.bot,
	}
}
