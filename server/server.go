package main

import (
	"log"
	"net"
)

func main() {
	// Слушаем порт 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()
	
	log.Println("Server listening on port 8080")
	
	// Бесконечный цикл для принятия подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		
		// Обрабатываем подключение в горутине
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	// Отправляем "OK\n" клиенту
	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
		return
	}
	
	// Логируем успешную обработку
	remoteAddr := conn.RemoteAddr().String()
	log.Printf("Sent 'OK\\n' to client %s and closed connection", remoteAddr)
}