package main

/* This function implements a Negamax search (see https://en.wikipedia.org/wiki/Negamax)
with alpha beta pruning */
func search(b *board, depth int, origDepth int, colour int, alpha int, beta int, bestMove *move, moveStack *[]move, movesProcessed int) int {
	if depth == 0 {
		return evaluateBoard(b, colour)
	}

	score := 0

	// Am I in check?
	moves := []move{}
	if inCheck(b, colour) {
		moves = generateAllLegalMovesInCheck(b, colour)
	} else {
		moves = generateAllLegalMoves(b, colour)
	}

	if len(moves) == 0 {
		return evaluateBoard(b, colour)
	}

	for _, move := range moves {
		// Make the possible move
		makeMove(b, &move)
		movesProcessed++

		// Run search again to minimise the score of the opposite colour
		score = -search(b, depth-1, origDepth, -colour, -beta, -alpha, bestMove, moveStack, movesProcessed)

		unmakeMove(b, &move)

		if score > alpha {
			alpha = score

			// If we are in the root search call and this is the new best score, update bestMove
			if depth == origDepth {
				*bestMove = move
			}

			if score >= beta {
				break
			}
		}
	}

	return alpha
}

func findBestMove(b *board, colour int) move {
	depth := 6
	bestMove := move{0, 0, 0, 0, 0}
	moveStack := make([]move, depth)
	search(b, depth, depth, colour, -99999, 99999, &bestMove, &moveStack, 0)

	return bestMove
}
