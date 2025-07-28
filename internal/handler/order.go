package handler

import (
	"net/http"
	"order-tracking/internal/dto"
	"order-tracking/internal/model"
	"order-tracking/internal/service"
	"order-tracking/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
		return
	}

	order, err := h.orderService.CreateOrder(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
		return
	}

	res := dto.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		Status:      string(order.Status),
		Description: order.Description,
		CreatedAt:   order.CreatedAt.String(),
		UpdatedAt:   order.UpdatedAt.String(),
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(res))
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid order ID"))
		return
	}

	order, err := h.orderService.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewErrorResponse("Order not found"))
		return
	}

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

func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	orders, err := h.orderService.GetUserOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
		return
	}

	var res []dto.OrderResponse
	for _, order := range orders {
		res = append(res, dto.OrderResponse{
			ID:          order.ID,
			UserID:      order.UserID,
			Status:      string(order.Status),
			Description: order.Description,
			CreatedAt:   order.CreatedAt.String(),
			UpdatedAt:   order.UpdatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
		return
	}

	var res []dto.OrderResponse
	for _, order := range orders {
		res = append(res, dto.OrderResponse{
			ID:          order.ID,
			UserID:      order.UserID,
			Status:      string(order.Status),
			Description: order.Description,
			CreatedAt:   order.CreatedAt.String(),
			UpdatedAt:   order.UpdatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
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

func (h *OrderHandler) GetOrderHistory(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid order ID"))
		return
	}

	histories, err := h.orderService.GetOrderHistory(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
		return
	}

	var res []dto.OrderHistoryResponse
	for _, history := range histories {
		res = append(res, dto.OrderHistoryResponse{
			ID:        history.ID,
			OrderID:   history.OrderID,
			Status:    string(history.Status),
			ChangedBy: history.ChangedBy,
			CreatedAt: history.CreatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(res))
}
