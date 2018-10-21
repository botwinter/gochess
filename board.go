package main

const (
	// Castling flags
	whiteRookKingSideMoved = 1 << iota
	whiteRookQueenSideMoved
	whiteKingMoved
	blackRookKingSideMoved
	blackRookQueenSideMoved
	blackKingMoved
	numBoardFlags
)

type board struct {
	squares [][]int

	flags uint64
}

func newBoard() *board {
	b := board{}

	return &b
}

func newBoardWithPieces(pieces [][]int) *board {
	b := board{pieces, 0}

	// Set flags based on piece layout
	if b.squares[0][0] != whiteRook {
		b.flags = setFlag(b.flags, whiteRookQueenSideMoved)
	}
	if b.squares[4][0] != whiteKing {
		b.flags = setFlag(b.flags, whiteKingMoved)
	}
	if b.squares[7][0] != whiteRook {
		b.flags = setFlag(b.flags, whiteRookKingSideMoved)
	}
	if b.squares[0][7] != blackRook {
		b.flags = setFlag(b.flags, blackRookQueenSideMoved)
	}
	if b.squares[4][7] != blackKing {
		b.flags = setFlag(b.flags, blackKingMoved)
	}
	if b.squares[7][7] != blackRook {
		b.flags = setFlag(b.flags, blackRookKingSideMoved)
	}

	return &b
}

func newBoardWithPiecesAndFlags(pieces [][]int, flags uint64) *board {
	b := board{pieces, flags}

	// Set flags based on piece layout
	if b.squares[0][0] != whiteRook {
		b.flags = setFlag(b.flags, whiteRookQueenSideMoved)
	}
	if b.squares[4][0] != whiteKing {
		b.flags = setFlag(b.flags, whiteKingMoved)
	}
	if b.squares[7][0] != whiteRook {
		b.flags = setFlag(b.flags, whiteRookKingSideMoved)
	}
	if b.squares[0][7] != blackRook {
		b.flags = setFlag(b.flags, blackRookQueenSideMoved)
	}
	if b.squares[4][7] != blackKing {
		b.flags = setFlag(b.flags, blackKingMoved)
	}
	if b.squares[7][7] != blackRook {
		b.flags = setFlag(b.flags, blackRookKingSideMoved)
	}

	return &b
}

