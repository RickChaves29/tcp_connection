package usecases

import (
	"bytes"
	"log"
	"regexp"
)

type ClientEntity struct {
	ID       string
	HostName string
}

func RemountPayload(payload []byte) []byte {
	rx, err := regexp.Compile(`^[\s]+|[\s]+$`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	newPayload := rx.ReplaceAll(bytes.Trim(payload, "\x00"), []byte(""))

	return newPayload
}
