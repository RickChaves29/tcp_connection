package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"regexp"

	"github.com/google/uuid"
)

type Client struct {
	ID      string
	Host    string
	Payload string
}

func main() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handlerConnection(conn)
	}

}

func handlerConnection(conn net.Conn) {
	defer conn.Close()
	payload := make([]byte, 1024)
	rx, err := regexp.Compile(`[\n\s]`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := uuid.New().String()
	fmt.Printf("new connection from host: %s id: %s\n", conn.RemoteAddr(), id)
	for {
		_, err := conn.Read(payload)
		if err != nil {
			log.Fatalln(err.Error())
		}
		data := rx.ReplaceAll(bytes.Trim(payload, "\x00"), []byte(""))
		if string(data) == "LIST" {
			fmt.Println("list data: ", string(data))
		} else if string(data) == "RELAY" {
			fmt.Println("replay data: ", string(data))
		} else {
			fmt.Println("any data: ", string(data))
		}
	}
}
