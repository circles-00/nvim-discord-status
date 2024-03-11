package main

import (
	"fmt"
	"net"
	"time"

	tcp "discord_status/tcp"
)

func main() {
	// Listen for incoming connections
	const port = 49069
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port %d\n", port)

	t := time.Now()

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		tcp.NumberOfClients = tcp.NumberOfClients + 1
		// Handle client connection in a goroutine
		go tcp.HandleTCPClient(conn, t)
	}
}
