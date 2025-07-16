package discord

import (
	"context"
	"log"
	"time"
)

// ExampleOverlayClient_OpenVoiceSettingsWithContext demonstrates how to use OpenVoiceSettingsWithContext with a timeout.
// This example is for documentation only and requires a real, initialized OverlayClient.
func ExampleOverlayClient_OpenVoiceSettingsWithContext() {
	var overlayClient *OverlayClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := overlayClient.OpenVoiceSettingsWithContext(ctx)
	if err != nil {
		log.Fatalf("failed to open voice settings: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleOverlayClient_SetLockedWithContext demonstrates how to use SetLockedWithContext with a timeout.
// This example is for documentation only and requires a real, initialized OverlayClient.
func ExampleOverlayClient_SetLockedWithContext() {
	var overlayClient *OverlayClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := overlayClient.SetLockedWithContext(ctx, true)
	if err != nil {
		log.Fatalf("failed to set locked: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleOverlayClient_OpenActivityInviteWithContext demonstrates how to use OpenActivityInviteWithContext with a timeout.
// This example is for documentation only and requires a real, initialized OverlayClient.
func ExampleOverlayClient_OpenActivityInviteWithContext() {
	var overlayClient *OverlayClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := overlayClient.OpenActivityInviteWithContext(ctx, 1) // 1 = core.ActivityActionTypeJoin
	if err != nil {
		log.Fatalf("failed to open activity invite: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleOverlayClient_OpenGuildInviteWithContext demonstrates how to use OpenGuildInviteWithContext with a timeout.
// This example is for documentation only and requires a real, initialized OverlayClient.
func ExampleOverlayClient_OpenGuildInviteWithContext() {
	var overlayClient *OverlayClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := overlayClient.OpenGuildInviteWithContext(ctx, "inviteCode")
	if err != nil {
		log.Fatalf("failed to open guild invite: %v", err)
	}
	// No Output: (documentation only)
}
