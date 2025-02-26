package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("server listening on port 4221")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("error accepting connection", err.Error())
			os.Exit(1)
		}

		msg := make([]byte, 1024)

		conn.Read(msg)

		fmt.Println("message recieved: " + string(msg))

		conn.Write([]byte("server recieved your message: " + string(msg)))
	}
}
