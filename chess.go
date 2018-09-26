package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	b := newDebugBoard()

	err := renderBoard(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	findBestMove(b, white)
}
