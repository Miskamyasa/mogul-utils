package utils

import "github.com/gofiber/fiber/v2/log"

func Notify(msg string, err error) {
	// TODO: send admin notification
	log.Debugf("%s: %v", msg, err)
}

func Fatal(msg string, err error) {
	Notify(msg, err)
	log.Fatalf("%s: %v", msg, err)
}
