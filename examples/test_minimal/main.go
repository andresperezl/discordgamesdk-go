package main

import (
	"fmt"

	discord "github.com/andresperezl/discordctl"
)

func main() {
	fmt.Println("Testing Discord SDK initialization...")

	// Try to create a Discord core instance
	core, result := discord.Create(1311711649018941501, 0, nil)
	if result != discord.ResultOk {
		fmt.Printf("Failed to create Discord core: %v\n", result)
		return
	}

	fmt.Println("Discord core created successfully!")

	// Try to run callbacks
	callbackResult := core.RunCallbacks()
	fmt.Printf("RunCallbacks result: %v\n", callbackResult)

	// Clean up
	core.Destroy()

	fmt.Println("Test completed successfully!")
}
