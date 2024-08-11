package utils

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func Notify(msg string, err error) {
	// TODO: send admin notification
	log.Debugf("%s: %v", msg, err)
}

func Fatal(msg string, err error) {
	Notify(msg, err)
	os.Exit(1)
}
