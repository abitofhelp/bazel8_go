#!/bin/bash

# Ensure GCov is available
if ! command -v gcov >/dev/null 2>&1; then
  echo "GCov not found. Please install GCov."
  exit 1
fi

# Ensure LCOV is available
if ! command -v genhtml >/dev/null 2>&1; then
  echo "genhtml command not found. Installing LCOV..."
  if [ "$(uname)" = "Darwin" ]; then
    brew install lcov
  elif [ "$(uname)" = "Linux" ]; then
    sudo apt-get update && sudo apt-get install -y lcov
  else
    echo "Unsupported OS. Please install LCOV manually."
    exit 1
  fi
fi

# Create coverage directory
mkdir -p coverage-reports

# Run Bazel coverage with explicit GCov path
# Note: The --gcov_tool option is not used as it's not recognized by Bazel
# The issue seems to be resolved by installing LCOV
bazel coverage --test_output=all --combined_report=lcov //...

# Generate HTML report
genhtml -o coverage-reports/bazel "$(bazel info output_path)/_coverage/_coverage_report.dat"

echo "Coverage report generated at coverage-reports/bazel/index.html"
