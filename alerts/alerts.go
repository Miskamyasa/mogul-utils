package alerts

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		With().
		Timestamp().
		Logger()
)

func Send(msg string, err error) {
	logger := CreateLogger().With().Str("component", "alerts").Logger()
	if err != nil {
		logger.Err(err).Msg(msg)
		return
	}
	logger.Info().Msg(msg)
	// TODO: send admin notification
}

func Fatal(msg string, err error) {
	Send(msg, err)
	os.Exit(1)
}

func CreateLogger() zerolog.Logger {
	return log.
		With().
		Str("service", os.Getenv("SERVICE_NAME")).
		Str("version", os.Getenv("SERVICE_VERSION")).
		Str("env", os.Getenv("ENV")).
		Logger()
}
