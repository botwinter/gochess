package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Board flags
const (
	// Castling flags
	// IMPORTANT - these are actually COULD castle, not CAN. They don't track things like
	// pieces in the way.
	whiteCanCastleKingSide = 1 << iota
	whiteCanCastleQueenSide
	blackCanCastleKingSide
	blackCanCastleQueenSide

	numBoardFlags
)

type board struct {
	squares      [][]int
	flags        uint64
	colourToMove int
}

func newBoard() *board {
	return &board{
		squares: [][]int{
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		},
	}
}

func prettyPrintBoard(b *board) {
	for colNum := 7; colNum >= 0; colNum-- {
		for rowNum := 0; rowNum < 8; rowNum++ {
			if b.squares[rowNum][colNum] == whiteRook {
				fmt.Printf(" R ")
			} else if b.squares[rowNum][colNum] == whitePawn {
				fmt.Printf(" P ")
			} else if b.squares[rowNum][colNum] == whiteKnight {
				fmt.Printf(" N ")
			} else if b.squares[rowNum][colNum] == whiteBishop {
				fmt.Printf(" B ")
			} else if b.squares[rowNum][colNum] == whiteKing {
				fmt.Printf(" K ")
			} else if b.squares[rowNum][colNum] == whiteQueen {
				fmt.Printf(" Q ")
			} else if b.squares[rowNum][colNum] == blackRook {
				fmt.Printf(" r ")
			} else if b.squares[rowNum][colNum] == blackPawn {
				fmt.Printf(" p ")
			} else if b.squares[rowNum][colNum] == blackKnight {
				fmt.Printf(" n ")
			} else if b.squares[rowNum][colNum] == blackBishop {
				fmt.Printf(" b ")
			} else if b.squares[rowNum][colNum] == blackKing {
				fmt.Printf(" k ")
			} else if b.squares[rowNum][colNum] == blackQueen {
				fmt.Printf(" q ")
			} else {
				fmt.Printf("   ")
			}

			if rowNum != 7 {
				fmt.Printf("|")
			} else {
				fmt.Printf("\n")
			}
		}
		if colNum != 0 {
			fmt.Println("-------------------------------")
		}
	}
}

func newBoardFromCoords(pieces [][]int, flags uint64, colourToMove int) *board {
	return &board{pieces, flags, colourToMove}
}

func newBoardFromFen(fen string) *board {
	b := newBoard()

	fenFields := strings.Split(fen, " ")

	fenPiecesSplit := strings.Split(fenFields[0], "/")
	fenActiveColour := fenFields[1]
	fenCastling := fenFields[2]
	//fenEnPassant := fenFields[3]
	//fenHalfMoveClock := fenFields[4]
	//fenFullMoveNo := fenFields[5]

	// First, populate board based on pieces
	for rowNum, row := range fenPiecesSplit {
		colIdx := 0
		for _, fenValue := range row {
			if fenValue == 'R' {
				b.squares[colIdx][7-rowNum] = whiteRook
				colIdx++
			} else if fenValue == 'N' {
				b.squares[colIdx][7-rowNum] = whiteKnight
				colIdx++
			} else if fenValue == 'B' {
				b.squares[colIdx][7-rowNum] = whiteBishop
				colIdx++
			} else if fenValue == 'K' {
				b.squares[colIdx][7-rowNum] = whiteKing
				colIdx++
			} else if fenValue == 'Q' {
				b.squares[colIdx][7-rowNum] = whiteQueen
				colIdx++
			} else if fenValue == 'P' {
				b.squares[colIdx][7-rowNum] = whitePawn
				colIdx++
			} else if fenValue == 'r' {
				b.squares[colIdx][7-rowNum] = blackRook
				colIdx++
			} else if fenValue == 'n' {
				b.squares[colIdx][7-rowNum] = blackKnight
				colIdx++
			} else if fenValue == 'b' {
				b.squares[colIdx][7-rowNum] = blackBishop
				colIdx++
			} else if fenValue == 'k' {
				b.squares[colIdx][7-rowNum] = blackKing
				colIdx++
			} else if fenValue == 'q' {
				b.squares[colIdx][7-rowNum] = blackQueen
				colIdx++
			} else if fenValue == 'p' {
				b.squares[colIdx][7-rowNum] = blackPawn
				colIdx++
			} else if unicode.IsNumber(fenValue) {
				number := int(fenValue - '0')
				for i := colIdx; i < (colIdx + number); i++ {
					b.squares[i][7-rowNum] = empty
				}
				colIdx = colIdx + number
			}
		}
	}

	// Set active colour
	if fenActiveColour == "w" {
		b.colourToMove = white
	} else {
		b.colourToMove = black
	}

	// Set castling flags
	for _, flag := range fenCastling {
		if flag == 'K' {
			b.flags = setFlag(b.flags, whiteCanCastleKingSide)
		} else if flag == 'Q' {
			b.flags = setFlag(b.flags, whiteCanCastleQueenSide)
		} else if flag == 'k' {
			b.flags = setFlag(b.flags, blackCanCastleKingSide)
		} else if flag == 'q' {
			b.flags = setFlag(b.flags, blackCanCastleQueenSide)
		}
	}

	// TODO half move clock
	// TODO full move number

	return b
}

