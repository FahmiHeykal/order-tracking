package service

import (
	"order-tracking/internal/dto"
	"order-tracking/internal/model"
	"order-tracking/internal/repository"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	historyRepo *repository.HistoryRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, historyRepo *repository.HistoryRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		historyRepo: historyRepo,
	}
}

func (s *OrderService) CreateOrder(userID uint, req dto.CreateOrderRequest) (*model.Order, error) {
	order := &model.Order{
		UserID:      userID,
		Status:      model.StatusPending,
		Description: req.Description,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrderByID(id uint) (*model.Order, error) {
	return s.orderRepo.FindByID(id)
}

func (s *OrderService) GetUserOrders(userID uint) ([]model.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}

func (s *OrderService) GetAllOrders() ([]model.Order, error) {
	return s.orderRepo.FindAll()
}

func (s *OrderService) UpdateOrderStatus(id uint, status model.OrderStatus, changedBy uint) (*model.Order, error) {
	if err := s.orderRepo.UpdateStatus(id, status, changedBy); err != nil {
		return nil, err
	}

	return s.orderRepo.FindByID(id)
}

func (s *OrderService) GetOrderHistory(orderID uint) ([]model.OrderStatusHistory, error) {
	return s.historyRepo.FindByOrderID(orderID)
}
