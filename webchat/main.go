package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 全局连接池
var connections = make(map[*websocket.Conn]struct{})
var connectionsMutex sync.Mutex

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("webchat/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the chat application!",
		})
	})

	router.GET("/chat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./templates/chat.html", nil)
	})

	// 处理WebSocket连接
	router.GET("/ws", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 添加连接到连接池
		connectionsMutex.Lock()
		connections[ws] = struct{}{}
		connectionsMutex.Unlock()

		// 在这里处理WebSocket连接
		go handleWebSocketConnection(ws)

		defer func() {
			// 在连接关闭时从连接池中移除
			connectionsMutex.Lock()
			delete(connections, ws)
			connectionsMutex.Unlock()
		}()
	})

	err := router.Run(":8088")
	if err != nil {
		return
	}
}

// 处理WebSocket连接
func handleWebSocketConnection(ws *websocket.Conn) {
	for {
		// 读取消息
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		// 广播消息给所有连接
		connectionsMutex.Lock()
		for conn := range connections {
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				err := conn.Close()
				if err != nil {
					return
				}
				delete(connections, conn)
			}
		}
		connectionsMutex.Unlock()
	}
}
