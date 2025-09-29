package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	// Create a logger with INFO level
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	for {
		logger.Info("Application start up")
		time.Sleep(2 * time.Second)
	}
}
