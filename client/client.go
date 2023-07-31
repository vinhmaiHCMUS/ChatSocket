package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Kết nối tới server trên cổng 8080
	conn, err := net.Dial("tcp", "app-server:10010")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Nhập tên định danh từ client
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	name = strings.TrimSpace(name)

	// Gửi tên định danh tới server
	_, err = conn.Write([]byte(name + "\n"))
	if err != nil {
		fmt.Println("Error sending name:", err)
		return
	}

	// Xử lý việc đọc và gửi tin nhắn
	go readMessages(conn)

	for {
		// Nhập tin nhắn từ client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Gửi tin nhắn tới server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func readMessages(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		fmt.Printf("%s", buf[:n])
	}
}
