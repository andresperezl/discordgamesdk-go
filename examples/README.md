# Discord Game SDK Go Examples

This directory contains various examples demonstrating how to use the Discord Game SDK Go bindings.

## Examples

### Basic Example (`basic/`)
Demonstrates basic SDK initialization and manager access:
- SDK initialization
- Manager retrieval (Application, User, Activity)
- Callback processing

### Activity Example (`activity/`)
Demonstrates rich presence functionality:
- Rich presence creation and updates
- Activity timestamps and assets
- Party information and secrets
- Activity clearing

### User Example (`user/`)
Demonstrates user information retrieval:
- Current user information
- Premium type checking
- User flag verification

### Storage Example (`storage/`)
Demonstrates local storage functionality:
- Writing data to local storage
- Reading data from local storage
- Checking if data exists
- Getting storage statistics
- Listing storage files
- Getting storage path

## Building the Examples

### Prerequisites

1. **Go 1.19 or later** - [Download from golang.org](https://golang.org/dl/)
2. **Discord client running** - Make sure Discord is running and you're logged in
3. **Valid Discord Application** - You need a Discord application with a valid Client ID

### Quick Build

#### Windows
```cmd
# Using batch file (recommended)
build.bat

# Or using PowerShell directly
powershell -ExecutionPolicy Bypass -File build.ps1

# Build in debug mode
build.bat Debug
```

#### Manual Build
```cmd
# Navigate to project root
cd ..

# Build a specific example
go build -o examples/bin/basic.exe examples/basic/main.go

# Copy DLL to output directory
copy lib\discord_game_sdk.dll examples\bin\
```

### Build Output

The build script creates a `bin/` directory containing:
- All compiled executables (`.exe` files)
- `discord_game_sdk.dll` (copied from `lib/`)

## Running the Examples

1. **Make sure Discord is running** and you're logged in
2. **Navigate to the `bin/` directory**
3. **Run any example executable**:
   ```cmd
   cd examples\bin
   basic.exe
   activity.exe
   user.exe
   storage.exe
   ```

## Configuration

### Client ID
Each example uses a default Client ID (`1311711649018941501`). To use your own:

1. Create a Discord application at [Discord Developer Portal](https://discord.com/developers/applications)
2. Get your Application ID (Client ID)
3. Replace the Client ID in the example files:
   ```go
   clientID := int64(YOUR_CLIENT_ID_HERE)
   ```

### Discord Application Setup
For full functionality, your Discord application should have:
- **Rich Presence** enabled
- **Bot** user created (for some features)
- **OAuth2** configured (for user features)

## Troubleshooting

### Common Issues

1. **"Failed to create Discord core: 4"**
   - Make sure Discord client is running and logged in
   - Verify your Client ID is correct
   - Check that your Discord application is properly configured

2. **"DLL not found"**
   - Ensure `discord_game_sdk.dll` is in the same directory as the executable
   - The build script should copy it automatically

3. **"Permission denied"**
   - Run as administrator if needed
   - Check Windows Defender/antivirus settings

4. **Build errors**
   - Ensure Go is properly installed and in PATH
   - Check that all dependencies are available
   - Verify the project structure is correct

### Debug Mode
Build in debug mode for more detailed error information:
```cmd
build.bat Debug
```

## Example Features

### Basic Example
- ✅ SDK initialization
- ✅ Manager access
- ✅ Callback processing

### Activity Example
- ✅ Rich presence creation
- ✅ Activity updates with timestamps
- ✅ Party information
- ✅ Activity secrets
- ✅ Activity clearing

### User Example
- ✅ Current user information
- ✅ Premium type checking
- ✅ User flag verification

### Storage Example
- ✅ Local data storage
- ✅ Data persistence
- ✅ Storage statistics
- ✅ File management

## Contributing

To add new examples:

1. Create a new directory under `examples/`
2. Add a `main.go` file
3. Follow the existing pattern for SDK initialization
4. Test with the build script

## License

These examples are part of the Discord Game SDK Go bindings project. 