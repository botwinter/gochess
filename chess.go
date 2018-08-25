package main

import "fmt"

func main() {
	fmt.Printf("chess!\n")

	b := newBoardWithPieces([][]int{
		[]int{empty, whitePawn, whiteRook, empty, empty, empty, empty, whitePawn},
		[]int{empty, empty, blackRook, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, blackRook, empty},
		[]int{empty, whitePawn, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, blackPawn, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
	})

	moves := generateAllLegalMoves(b)
	for _, m := range moves {
		fmt.Printf("%d:%d %d:%d\n", m.fromX, m.fromY, m.toX, m.toY)
	}
}
