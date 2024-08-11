package utils

import "log"

func Notify(msg string, err error) {
	// TODO: send admin notification
	log.Printf("%s: %v", msg, err)
}

func Fatal(msg string, err error) {
	Notify(msg, err)
	log.Fatalf("%s: %v", msg, err)
}
