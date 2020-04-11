package main

import "testing"

func Test_evaluateBoard(t *testing.T) {
	type args struct {
		b *board
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test starting pos",
			args{
				newBoardFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			},
			0,
		},
		{
			"test white taken all",
			args{
				newBoardFromFen("8/8/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			},
			23905,
		},
		{
			"test black taken all",
			args{
				newBoardFromFen("rnbqkbnr/pppppppp/8/8/8/8/8/8 w KQkq - 0 1"),
			},
			-23905,
		},
		{
			"test mix",
			args{
				newBoardFromFen("r1bqkb1r/p1pp1ppp/8/8/3pP3/8/PP3PPP/RN1QK1N1 w KQkq - 0 1"),
			},
			-665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateBoard(tt.args.b); got != tt.want {
				t.Errorf("evaluateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
