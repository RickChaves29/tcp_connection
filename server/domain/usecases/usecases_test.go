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
