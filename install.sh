#!/bin/bash

# Notion Todo CLI Installer for Linux
set -e

# Variables
REPO="caffeines/notion-todo"
BINARY_NAME="todo"
INSTALL_DIR="/usr/local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        ARCH="x86_64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}Unsupported architecture: $ARCH${NC}"
        exit 1
        ;;
esac

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
if [[ "$OS" != "linux" ]]; then
    echo -e "${RED}This installer is for Linux only. Current OS: $OS${NC}"
    exit 1
fi

echo -e "${GREEN}Installing Notion Todo CLI for Linux ${ARCH}...${NC}"

# Get latest release info
echo "Fetching latest release information..."
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest")
VERSION=$(echo "$LATEST_RELEASE" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [[ -z "$VERSION" ]]; then
    echo -e "${RED}Failed to get latest version${NC}"
    exit 1
fi

echo -e "${GREEN}Latest version: $VERSION${NC}"

# Construct download URL
FILENAME="notion-todo_Linux_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "Downloading from: $DOWNLOAD_URL"

# Create temporary directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download and extract
curl -L -o "$FILENAME" "$DOWNLOAD_URL" -s
if [[ $? -ne 0 ]]; then
    echo -e "${RED}Download failed${NC}"
    exit 1
fi

tar -xzf "$FILENAME"

# Check if binary exists
if [[ ! -f "$BINARY_NAME" ]]; then
    echo -e "${RED}Binary not found in archive${NC}"
    exit 1
fi

# Make binary executable
chmod +x "$BINARY_NAME"

# Install binary
echo "Installing to $INSTALL_DIR..."
if [[ -w "$INSTALL_DIR" ]]; then
    cp "$BINARY_NAME" "$INSTALL_DIR/"
else
    echo "Need sudo privileges to install to $INSTALL_DIR"
    sudo cp "$BINARY_NAME" "$INSTALL_DIR/"
fi

# Cleanup
cd /
rm -rf "$TMP_DIR"

# Verify installation
if command -v "$BINARY_NAME" &> /dev/null; then
    echo -e "${GREEN}✅ Installation successful!${NC}"
    echo ""
    echo "Get started with:"
    echo -e "${YELLOW}  $BINARY_NAME guide${NC}"
    echo ""
    echo "For help:"
    echo -e "${YELLOW}  $BINARY_NAME --help${NC}"
else
    echo -e "${RED}Installation may have failed. Binary not found in PATH.${NC}"
    echo "Try running: export PATH=\$PATH:$INSTALL_DIR"
fi
