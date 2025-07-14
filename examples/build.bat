@echo off
REM Discord Game SDK Go Examples Build Script (Batch Wrapper)
REM This script calls the PowerShell build script

echo === Discord Game SDK Go Examples Build Script ===

REM Check if PowerShell is available
powershell -Command "Get-Host" >nul 2>&1
if %errorlevel% neq 0 (
    echo Error: PowerShell is not available
    pause
    exit /b 1
)

REM Check if build.ps1 exists
if not exist "build.ps1" (
    echo Error: build.ps1 not found
    pause
    exit /b 1
)

REM Run the PowerShell script
echo Running PowerShell build script...
powershell -ExecutionPolicy Bypass -File "build.ps1" %*

REM Check if build was successful
if %errorlevel% neq 0 (
    echo.
    echo Build failed with error code %errorlevel%
    pause
    exit /b %errorlevel%
)

echo.
echo Build completed successfully!
echo.
echo To run an example:
echo 1. Navigate to the bin directory
echo 2. Run any .exe file
echo.
pause 