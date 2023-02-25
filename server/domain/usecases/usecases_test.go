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
