package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/nsf/termbox-go"
)

var termWidth int
var termHeight int

const (
	squareXSize = 5
	squareYSize = 3
	statusYPos  = 8*squareYSize + 5
	scoreYPos   = statusYPos + 1
	moveYPos    = statusYPos + 2
)

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
	cursesXStart := squareXSize * x
	cursesYStart := squareYSize * (8 - y)

	for cursesX := cursesXStart; (cursesX - cursesXStart) < squareXSize; cursesX++ {
		for cursesY := cursesYStart; (cursesY - cursesYStart) < squareYSize; cursesY++ {
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

	termWidth, termHeight = termbox.Size()

	for x, col := range b.squares {
		for y, sq := range col {
			renderSquare(x, y, sq)
		}
	}

	return termbox.Flush()
}

func renderStatusLine(line string) error {
	return renderTextLine(statusYPos, line)
}

func renderScoreLine(line string) error {
	return renderTextLine(scoreYPos, line)
}

func renderTextLine(ypos int, line string) error {
	for pos, char := range line {
		termbox.SetCell(pos, ypos, char, termbox.ColorDefault, termbox.ColorDefault)
	}

	for pos := len(line); pos < termWidth; pos++ {
		termbox.SetCell(pos, ypos, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}

	return termbox.Flush()
}

func handleKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); {
		case ev.Key == termbox.KeyEsc:
			os.Exit(0)
		case ev.Key == termbox.KeyArrowLeft:
			return
		case ev.Key == termbox.KeyArrowRight:
			return
		default:
		}
	}
}

func handleMoveInput(b *board) move {
	curPos := 0
	ret := move{}

	renderStatusLine("Make move:")
	renderTextLine(moveYPos, fmt.Sprintf("  [%d,%d] -> [%d,%d]", ret.fromX, ret.fromY, ret.toX, ret.toY))

	for {
		switch ev := termbox.PollEvent(); {
		case ev.Key == termbox.KeyEsc:
			os.Exit(0)
		case ev.Key == termbox.KeyEnter:
			if !validMove(b, &ret, white) {
				ret = move{}
				curPos = 0

				renderStatusLine("Invalid move, try again:")
				renderTextLine(moveYPos, fmt.Sprintf("  [%d,%d] -> [%d,%d]", ret.fromX, ret.fromY, ret.toX, ret.toY))
				break
			}

			return ret
		case unicode.IsDigit(ev.Ch): // Only deal with coords for now
			val := int(ev.Ch) - '0' // Convert UTF-8 integer char to int by subtracting UTF-8 '0'

			switch curPos {
			case 0:
				ret.fromX = val
			case 1:
				ret.fromY = val
			case 2:
				ret.toX = val
			case 3:
				ret.toY = val
			default:
				continue
			}

			if curPos > 3 {
				continue
			}

			curPos++
			renderTextLine(moveYPos, fmt.Sprintf("  [%d,%d] -> [%d,%d]", ret.fromX, ret.fromY, ret.toX, ret.toY))
			termbox.Flush()
		}
	}
}
