package main

import (
	"fmt"
	"reflect"
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
				newBoardFromFen("8/2q5/3P4/4K3/8/8/8 w - - 0 1"),
				[2]int{4, 4},
				[2]int{3, 5},
				white,
				topLeftDiagonal,
			},
			true,
		},
		{
			"test Rook on diagonal",
			args{
				newBoardFromFen("8/2r5/3P4/4K3/8/8/8 w - - 0 1"),
				[2]int{4, 4},
				[2]int{3, 5},
				white,
				topLeftDiagonal,
			},
			false,
		},
		{
			"test another piece in way",
			args{
				newBoardFromFen("1q6/2p5/3P4/4K3/8/8/8 w - - 0 1"),
				[2]int{4, 4},
				[2]int{3, 5},
				white,
				topLeftDiagonal,
			},
			false,
		},
		{
			"test bishop",
			args{
				newBoardFromFen("8/8/8/4K3/5P2/8/7B/8 w - - 0 1"),
				[2]int{4, 4},
				[2]int{5, 3},
				white,
				topLeftDiagonal,
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
				newBoardFromFen("8/8/8/1q1PK3/8/8/8 w - - 0 1"),
				[2]int{4, 4},
				[2]int{3, 4},
				white,
				left,
			},
			true,
		},
		{
			"test Rook up",
			args{
				newBoardFromFen("8/4r3/4P3/4K3/8/8/8 w - - 0 1"),
				[2]int{4, 4},
				[2]int{4, 5},
				white,
				up,
			},
			true,
		},
		{
			"test Rook left with other piece in way",
			args{
				newBoardFromFen("8/8/8/4KPbr/8/8/8 w - - 0 1"),
				[2]int{4, 4},
				[2]int{5, 4},
				white,
				right,
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

func Test_makeMove(t *testing.T) {
	type args struct {
		b *board
		m *move
	}
	tests := []struct {
		name string
		args args
		want *board
	}{
		{
			"test simple",
			args{
				newBoardFromFen("8/8/8/8/8/8/3P4/8 w - - 0 1"),
				&move{3, 1, 3, 2, empty, 0, 0},
			},
			newBoardFromFen("8/8/8/8/8/3P4/8/8 w - - 0 1"),
		},
		{
			"test take",
			args{
				newBoardFromFen("8/8/8/8/8/8/3R1p2/8 w - - 0 1"),
				&move{3, 1, 5, 1, empty, 0, 0},
			},
			newBoardFromFen("8/8/8/8/8/8/5R2/8 w - - 0 1"),
		},
		{
			"test castle king side",
			args{
				newBoardFromFen("r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1"),
				&move{4, 0, 6, 0, empty, kingCastle, 0},
			},
			newBoardFromFen("r3k2r/8/8/8/8/8/8/R4RK1 w kq - 0 1"),
		},
		{
			"test castle queen side",
			args{
				newBoardFromFen("r3k2r/8/8/8/8/8/8/R3K2R w - - 0 1"),
				&move{4, 7, 2, 7, empty, queenCastle, 0},
			},
			newBoardFromFen("2kr3r/8/8/8/8/8/8/R3K2R w - - 0 1"),
		},
		{
			"test pawn promotion",
			args{
				newBoardFromFen("r3k2r/1P6/8/8/8/8/8/R3K2R w - - 0 1"),
				&move{1, 6, 0, 7, empty, queenPromotion, 0},
			},
			newBoardFromFen("Q3k2r/8/8/8/8/8/8/R3K2R w - - 0 1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeMove(tt.args.b, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeMove() = %v, want %v", got, tt.want)
				fmt.Println("Got:")
				prettyPrintBoard(got)
				fmt.Println("Want:")
				prettyPrintBoard(tt.want)
			}
		})
	}
}

