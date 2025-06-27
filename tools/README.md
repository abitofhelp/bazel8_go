# Tools

This directory contains utility scripts and tools for the Bazel8_Go project.

## Overview

The tools in this directory help with various development tasks, such as generating code coverage reports and installing dependencies.

## Available Tools

### Install LCOV

`install_lcov.sh` - A script to install LCOV, which is used for generating HTML code coverage reports.

#### Usage

```bash
./tools/install_lcov.sh
```

### Run Coverage

`run_coverage.sh` - A script to generate code coverage reports for the project.

#### Usage

```bash
./tools/run_coverage.sh
```

This script:
1. Checks if GCov is installed
2. Checks if LCOV is installed, and installs it if not (using Homebrew on macOS or apt on Linux)
3. Runs the Bazel coverage command
4. Generates an HTML report using genhtml

The coverage report will be available at `coverage-reports/bazel/index.html`.

## More Information

For more details on code coverage, see the [coverage documentation](../DOCS/COVERAGE.md).
