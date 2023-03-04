package presenter

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
	"github.com/google/uuid"
)

type presenter struct {
	Usecase  *usecases.Usecase
	Listener net.Listener
}

func NewPresenter(uc *usecases.Usecase) (*presenter, error) {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		return nil, err
	}
	return &presenter{
		Usecase:  uc,
		Listener: l,
	}, nil
}

func (p *presenter) Start() error {
	defer p.Listener.Close()
	return p.handlerConnections()
}

func (p *presenter) handlerConnections() error {
	for {
		conn, err := p.Listener.Accept()
		if err != nil {
			return err
		}
		if conn == nil {
			return errors.New("connection is null")
		}
		uuid := uuid.New().String()
		id, err := p.Usecase.AddNewClient(uuid, conn.RemoteAddr().String())
		if err != nil {
			log.Printf("LOG - [create-id-error]: %v", err.Error())
		}
		go p.HandlerConnection(conn, id)
	}
}

func (p *presenter) HandlerConnection(conn net.Conn, id string) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	fmt.Fprintf(conn, "Welcome Client Your ID is %v\n", id)
	log.Printf("LOG - [client]: new client connected\n")
	idPayload, err := r.ReadString('\n')
	if err != nil {
		log.Printf("LOG - [client-id-error]: %v\n", err.Error())
	}
	newID := usecases.RemountPayload([]byte(idPayload))
	actionPayload, err := r.ReadString('\n')
	if err != nil {
		log.Printf("LOG - [client-action-error]: %v\n", err.Error())
	}
	newAction := usecases.RemountPayload([]byte(actionPayload))
	_, err = r.ReadString('\n')
	if err != nil {
		log.Printf("LOG - [client-body-error]: %v\n", err.Error())
	}
	switch string(newAction) {
	case "LIST":
		idExists := p.Usecase.FindClientByID(string(newID))
		if !idExists {
			fmt.Fprintf(conn, "id incorrect")
		} else {
			clients := p.Usecase.ListAllClientsID()
			for n, id := range clients {
				fmt.Fprintf(conn, "client %v -> id: %v\n", n, id)
			}
		}
	}
}
