package main

import (
	"fmt"
	"net"
	"os"
)

// socket-server project main.go

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9090"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running")

	// TCP 소켓을 생성하고 커넥션을 기다리도록(listen 설정)
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	// buffer에 값 읽어오기
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))

	// client에게 응답 메시지 적어주기
	_, err = connection.Write([]byte("Thanks! Got your message: " + string(buffer[:mLen])))
	connection.Close()
}
