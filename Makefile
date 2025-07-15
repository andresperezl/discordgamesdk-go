# Makefile for Discord Game SDK Go Wrapper

.PHONY: help sdk-download sdk-check build build-win-dll examples examples-win-dll clean test install quickstart example

# Default target
help:
	@echo "Available targets:"
	@echo "  help         - Show this help message"
	@echo "  sdk-download - Download and extract Discord SDK files"
	@echo "  sdk-check    - Check if Discord SDK files are present"
	@echo "  build        - Build the main package"
	@echo "  examples     - Build all examples"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  install      - Download SDK and build main package"
	@echo "  quickstart   - Download SDK and build examples"
	@echo "  example      - Build a single example: make example NAME=activity"

# Detect system architecture
ifeq ($(OS),Windows_NT)
    # Windows architecture detection
    ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
        ARCH := x86_64
    else ifeq ($(PROCESSOR_ARCHITECTURE),x86)
        ARCH := x86
    else
        ARCH := x86_64
    endif
else
    # Unix-like system architecture detection
    ARCH := $(shell uname -m)
    ifeq ($(ARCH),x86_64)
        ARCH := x86_64
    else ifeq ($(ARCH),amd64)
        ARCH := x86_64
    else ifeq ($(ARCH),i386)
        ARCH := x86
    else ifeq ($(ARCH),i686)
        ARCH := x86
    else ifeq ($(ARCH),aarch64)
        ARCH := aarch64
    else ifeq ($(ARCH),arm64)
        ARCH := aarch64
    else
        ARCH := x86_64
    endif
endif

# Download Discord SDK files
sdk-download:
	@echo "Downloading Discord SDK files for architecture: $(ARCH)..."
ifeq ($(OS),Windows_NT)
	@powershell -ExecutionPolicy Bypass -File scripts/download_sdk.ps1
else
	@chmod +x scripts/download_sdk.sh
	@./scripts/download_sdk.sh
endif

# Check if SDK files are present
sdk-check:
	@echo "Checking Discord SDK files for architecture: $(ARCH)..."
ifeq ($(OS),Windows_NT)
	@powershell -ExecutionPolicy Bypass -File scripts/download_sdk.ps1
else
	@chmod +x scripts/download_sdk.sh
	@./scripts/download_sdk.sh
endif

# Windows-specific DLL copy for main build
build-win-dll:
	@echo "Copying DLL files for Windows ($(ARCH))..."
	@if exist lib\discord_game_sdk.dll (copy lib\discord_game_sdk.dll .) else (echo Warning: lib\discord_game_sdk.dll not found)
	@if exist lib\discord_game_sdk.dll.lib (copy lib\discord_game_sdk.dll.lib .) else (echo Warning: lib\discord_game_sdk.dll.lib not found)

# Build the main package
build: sdk-check
	@echo "Building main package for architecture: $(ARCH)..."
ifeq ($(OS),Windows_NT)
	@go build -o discordctl.exe .
	@$(MAKE) build-win-dll
else
	@go build -o discordctl .
endif

# Windows-specific DLL copy for examples
examples-win-dll:
	@echo "Copying DLL files for Windows examples ($(ARCH))..."
	@if exist lib\discord_game_sdk.dll (copy lib\discord_game_sdk.dll examples\bin\) else (echo Warning: lib\discord_game_sdk.dll not found)
	@if exist lib\discord_game_sdk.dll.lib (copy lib\discord_game_sdk.dll.lib examples\bin\) else (echo Warning: lib\discord_game_sdk.dll.lib not found)

# Build all examples
examples: sdk-check
ifeq ($(OS),Windows_NT)
	@echo "Building examples (Windows - $(ARCH))..."
	@if not exist examples\bin mkdir examples\bin
	@for %%d in (examples\activity examples\activity_simple examples\basic examples\callback_test examples\configuration_test examples\diagnostic examples\find_client_id examples\storage examples\test_minimal examples\user) do @if exist %%d\main.go ( \
		echo Building %%~nxd... && \
		go build -o examples\bin\%%~nxd.exe %%d\main.go \
	)
	@$(MAKE) examples-win-dll
else
	@echo "Building examples (Unix - $(ARCH))..."
	@mkdir -p examples/bin
	@for example in examples/*/main.go; do \
		if [ -f "$$example" ]; then \
			name=$$(basename $$(dirname "$$example")); \
			echo "Building $$name..."; \
			go build -o "examples/bin/$$name" "$$example"; \
		fi \
	done
endif

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
ifeq ($(OS),Windows_NT)
	@if exist discordctl.exe del discordctl.exe
	@if exist discordctl del discordctl
	@if exist discord_game_sdk.dll del discord_game_sdk.dll
	@if exist discord_game_sdk.dll.lib del discord_game_sdk.dll.lib
	@if exist examples\bin rmdir /s /q examples\bin
else
	@rm -f discordctl
	@rm -rf examples/bin
endif
	@go clean

# Run tests
test: sdk-check
	@echo "Running tests..."
	@go test ./...

# Install dependencies and build
install: sdk-download build
	@echo "Installation complete!"

# Quick start - download SDK and build examples
quickstart: sdk-download examples
	@echo "Quick start complete! Examples built in examples/bin/" 

# Build a single example by name
example: sdk-check
ifeq ($(OS),Windows_NT)
	@if not exist examples\bin mkdir examples\bin
	@echo Building example $(NAME)...
	@go build -o examples\bin\$(NAME).exe examples\$(NAME)\main.go
	@$(MAKE) examples-win-dll
else
	@mkdir -p examples/bin
	@echo "Building example $(NAME)..."
	@go build -o examples/bin/$(NAME) examples/$(NAME)/main.go
endif 