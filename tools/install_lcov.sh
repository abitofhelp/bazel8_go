#!/bin/bash

# Script to install lcov package which provides the genhtml command
# This is required for generating HTML coverage reports from Bazel coverage data

echo "Checking for genhtml command..."

if command -v genhtml >/dev/null 2>&1; then
    echo "genhtml command found. lcov is already installed."
else
    echo "genhtml command not found. Installing lcov package..."
    
    # Detect OS
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        if command -v brew >/dev/null 2>&1; then
            echo "Installing lcov using Homebrew..."
            brew install lcov
        else
            echo "Homebrew not found. Please install Homebrew first:"
            echo "  /bin/bash -c \"\$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
            exit 1
        fi
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        if command -v apt-get >/dev/null 2>&1; then
            # Debian/Ubuntu
            echo "Installing lcov using apt-get..."
            sudo apt-get update
            sudo apt-get install -y lcov
        elif command -v yum >/dev/null 2>&1; then
            # CentOS/RHEL
            echo "Installing lcov using yum..."
            sudo yum install -y lcov
        elif command -v dnf >/dev/null 2>&1; then
            # Fedora
            echo "Installing lcov using dnf..."
            sudo dnf install -y lcov
        else
            echo "Unsupported Linux distribution. Please install lcov manually."
            exit 1
        fi
    else
        echo "Unsupported operating system: $OSTYPE"
        echo "Please install lcov manually."
        exit 1
    fi
fi

# Verify installation
if command -v genhtml >/dev/null 2>&1; then
    echo "lcov installation successful. genhtml command is now available."
    echo "You can now run 'make bazel-test-coverage' to generate coverage reports."
else
    echo "Failed to install lcov or genhtml command not found after installation."
    echo "Please install lcov manually."
    exit 1
fi