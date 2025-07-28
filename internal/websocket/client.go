package websocket

import (
	"log"
	"order-tracking/internal/dto"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	send    chan dto.WebSocketMessage
	orderID string
	userID  uint
}

func NewClient(hub *Hub, conn *websocket.Conn, orderID string, userID uint) *Client {
	return &Client{
		hub:     hub,
		conn:    conn,
		send:    make(chan dto.WebSocketMessage, 256),
		orderID: orderID,
		userID:  userID,
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for message := range c.send {
		err := c.conn.WriteJSON(message)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
	}
}
