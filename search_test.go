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
			"test queen moving to defend king",
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
			move{2, 1, 3, 1, empty},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestMove(tt.args.b, tt.args.colour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBestMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
