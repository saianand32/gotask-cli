APP_NAME := gotask
SRC_DIR := ./cmd/app
BIN_DIR := ./build/bin
GO := go
GO_BUILD := $(GO) build -o $(BIN_DIR)/$(APP_NAME)
VERSION ?= dev
LDFLAGS := -X main.version=$(VERSION)

# Default target
.PHONY: all build clean help build-linux build-mac build-windows

all: build

# Build for current OS
build:
	@echo "Building $(APP_NAME) ($(VERSION)) for current OS..."
	@mkdir -p $(BIN_DIR)
	@$(GO_BUILD) -ldflags "$(LDFLAGS)" $(SRC_DIR)/main.go
	@echo "Build complete: $(BIN_DIR)/$(APP_NAME)"

# OS-specific builds
build-linux:
	@echo "Building $(APP_NAME) ($(VERSION)) for Linux..."
	@mkdir -p $(BIN_DIR)
	@GOOS=linux GOARCH=amd64 $(GO_BUILD)-linux -ldflags "$(LDFLAGS)" $(SRC_DIR)/main.go
	@echo "linux build: $(BIN_DIR)/$(APP_NAME)-linux"

build-mac:
	@echo "Building $(APP_NAME) ($(VERSION)) for macOS..."
	@mkdir -p $(BIN_DIR)
	@GOOS=darwin GOARCH=amd64 $(GO_BUILD)-mac -ldflags "$(LDFLAGS)" $(SRC_DIR)/main.go
	@echo "mac build: $(BIN_DIR)/$(APP_NAME)-mac"

build-win:
	@echo "Building $(APP_NAME) ($(VERSION)) for Windows..."
	@mkdir -p $(BIN_DIR)
	@GOOS=windows GOARCH=amd64 $(GO_BUILD).exe -ldflags "$(LDFLAGS)" $(SRC_DIR)/main.go
	@echo "windows build: $(BIN_DIR)/$(APP_NAME).exe"

# Clean
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BIN_DIR)
	@echo "Clean complete"

# Help
help:
	@echo ""
	@echo "Available commands:"
	@echo "  make                Build for current OS"
	@echo "  make build-linux    Build for Linux"
	@echo "  make build-mac      Build for macOS"
	@echo "  make build-win      Build for Windows"
	@echo "  make clean          Remove built files"
	@echo "  make help           Show this help message"
	@echo ""
	@echo "  Override build version by setting explicitly: make build VERSION=1.0.0"
	@echo ""
