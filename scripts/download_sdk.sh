#!/bin/bash

# Discord SDK Download Script for Unix-like systems
# Downloads and extracts Discord Game SDK files to the appropriate locations

set -e

# Configuration
SDK_VERSION="3.2.1"
DOWNLOAD_URL="https://dl-game-sdk.discordapp.net/3.2.1/discord_game_sdk.zip"
LIB_DIR="lib"
DISCORD_CGO_DIR="discordcgo"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
WHITE='\033[1;37m'
NC='\033[0m' # No Color

# Function to print colored output
print_color() {
    local color=$1
    shift
    echo -e "${color}$*${NC}"
}

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to detect system architecture
get_system_architecture() {
    local arch=$(uname -m)
    case $arch in
        x86_64|amd64)
            echo "x86_64"
            ;;
        i386|i686)
            echo "x86"
            ;;
        aarch64|arm64)
            echo "aarch64"
            ;;
        *)
            echo "x86_64"  # Default to x86_64
            ;;
    esac
}

# Function to check required files
check_required_files() {
    local missing_files=()
    local architecture=$(get_system_architecture)
    local required_files=()
    
    # Always require the header file
    required_files+=("$LIB_DIR/discord_game_sdk.h")
    
    # Add architecture-specific files
    if [[ "$architecture" == "x86_64" ]]; then
        required_files+=(
            "$LIB_DIR/discord_game_sdk.dll"
            "$LIB_DIR/discord_game_sdk.dll.lib"
            "$LIB_DIR/discord_game_sdk.so"
        )
    elif [[ "$architecture" == "x86" ]]; then
        required_files+=(
            "$LIB_DIR/discord_game_sdk.dll"
            "$LIB_DIR/discord_game_sdk.dll.lib"
        )
    elif [[ "$architecture" == "aarch64" ]]; then
        required_files+=(
            "$LIB_DIR/discord_game_sdk.dylib"
        )
    fi
    
    for file in "${required_files[@]}"; do
        if [[ ! -f "$file" ]]; then
            missing_files+=("$file")
        fi
    done
    
    echo "${missing_files[@]}"
}

