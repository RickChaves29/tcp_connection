package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatalf("LOG - [error]: %v", err.Error())
	}
	defer conn.Close()
	c := make(chan string)
	go func(c chan string) {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Printf("LOG - [io-error]: %v\n", err.Error())
		}

		c <- "receive payload\n"
	}(c)
	go func(c chan string) {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			log.Printf("LOG - [io-error]: %v\n", err.Error())
		}

		c <- "send payload\n"
	}(c)

	fmt.Printf(<-c)

}
