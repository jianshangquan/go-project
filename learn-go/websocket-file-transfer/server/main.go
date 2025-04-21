package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default
		return true
	},
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not get caller info")
	}

	dir := filepath.Dir(filename)
	var dist, _ = os.Create(path.Join(dir, "test-transfered.mp4"))
	defer dist.Close()

	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		// fmt.Printf("Received: %s\n", string(msg))
		dist.Write(msg)
	}
}

func main() {
	r := gin.Default()

	r.GET("/ws", wsHandler)

	fmt.Println("WebSocket server running at :8080")
	r.Run(":8080")
}
