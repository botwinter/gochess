package main

import "testing"

func Test_evaluateBoard(t *testing.T) {
	type args struct {
		b      *board
		colour int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test starting pos",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteQueen, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, blackKing},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				}), white,
			},
			0,
		},
		{
			"test white taken all",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteQueen, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			23905,
		},
		{
			"test black taken all",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackQueen},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackRook},
				}), white,
			},
			-23905,
		},
		{
			"test mix",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteQueen, empty, empty, blackPawn, empty, empty, blackPawn, blackQueen},
					[]int{whiteKing, empty, empty, whitePawn, empty, empty, empty, blackKing},
					[]int{empty, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				}), white,
			},
			-665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateBoard(tt.args.b, tt.args.colour); got != tt.want {
				t.Errorf("evaluateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
