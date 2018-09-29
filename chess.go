package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

// Global flags
var cursesEnabled bool

func main() {
	b := newDefaultBoard()
	c := initCursesBoard()

	cursesEnabled = true

	termbox.Init()
	defer termbox.Close()

	err := renderBoard(c, b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Play the game - player is always white
	for {
		// Wait for player move
		move := handleMoveInput(c, b)

		// Make player move
		makeMove(b, move)

		// Find AI move
		renderBoard(c, b)
		renderStatusLine(c, "Computer is thinking...")
		bestAIMove := findBestMove(b, black)

		// Make AI move
		makeMove(b, &bestAIMove)

		// Update board view
		renderBoard(c, b)
	}
}
