package main

import "fmt"

// Global flags
var searchDepth = 5

var mainBoard *board

func main() {
	mainBoard = nil

	// Set deeper search depth for real game
	searchDepth = 6

	// Serve ChessboardJS and start listening for requests
	fmt.Println("Running at http://localhost:8080...")
	startServer()
}
