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
				newBoardFromFen("3k4/8/8/3q4/8/3Q4/8/3K4 w - - 0 1"),
				white,
			},
			move{3, 2, 3, 4, blackQueen, 0, 0},
		},
		{
			"test basic queen taking king",
			args{
				newBoardFromFen("3k4/8/1Q6/8/8/8/8/3K4 w - - 0 1"),
				white,
			},
			move{1, 5, 3, 7, blackKing, 0, 0},
		},
		{
			"test 4 move checkmate",
			args{
				newBoardFromFen("r1bqk1nr/pppp1ppp/2n5/2b1p3/2B1P3/5Q2/PPPP1PPP/RNB1K1NR w - - 0 1"),
				white,
			},
			move{5, 2, 5, 6, blackPawn, 0, 0},
		},
		{
			"test queens and kings",
			args{
				newBoardFromFen("3k4/8/8/3q4/8/8/2Q5/3K4 w - - 0 1"),
				white,
			},
			move{3, 0, 2, 0, empty, 0, 0},
		},
		{
			"test mate in one (1)",
			args{
				newBoardFromFen("r2q1rk1/1p2bp2/ppn5/2pP4/2P5/P2B4/1BQ2PPP/R4RK1 w - - 0 1"),
				white,
			},
			move{3, 2, 7, 6, empty, 0, 0},
		},
		{
			"test pawn promotion to queen",
			args{
				newBoardFromFen("3k4/P7/8/8/8/8/3B4/3K4 w - - 0 1"),
				white,
			},
			move{0, 6, 0, 7, empty, queenPromotion, 0},
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
