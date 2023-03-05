package presenter_test

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
	"github.com/RickChaves29/tcp_service/server/presenter"
)

var uc = usecases.NewUsecase([]usecases.ClientEntity{})

func init() {
	server, err := presenter.NewPresenter(uc)
	if err != nil {
		log.Printf("LOG - [test-start-error]: %v", err)
		return
	}
	go func() {
		for {
			conn, err := server.Listener.Accept()
			if err != nil {
				log.Printf("LOG - [test-listener-error]: %v", err)

			}
			if conn == nil {
				log.Printf("LOG - [test-connction-error]: %v", err)
			}
			id, err := server.Usecase.AddNewClient("9e32bfb4-913a-472a-90d1-e4b4da3e09af", conn.RemoteAddr().String())
			if err != nil {
				log.Printf("LOG - [create-id-error]: %v", err.Error())
			}
			connections := make(map[string]net.Conn)
			connections[id] = conn
			server.Connections = connections
			go server.HandlerConnection(conn, id)
		}
	}()
}

func TestIfConnectionIsOK(t *testing.T) {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		t.Error("could not connected to server: ", err.Error())
	}
	defer conn.Close()
}

func TestIfReturnListOfCLientIsCorrect(t *testing.T) {
	testCasesWant := struct {
		msgServer string
		listWant  string
	}{
		msgServer: "Welcome Client Your ID is 9e32bfb4-913a-472a-90d1-e4b4da3e09af\n",
		listWant:  "client 0 -> id: 9e32bfb4-913a-472a-90d1-e4b4da3e09af\n",
	}
	conn, _ := net.Dial("tcp", ":3000")
	r := bufio.NewReader(conn)
	msgServerHave, _ := r.ReadString('\n')
	fmt.Fprintf(conn, "9e32bfb4-913a-472a-90d1-e4b4da3e09af\n")
	fmt.Fprintf(conn, "LIST\n")
	fmt.Fprintf(conn, "\n")
	for range uc.Repository {
		clients, _ := r.ReadString('\n')
		t.Run("if message welcome is correct receive", func(t *testing.T) {
			if strings.EqualFold(msgServerHave, testCasesWant.msgServer) == false {
				t.Errorf("have %v want %v\n", msgServerHave, testCasesWant.msgServer)
			}
		})
		t.Run("if receive list of clients id", func(t *testing.T) {
			if clients != testCasesWant.listWant {
				t.Errorf("have %v want %v", clients, testCasesWant.listWant)
			}
		})
	}
}

func TestIfReturnBodyWhenClientSendActionRelay(t *testing.T) {
	testCasesWant := struct {
		msgServer string
		bodyWant  string
	}{
		msgServer: "Welcome Client Your ID is 9e32bfb4-913a-472a-90d1-e4b4da3e09af\n",
		bodyWant:  "body receive -> any data\n",
	}
	conn, _ := net.Dial("tcp", ":3000")
	r := bufio.NewReader(conn)
	_, _ = r.ReadString('\n')
	fmt.Fprintf(conn, "9e32bfb4-913a-472a-90d1-e4b4da3e09af\n")
	fmt.Fprintf(conn, "RELAY\n")
	fmt.Fprintf(conn, "any data\n")
	bodyHave, _ := r.ReadString('\n')
	t.Run("if body is correct receive", func(t *testing.T) {
		if strings.EqualFold(bodyHave, testCasesWant.bodyWant) == false {
			t.Errorf("have %v want %v\n", bodyHave, testCasesWant.bodyWant)
		}
	})
}
