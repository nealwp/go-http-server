package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
)

func StartServer() {

	listener, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("failed to bind to port 4221")
		os.Exit(1)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("error accepting connection", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	allowedMethods := []string{"GET", "POST", "HEAD"}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')

	if err != nil {
		return // ?
	}

	parts := strings.Fields(requestLine)

	if len(parts) < 3 || parts[2] != "HTTP/1.0" {
		response := "HTTP/1.0 400 Bad Request\r\n\r\n"
		_, _ = conn.Write([]byte(response))
		return
	}

	requestMethod := parts[0]

	if !slices.Contains(allowedMethods, requestMethod) {
		response := "HTTP/1.0 501 Not Implemented\r\n\r\n"
		_, _ = conn.Write([]byte(response))
		return
	}

	response := "HTTP/1.0 200 OK\r\n\r\n"

	_, _ = conn.Write([]byte(response))

}
