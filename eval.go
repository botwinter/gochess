package main

const (
	kingWeight   = 20000
	queenWeight  = 900
	rookWeight   = 500
	bishopWeight = 330
	knightWeight = 320
	pawnWeight   = 100
)

var pawnEvalTable = [][]int{
	[]int{0, 0, 0, 0, 0, 0, 0, 0},
	[]int{50, 50, 50, 50, 50, 50, 50, 50},
	[]int{10, 10, 20, 30, 30, 20, 10, 10},
	[]int{5, 5, 10, 25, 25, 10, 5, 5},
	[]int{0, 0, 0, 20, 20, 0, 0, 0},
	[]int{5, -5, -10, 0, 0, -10, -5, 5},
	[]int{5, 10, 10, -20, -20, 10, 10, 5},
	[]int{0, 0, 0, 0, 0, 0, 0, 0},
}

var knightEvalTable = [][]int{
	[]int{-50, -40, -30, -30, -30, -30, -40, -50},
	[]int{-40, -20, 0, 0, 0, 0, -20, -40},
	[]int{-30, 0, 10, 15, 15, 10, 0, -30},
	[]int{-30, 5, 15, 20, 20, 15, 5, -30},
	[]int{-30, 0, 15, 20, 20, 15, 0, -30},
	[]int{-30, 5, 10, 15, 15, 10, 5, -30},
	[]int{-40, -20, 0, 5, 5, 0, -20, -40},
	[]int{-50, -40, -30, -30, -30, -30, -40, -50},
}

var bishopEvalTable = [][]int{
	[]int{-20, -10, -10, -10, -10, -10, -10, -20},
	[]int{-10, 0, 0, 0, 0, 0, 0, -10},
	[]int{-10, 0, 5, 10, 10, 5, 0, -10},
	[]int{-10, 5, 5, 10, 10, 5, 5, -10},
	[]int{-10, 0, 10, 10, 10, 10, 0, -10},
	[]int{-10, 10, 10, 10, 10, 10, 10, -10},
	[]int{-10, 5, 0, 0, 0, 0, 5, -10},
	[]int{-20, -10, -10, -10, -10, -10, -10, -20},
}

var rookEvalTable = [][]int{
	[]int{0, 0, 0, 0, 0, 0, 0, 0},
	[]int{5, 10, 10, 10, 10, 10, 10, 5},
	[]int{-5, 0, 0, 0, 0, 0, 0, -5},
	[]int{-5, 0, 0, 0, 0, 0, 0, -5},
	[]int{-5, 0, 0, 0, 0, 0, 0, -5},
	[]int{-5, 0, 0, 0, 0, 0, 0, -5},
	[]int{-5, 0, 0, 0, 0, 0, 0, -5},
	[]int{0, 0, 0, 5, 5, 0, 0, 0},
}

var queenEvalTable = [][]int{
	[]int{-20, -10, -10, -5, -5, -10, -10, -20},
	[]int{-10, 0, 0, 0, 0, 0, 0, -10},
	[]int{-10, 0, 5, 5, 5, 5, 0, -10},
	[]int{-5, 0, 5, 5, 5, 5, 0, -5},
	[]int{0, 0, 5, 5, 5, 5, 0, -5},
	[]int{-10, 5, 5, 5, 5, 5, 0, -10},
	[]int{-10, 0, 5, 0, 0, 0, 0, -10},
	[]int{-20, -10, -10, -5, -5, -10, -10, -20},
}

var kingEvalTable = [][]int{
	[]int{-30, -40, -40, -50, -50, -40, -40, -30},
	[]int{-30, -40, -40, -50, -50, -40, -40, -30},
	[]int{-30, -40, -40, -50, -50, -40, -40, -30},
	[]int{-30, -40, -40, -50, -50, -40, -40, -30},
	[]int{-20, -30, -30, -40, -40, -30, -30, -20},
	[]int{-10, -20, -20, -20, -20, -20, -20, -10},
	[]int{20, 20, 0, 0, 0, 0, 20, 20},
	[]int{20, 30, 10, 0, 0, 10, 30, 20},
}

/* This is a slightly less basic material + position evaluation */
func evaluateBoard(b *board, colour int) int {
	score := 0

	for x, col := range b.squares {
		for y, sq := range col {
			if sq == whiteKing {
				if colour == white {
					score += (kingWeight + kingEvalTable[x][y])
				} else {
					score -= (kingWeight + kingEvalTable[x][y])
				}
			}

			if sq == blackKing {
				if colour == black {
					score += (kingWeight + kingEvalTable[x][7-y])
				} else {
					score -= (kingWeight + kingEvalTable[x][7-y])
				}
			}

			if sq == whiteQueen {
				if colour == white {
					score += (queenWeight + queenEvalTable[x][y])
				} else {
					score -= (queenWeight + queenEvalTable[x][y])
				}
			}

			if sq == blackQueen {
				if colour == black {
					score += (queenWeight + queenEvalTable[x][7-y])
				} else {
					score -= (queenWeight + queenEvalTable[x][7-y])
				}
			}

			if sq == whiteRook {
				if colour == white {
					score += (rookWeight + rookEvalTable[x][y])
				} else {
					score -= (rookWeight + rookEvalTable[x][y])
				}
			}

			if sq == blackRook {
				if colour == black {
					score += (rookWeight + rookEvalTable[x][7-y])
				} else {
					score -= (rookWeight + rookEvalTable[x][7-y])
				}
			}

			if sq == whiteBishop {
				if colour == white {
					score += (bishopWeight + bishopEvalTable[x][y])
				} else {
					score -= (bishopWeight + bishopEvalTable[x][y])
				}
			}

			if sq == blackBishop {
				if colour == black {
					score += (bishopWeight + bishopEvalTable[x][7-y])
				} else {
					score -= (bishopWeight + bishopEvalTable[x][7-y])
				}
			}

			if sq == whiteKnight {
				if colour == white {
					score += (knightWeight + knightEvalTable[x][y])
				} else {
					score -= (knightWeight + knightEvalTable[x][y])
				}
			}

			if sq == blackKnight {
				if colour == black {
					score += (knightWeight + knightEvalTable[x][7-y])
				} else {
					score -= (knightWeight + knightEvalTable[x][7-y])
				}
			}

			if sq == whitePawn {
				if colour == white {
					score += (pawnWeight + pawnEvalTable[x][y])
				} else {
					score -= (pawnWeight + pawnEvalTable[x][y])
				}
			}

			if sq == blackPawn {
				if colour == black {
					score += (pawnWeight + pawnEvalTable[x][7-y])
				} else {
					score -= (pawnWeight + pawnEvalTable[x][7-y])
				}
			}
		}
	}

	return score
}
