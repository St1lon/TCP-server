package main

import (
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	log.Println("Starting TCPserver")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	handleErrors(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		handleErrors(err)
		go handleRequests(conn)
	}
}

func handleRequests(conn net.Conn) {
	defer conn.Close()

	log.Println("New client connected")

	// Небольшая задержка, чтобы убедиться, что клиент готов
	responseString := "OK\n"
	_, err := conn.Write([]byte(responseString))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
		return
	}

	log.Println("Response sent to client, closing connection")
}

func handleClientErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}
func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
