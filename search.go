package main

import (
	"fmt"
	"os"
)

func search(b *board, depth int, origDepth int, colour int, bestMove *move, moveStack *[]move) int {
	if depth == 0 {
		return evaluateBoard(b, colour)
	}

	max := -999
	score := 0
	moves := generateAllLegalMoves(b, colour)
	if len(moves) == 0 {
		return evaluateBoard(b, colour)
	}

	err := renderBoard(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, move := range moves {
		// Make the possible move
		makeMove(b, &move)

		// Run search again to minimise the score of the opposite colour
		score = -search(b, depth-1, origDepth, -colour, bestMove, moveStack)
		if score > max {
			max = score

			// If we are in the root search call and this is the new best score, update bestMove
			if depth == origDepth {
				oldBestMove := *bestMove
				*bestMove = move

				err = renderBoard(b)
				err = renderStatusLine(fmt.Sprintf("Updated best move! Old best move: %s    New best move: %s with score %d", oldBestMove.toString(), bestMove.toString(), max))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				handleKeyEvent()
			}
		} else {
			if depth == origDepth {
				err = renderBoard(b)
				err = renderStatusLine(fmt.Sprintf("Move %s has score: %d    not better than current best move %s with score %d", move.toString(), score, bestMove.toString(), max))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				handleKeyEvent()
			}
		}
		unmakeMove(b, &move)
	}

	return max
}

func findBestMove(b *board, colour int) move {
	depth := 5
	bestMove := move{0, 0, 0, 0, 0}
	moveStack := make([]move, depth)
	search(b, depth, depth, colour, &bestMove, &moveStack)

	err := renderBoard(b)
	err = renderStatusLine(fmt.Sprintf("Found best move: %s", bestMove.toString()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	handleKeyEvent()
	return bestMove
}
