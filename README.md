# Discord Game SDK Go Wrapper

This project provides Go bindings for the Discord Game SDK using a wrapper approach to overcome CGO limitations with complex C APIs.

## Architecture

The bindings use a two-layer approach:

1. **C Wrapper Layer** (`discordcgo/` package):
   - `discord_wrappers.c` - C wrapper functions that call the Discord SDK
   - `discord_wrappers.h` - C header declarations
   - `bindings.go` - Go wrapper functions that call the C wrappers

2. **Go API Layer** (main package):
   - Clean Go API that calls the wrapper functions
   - No direct C imports in the main package
   - Type-safe Go interfaces

## Features

- ✅ **Core SDK initialization and management**
- ✅ **Application manager** - OAuth2 tokens, locale, branch info
- ✅ **User manager** - Current user info, premium status, flags
- ✅ **Activity manager** - Rich presence, invites, join requests
- ✅ **Lobby manager** - Lobby creation, management, networking
- ✅ **Network manager** - Peer-to-peer networking
- ✅ **Storage manager** - File storage operations
- ✅ **Overlay manager** - Discord overlay integration
- ✅ **Relationship manager** - User relationships
- ✅ **Image manager** - Avatar and image handling
- ✅ **Voice manager** - Voice settings and controls
- ✅ **Store manager** - Entitlements and purchases
- ✅ **Achievement manager** - User achievements

## Installation

### Option 1: Automatic SDK Download (Recommended)

Use the provided helper scripts to automatically download and extract the Discord Game SDK:

**Windows:**
```powershell
# PowerShell
.\scripts\download_sdk.ps1

# Or batch file
scripts\download_sdk.bat
```

**Linux/macOS:**
```bash
# Make executable (first time only)
chmod +x scripts/download_sdk.sh

# Run the script
./scripts/download_sdk.sh
```

**Cross-platform (Go):**
```bash
go run scripts/download_sdk.go
```

**Using Makefile (recommended):**
```bash
# Download SDK files
make sdk-download

# Check if SDK files are present
make sdk-check

# Quick start (download SDK and build examples)
make quickstart
```

### Option 2: Manual Installation

1. Download the Discord Game SDK from the [Discord Developer Portal](https://discord.com/developers/docs/game-sdk/sdk-starter-guide)
2. Place the SDK files in the `lib/` directory:
   - `discord_game_sdk.dll` (Windows x86_64)
   - `discord_game_sdk.dll.lib` (Windows x86_64)
   - `discord_game_sdk.so` (Linux x86_64)
   - `discord_game_sdk.h`

3. Install the Go package:
```bash
go get github.com/andresperezl/discordctl
```

## Usage

```go
package main

import (
    "log"
    "github.com/andresperezl/discordctl"
    core "github.com/andresperezl/discordctl/core"
)

func main() {
    // Initialize the Discord SDK
    clientID := int64(1234567890123456789) // Your Discord application ID
    coreObj, err := discord.Create(clientID, core.CreateFlagsDefault, nil)
    if err != core.ResultOk {
        log.Fatalf("Failed to create Discord core: %v", err)
    }
    coreObj.Start()
    defer coreObj.Shutdown()

    // Get managers
    appManager := core.GetApplicationManager()
    userManager := core.GetUserManager()
    activityManager := core.GetActivityManager()

    // Run the main loop
    for {
        result := core.RunCallbacks()
        if result != core.ResultOk {
            log.Printf("RunCallbacks returned: %v", result)
        }
        
        // Your game logic here
        // ...
    }
}
```

## Robust Event Processing and User Initialization

**Important:** The Discord Game SDK requires regular calls to `RunCallbacks()` to process events and update state. If you do not do this, user and manager queries may return default/zero values or errors, even if the Discord client is running.

### Best Practice: Automatic Callback Loop

This wrapper provides a robust pattern:

- **Start the callback loop:**
  ```go
  core.Start()
  defer core.Shutdown()
  ```
  This runs `RunCallbacks()` in a background goroutine for the lifetime of your SDK instance.

- **Wait for user info:**
  ```go
  user, result := core.WaitForUser(5 * time.Second)
  if result != core.ResultOk {
      log.Fatalf("Failed to get current user: %v", result)
  }
  ```
  This blocks until the Discord client provides user info, or times out.

- **Graceful shutdown:**
  Always call `core.Shutdown()` (or use `defer`) to stop the callback loop and clean up resources.

### Example Usage

```go
core, err := discord.Create(clientID, discord.CreateFlagsDefault, nil)
if err != discord.ResultOk {
    log.Fatalf("Failed to create Discord core: %v", err)
}
core.Start()
defer core.Shutdown()

user, result := core.WaitForUser(5 * time.Second)
if result != discord.ResultOk {
    log.Fatalf("Failed to get current user: %v", result)
}
fmt.Printf("Hello, %s!\n", user.Username)
```

### Why is this necessary?
- The Discord SDK is asynchronous and event-driven. It needs time to connect to the client and receive user data.
- Without regular callbacks, you will get NotFound or zero values for user and manager queries.
- This pattern ensures your app is always in sync with Discord state.

## Examples

See the `examples/` directory for working examples:

- `test_basic.go` - Basic SDK initialization and manager access

## Building

The project uses CGO to interface with the Discord Game SDK. Make sure you have:

- Go 1.18+ with CGO enabled
- GCC compiler (MinGW on Windows)
- Discord Game SDK files in `lib/` directory

```bash
# Build the main package
go build .

# Build an example
go build ./examples/test_basic.go

# Using Makefile (recommended)
make build
make examples
```

**Note for Windows users:** The build process automatically copies the required DLL files to the output directory. If you're building manually, make sure to copy `lib/discord_game_sdk.dll` to the same directory as your executable.

## Wrapper Approach Benefits

1. **Type Safety**: Go types instead of unsafe C pointers
2. **Error Handling**: Proper Go error types
3. **Memory Management**: Automatic garbage collection
4. **Cross-Platform**: Works on Windows, Linux, macOS
5. **Maintainability**: Clean separation between C and Go code

## Limitations

- Callback functions are currently stubbed out (return `ResultInternalError`)
- Some advanced features may need additional wrapper functions
- Requires the Discord Game SDK to be installed

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add wrapper functions in `discordcgo/` package
4. Update the main package to use the new wrappers
5. Add tests and examples
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 