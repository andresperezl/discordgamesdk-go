package discord

import (
	"context"
	"log"
	"time"
)

// ExampleUserClient_GetUserWithContext demonstrates how to use GetUserWithContext with a timeout.
// This example is for documentation only and requires a real, initialized UserClient.
func ExampleUserClient_GetUserWithContext() {
	var userClient *UserClient // Assume this is properly initialized
	var userID int64           // Assume this is properly set

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := userClient.GetUserWithContext(ctx, userID)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}
	log.Printf("Fetched user: %s", user.Username)
	// No Output: (documentation only)
}
