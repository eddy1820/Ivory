package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Logger zerolog.Logger

func Init(logLevel string) {
	// Unix timestamp
	zerolog.TimeFieldFormat = time.RFC3339

	// log level (EX: debug, info, warn, error)
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.InfoLevel // fallback
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	Logger = zerolog.New(consoleWriter).
		Level(level).
		With().
		Timestamp().
		Caller().
		Logger()
}
