package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleClientMessage(c *websocket.Conn, mt int, rawMsg []byte) {
	msg := strings.Split(string(rawMsg[:]), " ")

	if msg[0] == "new" {
		mainBoard = newDefaultBoard()
	} else if msg[0] == "player_move" {
		move := &move{}

		from := msg[1]
		to := msg[2]

		move.fromX = int(from[0]) - 97
		move.fromY = int(from[1]) - 49
		move.toX = int(to[0]) - 97
		move.toY = int(to[1]) - 49

		fmt.Println(move.toString())
		makeMove(mainBoard, move)

		bestAIMove := findBestMove(mainBoard, black)

		makeMove(mainBoard, &bestAIMove)

		fmt.Println(bestAIMove.toString())
		ret := fmt.Sprintf("ai_move %s%s %s%s", string(bestAIMove.fromX+97), string(bestAIMove.fromY+49), string(bestAIMove.toX+97), string(bestAIMove.toY+49))
		c.WriteMessage(mt, []byte(ret))
	} else {
		fmt.Println("bad request")
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s", string(msg))
		handleClientMessage(c, mt, msg)
	}
}

func startServer() {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("./js")))

	// Handle requests from client
	http.HandleFunc("/ws", serveWs)

	http.ListenAndServe(":8080", nil)
}
