package main

import (
	"os"

	"github.com/abir-anhad/titanlog"
)

func main() {
	// 1. Create logger
	logger := titanlog.New(titanlog.InfoLevel, os.Stdout)

	// 2. Switch to JSON Formatter
	logger.SetFormatter(&titanlog.JSONFormatter{})

	// 3. Log something with fields
	logger.WithFields(titanlog.Fields{
		"user_id":  12345,
		"action":   "payment",
		"currency": "USD",
	}).Info("Payment processed successfully")
}
