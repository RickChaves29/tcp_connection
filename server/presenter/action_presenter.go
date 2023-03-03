package presenter

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
	"github.com/google/uuid"
)

type presenter struct {
	uc *usecases.Usecase
	l  net.Listener
}

func NewPresenter(uc *usecases.Usecase) (*presenter, error) {
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		return nil, err
	}
	return &presenter{
		uc: uc,
		l:  l,
	}, nil
}

func (p *presenter) Start() error {
	defer p.l.Close()
	return p.handlerConnections()
}

func (p *presenter) handlerConnections() error {
	conn, err := p.l.Accept()
	if err != nil {
		return err
	}
	if conn == nil {
		return errors.New("connection is null")
	}
	uuid := uuid.New().String()
	id, err := p.uc.AddNewClient(uuid, conn.RemoteAddr().String())
	if err != nil {
		log.Printf("LOG - [create-id-error]: %v", err.Error())
	}
	go p.handlerConnection(conn, id)
	return nil
}

func (p *presenter) handlerConnection(conn net.Conn, id string) {
	defer conn.Close()
	fmt.Fprintf(conn, "Welcome Client\n")
	fmt.Fprintf(conn, "Your ID is %v\n", id)
	log.Printf("LOG - [client]: new client connected\n")

}