func Test_unmakeMove(t *testing.T) {
	type args struct {
		b *board
		m *move
	}
	tests := []struct {
		name string
		args args
		want *board
	}{
		{
			"test simple",
			args{
				newBoardFromFen("8/8/8/8/8/3P4/8/8 w - - 0 1"),
				&move{3, 1, 3, 2, empty, 0, 0},
			},
			newBoardFromFen("8/8/8/8/8/8/3P4/8 w - - 0 1"),
		},
		{
			"test take",
			args{
				newBoardFromFen("8/8/8/8/8/3q4/8/8 w - - 0 1"),
				&move{1, 0, 3, 2, whiteBishop, 0, 0},
			},
			newBoardFromFen("8/8/8/8/8/3B4/8/1q6 w - - 0 1"),
		},
		{
			"test castle king side",
			args{
				newBoardFromFen("r3k2r/8/8/8/8/8/8/R4RK1 w - - 0 1"),
				&move{4, 0, 6, 0, empty, kingCastle, 0},
			},
			newBoardFromFen("r3k2r/8/8/8/8/8/8/R3K2R w - - 0 1"),
		},
		{
			"test castle queen side",
			args{
				newBoardFromFen("2kr3r/8/8/8/8/8/8/R3K2R w - - 0 1"),
				&move{4, 7, 2, 7, empty, queenCastle, 0},
			},
			newBoardFromFen("r3k2r/8/8/8/8/8/8/R3K2R w - - 0 1"),
		},
		{
			"test pawn promotion",
			args{
				newBoardFromFen("r3k2r/8/8/8/8/8/8/R1r1K2R w - - 0 1"),
				&move{2, 1, 2, 0, empty, rookPromotion, 0},
			},
			newBoardFromFen("r3k2r/8/8/8/8/8/2p5/R3K2R w - - 0 1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unmakeMove(tt.args.b, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmakeMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeThenUnmake(t *testing.T) {
	b := newBoardFromFen("r3k2r/8/8/8/8/8/8/R3K2R w KQ - 0 1")

	want := newBoardFromFen("r3k2r/8/8/8/8/8/8/R3K2R w KQ - 0 1")

	m := move{4, 0, 6, 0, empty, kingCastle, 0}
	t.Run("test make then unmake castle", func(t *testing.T) {
		if !hasFlag(b.flags, whiteCanCastleKingSide) {
			t.Errorf("whiteCanCastleKingSide flag is not set")
		}

		if !hasFlag(b.flags, whiteCanCastleQueenSide) {
			t.Errorf("whiteCanCastleQueenSide flag is not set")
		}

		makeMove(b, &m)

		if hasFlag(b.flags, whiteCanCastleKingSide) {
			t.Errorf("whiteCanCastleKingSide flag is set")
		}

		if hasFlag(b.flags, whiteCanCastleQueenSide) {
			t.Errorf("whiteCanCastleQueenSide flag is set")
		}

		if !hasFlag(m.boardFlags, whiteCanCastleKingSide) {
			t.Errorf("whiteCanCastleKingSide move boardflag is not set")
		}

		if !hasFlag(m.boardFlags, whiteCanCastleQueenSide) {
			t.Errorf("whiteCanCastleQueenSide move boardflag is not set")
		}

		unmakeMove(b, &m)

		if !reflect.DeepEqual(b, want) {
			t.Errorf("got %v, want %v", b, want)
			fmt.Println("Got:")
			prettyPrintBoard(b)
			fmt.Println("Want:")
			prettyPrintBoard(want)
		}
	})
}

func Test_newBoardFromFen(t *testing.T) {
	type args struct {
		fen string
	}
	tests := []struct {
		name string
		args args
		want *board
	}{
		{
			"test Start position",
			args{
				"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			},
			newBoardFromCoords([][]int{
				[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
				[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
				[]int{whiteQueen, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
				[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, blackKing},
				[]int{whiteBishop, whitePawn, empty, empty, empty, empty, blackPawn, blackBishop},
				[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
				[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
			}, whiteCanCastleKingSide|whiteCanCastleQueenSide|blackCanCastleKingSide|blackCanCastleQueenSide, white),
		},
		{
			"another position",
			args{
				"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
			},
			newBoardFromCoords([][]int{
				[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
				[]int{whiteKnight, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
				[]int{whiteBishop, whitePawn, empty, empty, blackPawn, empty, empty, blackBishop},
				[]int{whiteQueen, whitePawn, empty, empty, empty, empty, blackPawn, blackQueen},
				[]int{whiteKing, empty, empty, whitePawn, empty, empty, blackPawn, blackKing},
				[]int{whiteBishop, whitePawn, whiteKnight, empty, empty, empty, blackPawn, blackBishop},
				[]int{empty, whitePawn, empty, empty, empty, empty, blackPawn, blackKnight},
				[]int{whiteRook, whitePawn, empty, empty, empty, empty, blackPawn, blackRook},
			}, whiteCanCastleKingSide|whiteCanCastleQueenSide|blackCanCastleKingSide|blackCanCastleQueenSide, black),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBoardFromFen(tt.args.fen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBoardFromFen() = %v, want %v", got, tt.want)
				fmt.Println("Got:")
				prettyPrintBoard(got)
				fmt.Println("Want:")
				prettyPrintBoard(tt.want)
			}
		})
	}
}
