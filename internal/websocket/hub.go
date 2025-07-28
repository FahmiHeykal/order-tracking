package websocket

import (
	"order-tracking/internal/dto"
	"sync"
)

type Hub struct {
	clients    map[string]map[*Client]bool
	broadcast  chan dto.WebSocketMessage
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan dto.WebSocketMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if h.clients[client.orderID] == nil {
				h.clients[client.orderID] = make(map[*Client]bool)
			}
			h.clients[client.orderID][client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.orderID]; ok {
				delete(h.clients[client.orderID], client)
				close(client.send)
				if len(h.clients[client.orderID]) == 0 {
					delete(h.clients, client.orderID)
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.Lock()
			clients := h.clients[message.OrderID]
			for client := range clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(clients, client)
					if len(clients) == 0 {
						delete(h.clients, message.OrderID)
					}
				}
			}
			h.mu.Unlock()
		}
	}
}

func (h *Hub) Broadcast(message dto.WebSocketMessage) {
	h.broadcast <- message
}
