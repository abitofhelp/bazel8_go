# Code Coverage with Bazel

This document explains how to generate code coverage reports for the project using Bazel.

## Prerequisites

- Bazel
- GCov (usually comes with GCC)
- LCOV (for generating HTML reports)

## Using the Coverage Script

We've provided a script that simplifies the process of generating coverage reports:

```bash
./tools/run_coverage.sh
```

This script:
1. Checks if GCov is installed
2. Checks if LCOV is installed, and installs it if not (using Homebrew on macOS or apt on Linux)
3. Runs the Bazel coverage command
4. Generates an HTML report using genhtml

The coverage report will be available at `coverage-reports/bazel/index.html`.

## Known Issues

You may see the following warning during the coverage generation:

```
GCov does not exist at the given path: ''
```

This warning is related to how Bazel tries to locate GCov. Despite this warning, the coverage report is still generated correctly.

## Manual Coverage Generation

If you prefer to run the commands manually:

1. Install LCOV if not already installed:
   ```bash
   # macOS
   brew install lcov

   # Ubuntu/Debian
   sudo apt-get install lcov
   ```

2. Run Bazel coverage:
   ```bash
   bazel coverage --test_output=all --combined_report=lcov //...
   ```

3. Generate the HTML report:
   ```bash
   mkdir -p coverage-reports
   genhtml -o coverage-reports/bazel "$(bazel info output_path)/_coverage/_coverage_report.dat"
   ```

## Troubleshooting

If you encounter issues:

1. Make sure GCov is installed and in your PATH
2. Make sure LCOV is installed and in your PATH
3. Check that you have the necessary permissions to create directories and files in the project
