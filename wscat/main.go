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
	defer c.Close()

	// サーバーからのメッセージを非同期で受信する
	go func() {
		fmt.Printf("Connected to %s\n", os.Args[1])
		fmt.Print("> ")

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				panic(err)
			}

			fmt.Print("\r")           // 最終行の"> "を消す
			fmt.Printf("< %s\n", msg) // 受信したメッセージを表示
			fmt.Print("> ")           // 最終行に"> "を表示
		}
	}()

	// 標準入力からメッセージを送信する
	for {
		var msg []byte
		fmt.Scanln(&msg)
		c.WriteMessage(websocket.TextMessage, msg)
	}
}
