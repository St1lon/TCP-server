package main

import (
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	TYPE = "tcp"
	PORT = "8080"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Println("Unable to resolve addres " + err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Printf("Unable to make a connection to the Server: %v\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close() // Закрыть соединение при завершении функции

	recievedBuffer := make([]byte, 1024)
	n, err := conn.Read(recievedBuffer)
	if err != nil {
		log.Printf("Could not receive data sent from server: %v", err.Error())
	}
	if string(recievedBuffer[:n]) != "OK" {
		log.Print("The message is not OK")
	}
	log.Printf("Recieved Message: %v", string(recievedBuffer))
}
