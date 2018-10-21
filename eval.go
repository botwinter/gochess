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

/* This function always evaluates from the PoV of white. It is up to the caller to
reverse it for black */
func evaluateBoard(b *board) int {
	score := 0

	for x, col := range b.squares {
		for y, sq := range col {
			if sq == whiteKing {
				score += (kingWeight + whiteKingEvalTable[x][y])
			}

			if sq == blackKing {
				score -= (kingWeight + blackKingEvalTable[x][y])
			}

			if sq == whiteQueen {
				score += (queenWeight + whiteQueenEvalTable[x][y])
			}

			if sq == blackQueen {
				score -= (queenWeight + blackQueenEvalTable[x][y])
			}

			if sq == whiteRook {
				score += (rookWeight + whiteRookEvalTable[x][y])
			}

			if sq == blackRook {
				score -= (rookWeight + blackRookEvalTable[x][y])
			}

			if sq == whiteBishop {
				score += (bishopWeight + whiteBishopEvalTable[x][y])
			}

			if sq == blackBishop {
				score -= (bishopWeight + blackBishopEvalTable[x][y])
			}

			if sq == whiteKnight {
				score += (knightWeight + whiteKnightEvalTable[x][y])
			}

			if sq == blackKnight {
				score -= (knightWeight + blackKnightEvalTable[x][y])
			}

			if sq == whitePawn {
				score += (pawnWeight + whitePawnEvalTable[x][y])
			}

			if sq == blackPawn {
				score -= (pawnWeight + blackPawnEvalTable[x][y])
			}
		}
	}

	return score
}
