package main

const (
	kingWeight   = 200
	queenWeight  = 9
	rookWeight   = 5
	bishopWeight = 3
	knightWeight = 3
	pawnWeight   = 1
)

/* This is a basic material evaluation for now */
func evaluateBoard(b *board, colour int) int {
	score := 0

	for _, col := range b.squares {
		for _, sq := range col {
			if sq == whiteKing {
				if colour == white {
					score += kingWeight
				} else {
					score -= kingWeight
				}
			}

			if sq == blackKing {
				if colour == black {
					score += kingWeight
				} else {
					score -= kingWeight
				}
			}

			if sq == whiteQueen {
				if colour == white {
					score += queenWeight
				} else {
					score -= queenWeight
				}
			}

			if sq == blackQueen {
				if colour == black {
					score += queenWeight
				} else {
					score -= queenWeight
				}
			}

			if sq == whiteRook {
				if colour == white {
					score += rookWeight
				} else {
					score -= rookWeight
				}
			}

			if sq == blackRook {
				if colour == black {
					score += rookWeight
				} else {
					score -= rookWeight
				}
			}

			if sq == whiteBishop {
				if colour == white {
					score += bishopWeight
				} else {
					score -= bishopWeight
				}
			}

			if sq == blackBishop {
				if colour == black {
					score += bishopWeight
				} else {
					score -= bishopWeight
				}
			}

			if sq == whiteKnight {
				if colour == white {
					score += knightWeight
				} else {
					score -= knightWeight
				}
			}

			if sq == blackKnight {
				if colour == black {
					score += knightWeight
				} else {
					score -= knightWeight
				}
			}

			if sq == whitePawn {
				if colour == white {
					score += pawnWeight
				} else {
					score -= pawnWeight
				}
			}

			if sq == blackPawn {
				if colour == black {
					score += pawnWeight
				} else {
					score -= pawnWeight
				}
			}
		}
	}

	return score
}
