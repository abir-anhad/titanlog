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
