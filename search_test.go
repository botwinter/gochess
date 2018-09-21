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
			"test basic",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, whiteQueen, empty, blackPawn, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			move{3, 2, 3, 4, blackPawn},
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
