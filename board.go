package main

type board struct {
	squares [][]int
}

func newBoard() *board {
	b := board{}

	return &b
}

func newBoardWithPieces(pieces [][]int) *board {
	b := board{pieces}

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

func newDebugBoard() *board {
	return newBoardWithPieces([][]int{
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{whiteKing, empty, whiteQueen, empty, blackQueen, empty, empty, blackKing},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
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
	m.taken = b.squares[m.toX][m.toY]
	b.squares[m.toX][m.toY] = b.squares[m.fromX][m.fromY]
	b.squares[m.fromX][m.fromY] = empty

	return b
}

func unmakeMove(b *board, m *move) *board {
	b.squares[m.fromX][m.fromY] = b.squares[m.toX][m.toY]
	b.squares[m.toX][m.toY] = m.taken

	return b
}

// Keep this in a separate function so we only do it for user-provided moves
func validMove(b *board, m *move, colour int) bool {
	// Basic checks
	if m.fromX < 0 || m.fromX > 7 || m.fromY < 0 || m.fromY > 7 || m.toX < 0 || m.toX > 7 || m.toY < 0 || m.toY > 7 {
		return false
	}

	validMoves := []move{}
	if inCheck(b, colour) {
		validMoves = generateAllLegalMovesInCheck(b, colour)
	} else {
		validMoves = generateAllLegalMoves(b, colour)
	}

	for _, move := range validMoves {
		if m.fromX == move.fromX && m.fromY == move.fromY && m.toX == move.toX && m.toY == move.toY {
			return true
		}
	}

	return false
}
