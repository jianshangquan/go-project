package chat

import (
	"encoding/json"
	"errors"
	"myapp/src/logger"
	"myapp/src/utils"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


const (
    // Time allowed to write a message to the peer.
    writeWait = 10 * time.Second

    // Time allowed to read the next pong message from the peer.
    pongWait = 60 * time.Second

    // Send pings to peer with this period. Must be less than pongWait.
    pingPeriod = (pongWait * 9) / 10
)


type ChatController struct{
	logger *logger.Logger
	connections map[string]*websocket.Conn
	connMutex sync.Mutex
	upgrader websocket.Upgrader
}

func NewChatController(logger *logger.Logger) *ChatController{
	println("Chat Controller crated")
	return &ChatController{
		logger: logger,
		connections: make(map[string]*websocket.Conn),
		connMutex: sync.Mutex{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins for simplicity (you can customize this)
			},
		},
	}
}



func (controller *ChatController) handleWebsocket(c *gin.Context){
	controller.logger.Log("ws requested")
	var conn, err = controller.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()


	var userId = c.GetHeader("userId")
	controller.registerWsUser(userId, conn)
	controller.startPingPong(conn)


	// Loop to read messages from WebSocket
	for {
		// Read message from WebSocket
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			controller.logger.Log("Error reading message:")
			break
		}
		
		var chatData, _ = utils.ParseJsonFromBytes[ChatRequestData](msg)
		userErrorNotFound := controller.sendToUser(chatData.ToUser, msgType, chatData.Message)
		if userErrorNotFound != nil {
			controller.sendObject(conn, msgType, &ChatErrorResponse{
				Message: "User not found",
			})
		}
	}
	controller.logger.Log("end request")
}


func (controller *ChatController) registerWsUser(userId string, conn *websocket.Conn){
	controller.connMutex.Lock();
	controller.connections[userId] = conn;
	controller.connMutex.Unlock();
}

func (controller *ChatController) sendToUser(userId string, msgType int, message string) error{
	var connection, exists = controller.connections[userId]
	if !exists {
		controller.logger.Log("Error: User connection not found for userId: " + userId)
		return errors.New("user connection not found")
	}

	err := controller.send(connection, msgType, message)
	return err;
}

func (controller *ChatController) send(conn *websocket.Conn, msgType int, message string) error {
	return controller.sendRaw(conn, msgType, []byte(message));
}
func (controller *ChatController) sendObject(conn *websocket.Conn, msgType int, data interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err  // can't convert struct to JSON
	}
	return controller.sendRaw(conn, msgType, jsonBytes);
}
func (controller *ChatController) sendRaw(conn *websocket.Conn, msgType int, message []byte) error {
	if err := conn.WriteMessage(msgType, []byte(message)); err != nil {
		controller.logger.Log("Error sending message: " + string(message))
		return err;
	}
	return nil;
}

func (controller *ChatController) findUserIdByConn(conn *websocket.Conn) string {
	for key, value := range controller.connections{
		if value == conn {
			return key;
		}
	}
	return "";
}

func (controller *ChatController) startPingPong(conn *websocket.Conn) {
    conn.SetReadDeadline(time.Now().Add(pongWait))
    conn.SetPongHandler(func(appData string) error {
        conn.SetReadDeadline(time.Now().Add(pongWait)) // extend deadline on pong
        return nil
    })

    ticker := time.NewTicker(pingPeriod)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				userId := controller.findUserIdByConn(conn);
				controller.logger.Log("Ping failed, closing user connection " + userId)
                conn.Close()
                return
            }
        }
    }
}


func (controller *ChatController) RegisterRoute(r *gin.Engine){
	r.GET("/ws", controller.handleWebsocket)
}
