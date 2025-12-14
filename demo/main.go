package main

import (
	"os"

	"github.com/abir-anhad/titanlog"
)

func main() {
	logger := titanlog.New(titanlog.InfoLevel, os.Stdout)

	logger.Debug("This should be ignored")
	logger.Info("This is an info message")
	logger.Error("Something went wrong")
}
