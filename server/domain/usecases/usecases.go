package usecases

import (
	"bytes"
	"errors"
	"log"
	"regexp"
)

type ClientEntity struct {
	ID       string
	HostName string
}
type Usecase struct {
	Repository []ClientEntity
}

func NewUsecase(repo []ClientEntity) *Usecase {
	return &Usecase{
		Repository: repo,
	}
}
func RemountPayload(payload []byte) []byte {
	rx, err := regexp.Compile(`^[\s]+|[\s]+$`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	newPayload := rx.ReplaceAll(bytes.Trim(payload, "\x00"), []byte(""))

	return newPayload
}

func (us *Usecase) AddNewClient(id string, clientHost string) (string, error) {
	if clientHost == "" {
		return "", errors.New("client host is empty")
	}
	us.Repository = append(us.Repository, ClientEntity{id, clientHost})
	return id, nil
}
