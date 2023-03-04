package usecases_test

import (
	"testing"

	"github.com/RickChaves29/tcp_service/server/domain/usecases"
)

func TestShoudReturnRemountPayloadWithCorrectID(t *testing.T) {
	data := []byte("         ab33bc35-a29a-49df-b1d6-75e1366fee6e           \x00\x00\x00\x00")
	resultWant := []byte("ab33bc35-a29a-49df-b1d6-75e1366fee6e")
	resultHave := usecases.RemountPayload(data)
	if string(resultHave) != string(resultWant) {
		t.Errorf("result have %v but want %v", resultHave, resultWant)
	}
}
func TestShoudReturnRemountPayloadWithCorrectAction(t *testing.T) {
	data := []byte("         LIST           \x00\x00\x00\x00")
	resultWant := []byte("LIST")
	resultHave := usecases.RemountPayload(data)
	if string(resultHave) != string(resultWant) {
		t.Errorf("result have %v but want %v", resultHave, resultWant)
	}
}

func TestShoudReturnRemountPayloadWithCorrectBody(t *testing.T) {
	data := []byte("{\"name\": \"any\"}      \x00\x00\x00\x00")
	resultWant := []byte(`{"name": "any"}`)
	resultHave := usecases.RemountPayload(data)
	if string(resultHave) != string(resultWant) {
		t.Errorf("result have %v but want %v", string(resultHave), string(resultWant))
	}
}

func TestShoudAddNewClientAndReturnID(t *testing.T) {
	repoTest := []usecases.ClientEntity{}
	uc := usecases.NewUsecase(repoTest)
	resultWant := "2f92634b-9eeb-4b84-8970-c47be1625952"
	resultHave, _ := uc.AddNewClient("2f92634b-9eeb-4b84-8970-c47be1625952", "localhost:40443")
	if resultHave != resultWant && len(uc.Repository) > 0 {
		t.Errorf("id have is %v but id want is %v", resultHave, resultWant)
	}
}
func TestShoudNotAddNewClientAndReturnIDAndError(t *testing.T) {
	repoTest := []usecases.ClientEntity{}
	uc := usecases.NewUsecase(repoTest)
	resultWant := ""
	resultHave, err := uc.AddNewClient("2f92634b-9eeb-4b84-8970-c47be1625952", "")
	t.Run("if message error is correct", func(t *testing.T) {
		if err.Error() != "client host is empty" {
			t.Errorf("error have is %v but error want is %v", err.Error(), "client host is empty")
		}
	})
	t.Run("if result is empty", func(t *testing.T) {
		if resultHave != resultWant {
			t.Errorf("result have is %v but result want is %v", resultHave, resultWant)
		}
	})
	t.Run("if repository lenght is 0", func(t *testing.T) {
		if len(uc.Repository) > 0 {
			t.Errorf("repository lenght have is %v but repository lenght want is %v", len(uc.Repository), 0)
		}
	})
}

func TestShoudListAllClientsIDWithError(t *testing.T) {
	repoTest := []usecases.ClientEntity{
		{
			ID:       "2f92634b-9eeb-4b84-8970-c47be1625952",
			HostName: "localhost:40443",
		},
		{
			ID:       "6d976bb0-9afd-4f39-bfcb-2450fce55200",
			HostName: "localhost:40444",
		},
		{
			ID:       "a08a59bf-a53d-4489-9c93-464f9ddc5928",
			HostName: "localhost:40445",
		},
		{
			ID:       "07704e40-f3ac-42fb-b2ca-03bd4f497e90",
			HostName: "localhost:40446",
		},
	}

	uc := usecases.NewUsecase(repoTest)

	t.Run("if list of clients receive is correct", func(t *testing.T) {
		clients := uc.ListAllClientsID()
		if len(clients) == 0 {
			t.Errorf("list client lenght have is %v but list client lenght want is 4", len(repoTest))
		}
	})
}
