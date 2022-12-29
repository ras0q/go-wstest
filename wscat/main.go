package main

import (
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

// 対話形式でメッセージを送受信できるコマンド
// 例: ./wscat ws://localhost:8080/ws
func main() {
	if len(os.Args) != 2 {
		panic("Usage: wscat <websocket url>")
	}

	c, _, err := websocket.DefaultDialer.Dial(os.Args[1], nil)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Print("> ")
		var msg []byte
		fmt.Scanln(&msg)
		c.WriteMessage(websocket.TextMessage, msg)

		_, msg, err = c.ReadMessage()
		if err != nil {
			panic(err)
		}

		fmt.Printf("< %s\n", msg)
	}
}
