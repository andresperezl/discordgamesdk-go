package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	SDKVersion    = "3.2.1"
	DownloadURL   = "https://dl-game-sdk.discordapp.net/3.2.1/discord_game_sdk.zip"
	LibDir        = "lib"
	DiscordCgoDir = "discordcgo"
	TempDir       = "temp_sdk_download"
)

// Colors for output
const (
	Red    = "\033[0;31m"
	Green  = "\033[0;32m"
	Yellow = "\033[1;33m"
	Blue   = "\033[0;34m"
	Cyan   = "\033[0;36m"
	White  = "\033[1;37m"
	NC     = "\033[0m" // No Color
)

func printColor(color, message string) {
	fmt.Printf("%s%s%s\n", color, message, NC)
}

func checkRequiredFiles() []string {
	requiredFiles := []string{
		filepath.Join(LibDir, "discord_game_sdk.dll"),
		filepath.Join(LibDir, "discord_game_sdk.dll.lib"),
		filepath.Join(LibDir, "discord_game_sdk.so"),
		filepath.Join(LibDir, "discord_game_sdk.h"),
	}

	var missingFiles []string
	for _, file := range requiredFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			missingFiles = append(missingFiles, file)
		}
	}
	return missingFiles
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func extractZip(zipPath, destPath string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		filePath := filepath.Join(destPath, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, file.Mode())
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		zipFile, err := file.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, zipFile)
		outFile.Close()
		zipFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func downloadSDK() error {
	printColor(Cyan, fmt.Sprintf("Downloading Discord Game SDK version %s...", SDKVersion))

	// Create temp directory
	if err := os.RemoveAll(TempDir); err != nil {
		return err
	}
	if err := os.MkdirAll(TempDir, 0755); err != nil {
		return err
	}

	zipPath := filepath.Join(TempDir, "discord_game_sdk.zip")

	// Download the SDK
	printColor(Yellow, fmt.Sprintf("Downloading from: %s", DownloadURL))
	if err := downloadFile(DownloadURL, zipPath); err != nil {
		return fmt.Errorf("failed to download SDK: %v", err)
	}

	printColor(Green, "Download completed successfully!")

	// Extract the SDK
	printColor(Yellow, "Extracting SDK files...")
	if err := extractZip(zipPath, TempDir); err != nil {
		return fmt.Errorf("failed to extract SDK: %v", err)
	}

	// Ensure lib directory exists
	if err := os.MkdirAll(LibDir, 0755); err != nil {
		return err
	}

	// Copy files to appropriate locations
	extractedPath := filepath.Join(TempDir, "discord_game_sdk")
	if _, err := os.Stat(extractedPath); err == nil {
		// Copy Windows files (x86_64)
		winDll := filepath.Join(extractedPath, "lib", "x86_64", "discord_game_sdk.dll")
		winLib := filepath.Join(extractedPath, "lib", "x86_64", "discord_game_sdk.dll.lib")

		if _, err := os.Stat(winDll); err == nil {
			if err := copyFile(winDll, filepath.Join(LibDir, "discord_game_sdk.dll")); err != nil {
				return err
			}
		}
		if _, err := os.Stat(winLib); err == nil {
			if err := copyFile(winLib, filepath.Join(LibDir, "discord_game_sdk.dll.lib")); err != nil {
				return err
			}
		}

		// Copy Linux files (x86_64)
		linuxSo := filepath.Join(extractedPath, "lib", "x86_64", "libdiscord_game_sdk.so")
		if _, err := os.Stat(linuxSo); err == nil {
			if err := copyFile(linuxSo, filepath.Join(LibDir, "discord_game_sdk.so")); err != nil {
				return err
			}
		}

		// Copy header files
		headerFile := filepath.Join(extractedPath, "c", "discord_game_sdk.h")
		if _, err := os.Stat(headerFile); err == nil {
			if err := copyFile(headerFile, filepath.Join(LibDir, "discord_game_sdk.h")); err != nil {
				return err
			}
		}
	}

	printColor(Green, "SDK files extracted successfully!")

	// Clean up temp directory
	return os.RemoveAll(TempDir)
}

func main() {
	printColor(Cyan, "Discord SDK Download Script")
	printColor(Cyan, "==========================")

	// Check if required files exist
	missingFiles := checkRequiredFiles()

	if len(missingFiles) == 0 {
		printColor(Green, "All required Discord SDK files are present!")
		printColor(Yellow, "Files found:")

		files, err := os.ReadDir(LibDir)
		if err == nil {
			for _, file := range files {
				if !file.IsDir() {
					printColor(White, fmt.Sprintf("  - %s", file.Name()))
				}
			}
		}
	} else {
		printColor(Yellow, "Missing Discord SDK files:")
		for _, file := range missingFiles {
			printColor(Red, fmt.Sprintf("  - %s", file))
		}
		printColor(Yellow, "Downloading and extracting SDK files...")

		if err := downloadSDK(); err != nil {
			printColor(Red, fmt.Sprintf("Error downloading/extracting SDK: %v", err))
			os.Exit(1)
		}

		// Verify files after download
		stillMissing := checkRequiredFiles()
		if len(stillMissing) == 0 {
			printColor(Green, "SDK setup completed successfully!")
		} else {
			printColor(Red, "Some files are still missing after download:")
			for _, file := range stillMissing {
				printColor(Red, fmt.Sprintf("  - %s", file))
			}
			os.Exit(1)
		}
	}

	printColor(Green, "Discord SDK setup complete!")
}
