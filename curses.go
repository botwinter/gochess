package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

const (
	start = iota
	userMoveFrom
	userMoveTo
	aiMove
)

const (
	squareXSize = 5
	squareYSize = 3
	statusYPos  = 8*squareYSize + 5
	scoreYPos   = statusYPos + 1
	moveYPos    = statusYPos + 2
)

type cursesBoard struct {
	termWidth  int
	termHeight int

	state int

	cursorXPos int
	cursorYPos int

	move *move
}

func initCursesBoard() *cursesBoard {
	c := cursesBoard{}

	c.state = start
	c.termWidth, c.termHeight = termbox.Size()
	c.move = &move{0, 0, 0, 0, empty}

	return &c
}

func renderSquare(c *cursesBoard, x int, y int, square int) {
	// Get square colour based on coord
	squareBgColour := termbox.ColorBlue
	if (x+y)%2 != 0 {
		squareBgColour = termbox.ColorYellow
	}

	if (c.state == userMoveFrom || c.state == userMoveTo) && x == c.cursorXPos && y == c.cursorYPos {
		squareBgColour = termbox.ColorCyan
	}

	if c.state == userMoveTo && x == c.move.fromX && y == c.move.fromY {
		squareBgColour = termbox.ColorRed
	}

	squarePiece := ' '
	squareFgColour := termbox.ColorWhite | termbox.AttrBold

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
		squareFgColour = termbox.ColorBlack | termbox.AttrBold
	case blackRook:
		squarePiece = 'R'
		squareFgColour = termbox.ColorBlack | termbox.AttrBold
	case blackKnight:
		squarePiece = 'N'
		squareFgColour = termbox.ColorBlack | termbox.AttrBold
	case blackBishop:
		squarePiece = 'B'
		squareFgColour = termbox.ColorBlack | termbox.AttrBold
	case blackKing:
		squarePiece = 'K'
		squareFgColour = termbox.ColorBlack | termbox.AttrBold
	case blackQueen:
		squarePiece = 'Q'
		squareFgColour = termbox.ColorBlack | termbox.AttrBold
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

func renderBoard(c *cursesBoard, b *board) error {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	c.termWidth, c.termHeight = termbox.Size()

	for x, col := range b.squares {
		for y, sq := range col {
			renderSquare(c, x, y, sq)
		}
	}

	return termbox.Flush()
}

func renderStatusLine(c *cursesBoard, line string) error {
	return renderTextLine(c, statusYPos, line)
}

func renderScoreLine(c *cursesBoard, line string) error {
	return renderTextLine(c, scoreYPos, line)
}

func renderTextLine(c *cursesBoard, ypos int, line string) error {
	for pos, char := range line {
		termbox.SetCell(pos, ypos, char, termbox.ColorDefault, termbox.ColorDefault)
	}

	for pos := len(line); pos < c.termWidth; pos++ {
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

func handleGameEnd() {
	for {
		switch ev := termbox.PollEvent(); {
		case ev.Key == termbox.KeyEsc:
			os.Exit(0)
		case ev.Key == termbox.KeyEnter:
			return
		default:
		}
	}
}

func handleMoveInput(c *cursesBoard, b *board) *move {
	c.state = userMoveFrom

	renderBoard(c, b)
	renderStatusLine(c, "Your turn")

	for {
		switch ev := termbox.PollEvent(); {
		case ev.Key == termbox.KeyEsc:
			os.Exit(0)
		case ev.Key == termbox.KeyEnter:
			if c.state == userMoveFrom {
				if isWhite(b.squares[c.cursorXPos][c.cursorYPos]) {
					c.move.fromX = c.cursorXPos
					c.move.fromY = c.cursorYPos
					c.state = userMoveTo
				}
			} else if c.state == userMoveTo {
				if c.cursorXPos == c.move.fromX && c.cursorYPos == c.move.fromY {
					c.state = userMoveFrom
					break
				}
				c.move.toX = c.cursorXPos
				c.move.toY = c.cursorYPos
				if validMove(b, c.move, white) {
					c.state = aiMove
					return c.move
				}
			}
		case ev.Key == termbox.KeyArrowDown:
			if c.cursorYPos > 0 {
				c.cursorYPos--
			}
			renderBoard(c, b)
			renderStatusLine(c, "Your turn")
		case ev.Key == termbox.KeyArrowLeft:
			if c.cursorXPos > 0 {
				c.cursorXPos--
			}
			renderBoard(c, b)
			renderStatusLine(c, "Your turn")
		case ev.Key == termbox.KeyArrowRight:
			if c.cursorXPos < 7 {
				c.cursorXPos++
			}
			renderBoard(c, b)
			renderStatusLine(c, "Your turn")
		case ev.Key == termbox.KeyArrowUp:
			if c.cursorYPos < 7 {
				c.cursorYPos++
			}
			renderBoard(c, b)
			renderStatusLine(c, "Your turn")
		}
	}
}
