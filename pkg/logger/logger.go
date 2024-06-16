package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

const PKG_NAME = "Logger PKG"
const PKG_VERSION = "1.0.0"

var log Logger

type Logger struct{
	*zerolog.Logger
}

// NewLogger creates a new Logger instance with the provided zerolog.Logger.
//
// Parameters:
// - logger: A pointer to a zerolog.Logger instance.
//
// Returns:
// - A pointer to the newly created Logger instance.
func NewLogger(logLevel string, pkgName string, pkgVersion string) *Logger{
	logger := ConfigLogger(logLevel, pkgName, pkgVersion)
	log = Logger{&logger}
	return &log
}

// ConfigLogger configures the logger with the specified log level, package name, and package version.
//
// Parameters:
// - logLevel: The log level to set for the logger. Valid values are "debug", "info", "warn", and "error".
// - pkgName: The name of the package.
// - pkgVersion: The version of the package.
//
// Returns:
// - zerolog.Logger: The configured logger.
func ConfigLogger(logLevel string, pkgName string, pkgVersion string) zerolog.Logger {
	var zlogLevel zerolog.Level
	switch logLevel {
	case "debug":
		zlogLevel = zerolog.DebugLevel
	case "info":
		zlogLevel = zerolog.InfoLevel
	case "warn":
		zlogLevel = zerolog.WarnLevel
	case "error":
		zlogLevel = zerolog.ErrorLevel
	default:
		zlogLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(zlogLevel)
	zl := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		    With().
			Timestamp().
			Str("app", fmt.Sprintf("%s v%s", pkgName, pkgVersion)).
			Logger()

	return zl
}

// GetLogger returns a pointer to the Logger instance.
//
// Returns:
// - *Logger: A pointer to the Logger instance.
func GetLogger() *Logger {
	return &log
}