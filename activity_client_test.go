package discord

import (
	"context"
	"fmt"
	"log"
	"time"

	core "github.com/andresperezl/discordctl/core"
)

// ExampleActivityClient_SetActivityWithContext demonstrates how to use SetActivityWithContext with a timeout.
// This example is for documentation only and requires a real, initialized ActivityClient.
func ExampleActivityClient_SetActivityWithContext() {
	var activityClient *ActivityClient // Assume this is properly initialized
	activity := &core.Activity{Name: "GoDoc Example"}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := activityClient.SetActivityWithContext(ctx, activity)
	if err != nil {
		log.Printf("failed to set activity: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleActivityClient_ActivityJoinRequests demonstrates how to use ActivityJoinRequests to receive join requests via a channel.
// This example is for documentation only and requires a real, initialized ActivityClient.
func ExampleActivityClient_ActivityJoinRequests() {
	var activityClient *ActivityClient // Assume this is properly initialized
	joinRequests := activityClient.ActivityJoinRequests()
	go func() {
		for user := range joinRequests {
			fmt.Printf("Received join request from: %s\n", user.Username)
		}
	}()
	// No Output: (documentation only)
}

// ExampleActivityClient_ClearActivityWithContext demonstrates how to use ClearActivityWithContext with a timeout.
// This example is for documentation only and requires a real, initialized ActivityClient.
func ExampleActivityClient_ClearActivityWithContext() {
	var activityClient *ActivityClient // Assume this is properly initialized

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := activityClient.ClearActivityWithContext(ctx)
	if err != nil {
		log.Printf("failed to clear activity: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleActivityClient_SendRequestReplyWithContext demonstrates how to use SendRequestReplyWithContext with a timeout.
// This example is for documentation only and requires a real, initialized ActivityClient.
func ExampleActivityClient_SendRequestReplyWithContext() {
	var activityClient *ActivityClient // Assume this is properly initialized
	userID := int64(123456789)
	reply := core.ActivityJoinRequestReplyYes

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := activityClient.SendRequestReplyWithContext(ctx, userID, reply)
	if err != nil {
		log.Printf("failed to send request reply: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleActivityClient_SendInviteWithContext demonstrates how to use SendInviteWithContext with a timeout.
// This example is for documentation only and requires a real, initialized ActivityClient.
func ExampleActivityClient_SendInviteWithContext() {
	var activityClient *ActivityClient // Assume this is properly initialized
	userID := int64(123456789)
	inviteType := core.ActivityActionTypeJoin
	content := "Let's play together!"

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := activityClient.SendInviteWithContext(ctx, userID, inviteType, content)
	if err != nil {
		log.Printf("failed to send invite: %v", err)
	}
	// No Output: (documentation only)
}

// ExampleActivityClient_AcceptInviteWithContext demonstrates how to use AcceptInviteWithContext with a timeout.
// This example is for documentation only and requires a real, initialized ActivityClient.
func ExampleActivityClient_AcceptInviteWithContext() {
	var activityClient *ActivityClient // Assume this is properly initialized
	userID := int64(123456789)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := activityClient.AcceptInviteWithContext(ctx, userID)
	if err != nil {
		log.Printf("failed to accept invite: %v", err)
	}
	// No Output: (documentation only)
}
