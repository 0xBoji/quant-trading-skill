#!/bin/bash
# Quant Trading Skill - Installation Script (Go version)
# Usage: curl -fsSL https://raw.githubusercontent.com/0xboji/quant-trading-skill/main/install.sh | bash

set -e

REPO_URL="https://github.com/0xboji/quant-trading-skill"
INSTALL_DIR="$HOME/.quant-trading-skill"
BINARY_NAME="qts"

echo "======================================================================================================"
echo " Quant Trading Skill - Installer (Go Edition)"
echo "======================================================================================================"
echo ""

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "❌ Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

echo "✅ Detected: $OS-$ARCH"

# Check if binary already exists
if command -v $BINARY_NAME &> /dev/null; then
    CURRENT_VERSION=$($BINARY_NAME --version 2>&1 || echo "unknown")
    echo ""
    echo "⚠️  qts is already installed: $CURRENT_VERSION"
    read -p "   Do you want to reinstall? [y/N] " -n 1 -r
    echo ""
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "❌ Installation cancelled"
        exit 0
    fi
fi

# Create install directory
mkdir -p "$INSTALL_DIR"

# Download latest release
echo ""
echo "📦 Downloading latest release..."

DOWNLOAD_URL="$REPO_URL/releases/latest/download/${BINARY_NAME}-${OS}-${ARCH}"

if command -v curl &> /dev/null; then
    curl -fsSL "$DOWNLOAD_URL" -o "$INSTALL_DIR/$BINARY_NAME"
elif command -v wget &> /dev/null; then
    wget -q "$DOWNLOAD_URL" -O "$INSTALL_DIR/$BINARY_NAME"
else
    echo "❌ Error: Neither curl nor wget found. Please install one of them."
    exit 1
fi

# Fallback to git clone if binary not available
if [ ! -f "$INSTALL_DIR/$BINARY_NAME" ] || [ ! -s "$INSTALL_DIR/$BINARY_NAME" ]; then
    echo "📦 Binary not available, building from source..."
    
    # Check for Go
    if ! command -v go &> /dev/null; then
        echo "❌ Error: Go is not installed"
        echo "   Please install Go 1.21+ from: https://go.dev/dl/"
        exit 1
    fi
    
    # Clone and build
    TEMP_DIR=$(mktemp -d)
    git clone --quiet "$REPO_URL" "$TEMP_DIR"
    cd "$TEMP_DIR"
    go build -o "$INSTALL_DIR/$BINARY_NAME" ./cmd/qts
    cd - > /dev/null
    rm -rf "$TEMP_DIR"
fi

# Make executable
chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Download data directory
echo "📊 Downloading knowledge base..."
DATA_DIR="$INSTALL_DIR/data"
mkdir -p "$DATA_DIR"

for csv in strategies indicators risk-management data-sources anti-patterns; do
    curl -fsSL "$REPO_URL/raw/main/data/${csv}.csv" -o "$DATA_DIR/${csv}.csv"
done

# Add to PATH if not already there
SHELL_RC=""
if [ -n "$BASH_VERSION" ]; then
    SHELL_RC="$HOME/.bashrc"
elif [ -n "$ZSH_VERSION" ]; then
    SHELL_RC="$HOME/.zshrc"
fi

if [ -n "$SHELL_RC" ] && ! grep -q "$INSTALL_DIR" "$SHELL_RC" 2>/dev/null; then
    echo "" >> "$SHELL_RC"
    echo "# Quant Trading Skill" >> "$SHELL_RC"
    echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_RC"
    echo "✅ Added to PATH in $SHELL_RC"
fi

# Verify installation
CSV_COUNT=$(ls -1 "$DATA_DIR"/*.csv 2>/dev/null | wc -l | tr -d ' ')

echo ""
echo "======================================================================================================"
echo "✅ Installation successful!"
echo "======================================================================================================"
echo ""
echo "📁 Installed at: $INSTALL_DIR"
echo "📊 Knowledge base: $CSV_COUNT CSV files (122 entries)"
echo "🔍 Binary: $BINARY_NAME"
echo ""
echo "Quick Start:"
echo "------------"
echo ""
echo "# Search strategies (auto-detect domain)"
echo "$BINARY_NAME \"order flow crypto\""
echo ""
echo "# Search specific domain"
echo "$BINARY_NAME \"stop loss\" -d risk"
echo ""
echo "# Get more results"
echo "$BINARY_NAME \"rsi bollinger\" -d indicator -n 5"
echo ""
echo "Available domains: strategy, indicator, risk, data, anti-pattern"
echo ""
echo "Restart your shell or run: source $SHELL_RC"
echo ""
echo "======================================================================================================"
