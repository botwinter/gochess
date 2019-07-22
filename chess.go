package main

import "fmt"

type engine struct {
	board  *board
	server *server
}

// Global flags
var searchDepth = 5

func main() {
	mainEngine := engine{}
	mainServer := server{}
	mainServer.engine = &mainEngine
	mainEngine.server = &mainServer

	// Set deeper search depth for real game
	searchDepth = 6

	// Serve ChessboardJS and start listening for requests
	fmt.Println("Running at http://localhost:8080...")
	startServer(&mainServer)
}
