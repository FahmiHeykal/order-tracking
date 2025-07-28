package dto

type CreateOrderRequest struct {
	Description string `json:"description" binding:"required"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type OrderResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type OrderHistoryResponse struct {
	ID        uint   `json:"id"`
	OrderID   uint   `json:"order_id"`
	Status    string `json:"status"`
	ChangedBy uint   `json:"changed_by"`
	CreatedAt string `json:"created_at"`
}