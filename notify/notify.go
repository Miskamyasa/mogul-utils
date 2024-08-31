package notify

import (
	"log"
	"os"
)

func Send(msg string, err error) {
	// TODO: send admin notification
	log.Printf("%s: %v", msg, err)
}

func Fatal(msg string, err error) {
	Send(msg, err)
	os.Exit(1)
}
