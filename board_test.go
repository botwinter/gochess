package main

import (
	"testing"
)

func Test_isEnemyOnDiagonal(t *testing.T) {
	type args struct {
		b           *board
		kingCoords  [2]int
		pieceCoords [2]int
		colour      int
		diagonal    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test Queen on diagonal",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, blackQueen, empty},
					[]int{empty, empty, empty, empty, empty, whitePawn, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{3, 5}, white, topLeftDiagonal,
			},
			true,
		},
		{
			"test Rook on diagonal",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, blackRook, empty},
					[]int{empty, empty, empty, empty, empty, whitePawn, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{3, 5}, white, topLeftDiagonal,
			},
			false,
		},
		{
			"test another piece in way",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackQueen},
					[]int{empty, empty, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, empty, empty, whitePawn, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{3, 5}, white, topLeftDiagonal,
			},
			false,
		},
		{
			"test bishop",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
					[]int{empty, empty, empty, whitePawn, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteBishop, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{5, 3}, white, topLeftDiagonal,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEnemyOnDiagonal(tt.args.b, tt.args.kingCoords, tt.args.pieceCoords, tt.args.colour, tt.args.diagonal); got != tt.want {
				t.Errorf("isEnemyOnDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEnemyInDirection(t *testing.T) {
	type args struct {
		b           *board
		kingCoords  [2]int
		pieceCoords [2]int
		colour      int
		direction   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test Queen left",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, blackQueen, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whitePawn, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{3, 4}, white, left,
			},
			true,
		},
		{
			"test Rook up",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, whitePawn, blackRook, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{4, 5}, white, up,
			},
			true,
		},
		{
			"test Rook left with other piece in way",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
					[]int{empty, empty, empty, empty, whitePawn, empty, empty, empty},
					[]int{empty, empty, empty, empty, blackBishop, empty, empty, empty},
					[]int{empty, empty, empty, empty, blackRook, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{5, 4}, white, right,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEnemyInDirection(tt.args.b, tt.args.kingCoords, tt.args.pieceCoords, tt.args.colour, tt.args.direction); got != tt.want {
				t.Errorf("isEnemyInDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
