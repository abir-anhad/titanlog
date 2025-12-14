package titanlog

// Level represents the severity of a log message.
// Higher values indicate more severe events.
type Level int

const (
	// DebugLevel is for verbose output, useful for developers.
	DebugLevel Level = iota
	// InfoLevel is for standard operational messages.
	InfoLevel
	// WarnLevel is for non-critical issues that should be looked at.
	WarnLevel
	// ErrorLevel is for runtime errors that require attention.
	ErrorLevel
	// FatalLevel is for severe errors that may cause the application to crash.
	FatalLevel
)

// String returns the string representation of the Level (e.g., "INFO").
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
