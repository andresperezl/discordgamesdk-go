package main

import (
	"fmt"

	core "github.com/andresperezl/discordctl/core"
)

func main() {
	fmt.Println("Testing Discord SDK initialization...")

	// Try to create a Discord core instance
	coreObj, result := core.Create(1311711649018941501, 0, nil)
	if result != core.ResultOk {
		fmt.Printf("Failed to create Discord core: %v\n", result)
		return
	}

	fmt.Println("Discord core created successfully!")

	// Try to run callbacks
	callbackResult := coreObj.RunCallbacks()
	fmt.Printf("RunCallbacks result: %v\n", callbackResult)

	// Clean up
	coreObj.Destroy()

	fmt.Println("Test completed successfully!")
}
