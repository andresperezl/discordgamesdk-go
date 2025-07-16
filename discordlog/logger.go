package discordlog

import (
	"log/slog"
	"sync"
)

var (
	loggerMu sync.RWMutex
	logger   *slog.Logger
)

// SetLogger allows users to set a custom slog.Logger for SDK logging
func SetLogger(l *slog.Logger) {
	loggerMu.Lock()
	defer loggerMu.Unlock()
	logger = l
}

// GetLogger returns the current logger, or a no-op logger if none is set
func GetLogger() *slog.Logger {
	loggerMu.RLock()
	defer loggerMu.RUnlock()
	if logger != nil {
		return logger
	}
	return slog.New(slog.NewTextHandler(nil, &slog.HandlerOptions{Level: slog.LevelError})) // no-op
}
