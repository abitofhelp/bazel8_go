# Packages

This directory contains reusable packages for the Bazel8_Go project.

## Overview

The packages in this directory are designed to be modular, reusable components that implement specific functionality for the project. They follow clean architecture principles and are thoroughly tested.

## Available Packages

### Greeting

The [greeting](./greeting/README.md) package provides functionality for generating personalized greeting messages.

### Logger

The [logger](./logger/README.md) package provides logging utilities for the application, with a focus on context-aware logging.

## Usage

Each package has its own README.md file with detailed information on how to use it. Please refer to the individual package documentation for specific usage instructions.

## Development Guidelines

When developing new packages:

1. Follow the project's architectural principles (DDD, Clean Architecture, Hexagonal Architecture)
2. Ensure thorough documentation with godoc comments
3. Implement comprehensive unit tests with a minimum of 80% coverage
4. Create a README.md file for the package
5. Keep packages focused on a single responsibility
