package main

import "fmt"

type engine struct {
	board  *board
	server *server
}

// Global flags
var searchDepth = 5

func main() {
	mainServer := server{
		responseChan: make(chan string),
	}
	mainEngine := engine{
		board:  newDefaultBoard(),
		server: &mainServer,
	}

	mainServer.engine = &mainEngine

	// Set deeper search depth for real game
	searchDepth = 6

	// Serve ChessboardJS and start listening for requests
	fmt.Println("Running at http://localhost:8080...")
	startServer(&mainServer)
}
