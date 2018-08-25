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
				newBoardWithPieces([][]int{
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
				newBoardWithPieces([][]int{
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
				move{4, 1, 4, 2},
				move{4, 1, 4, 3},
			},
		},
		{
			"test take",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 6, 4, 5},
				move{4, 6, 4, 4},
				move{4, 6, 3, 5},
			},
		},
		{
			"test take in king diagonal",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 6, 3, 5},
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
				newBoardWithPieces([][]int{
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
				move{3, 3, 2, 3},
				move{3, 3, 1, 3},
				move{3, 3, 4, 3},
				move{3, 3, 3, 2},
				move{3, 3, 3, 4},
			},
		},
		{
			"test basic 2",
			args{
				newBoardWithPieces([][]int{
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
				move{3, 3, 2, 3},
				move{3, 3, 1, 3},
				move{3, 3, 4, 3},
				move{3, 3, 5, 3},
				move{3, 3, 6, 3},
				move{3, 3, 7, 3},
				move{3, 3, 3, 2},
				move{3, 3, 3, 4},
			},
		},
		{
			"test queen on row",
			args{
				newBoardWithPieces([][]int{
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
				move{5, 1, 4, 1},
				move{5, 1, 6, 1},
				move{5, 1, 7, 1},
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
				newBoardWithPieces([][]int{
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
				move{4, 1, 2, 0},
				move{4, 1, 2, 2},
				move{4, 1, 3, 3},
				move{4, 1, 6, 2},
				move{4, 1, 6, 0},
			},
		},
		{
			"test bounds",
			args{
				newBoardWithPieces([][]int{
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
				move{0, 0, 1, 2},
				move{0, 0, 2, 1},
			},
		},
		{
			"test rook on row",
			args{
				newBoardWithPieces([][]int{
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
				newBoardWithPieces([][]int{
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
				newBoardWithPieces([][]int{
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
				newBoardWithPieces([][]int{
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
				move{3, 4, 2, 3},
				move{3, 4, 1, 2},
				move{3, 4, 2, 5},
				move{3, 4, 1, 6},
				move{3, 4, 0, 7},
				move{3, 4, 4, 5},
			},
		},
		{
			"test rook on col",
			args{
				newBoardWithPieces([][]int{
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
				newBoardWithPieces([][]int{
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
				move{6, 5, 5, 4},
				move{6, 5, 7, 6},
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
				newBoardWithPieces([][]int{
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
				move{2, 3, 1, 3},
				move{2, 3, 0, 3},
				move{2, 3, 3, 3},
				move{2, 3, 2, 2},
				move{2, 3, 2, 4},
				move{2, 3, 2, 5},
				move{2, 3, 2, 6},
				move{2, 3, 1, 2},
				move{2, 3, 0, 1},
				move{2, 3, 3, 2},
				move{2, 3, 4, 1},
				move{2, 3, 5, 0},
				move{2, 3, 1, 4},
				move{2, 3, 0, 5},
				move{2, 3, 3, 4},
				move{2, 3, 4, 5},
				move{2, 3, 5, 6},
				move{2, 3, 6, 7},
			},
		},
		{
			"test rook on row",
			args{
				newBoardWithPieces([][]int{
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
				move{6, 3, 5, 3},
				move{6, 3, 7, 3},
			},
		},
		{
			"test bishop on diag",
			args{
				newBoardWithPieces([][]int{
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
				move{2, 5, 1, 5},
				move{2, 5, 0, 5},
				move{2, 5, 2, 4},
				move{2, 5, 2, 3},
				move{2, 5, 2, 6},
				move{2, 5, 2, 7},
				move{2, 5, 3, 4},
				move{2, 5, 1, 6},
				move{2, 5, 0, 7},
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
				newBoardWithPieces([][]int{
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
				move{4, 2, 3, 2},
				move{4, 2, 3, 3},
				move{4, 2, 4, 3},
				move{4, 2, 5, 3},
				move{4, 2, 5, 2},
				move{4, 2, 5, 1},
				move{4, 2, 4, 1},
				move{4, 2, 3, 1},
			},
		},
		{
			"test bishop on diag",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 2, 3, 3},
				move{4, 2, 5, 3},
				move{4, 2, 5, 2},
				move{4, 2, 5, 1},
				move{4, 2, 4, 1},
				move{4, 2, 3, 1},
			},
		},
		{
			"test rook on row",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 2, 3, 2},
				move{4, 2, 3, 3},
				move{4, 2, 4, 3},
				move{4, 2, 5, 3},
				move{4, 2, 5, 2},
				move{4, 2, 5, 1},
				move{4, 2, 4, 1},
				move{4, 2, 3, 1},
			},
		},
		{
			"test queen on col",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 2, 3, 2},
				move{4, 2, 3, 3},
				move{4, 2, 4, 3},
				move{4, 2, 5, 3},
				move{4, 2, 5, 2},
				move{4, 2, 5, 1},
				move{4, 2, 4, 1},
				move{4, 2, 3, 1},
			},
		},
		{
			"test knight",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 2, 3, 2},
				move{4, 2, 3, 3},
				move{4, 2, 4, 3},
				move{4, 2, 5, 3},
				move{4, 2, 5, 2},
				move{4, 2, 5, 1},
				move{4, 2, 4, 1},
				move{4, 2, 3, 1},
			},
		},
		{
			"test pawns",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 2, 3, 2},
				move{4, 2, 3, 3},
				move{4, 2, 4, 3},
				move{4, 2, 5, 3},
				move{4, 2, 5, 2},
				move{4, 2, 5, 1},
				move{4, 2, 4, 1},
				move{4, 2, 3, 1},
			},
		},
		{
			"test other king",
			args{
				newBoardWithPieces([][]int{
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
				move{4, 2, 3, 2},
				move{4, 2, 3, 3},
				move{4, 2, 4, 3},
				move{4, 2, 5, 3},
				move{4, 2, 5, 2},
				move{4, 2, 5, 1},
				move{4, 2, 4, 1},
				move{4, 2, 3, 1},
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
