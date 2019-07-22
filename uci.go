package main

import (
	"fmt"
	"strings"
)

func findTokenIndex(split []string, token string) int {
	for i, value := range split {
		if value == token {
			return i
		}
	}
	return -1
}

// To convert a file letter to a coord, can just the ASCII value
// of the letter, and subtract 97 (ASCII a == 97)
func letterToCoord(letter byte) int {
	return int(letter) - 97
}

// UCI notation example moves:
// e2e4 e7e5
func uciNotationToMove(str string) move {
	return move{
		fromX: letterToCoord(str[0]),
		fromY: int(str[1]) - 1,
		toX:   letterToCoord(str[2]),
		toY:   int(str[3]) - 1,
	}
}

func handleUCINewGame(engine *engine) int {
	engine.board = newDefaultBoard()

	return 0
}

func handleUCIPosition(engine *engine, cmd []string) int {
	// Next should either be "startpos" or "fen"
	position := cmd[1]
	fen := ""
	if position == "startpos" {
		fen = startPositionFen
	} else {
		fen = strings.Join(cmd[1:7], "") // FEN string has 6 space-separated fields
	}

	// Setup position
	engine.board = newBoardFromFen(fen)

	// Make moves
	moveIdx := findTokenIndex(cmd, "moves")
	for _, moveStr := range cmd[moveIdx+1:] {
		move := uciNotationToMove(moveStr)
		makeMove(engine.board, &move)
	}

	return 0
}

func handleUCIGo(engine *engine) {
	// Find move
	bestAIMove := findBestMove(engine.board, black)
	makeMove(engine.board, &bestAIMove)
	fmt.Println(bestAIMove.toString())

	// Format move and pass to response channel to send to client
	ret := fmt.Sprintf("ai_move %s%s %s%s", string(bestAIMove.fromX+97), string(bestAIMove.fromY+49), string(bestAIMove.toX+97), string(bestAIMove.toY+49))
	handleUCIResponse(engine, ret)
}

func handleUCIResponse(engine *engine, msg string) int {
	engine.server.responseChan <- msg
	return 0
}

func handleUCICommand(engine *engine, msg string) int {
	split := strings.Split(msg, " ")
	cmd := split[0]

	switch cmd {
	case "uci":
		// Do nothing
	case "debug":
		// TODO
	case "isready":
		// TODO
	case "setoption":
		// TODO
	case "register":
		// TODO
	case "ucinewgame":
		handleUCINewGame(engine)
	case "position":
		handleUCIPosition(engine, split)
	case "go":
		handleUCIGo(engine)
	case "stop":
		// TODO
	case "ponderhit":
		// TODO
	case "quit":
		// TODO
	}

	return 0
}
