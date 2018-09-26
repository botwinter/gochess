package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

// Global flags
var cursesEnabled bool

func main() {
	b := newDebugBoard()

	// CLI flags
	cursesPtr := flag.Bool("curses", true, "bool")
	flag.Parse()

	cursesEnabled = *cursesPtr

	if cursesEnabled {
		termbox.Init()
		defer termbox.Close()

		err := renderBoard(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	findBestMove(b, white)
}
