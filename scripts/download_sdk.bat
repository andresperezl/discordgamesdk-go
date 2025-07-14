@echo off
REM Discord SDK Download Script for Windows (Batch version)
REM Downloads and extracts Discord Game SDK files to the appropriate locations

setlocal enabledelayedexpansion

REM Configuration
set SDK_VERSION=3.2.1
set DOWNLOAD_URL=https://dl-game-sdk.discordapp.net/3.2.1/discord_game_sdk.zip
set LIB_DIR=lib
set DISCORD_CGO_DIR=discordcgo

REM Colors for output (Windows 10+)
if "%TERM%"=="xterm" (
    set "GREEN=[92m"
    set "YELLOW=[93m"
    set "RED=[91m"
    set "CYAN=[96m"
    set "WHITE=[97m"
    set "NC=[0m"
) else (
    set "GREEN="
    set "YELLOW="
    set "RED="
    set "CYAN="
    set "WHITE="
    set "NC="
)

REM Function to print colored output
:print_color
echo %~1%~2%NC%
goto :eof

REM Function to detect system architecture
:get_system_architecture
if "%PROCESSOR_ARCHITECTURE%"=="AMD64" (
    echo x86_64
) else if "%PROCESSOR_ARCHITECTURE%"=="x86" (
    echo x86
) else (
    echo x86_64
)
goto :eof

REM Function to check required files
:check_required_files
set missing_files=
call :get_system_architecture
set architecture=%ERRORLEVEL%

REM Always require the header file
if not exist "%LIB_DIR%\discord_game_sdk.h" (
    set missing_files=!missing_files! %LIB_DIR%\discord_game_sdk.h
)

REM Add architecture-specific files
if "%architecture%"=="x86_64" (
    if not exist "%LIB_DIR%\discord_game_sdk.dll" set missing_files=!missing_files! %LIB_DIR%\discord_game_sdk.dll
    if not exist "%LIB_DIR%\discord_game_sdk.dll.lib" set missing_files=!missing_files! %LIB_DIR%\discord_game_sdk.dll.lib
    if not exist "%LIB_DIR%\discord_game_sdk.so" set missing_files=!missing_files! %LIB_DIR%\discord_game_sdk.so
) else if "%architecture%"=="x86" (
    if not exist "%LIB_DIR%\discord_game_sdk.dll" set missing_files=!missing_files! %LIB_DIR%\discord_game_sdk.dll
    if not exist "%LIB_DIR%\discord_game_sdk.dll.lib" set missing_files=!missing_files! %LIB_DIR%\discord_game_sdk.dll.lib
)

echo %missing_files%
goto :eof

REM Function to download and extract SDK
:download_sdk
call :print_color %CYAN% "Downloading Discord Game SDK version %SDK_VERSION%..."

set temp_dir=temp_sdk_download
set zip_path=%temp_dir%\discord_game_sdk.zip

REM Create temp directory
if exist "%temp_dir%" rmdir /s /q "%temp_dir%"
mkdir "%temp_dir%"

REM Download the SDK
call :print_color %YELLOW% "Downloading from: %DOWNLOAD_URL%"

REM Try to use PowerShell for download if available
powershell -Command "& {[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; Invoke-WebRequest -Uri '%DOWNLOAD_URL%' -OutFile '%zip_path%'}" 2>nul
if not exist "%zip_path%" (
    call :print_color %RED% "Failed to download SDK"
    exit /b 1
)

call :print_color %GREEN% "Download completed successfully!"

REM Extract the SDK
call :print_color %YELLOW% "Extracting SDK files..."

powershell -Command "Expand-Archive -Path '%zip_path%' -DestinationPath '%temp_dir%' -Force" 2>nul
if errorlevel 1 (
    call :print_color %RED% "Failed to extract SDK"
    exit /b 1
)

REM Ensure lib directory exists
if not exist "%LIB_DIR%" mkdir "%LIB_DIR%"

REM Copy files to appropriate locations
set extracted_path=%temp_dir%
call :get_system_architecture
set architecture=%ERRORLEVEL%

call :print_color %YELLOW% "Detected architecture: %architecture%"

if exist "%extracted_path%" (
    REM Copy header file
    if exist "%extracted_path%\c\discord_game_sdk.h" (
        copy "%extracted_path%\c\discord_game_sdk.h" "%LIB_DIR%\" >nul
        call :print_color %GREEN% "Copied header file: discord_game_sdk.h"
    )
    
    REM Copy architecture-specific shared libraries
    if "%architecture%"=="x86_64" (
        REM Copy Windows x86_64 files
        if exist "%extracted_path%\lib\x86_64\discord_game_sdk.dll" (
            copy "%extracted_path%\lib\x86_64\discord_game_sdk.dll" "%LIB_DIR%\" >nul
            call :print_color %GREEN% "Copied: discord_game_sdk.dll (x86_64)"
        )
        if exist "%extracted_path%\lib\x86_64\discord_game_sdk.dll.lib" (
            copy "%extracted_path%\lib\x86_64\discord_game_sdk.dll.lib" "%LIB_DIR%\" >nul
            call :print_color %GREEN% "Copied: discord_game_sdk.dll.lib (x86_64)"
        )
        REM Copy Linux x86_64 shared library
        if exist "%extracted_path%\lib\x86_64\discord_game_sdk.so" (
            copy "%extracted_path%\lib\x86_64\discord_game_sdk.so" "%LIB_DIR%\" >nul
            call :print_color %GREEN% "Copied: discord_game_sdk.so (x86_64)"
        )
    ) else if "%architecture%"=="x86" (
        REM Copy Windows x86 files
        if exist "%extracted_path%\lib\x86\discord_game_sdk.dll" (
            copy "%extracted_path%\lib\x86\discord_game_sdk.dll" "%LIB_DIR%\" >nul
            call :print_color %GREEN% "Copied: discord_game_sdk.dll (x86)"
        )
        if exist "%extracted_path%\lib\x86\discord_game_sdk.dll.lib" (
            copy "%extracted_path%\lib\x86\discord_game_sdk.dll.lib" "%LIB_DIR%\" >nul
            call :print_color %GREEN% "Copied: discord_game_sdk.dll.lib (x86)"
        )
    )
)

call :print_color %GREEN% "SDK files extracted successfully!"

REM Clean up temp directory
if exist "%temp_dir%" rmdir /s /q "%temp_dir%"

goto :eof

REM Main script execution
call :print_color %CYAN% "Discord SDK Download Script"
call :print_color %CYAN% "=========================="

REM Check if required files exist
for /f "tokens=*" %%i in ('call :check_required_files') do set missing_files=%%i

if "%missing_files%"=="" (
    call :print_color %GREEN% "All required Discord SDK files are present!"
    call :print_color %YELLOW% "Files found:"
    for %%f in ("%LIB_DIR%\*") do (
        call :print_color %WHITE% "  - %%~nxf"
    )
) else (
    call :print_color %YELLOW% "Missing Discord SDK files:"
    for %%f in (%missing_files%) do (
        call :print_color %RED% "  - %%f"
    )
    call :print_color %YELLOW% "Downloading and extracting SDK files..."
    call :download_sdk
    
    REM Verify files after download
    for /f "tokens=*" %%i in ('call :check_required_files') do set still_missing=%%i
    if "%still_missing%"=="" (
        call :print_color %GREEN% "SDK setup completed successfully!"
    ) else (
        call :print_color %RED% "Some files are still missing after download:"
        for %%f in (%still_missing%) do (
            call :print_color %RED% "  - %%f"
        )
        exit /b 1
    )
)

call :print_color %GREEN% "Discord SDK setup complete!" 