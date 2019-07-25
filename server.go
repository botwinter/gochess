package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type server struct {
	engine       *engine
	responseChan chan string
	connection   *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func readFromConnection(readChan chan string, conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s", string(msg))
		readChan <- string(msg)
	}
}

func serveWs(server *server, w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer c.Close()

	readChan := make(chan (string))
	go readFromConnection(readChan, c)

	for {
		select {
		case msg := <-readChan:
			handleUCICommand(server.engine, msg)
		case msg := <-server.responseChan:
			c.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}
}

func startServer(server *server) {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("./js")))
	// Handle requests from client
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(server, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