func newDefaultBoard() *board {
	return newBoardWithPieces([][]int{
		[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
		[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
		[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
		[]int{whiteQueen, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
		[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, blackKing},
		[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
		[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
		[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
	})
}

/* This function looks along a specified diagonal, and returns true if the current piece is the only
thing standing between the king and an enemy bishop or queen.
If pieceCoords[0] is -1, then this function assumes the king is the current piece */
func isEnemyOnDiagonal(b *board, kingCoords [2]int, pieceCoords [2]int, colour int, diagonal int) bool {
	enemyInWay := false
	xcoord := 0
	ycoord := 0

	// Set initial coords to one square from king in target diagonal
	if diagonal == bottomLeftDiagonal {
		xcoord = kingCoords[0] - 1
		ycoord = kingCoords[1] - 1
	} else if diagonal == topLeftDiagonal {
		xcoord = kingCoords[0] - 1
		ycoord = kingCoords[1] + 1
	} else if diagonal == topRightDiagonal {
		xcoord = kingCoords[0] + 1
		ycoord = kingCoords[1] + 1
	} else if diagonal == bottomRightDiagonal {
		xcoord = kingCoords[0] + 1
		ycoord = kingCoords[1] - 1
	}

	for xcoord >= 0 && xcoord < 8 && ycoord >= 0 && ycoord < 8 {
		if (colour == white && (b.squares[xcoord][ycoord] == blackBishop || b.squares[xcoord][ycoord] == blackQueen)) || (colour == black && (b.squares[xcoord][ycoord] == whiteBishop || b.squares[xcoord][ycoord] == whiteQueen)) {
			// Found an enemy bishop/queen, so we are the only thing in the way
			enemyInWay = true
			break
		} else if b.squares[xcoord][ycoord] != empty && !(pieceCoords[0] != -1 && xcoord == pieceCoords[0] && ycoord == pieceCoords[1]) {
			// Found a non-empty square that is NOT an enemy queen/bishop and not our piece, so we're safe
			break
		}

		// Move coords along diagonal
		if diagonal == bottomLeftDiagonal {
			xcoord--
			ycoord--
		} else if diagonal == topLeftDiagonal {
			xcoord--
			ycoord++
		} else if diagonal == topRightDiagonal {
			xcoord++
			ycoord++
		} else if diagonal == bottomRightDiagonal {
			xcoord++
			ycoord--
		}
	}

	return enemyInWay
}

/* This function looks along a specified direction, and returns true if the current piece is the only
thing standing between the king and an enemy rook or queen.
If pieceCoords[0] is -1, then this function assumes the king is the current piece */
func isEnemyInDirection(b *board, kingCoords [2]int, pieceCoords [2]int, colour int, direction int) bool {
	enemyInWay := false
	xcoord := 0
	ycoord := 0

	// Set initial coords to one square from king in target direction
	if direction == left {
		xcoord = kingCoords[0] - 1
		ycoord = kingCoords[1]
	} else if direction == up {
		xcoord = kingCoords[0]
		ycoord = kingCoords[1] + 1
	} else if direction == right {
		xcoord = kingCoords[0] + 1
		ycoord = kingCoords[1]
	} else if direction == down {
		xcoord = kingCoords[0]
		ycoord = kingCoords[1] - 1
	}

	for xcoord >= 0 && xcoord < 8 && ycoord >= 0 && ycoord < 8 {
		if (colour == white && (b.squares[xcoord][ycoord] == blackRook || b.squares[xcoord][ycoord] == blackQueen)) || (colour == black && (b.squares[xcoord][ycoord] == whiteRook || b.squares[xcoord][ycoord] == whiteQueen)) {
			// Found an enemy rook/queen, so we are the only thing in the way
			enemyInWay = true
			break
		} else if b.squares[xcoord][ycoord] != empty && !(pieceCoords[0] != -1 && xcoord == pieceCoords[0] && ycoord == pieceCoords[1]) {
			// Found a non-empty square that is NOT an enemy queen/bishop and not our piece, so we're safe
			break
		}

		// Move coords along direction
		if direction == left {
			xcoord--
		} else if direction == up {
			ycoord++
		} else if direction == right {
			xcoord++
		} else if direction == down {
			ycoord--
		}
	}

	return enemyInWay
}

/* This function assumes the move is valid */
func makeMove(b *board, m *move) *board {
	fromPiece := b.squares[m.fromX][m.fromY]
	toPiece := fromPiece

	// Check for special move flags
	if hasFlag(m.flags, queenPromotion) {
		if isWhite(fromPiece) {
			toPiece = whiteQueen
		} else {
			toPiece = blackQueen
		}
	} else if hasFlag(m.flags, rookPromotion) {
		if isWhite(fromPiece) {
			toPiece = whiteRook
		} else {
			toPiece = blackRook
		}
	} else if hasFlag(m.flags, knightPromotion) {
		if isWhite(fromPiece) {
			toPiece = whiteKnight
		} else {
			toPiece = blackKnight
		}
	} else if hasFlag(m.flags, bishopPromotion) {
		if isWhite(fromPiece) {
			toPiece = whiteBishop
		} else {
			toPiece = blackBishop
		}
	} else if hasFlag(m.flags, kingCastle) {
		b.squares[4][m.fromY] = empty
		b.squares[7][m.fromY] = empty
		if m.fromY == 0 {
			b.flags = setFlag(b.flags, whiteRookKingSideMoved)
			b.flags = setFlag(b.flags, whiteKingMoved)
			m.boardFlags = setFlag(m.boardFlags, whiteRookKingSideMoved)
			m.boardFlags = setFlag(m.boardFlags, whiteKingMoved)
			b.squares[6][m.fromY] = whiteKing
			b.squares[5][m.fromY] = whiteRook
		} else {
			b.flags = setFlag(b.flags, blackRookKingSideMoved)
			b.flags = setFlag(b.flags, blackKingMoved)
			m.boardFlags = setFlag(m.boardFlags, blackRookKingSideMoved)
			m.boardFlags = setFlag(m.boardFlags, blackKingMoved)
			b.squares[6][m.fromY] = blackKing
			b.squares[5][m.fromY] = blackRook
		}

		return b
	} else if hasFlag(m.flags, queenCastle) {
		b.squares[4][m.fromY] = empty
		b.squares[0][m.fromY] = empty
		if m.fromY == 0 {
			b.flags = setFlag(b.flags, whiteRookQueenSideMoved)
			b.flags = setFlag(b.flags, whiteKingMoved)
			m.boardFlags = setFlag(m.boardFlags, whiteRookQueenSideMoved)
			m.boardFlags = setFlag(m.boardFlags, whiteKingMoved)
			b.squares[2][m.fromY] = whiteKing
			b.squares[3][m.fromY] = whiteRook
		} else {
			b.flags = setFlag(b.flags, blackRookQueenSideMoved)
			b.flags = setFlag(b.flags, blackKingMoved)
			m.boardFlags = setFlag(m.boardFlags, blackRookQueenSideMoved)
			m.boardFlags = setFlag(m.boardFlags, blackKingMoved)
			b.squares[2][m.fromY] = blackKing
			b.squares[3][m.fromY] = blackRook
		}

		return b
	}

	// If not castling, then set any flags for future castling.
	// Also set them in the move so we can unset them later.
	if fromPiece == whiteRook {
		if m.fromX == 0 && m.fromY == 0 {
			if !hasFlag(b.flags, whiteRookQueenSideMoved) {
				b.flags = setFlag(b.flags, whiteRookQueenSideMoved)
				m.boardFlags = setFlag(m.boardFlags, whiteRookQueenSideMoved)
			}
		} else if m.fromX == 7 && m.fromY == 0 {
			if !hasFlag(b.flags, whiteRookKingSideMoved) {
				b.flags = setFlag(b.flags, whiteRookKingSideMoved)
				m.boardFlags = setFlag(m.boardFlags, whiteRookKingSideMoved)
			}
		}
	} else if fromPiece == blackRook {
		if m.fromX == 0 && m.fromY == 7 {
			if !hasFlag(b.flags, blackRookQueenSideMoved) {
				b.flags = setFlag(b.flags, blackRookQueenSideMoved)
				m.boardFlags = setFlag(m.boardFlags, blackRookQueenSideMoved)
			}
		} else if m.fromX == 7 && m.fromY == 7 {
			if !hasFlag(b.flags, blackRookKingSideMoved) {
				b.flags = setFlag(b.flags, blackRookKingSideMoved)
				m.boardFlags = setFlag(m.boardFlags, blackRookKingSideMoved)
			}
		}
	} else if fromPiece == whiteKing {
		if !hasFlag(b.flags, whiteKingMoved) {
			b.flags = setFlag(b.flags, whiteKingMoved)
			m.boardFlags = setFlag(m.boardFlags, whiteKingMoved)
		}
	} else if fromPiece == blackKing {
		if !hasFlag(b.flags, blackKingMoved) {
			b.flags = setFlag(b.flags, blackKingMoved)
			m.boardFlags = setFlag(m.boardFlags, blackKingMoved)
		}
	}

	// Make the standard move
	m.taken = b.squares[m.toX][m.toY]
	b.squares[m.toX][m.toY] = toPiece
	b.squares[m.fromX][m.fromY] = empty

	return b
}

/* This function unmakes a move, including any board-level flags set during a move. */
func unmakeMove(b *board, m *move) *board {
	fromPiece := b.squares[m.toX][m.toY]

	// Unset any board-level flags that might have been set
	if hasFlag(m.boardFlags, whiteRookKingSideMoved) {
		b.flags = clearFlag(b.flags, whiteRookKingSideMoved)
	}
	if hasFlag(m.boardFlags, whiteRookQueenSideMoved) {
		b.flags = clearFlag(b.flags, whiteRookQueenSideMoved)
	}
	if hasFlag(m.boardFlags, blackRookKingSideMoved) {
		b.flags = clearFlag(b.flags, blackRookKingSideMoved)
	}
	if hasFlag(m.boardFlags, blackRookQueenSideMoved) {
		b.flags = clearFlag(b.flags, blackRookQueenSideMoved)
	}
	if hasFlag(m.boardFlags, whiteKingMoved) {
		b.flags = clearFlag(b.flags, whiteKingMoved)
	}
	if hasFlag(m.boardFlags, blackKingMoved) {
		b.flags = clearFlag(b.flags, blackKingMoved)
	}

	// Check for special move flags
	if hasFlag(m.flags, queenPromotion) || hasFlag(m.flags, rookPromotion) || hasFlag(m.flags, knightPromotion) || hasFlag(m.flags, bishopPromotion) {
		if isWhite(fromPiece) {
			fromPiece = whitePawn
		} else {
			fromPiece = blackPawn
		}
	}

	if hasFlag(m.flags, kingCastle) {
		if m.fromY == 0 {
			b.squares[4][m.fromY] = whiteKing
			b.squares[7][m.fromY] = whiteRook
		} else {
			b.squares[4][m.fromY] = blackKing
			b.squares[7][m.fromY] = blackRook
		}

		b.squares[6][m.fromY] = empty
		b.squares[5][m.fromY] = empty

		return b
	} else if hasFlag(m.flags, queenCastle) {
		if m.fromY == 0 {
			b.squares[4][m.fromY] = whiteKing
			b.squares[0][m.fromY] = whiteRook
		} else {
			b.squares[4][m.fromY] = blackKing
			b.squares[0][m.fromY] = blackRook
		}
		b.squares[2][m.fromY] = empty
		b.squares[3][m.fromY] = empty

		return b
	}

	b.squares[m.fromX][m.fromY] = fromPiece
	b.squares[m.toX][m.toY] = m.taken

	return b
}

// Keep this in a separate function so we only do it for user-provided moves
func validMove(b *board, m *move, colour int) *move {
	// Basic checks
	if m.fromX < 0 || m.fromX > 7 || m.fromY < 0 || m.fromY > 7 || m.toX < 0 || m.toX > 7 || m.toY < 0 || m.toY > 7 {
		return nil
	}

	validMoves := []move{}
	if inCheck(b, colour) {
		validMoves = generateAllLegalMovesInCheck(b, colour)
	} else {
		validMoves = generateAllLegalMoves(b, colour)
	}

	for _, move := range validMoves {
		if m.fromX == move.fromX && m.fromY == move.fromY && m.toX == move.toX && m.toY == move.toY {
			return &move
		}
	}

	return nil
}
