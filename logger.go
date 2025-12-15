// Package titanlog provides a simple, leveled structured logging system.
// It supports log levels, custom outputs, and structured fields.
package titanlog

import (
	"fmt"
	"io"
	"maps"
	"sync"
)

// Fields is a type alias for a map of key-value pairs used in structured logging.
// Keys should be strings, and values can be of any type.
type Fields map[string]interface{}

// Logger is the main struct that holds the configuration for logging.
// It is safe for concurrent use.
type Logger struct {
	threshold Level
	output    io.Writer
	fields    Fields
	formatter Formatter
	mu        sync.Mutex // Ensures thread safety
}

// New creates a new Logger instance.
// The threshold determines the minimum log level to output (e.g., if set to InfoLevel, DebugLevel logs are ignored).
// The output parameter specifies where logs should be written (e.g., os.Stdout or a file).
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
		fields:    make(Fields),
		formatter: &TextFormatter{},
	}
}

// SetFormatter allows changing the logging format (e.g., JSON or Text) at runtime.
func (l *Logger) SetFormatter(f Formatter) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.formatter = f
}

// WithFields creates a new Logger instance with the provided fields added to the existing context.
//
// This method returns a copy of the logger; the original logger is not modified.
// This allows for creating context-specific loggers (e.g., per-request or per-user)
// without affecting the global logger.
func (l *Logger) WithFields(f Fields) *Logger {
	newFields := make(Fields)

	// Copy existing fields from the parent logger to the new map
	maps.Copy(newFields, l.fields)

	// Add the new fields to the new map
	maps.Copy(newFields, f)

	// Note: We don't copy the mutex. The new logger gets a fresh, unlocked mutex.
	return &Logger{
		threshold: l.threshold,
		output:    l.output,
		fields:    newFields,
		formatter: l.formatter,
	}
}

// log is a private helper that formats and writes the log message to the output.
func (l *Logger) log(lvl Level, message string) {
	// 1. Check threshold BEFORE locking.
	// This is a performance optimization. If the log is ignored, we don't pay the cost of locking.
	if lvl < l.threshold {
		return
	}

	// 2. Lock the mutex to ensure no other goroutine writes at the same time.
	l.mu.Lock()
	// 3. Defer the Unlock.
	// This ensures that l.mu.Unlock() is ALWAYS called when this function exits,
	// even if we return early due to an error.
	defer l.mu.Unlock()

	// Use the formatter
	serialized, err := l.formatter.Format(lvl, message, l.fields)
	if err != nil {
		fmt.Printf("Failed to format log: %v\n", err)
		return
	}

	// Write the bytes to the output
	l.output.Write(serialized)
}

// Debug logs a message at DebugLevel.
// These are typically used for verbose output during development.
func (l *Logger) Debug(message string) {
	l.log(DebugLevel, message)
}

// Info logs a message at InfoLevel.
// These are used for standard operational events.
func (l *Logger) Info(message string) {
	l.log(InfoLevel, message)
}

// Warn logs a message at WarnLevel.
// These indicate non-critical issues that should be reviewed.
func (l *Logger) Warn(message string) {
	l.log(WarnLevel, message)
}

// Error logs a message at ErrorLevel.
// These indicate runtime errors that require attention.
func (l *Logger) Error(message string) {
	l.log(ErrorLevel, message)
}

// Fatal logs a message at FatalLevel.
// These indicate severe errors that may cause the application to crash or become unusable.
func (l *Logger) Fatal(message string) {
	l.log(FatalLevel, message)
}
