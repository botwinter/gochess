package main

const (
	kingWeight   = 20000
	queenWeight  = 900
	rookWeight   = 500
	bishopWeight = 330
	knightWeight = 320
	pawnWeight   = 100
)

var whitePawnEvalTable = [][]int{
	[]int{0, 5, 5, 0, 5, 10, 50, 0},
	[]int{0, 10, -5, 0, 5, 10, 50, 0},
	[]int{0, 10, -10, 0, 10, 20, 50, 0},
	[]int{0, -20, 0, 20, 25, 30, 50, 0},
	[]int{0, -20, 0, 20, 25, 30, 50, 0},
	[]int{0, 10, -10, 0, 10, 20, 50, 0},
	[]int{0, 10, -5, 0, 5, 10, 50, 0},
	[]int{0, 5, 5, 0, 5, 10, 50, 0},
}

var whiteKnightEvalTable = [][]int{
	[]int{-50, -40, -30, -30, -30, -30, -40, -50},
	[]int{-40, -20, -5, 0, 5, 0, -20, -40},
	[]int{-30, 0, 10, 15, 15, 10, 0, -30},
	[]int{-30, 5, 15, 20, 20, 15, 0, -30},
	[]int{-30, 5, 15, 20, 20, 15, 0, -30},
	[]int{-30, 0, 10, 15, 15, 10, 0, -30},
	[]int{-40, -20, -5, 0, 5, 0, -20, -40},
	[]int{-50, -40, -30, -30, -30, -30, -40, -50},
}

var whiteBishopEvalTable = [][]int{
	[]int{-20, -10, -10, -10, -10, -10, -10, -20},
	[]int{-10, 5, 10, 0, 5, 0, 0, -10},
	[]int{-10, 0, 10, 10, 5, 5, 0, -10},
	[]int{-10, 0, 10, 10, 10, 10, 0, -10},
	[]int{-10, 0, 10, 10, 10, 10, 0, -10},
	[]int{-10, 0, 10, 10, 5, 5, 0, -10},
	[]int{-10, 5, 10, 0, 5, 0, 0, -10},
	[]int{-20, -10, -10, -10, -10, -10, -10, -20},
}

var whiteRookEvalTable = [][]int{
	[]int{0, -5, -5, -5, -5, -5, 5, 0},
	[]int{0, 0, 0, 0, 0, 0, 10, 0},
	[]int{0, 0, 0, 0, 0, 0, 10, 0},
	[]int{5, 0, 0, 0, 0, 0, 10, 0},
	[]int{5, 0, 0, 0, 0, 0, 10, 0},
	[]int{0, 0, 0, 0, 0, 0, 10, 0},
	[]int{0, 0, 0, 0, 0, 0, 10, 0},
	[]int{0, -5, -5, -5, -5, -5, 5, 0},
}

var whiteQueenEvalTable = [][]int{
	[]int{-20, -10, -10, 0, -5, -10, -10, -20},
	[]int{-10, 0, 5, 0, 0, 0, 0, -10},
	[]int{-10, 5, 5, 5, 5, 5, 0, -10},
	[]int{-5, 0, 5, 5, 5, 5, 0, -5},
	[]int{-5, 0, 5, 5, 5, 5, 0, -5},
	[]int{-10, 5, 5, 5, 5, 5, 0, -10},
	[]int{-10, 0, 0, 0, 0, 0, 0, -10},
	[]int{-20, -10, -10, -5, -5, -10, -10, -20},
}

var whiteKingEvalTable = [][]int{
	[]int{20, 20, -10, -20, -30, -30, -30, -30},
	[]int{30, 20, -20, -30, -40, -40, -40, -40},
	[]int{10, 0, -20, -30, -40, -40, -40, -40},
	[]int{0, 0, -20, -40, -50, -50, -50, -50},
	[]int{0, 0, -20, -40, -50, -50, -50, -50},
	[]int{10, 0, -20, -30, -40, -40, -40, -40},
	[]int{30, 20, -20, -30, -40, -40, -40, -40},
	[]int{20, 20, -10, -20, -30, -30, -30, -30},
}

var blackPawnEvalTable = reverseBoardArray(whitePawnEvalTable)
var blackKnightEvalTable = reverseBoardArray(whiteKnightEvalTable)
var blackBishopEvalTable = reverseBoardArray(whiteBishopEvalTable)
var blackRookEvalTable = reverseBoardArray(whiteRookEvalTable)
var blackQueenEvalTable = reverseBoardArray(whiteQueenEvalTable)
var blackKingEvalTable = reverseBoardArray(whiteKingEvalTable)

/* This is a slightly less basic material + position evaluation */
func evaluateBoard(b *board, colour int) int {
	score := 0

	for x, col := range b.squares {
		for y, sq := range col {
			if sq == whiteKing {
				if colour == white {
					score += (kingWeight + whiteKingEvalTable[x][y])
				} else {
					score -= (kingWeight + whiteKingEvalTable[x][y])
				}
			}

			if sq == blackKing {
				if colour == black {
					score += (kingWeight + blackKingEvalTable[x][y])
				} else {
					score -= (kingWeight + blackKingEvalTable[x][y])
				}
			}

			if sq == whiteQueen {
				if colour == white {
					score += (queenWeight + whiteQueenEvalTable[x][y])
				} else {
					score -= (queenWeight + whiteQueenEvalTable[x][y])
				}
			}

			if sq == blackQueen {
				if colour == black {
					score += (queenWeight + blackQueenEvalTable[x][y])
				} else {
					score -= (queenWeight + blackQueenEvalTable[x][y])
				}
			}

			if sq == whiteRook {
				if colour == white {
					score += (rookWeight + whiteRookEvalTable[x][y])
				} else {
					score -= (rookWeight + whiteRookEvalTable[x][y])
				}
			}

			if sq == blackRook {
				if colour == black {
					score += (rookWeight + blackRookEvalTable[x][y])
				} else {
					score -= (rookWeight + blackRookEvalTable[x][y])
				}
			}

			if sq == whiteBishop {
				if colour == white {
					score += (bishopWeight + whiteBishopEvalTable[x][y])
				} else {
					score -= (bishopWeight + whiteBishopEvalTable[x][y])
				}
			}

			if sq == blackBishop {
				if colour == black {
					score += (bishopWeight + blackBishopEvalTable[x][y])
				} else {
					score -= (bishopWeight + blackBishopEvalTable[x][y])
				}
			}

			if sq == whiteKnight {
				if colour == white {
					score += (knightWeight + whiteKnightEvalTable[x][y])
				} else {
					score -= (knightWeight + whiteKnightEvalTable[x][y])
				}
			}

			if sq == blackKnight {
				if colour == black {
					score += (knightWeight + blackKnightEvalTable[x][y])
				} else {
					score -= (knightWeight + blackKnightEvalTable[x][y])
				}
			}

			if sq == whitePawn {
				if colour == white {
					score += (pawnWeight + whitePawnEvalTable[x][y])
				} else {
					score -= (pawnWeight + whitePawnEvalTable[x][y])
				}
			}

			if sq == blackPawn {
				if colour == black {
					score += (pawnWeight + blackPawnEvalTable[x][y])
				} else {
					score -= (pawnWeight + blackPawnEvalTable[x][y])
				}
			}
		}
	}

	return score
}
