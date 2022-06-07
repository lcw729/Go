package main

import (
	"fmt"
	"net"
)

// socket-server project main.go

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9090"
	SERVER_TYPE = "tcp"
)

func main() {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	// send some data
	_, err = connection.Write([]byte("Hello Server! Greetings."))
	// 응답 메시지 읽어오기
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
}
