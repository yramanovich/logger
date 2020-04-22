package main

import (
	"logger"
	"os"
)

func main() {
	log := logger.New(
		logger.SetLogLevel(logger.LevelDebug),
		logger.SetFormatter(logger.JSONFormatter{Pretty: true}),
		logger.SetWriter(os.Stdout),
		)

	log.Info("msg", "info message")
	log.Warn("msg", "warn message")
	log.Debug("msg", "debug message")
	log.Error("msg", "error message")
}
