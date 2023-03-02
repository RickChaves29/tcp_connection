package main

import (
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":"+os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("LOG - [ERROR]: %v", err.Error())
	}
	defer conn.Close()
}
