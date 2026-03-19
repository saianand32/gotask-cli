#!/bin/bash

# gotask-cli macOS Installer

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$DIR/../.." && pwd )"
INSTALL_DIR="/usr/local/bin"
DEST_BIN="$INSTALL_DIR/gotask"

print_help() {
    echo "==============================="
    echo "   gotask-cli macOS Installer  "
    echo "==============================="
    echo "Usage: $0 [command]"
    echo ""
    echo "Commands:"
    echo "  --install    Install gotask-cli"
    echo "  --uninstall  Uninstall gotask-cli"
    echo "  --update     Update gotask-cli to the latest built version"
    echo "  --help       Show this help message"
}

build_and_install() {
    local action=$1
    echo "Changing to project root: $PROJECT_ROOT"
    cd "$PROJECT_ROOT" || { echo "Failed to change directory to $PROJECT_ROOT"; exit 1; }

    echo "Building the macOS binary..."
    if ! make build-mac; then
        echo "Error: Build failed."
        exit 1
    fi

    local BINARY_PATH="build/bin/gotask-mac"

    if [ ! -f "$BINARY_PATH" ]; then
        echo "Error: Binary not found at $BINARY_PATH"
        exit 1
    fi

    if [ "$action" = "install" ]; then
        echo "Installing gotask-cli to $INSTALL_DIR..."
    else
        echo "Updating gotask-cli in $INSTALL_DIR..."
    fi
    echo "You may be prompted for your password to copy the file to $INSTALL_DIR"

    sudo cp "$BINARY_PATH" "$DEST_BIN"
    sudo chmod +x "$DEST_BIN"

    if [ $? -eq 0 ]; then
        echo "==============================="
        if [ "$action" = "install" ]; then
            echo "Installation completed successfully!"
        else
            echo "Update completed successfully!"
        fi
        echo "You can now run 'gotask' from your terminal."
        echo "Try running: gotask help"
        echo "==============================="
    else
        echo "$action failed. Please check your permissions and try again."
        exit 1
    fi
}

uninstall() {
    if [ ! -f "$DEST_BIN" ]; then
        echo "gotask-cli is not installed at $DEST_BIN."
        exit 0
    fi

    echo "Uninstalling gotask-cli from $INSTALL_DIR..."
    echo "You may be prompted for your password to remove the file."
    
    sudo rm "$DEST_BIN"
    
    if [ $? -eq 0 ]; then
        echo "==============================="
        echo "Uninstallation completed successfully!"
        echo "==============================="
    else
        echo "Uninstallation failed. Please check your permissions and try again."
        exit 1
    fi
}

if [ $# -eq 0 ]; then
    print_help
    exit 0
fi

case "$1" in
    --install)
        build_and_install "install"
        ;;
    --uninstall)
        uninstall
        ;;
    --update)
        build_and_install "update"
        ;;
    --help)
        print_help
        ;;
    *)
        echo "Unknown command: $1"
        print_help
        exit 1
        ;;
esac
