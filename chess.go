package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	b := newDefaultBoard()

	err := renderBoard(b)
	if err != nil {
		fmt.Println(err)
	}

	handleKeyEvent()
}
