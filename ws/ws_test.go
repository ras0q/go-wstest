package ws_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/ras0q/go-wstest/ws"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

type wsHandler struct{}

// *wsHandlerはhttp.Handlerインターフェイスを満たす
func (h *wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.Serve(w, r)
}

func TestServe(t *testing.T) {
	t.Run("n回並行にWebsocketの相互通信を行う", func(t *testing.T) {
		t.Parallel()

		var (
			n  = 10
			wg = sync.WaitGroup{}
		)

		wg.Add(n)
		defer wg.Wait()

		for i := 0; i < n; i++ {
			go func(i int) {
				defer wg.Done()

				// httptestを無理やりws用に使う
				server := httptest.NewServer(&wsHandler{})
				server.URL = strings.Replace(server.URL, "http", "ws", 1)

				// クライアントを立ち上げる
				c, _, err := websocket.DefaultDialer.Dial(server.URL, nil)
				require.NoError(t, err)

				// "hello {i}"と送信すると、、、
				err = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("hello %d", i)))
				require.NoError(t, err)

				// "You said: hello {i}"と返ってくる
				_, msg, err := c.ReadMessage()
				require.NoError(t, err)
				require.Equal(t, []byte(fmt.Sprintf("You said: hello %d", i)), msg)
			}(i)
		}
	})
}
