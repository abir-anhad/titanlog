// Package titanlog provides a simple, leveled structured logging system.
// It supports log levels, custom outputs, and structured fields.
package titanlog

import (
	"fmt"
	"io"
	"time"
)

// Logger is the main struct that holds the configuration for logging.
// It is safe for concurrent use (we will add this later!).
type Logger struct {
	threshold Level
	output    io.Writer
}

// New creates a new Logger instance.
// The threshold determines the minimum log level to output.
// The output parameter specifies where logs should be written (e.g., os.Stdout).
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// log is a private helper that writes the log message to the output.
func (l *Logger) log(lvl Level, message string) {
	if lvl < l.threshold {
		return
	}

	timenow := time.Now().Format(time.RFC3339)
	// We use Fprintf to write to the specified output (file, stdout, etc.)
	fmt.Fprintf(l.output, "%v: %s - message: %s\n", timenow, lvl.String(), message)
}

// Debug logs a message at DebugLevel.
func (l *Logger) Debug(message string) {
	l.log(DebugLevel, message)
}

// Info logs a message at InfoLevel.
func (l *Logger) Info(message string) {
	l.log(InfoLevel, message)
}

// Warn logs a message at WarnLevel.
func (l *Logger) Warn(message string) {
	l.log(WarnLevel, message)
}

// Error logs a message at ErrorLevel.
func (l *Logger) Error(message string) {
	l.log(ErrorLevel, message)
}

// Fatal logs a message at FatalLevel.
func (l *Logger) Fatal(message string) {
	l.log(FatalLevel, message)
}
