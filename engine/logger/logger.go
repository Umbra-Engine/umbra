package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

// Log level constants for controlling logger verbosity.
const (
	LogLevelError = iota // Only for error messages.
	LogLevelWarn         // Warnings and errors.
	LogLevelInfo         // General informational messages.
	LogLevelDebug        // Debugging messages.
	LogLevelTrace        // Highly detailed trace messages.
)

var currentLevel = LogLevelInfo

// init sets default logger output and disables timestamp flags (we use custom timestamps).
func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

func timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Info logs a message at the INFO level.
func Info(format string, args ...any) {
	if currentLevel >= LogLevelInfo {
		prefix := color.New(color.FgCyan).Sprint("[INFO]")
		log.Printf("%s %s %s", timestamp(), prefix, fmt.Sprintf(format, args...))
	}
}

// Warn logs a message at the WARN level.
func Warn(format string, args ...any) {
	if currentLevel >= LogLevelWarn {
		prefix := color.New(color.FgYellow).Sprint("[WARN]")
		log.Printf("%s %s %s", timestamp(), prefix, fmt.Sprintf(format, args...))
	}
}

// Error logs a message at the ERROR level.
func Error(format string, args ...any) {
	if currentLevel >= LogLevelError {
		prefix := color.New(color.FgMagenta).Sprint("[ERROR]")
		log.Printf("%s %s %s", timestamp(), prefix, fmt.Sprintf(format, args...))
	}
}

// Debug logs a message at the DEBUG level. Useful for development diagnostics.
func Debug(format string, args ...any) {
	if currentLevel >= LogLevelDebug {
		prefix := color.New(color.FgGreen).Sprint("[DEBUG]")
		log.Printf("%s %s %s", timestamp(), prefix, fmt.Sprintf(format, args...))
	}
}

// Trace logs a message at the TRACE level. Use for fine-grained, verbose output.
func Trace(format string, args ...any) {
	if currentLevel >= LogLevelTrace {
		prefix := color.New(color.FgHiBlack).Sprint("[TRACE]")
		log.Printf("%s %s %s", timestamp(), prefix, fmt.Sprintf(format, args...))
	}
}

// Fatal logs a message with [FATAL] and immediately terminates the program with os.Exit(1).
func Fatal(format string, args ...any) {
	prefix := color.New(color.FgRed, color.Bold).Sprint("[FATAL]")
	log.Printf("%s %s %s", timestamp(), prefix, fmt.Sprintf(format, args...))
	os.Exit(1)
}

// SetLevel sets the current logging verbosity level.
// Accepted values: LogLevelError, LogLevelWarn, LogLevelInfo, LogLevelDebug, LogLevelTrace.
func SetLevel(level int) {
	currentLevel = level
}

// GetLevel returns the current logging verbosity level.
func GetLevel() int {
	return currentLevel
}
