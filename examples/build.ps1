# Discord Game SDK Go Examples Build Script
# This script compiles all examples and copies the necessary DLL

param(
    [string]$BuildType = "Release"
)

Write-Host "=== Discord Game SDK Go Examples Build Script ===" -ForegroundColor Green

# Set error action preference
$ErrorActionPreference = "Stop"

# Get the script directory
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = Split-Path -Parent $ScriptDir
$LibDir = Join-Path $ProjectRoot "lib"
$DllPath = Join-Path $LibDir "discord_game_sdk.dll"

# Check if DLL exists
if (-not (Test-Path $DllPath)) {
    Write-Error "Discord Game SDK DLL not found at: $DllPath"
    exit 1
}

Write-Host "Found Discord Game SDK DLL at: $DllPath" -ForegroundColor Yellow

# Create output directory
$OutputDir = Join-Path $ScriptDir "bin"
if (-not (Test-Path $OutputDir)) {
    New-Item -ItemType Directory -Path $OutputDir | Out-Null
    Write-Host "Created output directory: $OutputDir" -ForegroundColor Yellow
}

# Build flags
$BuildFlags = @()
if ($BuildType -eq "Debug") {
    $BuildFlags += "-gcflags=all=-N -l"
    Write-Host "Building in Debug mode" -ForegroundColor Yellow
} else {
    Write-Host "Building in Release mode" -ForegroundColor Yellow
}

# Get all example directories
$ExampleDirs = Get-ChildItem -Path $ScriptDir -Directory | Where-Object { $_.Name -ne "bin" }

Write-Host "Found example directories:" -ForegroundColor Cyan
foreach ($dir in $ExampleDirs) {
    Write-Host "  - $($dir.Name)" -ForegroundColor White
}

# Build each example
foreach ($exampleDir in $ExampleDirs) {
    $exampleName = $exampleDir.Name
    $examplePath = $exampleDir.FullName

    if (-not $examplePath) {
        Write-Warning "Directory $exampleName has no valid path, skipping..."
        continue
    }

    $mainFile = Join-Path $examplePath "main.go"

    if (-not (Test-Path $mainFile)) {
        Write-Warning "No main.go found in $exampleName, skipping..."
        continue
    }

    Write-Host "Building $exampleName..." -ForegroundColor Green

    # Set output path
    $exeName = "$exampleName.exe"
    $outputPath = Join-Path $OutputDir $exeName

    # Build command
    $buildCmd = "go build"
    if ($BuildFlags.Count -gt 0) {
        $buildCmd += " $($BuildFlags -join ' ')"
    }
    $buildCmd += " -o `"$outputPath`" `"$mainFile`""

    try {
        # Change to project root for build
        Push-Location $ProjectRoot

        # Execute build
        Invoke-Expression $buildCmd

        if ($LASTEXITCODE -eq 0) {
            Write-Host "Successfully built $exeName" -ForegroundColor Green

            # Copy DLL to output directory
            $dllDest = Join-Path $OutputDir "discord_game_sdk.dll"
            Copy-Item -Path $DllPath -Destination $dllDest -Force
            Write-Host "Copied DLL to output directory" -ForegroundColor Green
        } else {
            Write-Error "Failed to build $exampleName"
        }
    }
    catch {
        Write-Error "Error building $exampleName`: $_"
    }
    finally {
        Pop-Location
    }
}

Write-Host "=== Build Summary ===" -ForegroundColor Green
Write-Host "Output directory: $OutputDir" -ForegroundColor Yellow

# List built executables
$builtExes = Get-ChildItem -Path $OutputDir -Filter "*.exe"
if ($builtExes.Count -gt 0) {
    Write-Host "Built executables:" -ForegroundColor Cyan
    foreach ($exe in $builtExes) {
        $size = [math]::Round($exe.Length / 1KB, 2)
        Write-Host "  - $($exe.Name) ($size KB)" -ForegroundColor White
    }
} else {
    Write-Warning "No executables were built"
}

# Check if DLL was copied
$copiedDll = Join-Path $OutputDir "discord_game_sdk.dll"
if (Test-Path $copiedDll) {
    Write-Host "Discord Game SDK DLL copied to output directory" -ForegroundColor Green
} else {
    Write-Warning "Discord Game SDK DLL not found in output directory"
}

Write-Host "Build completed!" -ForegroundColor Green
Write-Host "To run an example, navigate to the bin directory and execute the .exe file." -ForegroundColor Yellow 