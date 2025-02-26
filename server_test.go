package main

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func formatRequest(request string) string {
	return request + "\r\n\r\n"
}

func sendMockRequest(conn net.Conn, request string) {
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, _ = writer.WriteString(formatRequest(request))
	_ = writer.Flush()
}

func TestStartServer(t *testing.T) {

	go func() {
		StartServer()
	}()

	time.Sleep(100 * time.Millisecond)

	t.Run("should accept a connection", func(t *testing.T) {
		conn, err := net.Dial("tcp", "localhost:4221")
		if err != nil {
			t.Fatal("could not connect to port")
		}
		defer conn.Close()
	})

}

func TestHandleConnection(t *testing.T) {
	t.Run("should respond to client", func(t *testing.T) {
		client, server := net.Pipe()

		go sendMockRequest(client, "GET / HTTP/1.0")

		go handleConnection(server)

		reader := bufio.NewReader(client)

		statusLine, err := reader.ReadString('\n')

		if err != nil {
			t.Fatal("could not read status from response")
		}

		expectedStatus := "HTTP/1.0 200 OK\r\n"

		if statusLine != expectedStatus {
			t.Fatalf("expected %s, got %s", expectedStatus, statusLine)
		}
	})

	t.Run("should reject requests without HTTP/1.0 header", func(t *testing.T) {
		client, server := net.Pipe()

		go sendMockRequest(client, "GET / HTTP/1.1")

		go handleConnection(server)

		reader := bufio.NewReader(client)

		statusLine, err := reader.ReadString('\n')

		if err != nil {
			t.Fatal("could not read status from response")
		}

		expectedStatus := "HTTP/1.0 400 Bad Request\r\n"

		if statusLine != expectedStatus {
			t.Fatalf("expected %s, got %s", expectedStatus, statusLine)
		}
	})
}
