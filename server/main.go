package main

import (
	"log"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
	"github.com/RickChaves29/tcp_service/server/presenter"
)

func main() {
	uc := usecases.NewUsecase([]usecases.ClientEntity{})
	server, err := presenter.NewPresenter(uc)
	if err != nil {
		log.Printf("LOG - [test-start-error]: %v", err)
		return
	}
	for {
		server.Start()
	}
}
