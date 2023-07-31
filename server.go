package main

import (
	"fmt"
	"net"
	"sync"
)

var clients = make(map[net.Conn]bool)
var mutex sync.Mutex

func main() {
	// Lắng nghe kết nối từ client trên cổng 10010
	listener, err := net.Listen("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 10010")

	for {
		// Chấp nhận kết nối từ client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		mutex.Lock()
		clients[conn] = true
		mutex.Unlock()

		fmt.Printf("Client connected: %s\n", conn.RemoteAddr())

		// Xử lý kết nối từ client trong một goroutine
		go handleClientConnection(conn)
	}
}

func handleClientConnection(conn net.Conn) {
	defer conn.Close()

	// Gửi tin nhắn chào mừng tới client
	conn.Write([]byte("Welcome to the chat server!\n"))

	// Xử lý việc nhận và gửi tin nhắn
	for {
		// Đọc tin nhắn từ client
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr())

			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()

			return
		}

		message := string(buf[:n])

		// Gửi tin nhắn từ client tới tất cả các client khác
		broadcastMessage(conn, message)
	}
}

func broadcastMessage(sender net.Conn, message string) {
	mutex.Lock()
	for client := range clients {
		if client != sender {
			client.Write([]byte(message))
		}
	}
	mutex.Unlock()
}
