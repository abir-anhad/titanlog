````markdown
# TitanLog

TitanLog is a simple, industry-level structured logging library for Go.  
It supports log levels, structured data (fields), and immutable context loggers, making it ideal for building observability into modern applications.

TitanLog is designed with clarity, safety, and composability in mind, following idiomatic Go patterns.

---

## Features

- **Leveled Logging**  
  Supports `Debug`, `Info`, `Warn`, `Error`, and `Fatal` log levels.

- **Structured Data (Fields)**  
  Attach key-value pairs to logs for better context and machine-friendly parsing.

- **Immutable Context Loggers**  
  `WithFields` creates a derived logger without mutating the original one, ensuring thread safety and preventing side effects.

- **Custom Output**  
  Write logs to `os.Stdout`, files, or any `io.Writer`.

---

## Installation

```bash
go get github.com/abir-anhad/titanlog
````

---

## Usage

### Basic Logging

```go
package main

import (
    "os"

    "github.com/abir-anhad/titanlog"
)

func main() {
    // Initialize logger with Info level and write to Stdout
    log := titanlog.New(titanlog.InfoLevel, os.Stdout)

    log.Info("Application started")
    log.Warn("Configuration file missing, using defaults")
    log.Debug("This will be ignored because threshold is Info")
}
```

---

### Structured Logging

Add context to your logs using `WithFields`.
This creates a derived logger that retains the fields for all subsequent calls.

```go
package main

import (
    "os"

    "github.com/abir-anhad/titanlog"
)

func main() {
    log := titanlog.New(titanlog.InfoLevel, os.Stdout)

    // Create a context-aware logger (e.g., for a specific request)
    // The original 'log' remains unchanged.
    reqLogger := log.WithFields(titanlog.Fields{
        "request_id": "req-12345",
        "ip":         "192.168.1.1",
    })

    reqLogger.Info("Handling request")
    // Example output:
    // 2023-10-01T12:00:00Z INFO request_id=req-12345 ip=192.168.1.1 message="Handling request"

    reqLogger.Error("Database connection failed")
    // Example output:
    // 2023-10-01T12:00:00Z ERROR request_id=req-12345 ip=192.168.1.1 message="Database connection failed"
}
```

---

## Design Principles

* Explicit over magic
* Immutable by default
* Interfaces for extensibility
* Minimal API surface

---

## License

MIT License

```

