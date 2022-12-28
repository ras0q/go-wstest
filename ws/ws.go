package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				panic(err)
			}

			c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("You said: %s", msg)))
		}
	}()
}
