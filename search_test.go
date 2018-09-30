package main

import (
	"reflect"
	"testing"
)

func Test_findBestMove(t *testing.T) {
	type args struct {
		b      *board
		colour int
	}
	tests := []struct {
		name string
		args args
		want move
	}{
		{
			"test basic queen taking queen",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, whiteQueen, empty, blackQueen, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			move{3, 2, 3, 4, blackQueen},
		},
		{
			"test basic queen taking king",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, whiteQueen, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			move{1, 5, 3, 7, blackKing},
		},
		{
			"test 4 move checkmate",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, empty},
					[]int{whiteBishop, whitePawn, empty, whiteBishop, blackBishop, blackKnight, blackPawn, blackBishop},
					[]int{empty, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
					[]int{whiteKing, empty, empty, whitePawn, blackPawn, empty, empty, blackKing},
					[]int{empty, whitePawn, whiteQueen, empty, empty, empty, blackPawn, empty},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				}), white,
			},
			move{5, 2, 5, 6, blackPawn},
		},
		{
			"test queens and kings",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteQueen, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, blackQueen, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			move{3, 0, 2, 0, empty},
		},
		/*{
			"test mate in one (1)",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, empty, whitePawn, empty, empty, blackPawn, empty, blackRook},
					[]int{empty, whiteBishop, empty, empty, empty, blackPawn, blackPawn, empty},
					[]int{empty, whiteQueen, empty, whitePawn, blackPawn, blackKnight, empty, empty},
					[]int{empty, empty, whiteBishop, empty, whitePawn, empty, empty, blackQueen},
					[]int{empty, empty, empty, empty, empty, empty, blackBishop, empty},
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, whitePawn, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			move{3, 2, 7, 6, empty},
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestMove(tt.args.b, tt.args.colour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBestMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
