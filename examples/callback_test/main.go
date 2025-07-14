package main

import (
	"fmt"
	"log"
	"time"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("=== Discord Game SDK Callback Test ===")
	fmt.Println("This example demonstrates enhanced callback handling and result tracking.")

	// Initialize Discord SDK
	clientID := int64(1311711649018941501)
	fmt.Printf("Initializing Discord SDK with client ID: %d\n", clientID)

	core, err := discord.Create(clientID, discord.CreateFlagsDefault, nil)
	if err != discord.ResultOk {
		log.Fatalf("Failed to create Discord core: %v", err)
	}

	// Start the callback loop
	fmt.Println("Starting callback loop...")
	core.Start()
	defer core.Shutdown()

	// Test 1: Wait for initialization
	fmt.Println("\n=== Test 1: Initialization Waiting ===")
	fmt.Println("Waiting for SDK initialization...")
	startTime := time.Now()
	if !core.WaitForInitialization(5 * time.Second) {
		log.Fatal("Failed to initialize Discord SDK within timeout")
	}
	initTime := time.Since(startTime)
	fmt.Printf("✓ SDK initialized successfully in %v\n", initTime)

	// Test 2: Wait for user with timeout
	fmt.Println("\n=== Test 2: User Connection Waiting ===")
	fmt.Println("Waiting for user connection...")
	userStartTime := time.Now()
	user, result := core.WaitForUser(5 * time.Second)
	if result != discord.ResultOk {
		log.Fatalf("Failed to get current user: %v", result)
	}
	userTime := time.Since(userStartTime)
	fmt.Printf("✓ Connected as user: %s (ID: %d) in %v\n", user.Username, user.ID, userTime)

	// Test 3: Get managers with core reference
	fmt.Println("\n=== Test 3: Manager Access with Core Reference ===")

	activityManager := core.GetActivityManager()
	if activityManager == nil {
		log.Fatal("Failed to get activity manager")
	}
	fmt.Println("✓ Activity manager retrieved with core reference")

	// Test 4: Callback result tracking
	fmt.Println("\n=== Test 4: Callback Result Tracking ===")

	// Generate a callback ID
	callbackID := core.GenerateCallbackID()
	fmt.Printf("Generated callback ID: %s\n", callbackID)

	// Add a test callback result
	core.AddCallbackResult(callbackID, discord.ResultOk, "test_data")

	// Retrieve the callback result
	if result, found := core.GetCallbackResult(callbackID); found {
		fmt.Printf("✓ Retrieved callback result: %v\n", result)
	} else {
		fmt.Println("✗ Failed to retrieve callback result")
	}

	// Test 5: Async activity operations
	fmt.Println("\n=== Test 5: Async Activity Operations ===")

	// Create a test activity
	activity := &discord.Activity{
		Type:          discord.ActivityTypePlaying,
		ApplicationID: clientID,
		Name:          "Discord Go SDK",
		State:         "Testing Callbacks",
		Details:       "Enhanced callback handling demo",
		Timestamps: discord.ActivityTimestamps{
			Start: time.Now().Unix(),
		},
		Assets: discord.ActivityAssets{
			LargeImage: "logo",
			LargeText:  "Discord Go SDK",
			SmallImage: "go",
			SmallText:  "Go Language",
		},
		Party: discord.ActivityParty{
			Size: discord.PartySize{
				CurrentSize: 1,
				MaxSize:     4,
			},
			Privacy: discord.ActivityPartyPrivacyPublic,
		},
		Secrets: discord.ActivitySecrets{
			Match:    "secret-match",
			Join:     "secret-join",
			Spectate: "secret-spectate",
		},
		Instance: true,
	}

	// Test async activity update
	fmt.Println("Testing async activity update...")
	updateStartTime := time.Now()
	updateChan := activityManager.UpdateActivityAsync(activity)

	select {
	case result := <-updateChan:
		updateTime := time.Since(updateStartTime)
		if result == discord.ResultOk {
			fmt.Printf("✓ Async activity update completed successfully in %v\n", updateTime)
		} else {
			fmt.Printf("⚠ Async activity update failed: %v in %v\n", result, updateTime)
		}
	case <-time.After(5 * time.Second):
		fmt.Println("⚠ Async activity update timed out")
	}

	// Test async activity clear
	fmt.Println("Testing async activity clear...")
	clearStartTime := time.Now()
	clearChan := activityManager.ClearActivityAsync()

	select {
	case result := <-clearChan:
		clearTime := time.Since(clearStartTime)
		if result == discord.ResultOk {
			fmt.Printf("✓ Async activity clear completed successfully in %v\n", clearTime)
		} else {
			fmt.Printf("⚠ Async activity clear failed: %v in %v\n", result, clearTime)
		}
	case <-time.After(5 * time.Second):
		fmt.Println("⚠ Async activity clear timed out")
	}

	// Test 6: Wait for callback result
	fmt.Println("\n=== Test 6: Wait for Callback Result ===")

	// Create a test callback ID
	testCallbackID := core.GenerateCallbackID()
	fmt.Printf("Testing wait for callback ID: %s\n", testCallbackID)

	// Simulate a callback result after a delay
	go func() {
		time.Sleep(1 * time.Second)
		core.AddCallbackResult(testCallbackID, discord.ResultOk, "delayed_data")
	}()

	// Wait for the callback result
	if result, found := core.WaitForCallbackResult(testCallbackID, 3*time.Second); found {
		fmt.Printf("✓ Successfully waited for callback result: %v\n", result)
	} else {
		fmt.Println("✗ Failed to wait for callback result")
	}

	// Test 7: Multiple concurrent operations
	fmt.Println("\n=== Test 7: Multiple Concurrent Operations ===")

	// Start multiple async operations
	operations := []chan discord.Result{}

	for i := 0; i < 3; i++ {
		opChan := activityManager.UpdateActivityAsync(&discord.Activity{
			Type:          discord.ActivityTypePlaying,
			ApplicationID: clientID,
			Name:          fmt.Sprintf("Concurrent Test %d", i+1),
			State:         "Testing concurrent operations",
			Details:       fmt.Sprintf("Operation %d", i+1),
			Timestamps: discord.ActivityTimestamps{
				Start: time.Now().Unix(),
			},
		})
		operations = append(operations, opChan)
	}

	// Wait for all operations to complete
	completed := 0
	for i, opChan := range operations {
		select {
		case result := <-opChan:
			if result == discord.ResultOk {
				fmt.Printf("✓ Concurrent operation %d completed successfully\n", i+1)
				completed++
			} else {
				fmt.Printf("⚠ Concurrent operation %d failed: %v\n", i+1, result)
			}
		case <-time.After(3 * time.Second):
			fmt.Printf("⚠ Concurrent operation %d timed out\n", i+1)
		}
	}

	fmt.Printf("Completed %d/%d concurrent operations\n", completed, len(operations))

	// Final cleanup
	fmt.Println("\n=== Final Cleanup ===")
	finalClearChan := activityManager.ClearActivityAsync()
	select {
	case result := <-finalClearChan:
		if result == discord.ResultOk {
			fmt.Println("✓ Final activity clear completed")
		} else {
			fmt.Printf("⚠ Final activity clear failed: %v\n", result)
		}
	case <-time.After(3 * time.Second):
		fmt.Println("⚠ Final activity clear timed out")
	}

	fmt.Println("\n🎉 Callback test completed successfully!")
	fmt.Println("\nThis demonstrates:")
	fmt.Println("- Proper SDK initialization waiting")
	fmt.Println("- User connection verification")
	fmt.Println("- Callback result tracking and retrieval")
	fmt.Println("- Async activity operations with timeouts")
	fmt.Println("- Concurrent operation handling")
	fmt.Println("- Enhanced error handling and reporting")
	fmt.Println("- Robust callback execution after initialization")
}
