package service

import (
	"order-tracking/internal/dto"
	"order-tracking/internal/websocket"
)

type WebSocketService struct {
	hub *websocket.Hub
}

func NewWebSocketService(hub *websocket.Hub) *WebSocketService {
	return &WebSocketService{
		hub: hub,
	}
}

func (s *WebSocketService) BroadcastStatusUpdate(orderID string, status string, updatedAt string) {
	message := dto.WebSocketMessage{
		OrderID:   orderID,
		Status:    status,
		UpdatedAt: updatedAt,
	}
	s.hub.Broadcast(message)
}
