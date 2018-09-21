package main

import "fmt"

func main() {
	fmt.Printf("chess!\n")

	b := newBoardWithPieces([][]int{
		[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
		[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
		[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
		[]int{whiteQueen, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
		[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, blackKing},
		[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
		[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
		[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
	})

	moves := generateAllLegalMoves(b, white)
	for _, m := range moves {
		fmt.Printf("white move: %d:%d %d:%d\n", m.fromX, m.fromY, m.toX, m.toY)
	}

	moves = generateAllLegalMoves(b, black)
	for _, m := range moves {
		fmt.Printf("black move: %d:%d %d:%d\n", m.fromX, m.fromY, m.toX, m.toY)
	}
}
