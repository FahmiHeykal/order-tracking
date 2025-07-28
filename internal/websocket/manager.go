package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketManager struct {
	hub *Hub
}

func NewWebSocketManager(hub *Hub) *WebSocketManager {
	return &WebSocketManager{
		hub: hub,
	}
}

func (m *WebSocketManager) HandleWebSocket(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.MustGet("userID").(uint)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := NewClient(m.hub, conn, orderID, userID)
	m.hub.register <- client

	go client.writePump()
	go client.readPump()
}
