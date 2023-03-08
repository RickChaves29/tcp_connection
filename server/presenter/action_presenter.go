package presenter

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
	"github.com/google/uuid"
)

type presenter struct {
	Usecase     *usecases.Usecase
	Listener    net.Listener
	Connections map[string]net.Conn
}

func NewPresenter(uc *usecases.Usecase) (*presenter, error) {
	l, err := net.Listen("tcp", ":"+os.Getenv("SERVER_PORT"))
	log.Printf("LOG - [server-start]: running on port -> %v", os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, err
	}
	return &presenter{
		Usecase:     uc,
		Listener:    l,
		Connections: make(map[string]net.Conn),
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
		p.Connections[id] = conn
		go p.HandlerConnection(conn, id)
	}
}

func (p *presenter) HandlerConnection(conn net.Conn, id string) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	fmt.Fprintf(conn, "Welcome Client Your ID is %v\n", id)
	log.Printf("LOG - [client]: new client connected\n")
	fmt.Fprintf(conn, "ID [your id] (require): ")
	idPayload, err := r.ReadString('\n')
	if err != nil {
		log.Printf("LOG - [client-id-error]: %v\n", err.Error())
	}
	newID := usecases.RemountPayload([]byte(idPayload))
	fmt.Fprintf(conn, "ACTION [LIST|RELAY] (require): ")
	actionPayload, err := r.ReadString('\n')
	if err != nil {
		log.Printf("LOG - [client-action-error]: %v\n", err.Error())
	}
	newAction := usecases.RemountPayload([]byte(actionPayload))
	fmt.Fprintf(conn, "BODY [any data] (optional): ")
	bodyPayload, err := r.ReadString('\n')
	if err != nil {
		log.Printf("LOG - [client-body-error]: %v\n", err.Error())
	}
	body := usecases.RemountPayload([]byte(bodyPayload))
	switch string(newAction) {
	case "LIST":
		idExists := p.Usecase.FindClientByID(string(newID))
		if !idExists {
			fmt.Fprintf(conn, "id incorrect")
		} else {
			clients := p.Usecase.ListAllClientsID()
			for n, id := range clients {
				fmt.Fprintf(conn, "\nclient %v -> id: %v\n", n, id)
			}
		}
	case "RELAY":
		idExists := p.Usecase.FindClientByID(string(newID))
		if !idExists {
			fmt.Fprintf(conn, "id incorrect")
		} else {
			for _, connection := range p.Connections {
				fmt.Fprintf(connection, "\nbody receive -> %v\n", string(body))
			}
		}

	}
}
