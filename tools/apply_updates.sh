#!/bin/bash

# Script to apply updates that couldn't be made directly due to permission issues

echo "Applying updates to files that couldn't be modified directly..."

# Check if the new files exist
if [ -f "pkg/greeting/greeting_test.go.new" ]; then
    echo "Updating pkg/greeting/greeting_test.go..."
    cp pkg/greeting/greeting_test.go.new pkg/greeting/greeting_test.go
else
    echo "Error: pkg/greeting/greeting_test.go.new not found"
fi

if [ -f "cmd/BUILD.bazel.new" ]; then
    echo "Updating cmd/BUILD.bazel..."
    cp cmd/BUILD.bazel.new cmd/BUILD.bazel
else
    echo "Error: cmd/BUILD.bazel.new not found"
fi

echo "Updates applied. Please run tests to verify the changes."