package main

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to the WebSocket server
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not get caller info")
	}
	dir := filepath.Dir(filename)
	var src, _ = os.Open(path.Join(dir, "test.mp4"))
	defer src.Close()

	var buffer = make([]byte, 1024) // 1MB buffer

	for {
		n, err := src.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

		var data = buffer[:n]
		// fmt.Printf("Received: %s\n", string(data))
		conn.WriteMessage(websocket.BinaryMessage, data)
		// time.Sleep(2 * time.Second)

		if err != nil {
			log.Fatal(err)
		}
	}

	// // Read message loop
	// for {
	// 	_, msg, err := conn.ReadMessage()
	// 	if err != nil {
	// 		log.Println("Read error:", err)
	// 		break
	// 	}
	// 	fmt.Printf("Received from server: %s\n", msg)

	// 	// Optional: send another message after receiving
	// 	time.Sleep(2 * time.Second)
	// 	conn.WriteMessage(websocket.TextMessage, []byte("Ping from client!"))
	// }
}
