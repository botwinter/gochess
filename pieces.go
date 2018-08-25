package main

const (
	empty = iota
	whitePawn
	whiteBishop
	whiteKnight
	whiteRook
	whiteQueen
	whiteKing
	blackPawn
	blackBishop
	blackKnight
	blackRook
	blackQueen
	blackKing
)

const (
	black = -1
	white = 1
)

const (
	left = iota
	right
	up
	down
)
const (
	bottomLeftDiagonal = iota
	topLeftDiagonal
	topRightDiagonal
	bottomRightDiagonal
)

const ()

func isWhite(piece int) bool {
	if piece == whitePawn || piece == whiteBishop || piece == whiteKnight || piece == whiteRook || piece == whiteQueen || piece == whiteKing {
		return true
	}
	return false
}

func isBlack(piece int) bool {
	if piece == blackPawn || piece == blackBishop || piece == blackKnight || piece == blackRook || piece == blackQueen || piece == blackKing {
		return true
	}
	return false
}
