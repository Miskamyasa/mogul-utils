package alerts

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var once sync.Once
var log zerolog.Logger

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
	once.Do(func() {
		var writer io.Writer = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		// if os.Getenv("ENV") == "production" {
		// 	writer = os.Stderr
		// }
		log = zerolog.New(writer).
			With().
			Timestamp().
			Str("service", os.Getenv("SERVICE_NAME")).
			Str("version", os.Getenv("SERVICE_VERSION")).
			Str("env", os.Getenv("ENV")).
			Logger()
	})

	return log
}
