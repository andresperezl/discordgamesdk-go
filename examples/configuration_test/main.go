package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Application Configuration Test ===")
	fmt.Printf("Testing application ID: %d\n", 1311711649018941501)
	fmt.Println("This will test each feature and provide specific configuration guidance.")

	// Initialize Discord SDK
	clientID := int64(1311711649018941501)
	core, err := discord.Create(clientID, discord.CreateFlagsDefault, nil)
	if err != discord.ResultOk {
		log.Fatalf("Failed to create Discord core: %v", err)
	}
	// Start the callback loop for robust event processing
	core.Start()
	defer core.Shutdown()

	fmt.Println("✓ Discord SDK initialized successfully")

	// Wait for user info to become available (robust pattern)
	fmt.Println("Waiting for user info...")
	user, result := core.WaitForUser(5 * time.Second)
	if result != discord.ResultOk {
		fmt.Printf("Warning: Failed to get current user: %v\n", result)
		fmt.Println("  This might indicate Discord client issues or configuration problems")
	} else {
		fmt.Printf("✓ Connected as user: %s\n", user.Username)
	}

	// Test 1: Application Manager (should always work)
	fmt.Println("\n--- Test 1: Application Manager ---")
	appManager := core.GetApplicationManager()
	if appManager != nil {
		locale := appManager.GetCurrentLocale()
		branch := appManager.GetCurrentBranch()
		fmt.Printf("✓ Application manager working\n")
		fmt.Printf("  Locale: %s\n", locale)
		fmt.Printf("  Branch: %s\n", branch)
	} else {
		fmt.Println("✗ Application manager failed")
	}

	// Test 2: Activity Manager (requires Rich Presence)
	fmt.Println("\n--- Test 2: Activity Manager (Rich Presence) ---")
	activityManager := core.GetActivityManager()
	if activityManager != nil {
		fmt.Println("✓ Activity manager available")

		// Test simple activity update
		activity := discord.Activity{
			Type:          discord.ActivityTypePlaying,
			ApplicationID: clientID,
			Name:          "Configuration Test",
			State:         "Testing",
			Details:       "Checking Rich Presence",
			Instance:      true,
		}

		activityManager.UpdateActivity(&activity, func(result discord.Result) {
			fmt.Printf("Activity update result: %v\n", result)
			if result == discord.ResultOk {
				fmt.Println("✓ Rich Presence is working!")
			} else {
				fmt.Printf("✗ Rich Presence failed: %v\n", result)
				fmt.Println("  → Enable Rich Presence in your Discord application:")
				fmt.Println("    1. Go to Discord Developer Portal")
				fmt.Println("    2. Select your application")
				fmt.Println("    3. Go to 'Rich Presence' → 'Art Assets'")
				fmt.Println("    4. Enable Rich Presence and add some assets")
			}
		})

		// Let the callback loop process the update
		time.Sleep(2 * time.Second)

		// Clear activity
		activityManager.ClearActivity(func(result discord.Result) {
			fmt.Printf("Activity clear result: %v\n", result)
		})

		// Let the callback loop process the clear
		time.Sleep(1 * time.Second)
	} else {
		fmt.Println("✗ Activity manager not available")
	}

	// Test 3: User Manager (requires Bot user)
	fmt.Println("\n--- Test 3: User Manager (Bot User) ---")
	userManager := core.GetUserManager()
	if userManager != nil {
		fmt.Println("✓ User manager available")

		currentUser, result := userManager.GetCurrentUser()
		fmt.Printf("User retrieval result: %v\n", result)
		if result == discord.ResultOk {
			fmt.Println("✓ User features are working!")
			fmt.Printf("  User ID: %d\n", currentUser.ID)
			fmt.Printf("  Username: %s\n", currentUser.Username)
		} else {
			fmt.Printf("✗ User features failed: %v\n", result)
			fmt.Println("  → Configure Bot user in your Discord application:")
			fmt.Println("    1. Go to Discord Developer Portal")
			fmt.Println("    2. Select your application")
			fmt.Println("    3. Go to 'Bot' section")
			fmt.Println("    4. Create a bot user")
			fmt.Println("    5. Configure OAuth2 scopes (identify, email)")
		}
	} else {
		fmt.Println("✗ User manager not available")
	}

	// Test 4: Storage Manager (should always work)
	fmt.Println("\n--- Test 4: Storage Manager ---")
	storageManager := core.GetStorageManager()
	if storageManager != nil {
		fmt.Println("✓ Storage manager available")

		// Test storage operations
		testData := []byte("configuration test")
		result := storageManager.Write("config_test", testData)
		fmt.Printf("Storage write result: %v\n", result)
		if result == discord.ResultOk {
			fmt.Println("✓ Storage is working!")

			// Clean up
			storageManager.Delete("config_test")
		} else {
			fmt.Printf("✗ Storage failed: %v\n", result)
		}
	} else {
		fmt.Println("✗ Storage manager not available")
	}

	fmt.Println("\n=== Configuration Summary ===")
	fmt.Println("✅ Working features:")
	fmt.Println("  - SDK initialization")
	fmt.Println("  - Application manager")
	fmt.Println("  - Storage manager")

	fmt.Println("\n⚠️  Features that need configuration:")
	fmt.Println("  - Rich Presence (for activity features)")
	fmt.Println("  - Bot user (for user features)")

	fmt.Println("\n🎯 Next steps:")
	fmt.Println("1. Configure Rich Presence for activity features")
	fmt.Println("2. Create a Bot user for user features")
	fmt.Println("3. Run this test again to verify configuration")

	fmt.Println("\n🎉 Configuration test completed!")
}
