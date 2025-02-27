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

func parseResponse(client net.Conn) (string, error) {
	reader := bufio.NewReader(client)
	statusLine, err := reader.ReadString('\n')
	return statusLine, err
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

type testCase struct {
	request  string
	expected string
}

func TestHandleConnection(t *testing.T) {
	t.Run("should respond to client", func(t *testing.T) {

		testCases := []testCase{
			{request: "GET / HTTP/1.0", expected: "HTTP/1.0 200 OK\r\n"},
			{request: "GET / HTTP/1.1", expected: "HTTP/1.0 400 Bad Request\r\n"},
			{request: "foobar", expected: "HTTP/1.0 400 Bad Request\r\n"},
		}

		for _, test := range testCases {
			client, server := net.Pipe()

			go sendMockRequest(client, test.request)
			go handleConnection(server)

			response, err := parseResponse(client)

			if err != nil {
				t.Fatal("could not read status from response")
			}

			if response != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, response)
			}
		}
	})
}
