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
	// Подключаемся к серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()
	
	log.Println("Connected to server successfully")
	
	// Читаем ответ от сервера
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	
	// Проверяем что получен ответ "OK\n"
	response = strings.TrimSpace(response)
	if response == "OK" {
		fmt.Println("SUCCESS: Received expected response 'OK' from server")
	} else {
		fmt.Printf("FAILURE: Expected 'OK' but received '%s'\n", response)
		os.Exit(1)
	}
	
	log.Println("Client finished successfully")
}