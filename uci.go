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

func handleUCIPosition(engine *engine, position string) int {
	split := strings.Split(position, "moves")
	// Next should either be "startpos" or "fen"
	fen := ""
	if strings.HasPrefix(split[0], "startpos") {
		fen = startPositionFen
	} else {
		fen = split[0]
	}

	// Setup position
	engine.board = newBoardFromFen(fen)

	// Make moves
	for i := 1; i < len(split); i++ {
		move := uciNotationToMove(split[i])
		makeMove(engine.board, &move)
	}

	return 0
}

func handleUCIGo(engine *engine) {
	// Find move
	bestAIMove := findBestMove(engine.board, black)
	makeMove(engine.board, &bestAIMove)

	// Format move and pass to response channel to send to client
	ret := fmt.Sprintf("bestmove %s%s%s%s", string(bestAIMove.fromX+97), string(bestAIMove.fromY+49), string(bestAIMove.toX+97), string(bestAIMove.toY+49))
	handleUCIResponse(engine, ret)
	prettyPrintBoard(engine.board)
}

func handleUCIResponse(engine *engine, msg string) int {
	engine.server.responseChan <- msg
	return 0
}

func handleUCICommand(engine *engine, msg string) int {
	switch {
	case strings.HasPrefix(msg, "uci"):
		if strings.HasPrefix(msg, "ucinewgame") {
			handleUCINewGame(engine)
		} else {
			// Do nothing
		}
	case strings.HasPrefix(msg, "debug"):
		// TODO
	case strings.HasPrefix(msg, "isready"):
		// TODO
	case strings.HasPrefix(msg, "setoption"):
		// TODO
	case strings.HasPrefix(msg, "register"):
		// TODO
	case strings.HasPrefix(msg, "position"):
		handleUCIPosition(engine, msg[9:])
	case strings.HasPrefix(msg, "go"):
		handleUCIGo(engine)
	case strings.HasPrefix(msg, "stop"):
		// TODO
	case strings.HasPrefix(msg, "ponderhit"):
		// TODO
	case strings.HasPrefix(msg, "uci"):
		// TODO
	}

	return 0
}
