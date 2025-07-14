#!/usr/bin/env pwsh

# Discord SDK Download Script for Windows
# Downloads and extracts Discord Game SDK files to the appropriate locations

param(
    [string]$SDKVersion = "3.2.1",
    [string]$DownloadUrl = "https://dl-game-sdk.discordapp.net/3.2.1/discord_game_sdk.zip",
    [string]$LibDir = "lib",
    [string]$DiscordCgoDir = "discordcgo"
)

function Write-ColorOutput($ForegroundColor) {
    $fc = $host.UI.RawUI.ForegroundColor
    $host.UI.RawUI.ForegroundColor = $ForegroundColor
    if ($args) {
        Write-Output $args
    }
    else {
        $input | Write-Output
    }
    $host.UI.RawUI.ForegroundColor = $fc
}

function Get-SystemArchitecture {
    if ([Environment]::Is64BitOperatingSystem) {
        return "x86_64"
    } else {
        return "x86"
    }
}

function Test-RequiredFiles {
    $architecture = Get-SystemArchitecture
    $requiredFiles = @()
    
    # Always require the header file
    $requiredFiles += "$LibDir/discord_game_sdk.h"
    
    # Add architecture-specific files
    if ($architecture -eq "x86_64") {
        $requiredFiles += @(
            "$LibDir/discord_game_sdk.dll",
            "$LibDir/discord_game_sdk.dll.lib",
            "$LibDir/discord_game_sdk.so"
        )
    } else {
        $requiredFiles += @(
            "$LibDir/discord_game_sdk.dll",
            "$LibDir/discord_game_sdk.dll.lib"
        )
    }
    
    $missingFiles = @()
    foreach ($file in $requiredFiles) {
        if (-not (Test-Path $file)) {
            $missingFiles += $file
        }
    }
    
    return $missingFiles
}

function Download-SDK {
    Write-ColorOutput Green "Downloading Discord Game SDK version $SDKVersion..."
    
    $tempDir = "temp_sdk_download"
    $zipPath = "$tempDir/discord_game_sdk.zip"
    
    # Create temp directory
    if (Test-Path $tempDir) {
        Remove-Item $tempDir -Recurse -Force
    }
    New-Item -ItemType Directory -Path $tempDir | Out-Null
    
    try {
        # Download the SDK
        Write-ColorOutput Yellow "Downloading from: $DownloadUrl"
        Invoke-WebRequest -Uri $DownloadUrl -OutFile $zipPath -UseBasicParsing
        
        if (-not (Test-Path $zipPath)) {
            throw "Failed to download SDK"
        }
        
        Write-ColorOutput Green "Download completed successfully!"
        
        # Extract the SDK
        Write-ColorOutput Yellow "Extracting SDK files..."
        Expand-Archive -Path $zipPath -DestinationPath $tempDir -Force
        
        # Ensure lib directory exists
        if (-not (Test-Path $LibDir)) {
            New-Item -ItemType Directory -Path $LibDir | Out-Null
        }
        
        # Copy files to appropriate locations
        $extractedPath = "$tempDir"
        if (Test-Path $extractedPath) {
            $architecture = Get-SystemArchitecture
            Write-ColorOutput Yellow "Detected architecture: $architecture"
            
            # Copy header file
            if (Test-Path "$extractedPath/c/discord_game_sdk.h") {
                Copy-Item "$extractedPath/c/discord_game_sdk.h" "$LibDir/" -Force
                Write-ColorOutput Green "Copied header file: discord_game_sdk.h"
            }
            
            # Copy architecture-specific shared libraries
            if ($architecture -eq "x86_64") {
                # Copy Windows x86_64 files
                if (Test-Path "$extractedPath/lib/x86_64/discord_game_sdk.dll") {
                    Copy-Item "$extractedPath/lib/x86_64/discord_game_sdk.dll" "$LibDir/" -Force
                    Write-ColorOutput Green "Copied: discord_game_sdk.dll (x86_64)"
                }
                if (Test-Path "$extractedPath/lib/x86_64/discord_game_sdk.dll.lib") {
                    Copy-Item "$extractedPath/lib/x86_64/discord_game_sdk.dll.lib" "$LibDir/" -Force
                    Write-ColorOutput Green "Copied: discord_game_sdk.dll.lib (x86_64)"
                }
                # Copy Linux x86_64 shared library
                if (Test-Path "$extractedPath/lib/x86_64/discord_game_sdk.so") {
                    Copy-Item "$extractedPath/lib/x86_64/discord_game_sdk.so" "$LibDir/" -Force
                    Write-ColorOutput Green "Copied: discord_game_sdk.so (x86_64)"
                }
            } else {
                # Copy Windows x86 files
                if (Test-Path "$extractedPath/lib/x86/discord_game_sdk.dll") {
                    Copy-Item "$extractedPath/lib/x86/discord_game_sdk.dll" "$LibDir/" -Force
                    Write-ColorOutput Green "Copied: discord_game_sdk.dll (x86)"
                }
                if (Test-Path "$extractedPath/lib/x86/discord_game_sdk.dll.lib") {
                    Copy-Item "$extractedPath/lib/x86/discord_game_sdk.dll.lib" "$LibDir/" -Force
                    Write-ColorOutput Green "Copied: discord_game_sdk.dll.lib (x86)"
                }
            }
        }
        
        Write-ColorOutput Green "SDK files extracted successfully!"
        
    }
    catch {
        Write-ColorOutput Red "Error downloading/extracting SDK: $($_.Exception.Message)"
        exit 1
    }
    finally {
        # Clean up temp directory
        if (Test-Path $tempDir) {
            Remove-Item $tempDir -Recurse -Force
        }
    }
}

# Main script execution
Write-ColorOutput Cyan "Discord SDK Download Script"
Write-ColorOutput Cyan "=========================="

# Check if required files exist
$missingFiles = Test-RequiredFiles

if ($missingFiles.Count -eq 0) {
    Write-ColorOutput Green "All required Discord SDK files are present!"
    Write-ColorOutput Yellow "Files found:"
    Get-ChildItem $LibDir -Name | ForEach-Object { Write-ColorOutput White "  - $_" }
} else {
    Write-ColorOutput Yellow "Missing Discord SDK files:"
    $missingFiles | ForEach-Object { Write-ColorOutput Red "  - $_" }
    Write-ColorOutput Yellow "Downloading and extracting SDK files..."
    Download-SDK
    
    # Verify files after download
    $stillMissing = Test-RequiredFiles
    if ($stillMissing.Count -eq 0) {
        Write-ColorOutput Green "SDK setup completed successfully!"
    } else {
        Write-ColorOutput Red "Some files are still missing after download:"
        $stillMissing | ForEach-Object { Write-ColorOutput Red "  - $_" }
        exit 1
    }
}

Write-ColorOutput Green "Discord SDK setup complete!" 