# Function to download and extract SDK
download_sdk() {
    print_color $CYAN "Downloading Discord Game SDK version $SDK_VERSION..."
    
    local temp_dir="temp_sdk_download"
    local zip_path="$temp_dir/discord_game_sdk.zip"
    
    # Create temp directory
    if [[ -d "$temp_dir" ]]; then
        rm -rf "$temp_dir"
    fi
    mkdir -p "$temp_dir"
    
    # Check for required tools
    if ! command_exists curl && ! command_exists wget; then
        print_color $RED "Error: Neither curl nor wget is available. Please install one of them."
        exit 1
    fi
    
    # Download the SDK
    print_color $YELLOW "Downloading from: $DOWNLOAD_URL"
    
    if command_exists curl; then
        if ! curl -L -o "$zip_path" "$DOWNLOAD_URL"; then
            print_color $RED "Failed to download SDK with curl"
            exit 1
        fi
    elif command_exists wget; then
        if ! wget -O "$zip_path" "$DOWNLOAD_URL"; then
            print_color $RED "Failed to download SDK with wget"
            exit 1
        fi
    fi
    
    if [[ ! -f "$zip_path" ]]; then
        print_color $RED "Failed to download SDK"
        exit 1
    fi
    
    print_color $GREEN "Download completed successfully!"
    
    # Extract the SDK
    print_color $YELLOW "Extracting SDK files..."
    
    if ! command_exists unzip; then
        print_color $RED "Error: unzip is not available. Please install unzip."
        exit 1
    fi
    
    if ! unzip -q "$zip_path" -d "$temp_dir"; then
        print_color $RED "Failed to extract SDK"
        exit 1
    fi
    
    # Ensure lib directory exists
    mkdir -p "$LIB_DIR"
    
    # Copy files to appropriate locations
    local extracted_path="$temp_dir"
    local architecture=$(get_system_architecture)
    
    print_color $YELLOW "Detected architecture: $architecture"
    
    if [[ -d "$extracted_path" ]]; then
        # Copy header file
        if [[ -f "$extracted_path/c/discord_game_sdk.h" ]]; then
            cp "$extracted_path/c/discord_game_sdk.h" "$LIB_DIR/"
            print_color $GREEN "Copied header file: discord_game_sdk.h"
        fi
        
        # Copy architecture-specific shared libraries
        if [[ "$architecture" == "x86_64" ]]; then
            # Copy Windows x86_64 files
            if [[ -f "$extracted_path/lib/x86_64/discord_game_sdk.dll" ]]; then
                cp "$extracted_path/lib/x86_64/discord_game_sdk.dll" "$LIB_DIR/"
                print_color $GREEN "Copied: discord_game_sdk.dll (x86_64)"
            fi
            if [[ -f "$extracted_path/lib/x86_64/discord_game_sdk.dll.lib" ]]; then
                cp "$extracted_path/lib/x86_64/discord_game_sdk.dll.lib" "$LIB_DIR/"
                print_color $GREEN "Copied: discord_game_sdk.dll.lib (x86_64)"
            fi
            # Copy Linux x86_64 shared library
            if [[ -f "$extracted_path/lib/x86_64/discord_game_sdk.so" ]]; then
                cp "$extracted_path/lib/x86_64/discord_game_sdk.so" "$LIB_DIR/"
                print_color $GREEN "Copied: discord_game_sdk.so (x86_64)"
            fi
        elif [[ "$architecture" == "x86" ]]; then
            # Copy Windows x86 files
            if [[ -f "$extracted_path/lib/x86/discord_game_sdk.dll" ]]; then
                cp "$extracted_path/lib/x86/discord_game_sdk.dll" "$LIB_DIR/"
                print_color $GREEN "Copied: discord_game_sdk.dll (x86)"
            fi
            if [[ -f "$extracted_path/lib/x86/discord_game_sdk.dll.lib" ]]; then
                cp "$extracted_path/lib/x86/discord_game_sdk.dll.lib" "$LIB_DIR/"
                print_color $GREEN "Copied: discord_game_sdk.dll.lib (x86)"
            fi
        elif [[ "$architecture" == "aarch64" ]]; then
            # Copy macOS aarch64 files
            if [[ -f "$extracted_path/lib/aarch64/discord_game_sdk.dylib" ]]; then
                cp "$extracted_path/lib/aarch64/discord_game_sdk.dylib" "$LIB_DIR/"
                print_color $GREEN "Copied: discord_game_sdk.dylib (aarch64)"
            fi
        fi
    fi
    
    print_color $GREEN "SDK files extracted successfully!"
    
    # Clean up temp directory
    rm -rf "$temp_dir"
}

# Main script execution
print_color $CYAN "Discord SDK Download Script"
print_color $CYAN "=========================="

# Check if required files exist
missing_files=($(check_required_files))

if [[ ${#missing_files[@]} -eq 0 ]]; then
    print_color $GREEN "All required Discord SDK files are present!"
    print_color $YELLOW "Files found:"
    for file in "$LIB_DIR"/*; do
        if [[ -f "$file" ]]; then
            print_color $WHITE "  - $(basename "$file")"
        fi
    done
else
    print_color $YELLOW "Missing Discord SDK files:"
    for file in "${missing_files[@]}"; do
        print_color $RED "  - $file"
    done
    print_color $YELLOW "Downloading and extracting SDK files..."
    download_sdk
    
    # Verify files after download
    still_missing=($(check_required_files))
    if [[ ${#still_missing[@]} -eq 0 ]]; then
        print_color $GREEN "SDK setup completed successfully!"
    else
        print_color $RED "Some files are still missing after download:"
        for file in "${still_missing[@]}"; do
            print_color $RED "  - $file"
        done
        exit 1
    fi
fi

print_color $GREEN "Discord SDK setup complete!" 