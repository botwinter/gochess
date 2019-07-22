package main

import "testing"

func Test_evaluateBoard(t *testing.T) {
	type args struct {
		b *board
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test starting pos",
			args{
				newBoardFromCoords([][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteQueen, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, blackKing},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				}),
			},
			0,
		},
		{
			"test white taken all",
			args{
				newBoardFromCoords([][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteQueen, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, empty, empty},
				}),
			},
			23905,
		},
		{
			"test black taken all",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackQueen},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackRook},
				}),
			},
			-23905,
		},
		{
			"test mix",
			args{
				newBoardFromCoords([][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteQueen, empty, empty, blackPawn, empty, empty, blackPawn, blackQueen},
					[]int{whiteKing, empty, empty, whitePawn, empty, empty, empty, blackKing},
					[]int{empty, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				}),
			},
			-665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateBoard(tt.args.b); got != tt.want {
				t.Errorf("evaluateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
