package titanlog

import (
	"encoding/json"
	"fmt"
	"maps"
	"time"
)

// Formatter defines the interface that all log formatters must implement.
type Formatter interface {
	Format(level Level, msg string, fields Fields) ([]byte, error)
}

// JSONFormatter formats logs as a JSON object.
type JSONFormatter struct{}

func (f *JSONFormatter) Format(level Level, msg string, fields Fields) ([]byte, error) {
	data := make(Fields)

	// 1. Copy user fields
	maps.Copy(data, fields)

	// 2. Add system fields
	data["level"] = level.String()
	data["time"] = time.Now().Format(time.RFC3339)
	data["msg"] = msg

	// 3. Serialize to JSON bytes
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}

	// 4. Append newline
	return append(serialized, '\n'), nil
}

// TextFormatter formats logs as "TIME LEVEL key=value msg"
type TextFormatter struct{}

func (f *TextFormatter) Format(level Level, msg string, fields Fields) ([]byte, error) {
	fieldsString := ""

	for k, v := range fields {
		// FIXED: Used %s for the key (string) instead of %k
		fieldsString += fmt.Sprintf("%s=%v ", k, v)
	}

	timestamp := time.Now().Format(time.RFC3339)

	// Combine into final string
	logLine := fmt.Sprintf("%s: %s - %smessage: %s\n", timestamp, level.String(), fieldsString, msg)

	// Return bytes
	return []byte(logLine), nil
}
