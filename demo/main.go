package main

import (
	"os"

	"github.com/abir-anhad/titanlog"
)

func main() {
	logger := titanlog.New(titanlog.InfoLevel, os.Stdout)

	requestLogger := logger.WithFields(titanlog.Fields{
		"request_id": "req-12345",
		"ip":         "192.168.1.1",
	})

	requestLogger.Info("Request started")
	requestLogger.WithFields(titanlog.Fields{"user_id": 101}).Info("User authenticated")

	// 4. Log using the base logger again to prove it has NO fields
	logger.Warn("This is a global warning (should have no fields)")
}
