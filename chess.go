package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

// Global flags
var cursesEnabled bool

var searchDepth = 5

func main() {
	b := newDefaultBoard()
	c := initCursesBoard()

	cursesEnabled = true

	// Set deeper search depth for real game
	searchDepth = 6

	termbox.Init()
	defer termbox.Close()

	err := renderBoard(c, b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Play the game - player is always white
	for {
		// Generate player's possible moves
		validMoves := []move{}
		if inCheck(b, white) {
			validMoves = generateAllLegalMovesInCheck(b, white)
		} else {
			validMoves = generateAllLegalMoves(b, white)
		}

		if len(validMoves) == 0 {
			renderStatusLine(c, "You lost! Press ENTER to play again")
			handleGameEnd()
			b = newDefaultBoard()
			c = initCursesBoard()
			continue
		}

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
