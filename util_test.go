package main

import (
	"reflect"
	"testing"
)

func Test_reverseBoardArray(t *testing.T) {
	type args struct {
		array [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test reversing a board",
			args{
				[][]int{
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteQueen, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
					[]int{whiteKing, empty, empty, whitePawn, empty, empty, blackPawn, blackKing},
					[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
					[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
					[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				},
			},
			[][]int{
				[]int{blackRook, blackPawn, empty, empty, empty, empty, whitePawn, whiteRook},
				[]int{blackKnight, blackPawn, empty, empty, empty, empty, whitePawn, whiteKnight},
				[]int{blackBishop, blackPawn, empty, empty, empty, empty, whitePawn, whiteBishop},
				[]int{blackKing, blackPawn, empty, empty, whitePawn, empty, empty, whiteKing},
				[]int{blackQueen, blackPawn, empty, empty, empty, empty, whitePawn, whiteQueen},
				[]int{blackBishop, blackPawn, empty, empty, empty, empty, whitePawn, whiteBishop},
				[]int{blackKnight, blackPawn, empty, empty, empty, empty, whitePawn, whiteKnight},
				[]int{blackRook, blackPawn, empty, empty, empty, empty, whitePawn, whiteRook},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBoardArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
