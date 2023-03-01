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

func (uc *Usecase) AddNewClient(id string, clientHost string) (string, error) {
	if clientHost == "" {
		return "", errors.New("client host is empty")
	}
	uc.Repository = append(uc.Repository, ClientEntity{id, clientHost})
	return id, nil
}

func (uc *Usecase) ListAllClientsID(id string, action string) ([]string, error) {
	allClientsID := []string{}
	if id == "" {
		return nil, errors.New("id is empty")
	}
	if action == "" {
		return nil, errors.New("action is empty")
	}
	_, err := uc.FindClientByID(id)
	if err != nil {
		return nil, err
	}
	for _, client := range uc.Repository {
		allClientsID = append(allClientsID, client.ID)
	}
	return allClientsID, nil
}

func (uc *Usecase) FindClientByID(id string) (string, error) {
	var clientID string
	for _, client := range uc.Repository {
		if id == client.ID {
			clientID = client.ID
			break
		}
	}
	if clientID == "" {
		return "", errors.New("id not found")
	}
	return clientID, nil
}
