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
		Str("component", "alerts")
)

func Send(msg string, err error) {
	logger := GetLogger()
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

func GetLogger() zerolog.Logger {
	return log.
		Str("service", os.Getenv("SERVICE_NAME")).
		Str("version", os.Getenv("SERVICE_VERSION")).
		Str("env", os.Getenv("ENV")).
		Logger()
}
