package main

import (
	"fmt"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Client ID Finder ===")
	fmt.Println("This will test common Client ID patterns to find your application.")

	// Common Client ID patterns to test
	testClientIDs := []int64{
		1311711649018941501, // Current one
		1234567890123456789, // Example pattern
		987654321098765432,  // Another pattern
	}

	fmt.Println("\nTesting Client IDs...")

	for i, clientID := range testClientIDs {
		fmt.Printf("\n--- Test %d: Client ID %d ---\n", i+1, clientID)

		core, err := discord.Create(clientID, discord.CreateFlagsDefault, nil)
		if err != discord.ResultOk {
			fmt.Printf("✗ Failed to create Discord core: %v\n", err)
			continue
		}

		// Start the callback loop for robust event processing
		core.Start()

		fmt.Println("✓ Discord SDK initialized successfully")

		// Wait for user info to become available (robust pattern)
		user, result := core.WaitForUser(3 * time.Second)
		if result == discord.ResultOk {
			fmt.Printf("✓ Connected as user: %s\n", user.Username)
		} else {
			fmt.Printf("⚠ User connection failed: %v\n", result)
		}

		// Test application manager
		appManager := core.GetApplicationManager()
		if appManager != nil {
			locale := appManager.GetCurrentLocale()
			branch := appManager.GetCurrentBranch()
			fmt.Printf("✓ Application manager working (locale: %s, branch: %s)\n", locale, branch)
		}

		// Test activity manager
		activityManager := core.GetActivityManager()
		if activityManager != nil {
			fmt.Println("✓ Activity manager working")
		}

		// Test user manager
		userManager := core.GetUserManager()
		if userManager != nil {
			_, result := userManager.GetCurrentUser()
			if result == discord.ResultOk {
				fmt.Println("✓ User manager working")
			} else {
				fmt.Printf("⚠ User manager available but user retrieval failed: %v\n", result)
			}
		}

		core.Shutdown()
		fmt.Printf("✓ Client ID %d appears to be valid!\n", clientID)
	}

	fmt.Println("\n=== Instructions ===")
	fmt.Println("1. Look for the Client ID that shows '✓ Client ID X appears to be valid!'")
	fmt.Println("2. If none work, you need to:")
	fmt.Println("   - Go to https://discord.com/developers/applications")
	fmt.Println("   - Select your application")
	fmt.Println("   - Copy the 'Application ID' (this is your Client ID)")
	fmt.Println("   - Update all examples with this Client ID")

	fmt.Println("\n🎉 Client ID finder completed!")
}
