package main

import (
	"reflect"
	"testing"
)

func Test_generatePawnMoves(t *testing.T) {
	type args struct {
		b          *board
		myCoords   [2]int
		kingCoords [2]int
		colour     int
		moves      []move
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{
			"test piece in way",
			args{
				newBoardFromFen("8/8/8/1q1PK3/8/8/8 w KQkq - 0 1"),
				[2]int{3, 4}, [2]int{4, 4}, white, []move{},
			},
			[]move{},
		},
		{
			"test initial state",
			args{
				newBoardFromFen("8/8/8/8/8/8/4P3/4K3 w KQkq - 0 1"),
				[2]int{4, 1}, [2]int{4, 0}, white, []move{},
			},
			[]move{
				move{4, 1, 4, 2, empty, 0, 0},
				move{4, 1, 4, 3, empty, 0, 0},
			},
		},
		{
			"test take",
			args{
				newBoardFromFen("4k3/4p3/3B4/8/8/8/4P3/4K3 w KQkq - 0 1"),
				[2]int{4, 6}, [2]int{4, 7}, black, []move{},
			},
			[]move{
				move{4, 6, 4, 5, empty, 0, 0},
				move{4, 6, 4, 4, empty, 0, 0},
				move{4, 6, 3, 5, empty, 0, 0},
			},
		},
		{
			"test take in king diagonal",
			args{
				newBoardFromFen("5k2/5p2/3B4/8/8/8/4P3/4K3 w KQkq - 0 1"),
				[2]int{4, 6}, [2]int{5, 7}, black, []move{},
			},
			[]move{
				move{4, 6, 3, 5, empty, 0, 0},
			},
		},
		{
			"test promotions",
			args{
				newBoardFromFen("5k2/1P2p3/3B4/8/8/8/4P3/4K3 w KQkq - 0 1"),
				[2]int{1, 6}, [2]int{5, 7}, white, []move{},
			},
			[]move{
				move{1, 6, 1, 7, empty, queenPromotion, 0},
				move{1, 6, 1, 7, empty, rookPromotion, 0},
				move{1, 6, 1, 7, empty, knightPromotion, 0},
				move{1, 6, 1, 7, empty, bishopPromotion, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePawnMoves(tt.args.b, tt.args.myCoords, tt.args.kingCoords, tt.args.colour, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePawnMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRookMoves(t *testing.T) {
	type args struct {
		b          *board
		myCoords   [2]int
		kingCoords [2]int
		colour     int
		moves      []move
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{
			"test basic",
			args{
				newBoardFromFen("8/8/3Q4/8/1p1R1B2/2p5/3K4/8 w KQkq - 0 1"),
				[2]int{3, 3}, [2]int{3, 1}, white, []move{},
			},
			[]move{
				move{3, 3, 2, 3, empty, 0, 0},
				move{3, 3, 1, 3, empty, 0, 0},
				move{3, 3, 4, 3, empty, 0, 0},
				move{3, 3, 3, 2, empty, 0, 0},
				move{3, 3, 3, 4, empty, 0, 0},
			},
		},
		{
			"test basic 2",
			args{
				newBoardFromFen("8/8/3Q4/8/1p1R4/2p5/3K4/8 w KQkq - 0 1"),
				[2]int{3, 3}, [2]int{3, 1}, white, []move{},
			},
			[]move{
				move{3, 3, 2, 3, empty, 0, 0},
				move{3, 3, 1, 3, empty, 0, 0},
				move{3, 3, 4, 3, empty, 0, 0},
				move{3, 3, 5, 3, empty, 0, 0},
				move{3, 3, 6, 3, empty, 0, 0},
				move{3, 3, 7, 3, empty, 0, 0},
				move{3, 3, 3, 2, empty, 0, 0},
				move{3, 3, 3, 4, empty, 0, 0},
			},
		},
		{
			"test queen on row",
			args{
				newBoardFromFen("8/8/3Q4/8/1p6/2p5/3K1R1q/8 w KQkq - 0 1"),
				[2]int{5, 1}, [2]int{3, 1}, white, []move{},
			},
			[]move{
				move{5, 1, 4, 1, empty, 0, 0},
				move{5, 1, 6, 1, empty, 0, 0},
				move{5, 1, 7, 1, empty, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRookMoves(tt.args.b, tt.args.myCoords, tt.args.kingCoords, tt.args.colour, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateRookMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateKnightMoves(t *testing.T) {
	type args struct {
		b          *board
		myCoords   [2]int
		kingCoords [2]int
		colour     int
		moves      []move
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{
			"test basic",
			args{
				newBoardFromFen("5k2/4p3/3B4/8/3p1P2/8/4N3/4K3 w KQkq - 0 1"),
				[2]int{4, 1}, [2]int{4, 0}, white, []move{},
			},
			[]move{
				move{4, 1, 2, 0, empty, 0, 0},
				move{4, 1, 2, 2, empty, 0, 0},
				move{4, 1, 3, 3, empty, 0, 0},
				move{4, 1, 6, 2, empty, 0, 0},
				move{4, 1, 6, 0, empty, 0, 0},
			},
		},
		{
			"test bounds",
			args{
				newBoardFromFen("5k2/4p3/3B4/8/3p1P2/8/8/N3K3 w KQkq - 0 1"),
				[2]int{0, 0}, [2]int{4, 0}, white, []move{},
			},
			[]move{
				move{0, 0, 1, 2, empty, 0, 0},
				move{0, 0, 2, 1, empty, 0, 0},
			},
		},
		{
			"test rook on row",
			args{
				newBoardFromFen("5k2/4p3/3B4/8/3p1P2/8/8/1r1NK3 w KQkq - 0 1"),
				[2]int{3, 0}, [2]int{4, 0}, white, []move{},
			},
			[]move{},
		},
		{
			"test queen on col",
			args{
				newBoardFromFen("5k2/4p3/3B4/8/4KP2/4N3/8/4q3 w KQkq - 0 1"),
				[2]int{4, 2}, [2]int{4, 3}, white, []move{},
			},
			[]move{},
		},
		{
			"test bishop on diag",
			args{
				newBoardFromFen("5k2/4p3/8/8/4kP2/3N4/8/1B2q3 w KQkq - 0 1"),
				[2]int{3, 2}, [2]int{4, 3}, black, []move{},
			},
			[]move{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateKnightMoves(tt.args.b, tt.args.myCoords, tt.args.kingCoords, tt.args.colour, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateKnightMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateBishopMoves(t *testing.T) {
	type args struct {
		b          *board
		myCoords   [2]int
		kingCoords [2]int
		colour     int
		moves      []move
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{
			"test basic",
			args{
				newBoardFromFen("8/5Q2/8/3B4/4K3/1p6/8/8 w KQkq - 0 1"),
				[2]int{3, 4}, [2]int{4, 3}, white, []move{},
			},
			[]move{
				move{3, 4, 2, 3, empty, 0, 0},
				move{3, 4, 1, 2, empty, 0, 0},
				move{3, 4, 2, 5, empty, 0, 0},
				move{3, 4, 1, 6, empty, 0, 0},
				move{3, 4, 0, 7, empty, 0, 0},
				move{3, 4, 4, 5, empty, 0, 0},
			},
		},
		{
			"test rook on col",
			args{
				newBoardFromFen("8/4r3/8/4B3/4K3/1p6/8/8 w KQkq - 0 1"),
				[2]int{4, 4}, [2]int{4, 3}, white, []move{},
			},
			[]move{},
		},
		{
			"test queen on diag",
			args{
				newBoardFromFen("8/7q/6B1/8/4K3/1p6/8/8 - 0 1"),
				[2]int{6, 5}, [2]int{4, 3}, white, []move{},
			},
			[]move{
				move{6, 5, 5, 4, empty, 0, 0},
				move{6, 5, 7, 6, empty, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateBishopMoves(tt.args.b, tt.args.myCoords, tt.args.kingCoords, tt.args.colour, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateBishopMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateQueenMoves(t *testing.T) {
	type args struct {
		b          *board
		myCoords   [2]int
		kingCoords [2]int
		colour     int
		moves      []move
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{
			"test basic",
			args{
				newBoardFromFen("8/2p5/8/8/2Q1K3/8/2P5/8 - 0 1"),
				[2]int{2, 3}, [2]int{4, 3}, white, []move{},
			},
			[]move{
				move{2, 3, 1, 3, empty, 0, 0},
				move{2, 3, 0, 3, empty, 0, 0},
				move{2, 3, 3, 3, empty, 0, 0},
				move{2, 3, 2, 2, empty, 0, 0},
				move{2, 3, 2, 4, empty, 0, 0},
				move{2, 3, 2, 5, empty, 0, 0},
				move{2, 3, 2, 6, empty, 0, 0},
				move{2, 3, 1, 2, empty, 0, 0},
				move{2, 3, 0, 1, empty, 0, 0},
				move{2, 3, 3, 2, empty, 0, 0},
				move{2, 3, 4, 1, empty, 0, 0},
				move{2, 3, 5, 0, empty, 0, 0},
				move{2, 3, 1, 4, empty, 0, 0},
				move{2, 3, 0, 5, empty, 0, 0},
				move{2, 3, 3, 4, empty, 0, 0},
				move{2, 3, 4, 5, empty, 0, 0},
				move{2, 3, 5, 6, empty, 0, 0},
				move{2, 3, 6, 7, empty, 0, 0},
			},
		},
		{
			"test rook on row",
			args{
				newBoardFromFen("8/8/8/8/4K1Qr/8/8/8 - 0 1"),
				[2]int{6, 3}, [2]int{4, 3}, white, []move{},
			},
			[]move{
				move{6, 3, 5, 3, empty, 0, 0},
				move{6, 3, 7, 3, empty, 0, 0},
			},
		},
		{
			"test bishop on diag",
			args{
				newBoardFromFen("b7/8/2QP4/8/2b1K3/8/8/8 - 0 1"),
				[2]int{2, 5}, [2]int{4, 3}, white, []move{},
			},
			[]move{
				move{2, 5, 1, 5, empty, 0, 0},
				move{2, 5, 0, 5, empty, 0, 0},
				move{2, 5, 2, 4, empty, 0, 0},
				move{2, 5, 2, 3, empty, 0, 0},
				move{2, 5, 2, 6, empty, 0, 0},
				move{2, 5, 2, 7, empty, 0, 0},
				move{2, 5, 3, 4, empty, 0, 0},
				move{2, 5, 1, 6, empty, 0, 0},
				move{2, 5, 0, 7, empty, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateQueenMoves(tt.args.b, tt.args.myCoords, tt.args.kingCoords, tt.args.colour, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateQueenMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateKingMoves(t *testing.T) {
	type args struct {
		b          *board
		kingCoords [2]int
		colour     int
		moves      []move
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{
			"test basic",
			args{
				newBoardFromFen("8/8/8/8/8/4K3/8/8 w - - 0 1"),
				[2]int{4, 2}, white, []move{},
			},
			[]move{
				move{4, 2, 3, 2, empty, 0, 0},
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 4, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test bishop on diag",
			args{
				newBoardFromFen("8/8/8/8/8/4K3/2b5/8 w - - 0 1"),
				[2]int{4, 2}, white, []move{},
			},
			[]move{
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test rook on row",
			args{
				newBoardFromFen("8/8/8/8/r7/4K3/8/8 w - - 0 1"),
				[2]int{4, 2}, white, []move{},
			},
			[]move{
				move{4, 2, 3, 2, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test queen on col",
			args{
				newBoardFromFen("8/3q4/8/8/8/4K3/8/8 w - - 0 1"),
				[2]int{4, 2}, white, []move{},
			},
			[]move{
				move{4, 2, 4, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
			},
		},
		{
			"test knight",
			args{
				newBoardFromFen("8/8/8/2n5/8/4K3/8/8 w - - 0 1"),
				[2]int{4, 2}, white, []move{},
			},
			[]move{
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test pawns",
			args{
				newBoardFromFen("8/8/8/6p1/4p3/4K3/8/8 w - - 0 1"),
				[2]int{4, 2}, white, []move{},
			},
			[]move{
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 4, 3, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test other king",
			args{
				newBoardFromFen("8/8/8/8/6k1/4K3/8/8 w - - 0 1"),
				[2]int{4, 2}, white, []move{},
			},
			[]move{
				move{4, 2, 3, 2, empty, 0, 0},
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 4, 3, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test white castle",
			args{
				newBoardFromFen("8/8/8/8/8/8/8/R3K2R w KQ - 0 1"),
				[2]int{4, 0}, white, []move{},
			},
			[]move{
				move{4, 0, 3, 0, empty, 0, 0},
				move{4, 0, 3, 1, empty, 0, 0},
				move{4, 0, 4, 1, empty, 0, 0},
				move{4, 0, 5, 1, empty, 0, 0},
				move{4, 0, 5, 0, empty, 0, 0},
				move{4, 0, 6, 0, empty, kingCastle, 0},
				move{4, 0, 2, 0, empty, queenCastle, 0},
			},
		},
		{
			"test black castle",
			args{
				newBoardFromFen("r3k2r/8/8/8/8/8/8/4K3 b kq - 0 1"),
				[2]int{4, 7}, black, []move{},
			},
			[]move{
				move{4, 7, 3, 7, empty, 0, 0},
				move{4, 7, 5, 7, empty, 0, 0},
				move{4, 7, 5, 6, empty, 0, 0},
				move{4, 7, 4, 6, empty, 0, 0},
				move{4, 7, 3, 6, empty, 0, 0},
				move{4, 7, 6, 7, empty, kingCastle, 0},
				move{4, 7, 2, 7, empty, queenCastle, 0},
			},
		},
		{
			"test castle through check",
			args{
				newBoardFromFen("4kr2/8/8/8/8/8/8/R3K2R w KQ - 0 1"),
				[2]int{4, 0}, white, []move{},
			},
			[]move{
				move{4, 0, 3, 0, empty, 0, 0},
				move{4, 0, 3, 1, empty, 0, 0},
				move{4, 0, 4, 1, empty, 0, 0},
				move{4, 0, 2, 0, empty, queenCastle, 0},
			},
		},
		{
			"test castle through check 2",
			args{
				newBoardFromFen("r3k2r/8/2N5/8/8/8/8/8 b kq - 0 1"),
				[2]int{4, 7}, black, []move{},
			},
			[]move{
				move{4, 7, 5, 7, empty, 0, 0},
				move{4, 7, 5, 6, empty, 0, 0},
				move{4, 7, 3, 6, empty, 0, 0},
				move{4, 7, 6, 7, empty, kingCastle, 0},
			},
		},
	}
	for _, tt := range tests {
		if tt.name == "test castle through check" {
			t.Run(tt.name, func(t *testing.T) {
				if got := generateKingMoves(tt.args.b, tt.args.kingCoords, tt.args.colour, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
					prettyPrintBoard(tt.args.b)
					t.Errorf("generateKingMoves() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func Test_inCheck(t *testing.T) {
	type args struct {
		b      *board
		colour int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test empty",
			args{
				newBoardFromFen("8/8/8/8/8/4K3/8/8 w - - 0 1"),
				white,
			},
			false,
		},
		{
			"test queen",
			args{
				newBoardFromFen("8/8/8/4q3/8/4K3/8/8 w - - 0 1"),
				white,
			},
			true,
		},
		{
			"test queen blocked",
			args{
				newBoardFromFen("8/8/8/4q3/4P3/4K3/8/8 w - - 0 1"),
				white,
			},
			false,
		},
		{
			"test bishop",
			args{
				newBoardFromFen("8/8/8/2B5/8/4k3/8/8 b - - 0 1"),
				black,
			},
			true,
		},
		{
			"test bishop blocked",
			args{
				newBoardFromFen("8/8/8/2B5/3R4/4k3/8/8 b - - 0 1"),
				black,
			},
			false,
		},
		{
			"test knight",
			args{
				newBoardFromFen("8/8/8/8/8/1n6/8/K7 w - - 0 1"),
				white,
			},
			true,
		},
		{
			"test knight adjacent",
			args{
				newBoardFromFen("8/8/8/8/8/1nK5/8/8 w - - 0 1"),
				white,
			},
			false,
		},
		{
			"test rook",
			args{
				newBoardFromFen("8/8/8/8/8/4k2R/8/8 b - - 0 1"),
				black,
			},
			true,
		},
		{
			"test rook blocked",
			args{
				newBoardFromFen("8/8/8/8/8/4kb1R/8/8 b - - 0 1"),
				black,
			},
			false,
		},
		{
			"test pawn",
			args{
				newBoardFromFen("8/8/8/8/3p4/4K3/8/8 w - - 0 1"),
				white,
			},
			true,
		},
		{
			"test pawn in front",
			args{
				newBoardFromFen("8/8/8/8/4p3/4K3/8/8 w - - 0 1"),
				white,
			},
			false,
		},
		{
			"test pawn behind",
			args{
				newBoardFromFen("8/8/8/8/8/4K3/3p4/8 w - - 0 1"),
				white,
			},
			false,
		},
		{
			"test other king",
			args{
				newBoardFromFen("8/8/8/4k3/8/4K3/8/8 w - - 0 1"),
				white,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inCheck(tt.args.b, tt.args.colour); got != tt.want {
				t.Errorf("inCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateAllLegalMovesInCheck(t *testing.T) {
	type args struct {
		b      *board
		colour int
	}
	tests := []struct {
		name string
		args args
		want []move
	}{
		{
			"test basic",
			args{
				newBoardFromFen("3kq3/8/8/8/8/4K3/8/8 w - - 0 1"),
				white,
			},
			[]move{
				move{4, 2, 3, 2, empty, 0, 0},
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test queen and rook",
			args{
				newBoardFromFen("3kq3/8/8/8/8/r3K3/8/8 w - - 0 1"),
				white,
			},
			[]move{
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test checkmate",
			args{
				newBoardFromFen("3k4/8/8/8/8/8/1qr5/K7 w - - 0 1"),
				white,
			},
			[]move{},
		},
		{
			"test bishop",
			args{
				newBoardFromFen("3k4/8/8/2b5/8/4K3/8/8 w - - 0 1"),
				white,
			},
			[]move{
				move{4, 2, 3, 2, empty, 0, 0},
				move{4, 2, 4, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test bishop adjacent",
			args{
				newBoardFromFen("3k4/8/8/8/3b4/4K3/8/8 w - - 0 1"),
				white,
			},
			[]move{
				move{4, 2, 3, 2, empty, 0, 0},
				move{4, 2, 3, 3, blackBishop, 0, 0},
				move{4, 2, 4, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
				move{4, 2, 3, 1, empty, 0, 0},
			},
		},
		{
			"test multiple pieces taking checking piece",
			args{
				newBoardFromFen("3k4/8/8/8/2n5/4K3/1NR5/8 w - - 0 1"),
				white,
			},
			[]move{
				move{1, 1, 2, 3, blackKnight, 0, 0},
				move{2, 1, 2, 3, blackKnight, 0, 0},
				move{4, 2, 3, 2, empty, 0, 0},
				move{4, 2, 3, 3, empty, 0, 0},
				move{4, 2, 4, 3, empty, 0, 0},
				move{4, 2, 5, 3, empty, 0, 0},
				move{4, 2, 5, 2, empty, 0, 0},
				move{4, 2, 5, 1, empty, 0, 0},
				move{4, 2, 4, 1, empty, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateAllLegalMovesInCheck(tt.args.b, tt.args.colour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateAllLegalMovesInCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
