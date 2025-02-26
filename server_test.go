package main

import (
	"net"
	"testing"
	"time"
)

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
