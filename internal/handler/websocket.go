package handler

import (
	"net/http"
	"strconv"
	"time"

	"order-tracking/internal/dto"
	"order-tracking/internal/model"
	"order-tracking/internal/service"
	"order-tracking/pkg/response"

	"github.com/gin-gonic/gin"
)

type WebSocketHandler struct {
	orderService     *service.OrderService
	websocketService *service.WebSocketService
}

func NewWebSocketHandler(orderService *service.OrderService, websocketService *service.WebSocketService) *WebSocketHandler {
	return &WebSocketHandler{
		orderService:     orderService,
		websocketService: websocketService,
	}
}

func (h *WebSocketHandler) UpdateOrderStatusAndNotify(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid order ID"))
		return
	}

	userID := c.MustGet("userID").(uint)

	var req dto.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
		return
	}

	status := model.OrderStatus(req.Status)
	if !status.IsValid() {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid status"))
		return
	}

	order, err := h.orderService.UpdateOrderStatus(uint(orderID), status, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
		return
	}

	h.websocketService.BroadcastStatusUpdate(
		strconv.FormatUint(uint64(order.ID), 10),
		string(order.Status),
		time.Now().Format(time.RFC3339),
	)

	res := dto.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		Status:      string(order.Status),
		Description: order.Description,
		CreatedAt:   order.CreatedAt.String(),
		UpdatedAt:   order.UpdatedAt.String(),
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}
