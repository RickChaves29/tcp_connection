package presenter

import (
	"log"
	"net"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
)

type ActionPresenter struct {
	conn    net.Conn
	usecase *usecases.Usecase
}

func NewActionPresenter(conn net.Conn, uc *usecases.Usecase) *ActionPresenter {
	return &ActionPresenter{
		conn:    conn,
		usecase: uc,
	}
}

func (ap *ActionPresenter) SetAction() {
	var (
		id, action string
	)
	var connections = make(map[string]net.Conn)

	idPayload := make([]byte, 1024)
	actionPayload := make([]byte, 1024)
	bodyPayload := make([]byte, 1024)
	_, err := ap.conn.Read(idPayload)
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}
	_, err = ap.conn.Read(actionPayload)
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}
	_, err = ap.conn.Read(bodyPayload)
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}
	id = string(usecases.RemountPayload(idPayload))
	action = string(usecases.RemountPayload(actionPayload))
	body := string(usecases.RemountPayload(bodyPayload))
	connections[id] = ap.conn

	switch action {
	case "LIST":
		clients, err := ap.usecase.ListAllClientsID(id, action)
		if err != nil {
			log.Printf("LOG - [ERROR]: %v", err.Error())
		}
		for _, client := range clients {
			_, err = ap.conn.Write([]byte("client -> " + client + "\n"))
			if err != nil {
				log.Printf("LOG - [ERROR]: %v", err.Error())
			}
		}
	case "RELAY":
		_, err := ap.usecase.ListAllClientsID(id, action)
		if err != nil {
			log.Printf("LOG - [ERROR]: %v", err.Error())
		}
		for _, conn := range connections {
			_, err := conn.Write([]byte(body))
			if err != nil {
				log.Printf("LOG - [ERROR]: %v", err.Error())
			}
		}
	default:
		_, err = ap.conn.Write([]byte("action " + action + " is incorrect or not exist\n"))
		if err != nil {
			log.Printf("LOG - [ERROR]: %v", err.Error())
		}
	}

}
