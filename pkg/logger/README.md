# Logger v1.0.0: A Simple Zerolog Wrapper for Package-Level Logging

This package provides a convenient wrapper for the popular Zerolog logging library, enabling you to create a single logger instance for your package and access it from various modules within the package.

## Features:

-   **Package-Level Logging**: Initializes a logger with your package name and version for clear identification.
-   **Centralized Configuration**: Sets the log level based on a provided string ("debug", "info", "warn", or "error"). Defaults to debug for development convenience.
-   **Easy Access**: Offers a `GetLogger()` function to retrieve the configured logger instance from anywhere within your package.
-   **Zerolog Integration**: Leverages the power and features of Zerolog for structured logging.

## Installation:

Use `go get` to install the package:

```bash
    go get github.com/frangdelsolar/go-games/pkg/logger
```

## Import the package in your Go files:

```go
    import (
        "github.com/frangdelsolar/go-games/pkg/logger"
    )
```

## Usage:

Create a new logger instance at the beginning of your package (e.g., init.go or a dedicated file):

```go
    func init() {
        logger.NewLogger("debug", "MyPackageName", "1.0.0") // Replace with your package name and version
    }
```

Access the logger from different modules within your package using `logger.GetLogger()`:

```go
    func SomeFunction() {
        log := logger.GetLogger()
        log.Info().Msg("This is an informational message")
    }
```

## Example:

Refer to the included main.go file for a demonstration of using the go-logger package.

## Benefits:

-   **Consistency**: Ensures all logs within your package share the same structure and configuration.
-   **Readability**: Makes logging calls more explicit by utilizing the package-specific logger instance.
-   **Maintainability**: Simplifies logger setup and avoids potential redundancy.

## Customization:

You can modify the ConfigLogger function to customize the log output format (e.g., JSON instead of console) by adjusting the Zerolog configuration options.
Consider adding support for environment variables to control the log level dynamically.
