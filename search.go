package main

func search(b *board, depth int, origDepth int, colour int, bestMove *move) int {
	if depth == 0 {
		return evaluateBoard(b, colour)
	}

	max := -999
	score := 0
	moves := generateAllLegalMoves(b, colour)
	if len(moves) == 0 {
		return evaluateBoard(b, colour)
	}

	for _, move := range moves {
		// Make the possible move
		makeMove(b, &move)

		// Run search again to minimise the score of the opposite colour
		score = -search(b, depth-1, origDepth, -colour, bestMove)
		if score > max {
			max = score

			// If we are in the root search call and this is the new best score, update bestMove
			if depth == origDepth {
				*bestMove = move
			}
		}
		unmakeMove(b, &move)
	}

	return max
}

func findBestMove(b *board, colour int) move {
	bestMove := move{0, 0, 0, 0, 0}

	search(b, 4, 4, colour, &bestMove)

	return bestMove
}
