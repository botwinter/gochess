package main

import "github.com/nsf/termbox-go"

func renderSquare(x int, y int, square int) {
	// Get square colour based on coord
	squareBgColour := termbox.ColorBlack
	if (x+y)%2 != 0 {
		squareBgColour = termbox.ColorWhite
	}

	squarePiece := ' '
	squareFgColour := termbox.ColorYellow | termbox.AttrBold

	switch square {
	case whiteRook:
		squarePiece = 'R'
	case whiteKnight:
		squarePiece = 'N'
	case whiteBishop:
		squarePiece = 'B'
	case whiteQueen:
		squarePiece = 'Q'
	case whiteKing:
		squarePiece = 'K'
	case whitePawn:
		squarePiece = 'P'
	case blackPawn:
		squarePiece = 'P'
		squareFgColour = termbox.ColorRed | termbox.AttrBold
	case blackRook:
		squarePiece = 'R'
		squareFgColour = termbox.ColorRed | termbox.AttrBold
	case blackKnight:
		squarePiece = 'N'
		squareFgColour = termbox.ColorRed | termbox.AttrBold
	case blackBishop:
		squarePiece = 'B'
		squareFgColour = termbox.ColorRed | termbox.AttrBold
	case blackKing:
		squarePiece = 'K'
		squareFgColour = termbox.ColorRed | termbox.AttrBold
	case blackQueen:
		squarePiece = 'Q'
		squareFgColour = termbox.ColorRed | termbox.AttrBold
	}

	/* A square will actually be rendered as a 5x3 grid, with the piece in the middle.
	Get curses coord start - curses coords start top left, chess board coords start bottom left */
	cursesXStart := 5 * x
	cursesYStart := 3 * (8 - y)

	for cursesX := cursesXStart; (cursesX - cursesXStart) < 5; cursesX++ {
		for cursesY := cursesYStart; (cursesY - cursesYStart) < 3; cursesY++ {
			// Only set the piece flag if in the middle of the square grid
			if (cursesX-cursesXStart) == 2 && (cursesY-cursesYStart) == 1 {
				termbox.SetCell(cursesX, cursesY, squarePiece, squareFgColour, squareBgColour)
			} else {
				termbox.SetCell(cursesX, cursesY, ' ', squareFgColour, squareBgColour)
			}
		}
	}
}

func renderBoard(b *board) error {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for x, col := range b.squares {
		for y, sq := range col {
			renderSquare(x, y, sq)
		}
	}

	return termbox.Flush()
}

func handleKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowLeft:
				//dispatch("left", &Message{data: "push left"})
				break
			case termbox.KeyArrowRight:
				//dispatch("right", &Message{data: "push left"})
				break
			default:
			}
		default:
		}
	}
}