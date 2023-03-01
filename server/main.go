package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
	"github.com/RickChaves29/tcp_service/server/presenter"
	"github.com/google/uuid"
)

var connections = make(map[string]net.Conn)

func main() {
	PORT := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	repository := []usecases.ClientEntity{}
	uc := usecases.NewUsecase(repository)

	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("LOG - [ERROR]: %v\n", err.Error())
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("LOG - [ERROR]: %v\n", err.Error())
		}
		id := uuid.New().String()
		connections[id] = conn
		clientHost := conn.RemoteAddr().String()
		clientID, err := uc.AddNewClient(id, clientHost)
		log.Printf("LOG - [CONNECTED] host: %s client_id: %s\n", clientHost, clientID)
		if err != nil {
			log.Printf("LOG - [ERROR]: %v\n", err.Error())
		}
		serverMessage := fmt.Sprintf("Welcome your id is %s\n", clientID)
		conn.Write([]byte(serverMessage))
		ap := presenter.NewActionPresenter(conn, uc)
		go ap.SetAction(connections)
	}

}
