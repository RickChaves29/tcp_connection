package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":"+os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("LOG - [ERROR]: %v", err.Error())
	}
	defer conn.Close()

	readConnection := bufio.NewReader(conn)

	response, err := readConnection.ReadString('\n')
	if err != nil {
		if err != nil {
			log.Printf("LOG - [ERROR]: %v", err.Error())
		}
	}
	log.Println(strings.TrimSpace(response))
	var id string
	var action string
	var body string
	_, err = fmt.Scan(&id)
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}

	_, err = fmt.Scan(&action)
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}

	_, err = fmt.Scan(&body)
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}

	_, err = conn.Write([]byte(strings.TrimSpace(id)))
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}
	_, err = conn.Write([]byte(strings.TrimSpace(action)))
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}
	_, err = conn.Write([]byte(strings.TrimSpace(body)))
	if err != nil {
		log.Printf("LOG - [ERROR]: %v", err.Error())
	}

}
