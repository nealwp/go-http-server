package main

import (
	"fmt"
	"net"
	"os"
)

func StartServer() {

	listener, err := net.Listen("tcp", "0.0.0.0:4221")

	defer listener.Close()

	if err != nil {
		fmt.Println("failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("server listening on port 4221")

	for {
		_, err := listener.Accept()

		if err != nil {
			fmt.Println("error accepting connection", err.Error())
			os.Exit(1)
		}

	}
}
