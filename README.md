[![Go Reference](https://pkg.go.dev/badge/github.com/abir-anhad/titanlog.svg)](https://pkg.go.dev/github.com/abir-anhad/titanlog)

# TitanLog

TitanLog is a high-performance, thread-safe, structured logging library for Go.
It provides a simple and idiomatic API for leveled logging with support for context-aware fields and pluggable output formatters (JSON, Text).

TitanLog is designed to be minimal, dependency-free, and safe for concurrent use in modern Go applications and pipelines.

---

## Features

* **Leveled Logging**
  Granular control with `Debug`, `Info`, `Warn`, `Error`, and `Fatal` levels.

* **Structured Data**
  Attach key-value pairs (Fields) to logs for machine-readable context.

* **Thread-Safe**
  Built-in concurrency safety using mutexes, making it safe across multiple goroutines.

* **Immutable Context**
  `WithFields` creates a shallow copy of the logger, ensuring parent loggers remain unchanged.

* **Pluggable Formatters**
  Switch between human-readable text and JSON format at runtime.

* **Zero Dependencies**
  Built entirely using the Go standard library.

---

## Installation

```bash
go get github.com/abir-anhad/titanlog
```

---

## Usage

### Basic Logging

Initialize the logger with a threshold level and an output destination.

```go
package main

import (
    "os"
    "github.com/abir-anhad/titanlog"
)

func main() {
    // Only logs at InfoLevel or higher will be printed
    log := titanlog.New(titanlog.InfoLevel, os.Stdout)

    log.Info("Service started")
    log.Warn("Config file missing, using defaults")

    // Ignored because the threshold is Info
    log.Debug("Debugging connection...")
}
```

---

### Structured Logging (Context)

Use `WithFields` to attach structured context to logs.
This is especially useful for tracing requests in production.

```go
func processRequest(reqID string) {
    log := titanlog.New(titanlog.InfoLevel, os.Stdout)

    // Create a context-specific logger
    reqLog := log.WithFields(titanlog.Fields{
        "request_id": reqID,
        "ip":         "192.168.1.50",
    })

    reqLog.Info("Processing payment")
    reqLog.Error("Payment gateway timeout")
}
```

**Example Output (Text Formatter):**

```
2025-12-15T10:00:00Z INFO  request_id=123 ip=192.168.1.50 message="Processing payment"
2025-12-15T10:00:05Z ERROR request_id=123 ip=192.168.1.50 message="Payment gateway timeout"
```

---

### JSON Formatting

For production environments such as AWS CloudWatch, Datadog, or ELK Stack, use JSON formatting.

```go
func main() {
    log := titanlog.New(titanlog.InfoLevel, os.Stdout)

    // Switch to JSON formatter
    log.SetFormatter(&titanlog.JSONFormatter{})

    log.WithFields(titanlog.Fields{
        "user": "admin",
        "id":   55,
    }).Info("User logged in")
}
```

**Output:**

```json
{
  "level": "INFO",
  "msg": "User logged in",
  "time": "2025-12-15T10:00:00Z",
  "user": "admin",
  "id": 55
}
```

---

### Concurrency Safety

TitanLog is safe for concurrent use.
A single logger instance can be shared across multiple goroutines.

```go
func main() {
    log := titanlog.New(titanlog.InfoLevel, os.Stdout)

    for i := 0; i < 10; i++ {
        go func(id int) {
            log.WithFields(titanlog.Fields{
                "worker_id": id,
            }).Info("Worker started")
        }(i)
    }
}
```

---

## Design Philosophy

* Simple, explicit API
* No hidden global state
* Immutable logger context
* Safe defaults for production
* Idiomatic Go patterns

---

## Contributing

Contributions are welcome.

* Fork the repository
* Create a feature branch
* Open a pull request

For major changes, please open an issue first to discuss your proposal.

---

## License

MIT License

---


