# Discord SDK Download Scripts

This directory contains helper scripts to automatically download and extract the Discord Game SDK files to the appropriate locations in the project.

## Available Scripts

### 1. PowerShell Script (Windows)
**File:** `download_sdk.ps1`

**Usage:**
```powershell
# Run from the project root directory
.\scripts\download_sdk.ps1

# Or with custom parameters
.\scripts\download_sdk.ps1 -SDKVersion "3.2.1" -DownloadUrl "https://dl-game-sdk.discordapp.net/3.2.1/discord_game_sdk.zip"
```

**Requirements:**
- PowerShell 5.1 or later
- Internet connection
- Write permissions to the project directory

### 2. Bash Script (Linux/macOS)
**File:** `download_sdk.sh`

**Usage:**
```bash
# Make the script executable (first time only)
chmod +x scripts/download_sdk.sh

# Run from the project root directory
./scripts/download_sdk.sh

# Or run directly with bash
bash scripts/download_sdk.sh
```

**Requirements:**
- Bash shell
- `curl` or `wget` for downloading
- `unzip` for extracting
- Internet connection
- Write permissions to the project directory

### 3. Batch Script (Windows)
**File:** `download_sdk.bat`

**Usage:**
```cmd
# Run from the project root directory
scripts\download_sdk.bat
```

**Requirements:**
- Windows Command Prompt
- PowerShell available (for download functionality)
- Internet connection
- Write permissions to the project directory

### 4. Go Script (Cross-platform)
**File:** `download_sdk.go`

**Usage:**
```bash
# Run directly with Go
go run scripts/download_sdk.go

# Or build and run
go build -o download_sdk scripts/download_sdk.go
./download_sdk
```

**Requirements:**
- Go 1.16 or later
- Internet connection
- Write permissions to the project directory

## What the Scripts Do

1. **Detect system architecture:** The scripts automatically detect your system architecture (x86_64, x86, or aarch64) to copy the appropriate shared libraries.

2. **Check for existing files:** The scripts first check if the required Discord SDK files are already present in the `lib/` directory.

3. **Download SDK:** If files are missing, the scripts download the Discord Game SDK from the official Discord CDN.

4. **Extract files:** The downloaded ZIP file is extracted to a temporary directory.

5. **Copy to correct locations:** Files are copied from the extracted archive to the appropriate locations based on your system architecture:
   - **x86_64 systems:**
     - `lib/discord_game_sdk.dll` (Windows)
     - `lib/discord_game_sdk.dll.lib` (Windows)
     - `lib/discord_game_sdk.so` (Linux)
   - **x86 systems:**
     - `lib/discord_game_sdk.dll` (Windows)
     - `lib/discord_game_sdk.dll.lib` (Windows)
   - **aarch64 systems:**
     - `lib/discord_game_sdk.dylib` (macOS)
   - **All architectures:**
     - `lib/discord_game_sdk.h` (C header file)

6. **Clean up:** Temporary files are removed after extraction.

7. **Verify:** The scripts verify that all required files are present after the download.

## Required Files by Architecture

The scripts check for and download these files based on your system architecture:

### x86_64 (64-bit)
- `lib/discord_game_sdk.dll` - Windows dynamic library
- `lib/discord_game_sdk.dll.lib` - Windows import library
- `lib/discord_game_sdk.so` - Linux shared library
- `lib/discord_game_sdk.h` - C header file

### x86 (32-bit)
- `lib/discord_game_sdk.dll` - Windows dynamic library
- `lib/discord_game_sdk.dll.lib` - Windows import library
- `lib/discord_game_sdk.h` - C header file

### aarch64 (ARM64)
- `lib/discord_game_sdk.dylib` - macOS dynamic library
- `lib/discord_game_sdk.h` - C header file

## Configuration

You can customize the scripts by modifying these variables:

- `SDK_VERSION` - Discord SDK version (default: "3.2.1")
- `DOWNLOAD_URL` - Download URL (default: Discord CDN URL)
- `LIB_DIR` - Target directory for SDK files (default: "lib")

## Error Handling

The scripts include comprehensive error handling:

- Network connectivity issues
- File permission problems
- Missing required tools (curl, wget, unzip)
- Corrupted downloads
- Extraction failures
- Architecture detection failures

## Output

The scripts provide colored output to indicate:
- 🟢 Green: Success messages
- 🟡 Yellow: Information and progress
- 🔴 Red: Errors and warnings
- 🔵 Cyan: Headers and titles
- ⚪ White: File listings

## Integration

These scripts can be integrated into your build process:

```bash
# Example: Run SDK download before building
./scripts/download_sdk.sh && go build
```

```powershell
# Example: PowerShell integration
.\scripts\download_sdk.ps1; if ($LASTEXITCODE -eq 0) { go build }
```

```cmd
# Example: Batch integration
scripts\download_sdk.bat && go build
```

## Troubleshooting

### Common Issues

1. **Permission denied:** Make sure you have write permissions to the project directory
2. **Network error:** Check your internet connection and firewall settings
3. **Missing tools:** Install required tools (curl/wget, unzip) on your system
4. **Antivirus blocking:** Some antivirus software may block the download - add the project directory to exclusions
5. **Architecture detection failure:** The scripts default to x86_64 if architecture detection fails

### Manual Download

If the scripts fail, you can manually download the SDK:
1. Visit: https://dl-game-sdk.discordapp.net/3.2.1/discord_game_sdk.zip
2. Extract the ZIP file
3. Copy the appropriate files to the `lib/` directory based on your architecture:
   - For x86_64: Copy files from `lib/x86_64/` and `c/discord_game_sdk.h`
   - For x86: Copy files from `lib/x86/` and `c/discord_game_sdk.h`
   - For aarch64: Copy files from `lib/aarch64/` and `c/discord_game_sdk.h`

## Architecture Detection

The scripts automatically detect your system architecture:

- **Windows:** Uses `PROCESSOR_ARCHITECTURE` environment variable
- **Unix-like systems:** Uses `uname -m` command
- **Fallback:** Defaults to x86_64 if detection fails

This ensures that only the appropriate shared libraries for your system are copied, reducing disk space usage and avoiding compatibility issues. 