package main

import (
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
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whitePawn, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), &move{3, 1, 3, 2, empty, 0, 0},
			}, newBoardWithPieces([][]int{
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, whitePawn, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			}),
		},
		{
			"test take",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, whiteRook, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, blackPawn, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), &move{3, 1, 5, 1, empty, 0, 0},
			}, newBoardWithPieces([][]int{
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, whiteRook, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			}),
		},
		{
			"test castle king side",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				}), &move{4, 0, 6, 0, empty, kingCastle, 0},
			}, &board{[][]int{
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
			}, whiteKingMoved | whiteRookKingSideMoved},
		},
		{
			"test castle queen side",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				}), &move{4, 7, 2, 7, empty, queenCastle, 0},
			}, &board{[][]int{
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
				[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
				[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
			}, blackKingMoved | blackRookQueenSideMoved},
		},
		{
			"test pawn promotion",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, whitePawn, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				}), &move{1, 6, 0, 7, empty, queenPromotion, 0},
			}, &board{[][]int{
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, whiteQueen},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
			}, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeMove(tt.args.b, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeMove() = %v, want %v", got, tt.want)
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
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, whitePawn, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), &move{3, 1, 3, 2, empty, 0, 0},
			}, newBoardWithPieces([][]int{
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, whitePawn, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			}),
		},
		{
			"test take",
			args{
				newBoardWithPieces([][]int{
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, blackQueen, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				}), &move{1, 0, 3, 2, whiteBishop, 0, 0},
			}, newBoardWithPieces([][]int{
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{blackQueen, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, whiteBishop, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
			}),
		},
		{
			"test castle king side",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
				}), &move{4, 0, 6, 0, empty, kingCastle, whiteRookKingSideMoved | whiteKingMoved},
			}, &board{[][]int{
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
			}, 0},
		},
		{
			"test castle queen side",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				}), &move{4, 7, 2, 7, empty, queenCastle, blackKingMoved | blackRookQueenSideMoved},
			}, &board{[][]int{
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
			}, 0},
		},
		{
			"test pawn promotion",
			args{
				newBoardWithPieces([][]int{
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{blackRook, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{empty, empty, empty, empty, empty, empty, empty, empty},
					[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				}), &move{2, 1, 2, 0, empty, rookPromotion, 0},
			}, &board{[][]int{
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, blackPawn, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{empty, empty, empty, empty, empty, empty, empty, empty},
				[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
			}, 0},
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
	b := newBoardWithPieces([][]int{
		[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
	})

	want := &board{[][]int{
		[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{whiteKing, empty, empty, empty, empty, empty, empty, blackKing},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{empty, empty, empty, empty, empty, empty, empty, empty},
		[]int{whiteRook, empty, empty, empty, empty, empty, empty, blackRook},
	}, 0}

	m := move{4, 0, 6, 0, empty, kingCastle, 0}
	t.Run("test make then unmake castle", func(t *testing.T) {
		if hasFlag(b.flags, whiteRookKingSideMoved) {
			t.Errorf("whiteRookKingSideMoved flag is set")
		}

		if hasFlag(b.flags, whiteKingMoved) {
			t.Errorf("whiteKingMoved flag is set")
		}

		makeMove(b, &m)

		if !hasFlag(m.boardFlags, whiteRookKingSideMoved) {
			t.Errorf("whiteRookKingSideMoved flag is not set")
		}

		if !hasFlag(m.boardFlags, whiteKingMoved) {
			t.Errorf("whiteKingMoved flag is not set")
		}

		unmakeMove(b, &m)

		if !reflect.DeepEqual(b, want) {
			t.Errorf("got %v, want %v", b, want)
		}
	})
}
