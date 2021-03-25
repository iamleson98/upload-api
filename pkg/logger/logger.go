package logger

import (
	"os"
	"time"

	"sync"

	"github.com/rs/zerolog"
)

var (
	once sync.Once
	// Logger represents zero logger
	Logger zerolog.Logger
)

func init() {
	// initialize Logger only once
	once.Do(func() {
		Logger = NewLogger()
	})
}

// NewLogger returns new zero logger
func NewLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return logger
}
