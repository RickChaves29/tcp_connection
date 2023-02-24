package main

import (
	"log"
	"net"

	"github.com/google/uuid"
)

func main() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		id := uuid.New().String()
		conn.Write([]byte(id))
	}
}
