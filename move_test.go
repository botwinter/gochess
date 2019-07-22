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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, blackQueen, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whitePawn, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteKing, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{3, 4}, [2]int{4, 4}, white, []move{},
			},
			[]move{},
		},
		{
			"test initial state",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 1}, [2]int{4, 0}, white, []move{},
			},
			[]move{
				move{4, 1, 4, 2, empty, 0, 0},
				move{4, 1, 4, 3, empty, 0, 0},
			},
		},
		{
			"test take",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, whiteBishop, empty, empty},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 6}, [2]int{4, 7}, black, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, whiteBishop, empty, empty},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 6}, [2]int{5, 7}, black, []move{},
			},
			[]move{
				move{4, 6, 3, 5, empty, 0, 0},
			},
		},
		{
			"test promotions",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, whitePawn, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, whiteBishop, empty, empty},
					[]int{whiteKing, whitePawn, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{1, 6}, [2]int{5, 7}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackPawn, empty, empty, empty, empty},
					[]int{empty, empty, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, whiteKing, empty, whiteRook, empty, whiteQueen, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, whiteBishop, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{3, 3}, [2]int{3, 1}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackPawn, empty, empty, empty, empty},
					[]int{empty, empty, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, whiteKing, empty, whiteRook, empty, whiteQueen, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{3, 3}, [2]int{3, 1}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackPawn, empty, empty, empty, empty},
					[]int{empty, empty, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, whiteKing, empty, empty, empty, whiteQueen, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteRook, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackQueen, empty, empty, empty, empty, empty, empty},
				}), [2]int{5, 1}, [2]int{3, 1}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackPawn, empty, whiteBishop, empty, empty},
					[]int{whiteKing, whiteKnight, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, whitePawn, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 1}, [2]int{4, 0}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{whiteKnight, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackPawn, empty, whiteBishop, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, whitePawn, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{0, 0}, [2]int{4, 0}, white, []move{},
			},
			[]move{
				move{0, 0, 1, 2, empty, 0, 0},
				move{0, 0, 2, 1, empty, 0, 0},
			},
		},
		{
			"test rook on row",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{blackRook, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKnight, empty, empty, blackPawn, empty, whiteBishop, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, whitePawn, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{3, 0}, [2]int{4, 0}, white, []move{},
			},
			[]move{},
		},
		{
			"test queen on col",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, whiteBishop, empty, empty},
					[]int{blackQueen, empty, whiteKnight, whiteKing, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, whitePawn, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, [2]int{4, 3}, white, []move{},
			},
			[]move{},
		},
		{
			"test bishop on diag",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteBishop, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackKnight, empty, empty, empty, empty, empty},
					[]int{blackQueen, empty, empty, blackKing, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, whitePawn, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{3, 2}, [2]int{4, 3}, black, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteBishop, empty, empty, empty},
					[]int{empty, empty, empty, whiteKing, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, whiteQueen, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{3, 4}, [2]int{4, 3}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, whiteKing, whiteBishop, empty, blackRook, empty},
					[]int{empty, empty, empty, empty, empty, empty, whiteQueen, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 4}, [2]int{4, 3}, white, []move{},
			},
			[]move{},
		},
		{
			"test queen on diag",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, whiteKing, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, whiteBishop, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, blackQueen, empty},
				}), [2]int{6, 5}, [2]int{4, 3}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whitePawn, empty, whiteQueen, empty, empty, blackPawn, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, whiteKing, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{2, 3}, [2]int{4, 3}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, whiteKing, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, whiteQueen, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackRook, empty, empty, empty, empty},
				}), [2]int{6, 3}, [2]int{4, 3}, white, []move{},
			},
			[]move{
				move{6, 3, 5, 3, empty, 0, 0},
				move{6, 3, 7, 3, empty, 0, 0},
			},
		},
		{
			"test bishop on diag",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, blackBishop},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackBishop, empty, whiteQueen, empty, empty},
					[]int{empty, empty, empty, empty, empty, whitePawn, empty, empty},
					[]int{empty, empty, empty, whiteKing, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{2, 5}, [2]int{4, 3}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackBishop, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, white, []move{},
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
			"test queen on col",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, white, []move{},
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
			"test knight",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, white, []move{},
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
			"test pawns",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, white, []move{},
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
			"test other king",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 2}, white, []move{},
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
			"test white castle",
			args{
				newBoardFromCoords([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 0}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
				}), [2]int{4, 7}, black, []move{},
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
				newBoardFromCoords([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, blackRook, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
				}), [2]int{4, 0}, white, []move{},
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, whiteKnight, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
				}), [2]int{4, 7}, black, []move{},
			},
			[]move{
				move{4, 7, 5, 7, empty, 0, 0},
				move{4, 7, 5, 6, empty, 0, 0},
				move{4, 7, 3, 6, empty, 0, 0},
				move{4, 7, 6, 7, empty, kingCastle, 0},
			},
		},
		{
			"test castle if king has moved",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
				}), [2]int{3, 7}, black, []move{},
			},
			[]move{
				move{3, 7, 2, 7, empty, 0, 0},
				move{3, 7, 4, 7, empty, 0, 0},
				move{3, 7, 4, 6, empty, 0, 0},
				move{3, 7, 3, 6, empty, 0, 0},
				move{3, 7, 2, 6, empty, 0, 0},
			},
		},
		{
			"test castle if rook taken",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, whiteBishop},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
				}), [2]int{4, 7}, black, []move{},
			},
			[]move{
				move{4, 7, 3, 7, empty, 0, 0},
				move{4, 7, 5, 7, empty, 0, 0},
				move{4, 7, 5, 6, empty, 0, 0},
				move{4, 7, 4, 6, empty, 0, 0},
				move{4, 7, 3, 6, empty, 0, 0},
				move{4, 7, 6, 7, empty, kingCastle, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateKingMoves(tt.args.b, tt.args.kingCoords, tt.args.colour, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateKingMoves() = %v, want %v", got, tt.want)
			}
		})
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			false,
		},
		{
			"test queen",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, blackQueen, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			true,
		},
		{
			"test queen blocked",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, whitePawn, blackQueen, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			false,
		},
		{
			"test bishop",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteBishop, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), black,
			},
			true,
		},
		{
			"test bishop blocked",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, whiteBishop, empty, empty, empty},
					[]int{empty, empty, empty, whiteRook, empty, empty, empty, empty},
					[]int{empty, empty, blackKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), black,
			},
			false,
		},
		{
			"test knight",
			args{
				newBoardFromCoords([][]int{
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackKnight, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			true,
		},
		{
			"test knight adjacent",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackKnight, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			false,
		},
		{
			"test rook",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackKing, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteRook, empty, empty, empty, empty, empty, empty},
				}), black,
			},
			true,
		},
		{
			"test rook blocked",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackKing, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackBishop, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteRook, empty, empty, empty, empty, empty, empty},
				}), black,
			},
			false,
		},
		{
			"test pawn",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, whiteKing, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			true,
		},
		{
			"test pawn in front",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteKing, blackPawn, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			false,
		},
		{
			"test pawn behind",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackPawn, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			false,
		},
		{
			"test other king",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whiteKing, empty, blackKing, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, blackQueen},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, blackRook, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, blackQueen},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
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
				newBoardFromCoords([][]int{
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackQueen, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackRook, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
			},
			[]move{},
		},
		{
			"test bishop",
			args{
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, blackBishop, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, blackBishop, empty, empty, empty, blackKing},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
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
				newBoardFromCoords([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteKnight, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteRook, empty, blackKnight, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, whiteKing, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), white,
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
