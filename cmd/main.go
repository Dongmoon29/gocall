package main

import (
	"fmt"
	"os"
)

func main() {
	// 명령어 인자 확인
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "server":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gocall server <port>")
			return
		}
		port := os.Args[2]
		startServer(port)
	
	case "client":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gocall client <host:port>")
			return
		}
		address := os.Args[2]
		startClient(address)
	
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("gocall - Simple P2P Communication Tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  gocall server <port>      - Start server on specified port")
	fmt.Println("  gocall client <host:port> - Connect to server")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  gocall server 8080")
	fmt.Println("  gocall client localhost:8080")
}

// startServer와 startClient는 각각 server.go와 client.go에서 구현됨
