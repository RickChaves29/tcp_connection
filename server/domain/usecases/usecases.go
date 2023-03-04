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
	rx, err := regexp.Compile(`^[\n\s]+|[\n\s]+$`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	newPayload := rx.ReplaceAll(bytes.Trim(payload, "\x00"), []byte(""))

	return newPayload
}

func (uc *Usecase) AddNewClient(id string, clientHost string) (string, error) {
	idRemounted := RemountPayload([]byte(id))
	if clientHost == "" {
		return "", errors.New("client host is empty")
	}
	uc.Repository = append(uc.Repository, ClientEntity{string(idRemounted), clientHost})
	return string(idRemounted), nil
}

func (uc *Usecase) ListAllClientsID() []string {
	allClientsID := []string{}
	for _, client := range uc.Repository {
		allClientsID = append(allClientsID, client.ID)
	}
	return allClientsID
}

func (uc *Usecase) FindClientByID(id string) bool {
	var clientID string
	for _, client := range uc.Repository {
		if id == client.ID {
			clientID = client.ID
			break
		}
	}
	if clientID != "" {
		return true
	} else {
		return false
	}
}
