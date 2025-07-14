package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Diagnostic Example ===")
	fmt.Println("This example will test each operation step-by-step to identify issues.")

	// Initialize Discord SDK
	clientID := int64(1311711649018941501) // Replace with your actual client ID
	fmt.Printf("Initializing Discord SDK with client ID: %d\n", clientID)

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

	// Test 1: Application Manager
	fmt.Println("\n--- Testing Application Manager ---")
	appManager := core.GetApplicationManager()
	if appManager == nil {
		fmt.Println("✗ Failed to get application manager")
	} else {
		fmt.Println("✓ Application manager retrieved")

		// Test locale and branch retrieval
		locale := appManager.GetCurrentLocale()
		branch := appManager.GetCurrentBranch()
		fmt.Printf("Current locale: %s\n", locale)
		fmt.Printf("Current branch: %s\n", branch)
	}

	// Test 2: User Manager
	fmt.Println("\n--- Testing User Manager ---")
	userManager := core.GetUserManager()
	if userManager == nil {
		fmt.Println("✗ Failed to get user manager")
	} else {
		fmt.Println("✓ User manager retrieved")

		// Test current user retrieval
		currentUser, result := userManager.GetCurrentUser()
		fmt.Printf("Current user retrieval result: %v\n", result)
		if result == discord.ResultOk {
			fmt.Println("✓ Current user retrieved successfully")
			fmt.Printf("  User ID: %d\n", currentUser.ID)
			fmt.Printf("  Username: %s\n", currentUser.Username)
		} else {
			fmt.Printf("✗ Current user retrieval failed: %v\n", result)
			fmt.Println("  This might indicate:")
			fmt.Println("  - User not properly authenticated")
			fmt.Println("  - Discord application not configured for user features")
		}
	}

	// Test 3: Activity Manager
	fmt.Println("\n--- Testing Activity Manager ---")
	activityManager := core.GetActivityManager()
	if activityManager == nil {
		fmt.Println("✗ Failed to get activity manager")
	} else {
		fmt.Println("✓ Activity manager retrieved")

		// Test activity registration
		result := activityManager.RegisterCommand("discordctl-diagnostic")
		fmt.Printf("Activity registration result: %v\n", result)
		if result == discord.ResultOk {
			fmt.Println("✓ Activity manager registered successfully")
		} else {
			fmt.Printf("✗ Activity manager registration failed: %v\n", result)
		}
	}

	// Test 4: Storage Manager
	fmt.Println("\n--- Testing Storage Manager ---")
	storageManager := core.GetStorageManager()
	if storageManager == nil {
		fmt.Println("✗ Failed to get storage manager")
	} else {
		fmt.Println("✓ Storage manager retrieved")

		// Test storage write
		testData := []byte("diagnostic test data")
		result := storageManager.Write("diagnostic_test", testData)
		fmt.Printf("Storage write result: %v\n", result)
		if result == discord.ResultOk {
			fmt.Println("✓ Storage write successful")

			// Test storage read
			readData := make([]byte, len(testData))
			bytesRead, result := storageManager.Read("diagnostic_test", readData)
			fmt.Printf("Storage read result: %v (bytes read: %d)\n", result, bytesRead)
			if result == discord.ResultOk {
				fmt.Println("✓ Storage read successful")
			} else {
				fmt.Printf("✗ Storage read failed: %v\n", result)
			}

			// Test storage delete
			result = storageManager.Delete("diagnostic_test")
			fmt.Printf("Storage delete result: %v\n", result)
			if result == discord.ResultOk {
				fmt.Println("✓ Storage delete successful")
			} else {
				fmt.Printf("✗ Storage delete failed: %v\n", result)
			}
		} else {
			fmt.Printf("✗ Storage write failed: %v\n", result)
		}
	}

	// Let the callback loop process any pending operations
	fmt.Println("\n--- Processing Callbacks ---")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Diagnostic Summary ===")
	fmt.Println("If you see InternalError (4) or NotFound (9) errors:")
	fmt.Println("1. Make sure Discord client is running and you're logged in")
	fmt.Println("2. Verify your Discord application is properly configured:")
	fmt.Println("   - Rich Presence enabled")
	fmt.Println("   - Bot user created (for user features)")
	fmt.Println("   - OAuth2 configured (for user features)")
	fmt.Println("3. Check that your Client ID is correct")
	fmt.Println("4. Ensure your Discord application has the necessary permissions")

	fmt.Println("\n🎉 Diagnostic completed!")
}