func newDefaultBoard() *board {
	return newBoardFromFen(startPositionFen)
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

	// Store the current board flags to restore on unmake
	m.boardFlags = b.flags

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
			b.flags = clearFlag(b.flags, whiteCanCastleKingSide)
			b.flags = clearFlag(b.flags, whiteCanCastleQueenSide)
			b.squares[6][m.fromY] = whiteKing
			b.squares[5][m.fromY] = whiteRook
		} else {
			b.flags = clearFlag(b.flags, blackCanCastleKingSide)
			b.flags = clearFlag(b.flags, blackCanCastleQueenSide)
			b.squares[6][m.fromY] = blackKing
			b.squares[5][m.fromY] = blackRook
		}

		return b
	} else if hasFlag(m.flags, queenCastle) {
		b.squares[4][m.fromY] = empty
		b.squares[0][m.fromY] = empty
		if m.fromY == 0 {
			b.flags = clearFlag(b.flags, whiteCanCastleKingSide)
			b.flags = clearFlag(b.flags, whiteCanCastleQueenSide)
			b.squares[2][m.fromY] = whiteKing
			b.squares[3][m.fromY] = whiteRook
		} else {
			b.flags = clearFlag(b.flags, blackCanCastleKingSide)
			b.flags = clearFlag(b.flags, blackCanCastleQueenSide)
			b.squares[2][m.fromY] = blackKing
			b.squares[3][m.fromY] = blackRook
		}

		return b
	}

	// If not castling, then set any flags for future castling.
	// Also set them in the move so we can unset them later.
	if fromPiece == whiteRook {
		if m.fromX == 0 && m.fromY == 0 {
			b.flags = clearFlag(b.flags, whiteCanCastleQueenSide)
		} else if m.fromX == 7 && m.fromY == 0 {
			b.flags = clearFlag(b.flags, whiteCanCastleKingSide)
		}
	} else if fromPiece == blackRook {
		if m.fromX == 0 && m.fromY == 7 {
			b.flags = clearFlag(b.flags, blackCanCastleQueenSide)
		} else if m.fromX == 7 && m.fromY == 7 {
			b.flags = clearFlag(b.flags, blackCanCastleKingSide)
		}
	} else if fromPiece == whiteKing {
		b.flags = clearFlag(b.flags, whiteCanCastleKingSide)
		b.flags = clearFlag(b.flags, whiteCanCastleQueenSide)
	} else if fromPiece == blackKing {
		b.flags = clearFlag(b.flags, blackCanCastleKingSide)
		b.flags = clearFlag(b.flags, blackCanCastleQueenSide)
	}

	// Make the standard move
	m.taken = b.squares[m.toX][m.toY]
	b.squares[m.toX][m.toY] = toPiece
	b.squares[m.fromX][m.fromY] = empty

	return b
}

/* This function unmakes a move, including any board-level flags set during a move. */
func unmakeMove(b *board, m *move) *board {
	defer func() {
		b.flags = m.boardFlags
	}()

	fromPiece := b.squares[m.toX][m.toY]

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
