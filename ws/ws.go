package ws

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

var latestClientID = atomic.Int64{}

func Serve(w http.ResponseWriter, r *http.Request) {
	clientID := latestClientID.Add(1)
	mylog := log.Default()
	mylog.SetPrefix(fmt.Sprintf("client%d: ", clientID))

	mylog.Println("connected")
	defer mylog.Println("disconnected")

	c, err := new(websocket.Upgrader).Upgrade(w, r, nil)
	if err != nil {
		mylog.Println("upgrade error:", err)
		return
	}
	defer c.Close()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			mylog.Println("read error:", err)
			break
		}

		if err := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("You said: %s", msg))); err != nil {
			mylog.Println("write error:", err)
			break
		}
	}
}
