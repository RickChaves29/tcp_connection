package presenter_test

import (
	"bufio"
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
		server.Start()
	}()
}

func TestIfConnectionIsOK(t *testing.T) {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		t.Error("could not connected to server: ", err.Error())
	}
	defer conn.Close()
}
func TestCaseConnection(t *testing.T) {
	tc := struct {
		test         string
		welcameWant  string
		clientIDWant string
		idWant       string
	}{
		test:         "if LIST return all clients id",
		welcameWant:  "Welcome Client\n",
		clientIDWant: "Your ID is 4a9e363f-6fd6-4b0e-bd96-e7d75e581b43\n",
		idWant:       "4a9e363f-6fd6-4b0e-bd96-e7d75e581b43\n",
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		t.Error("could not connected to server: ", err.Error())
	}
	defer conn.Close()
	r := bufio.NewReader(conn)
	welcomeHave, _ := r.ReadString('\n')
	if strings.Compare(welcomeHave, tc.welcameWant) > 0 {
		t.Errorf("welcome have %v welcome want %v\n", welcomeHave, tc.welcameWant)
	}

	idClientHave, _ := r.ReadString('\n')
	if len([]byte(idClientHave)) > len(tc.clientIDWant) {
		t.Errorf("id client have %v id client want %v\n", idClientHave, string(tc.clientIDWant))
	}
}
