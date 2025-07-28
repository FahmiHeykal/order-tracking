package main

import (
	"log"

	"order-tracking/config"
	"order-tracking/internal/handler"
	"order-tracking/internal/middleware"
	"order-tracking/internal/model"
	"order-tracking/internal/repository"
	"order-tracking/internal/service"
	"order-tracking/internal/websocket"
	"order-tracking/pkg/response"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	if err := cfg.DB.AutoMigrate(
		&model.User{},
		&model.Order{},
		&model.OrderStatusHistory{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	hub := websocket.NewHub()
	go hub.Run()

	userRepo := repository.NewUserRepository(cfg.DB)
	orderRepo := repository.NewOrderRepository(cfg.DB)
	historyRepo := repository.NewHistoryRepository(cfg.DB)

	userService := service.NewUserService(userRepo)
	orderService := service.NewOrderService(orderRepo, historyRepo)
	wsService := service.NewWebSocketService(hub)

	authHandler := handler.NewAuthHandler(userService, cfg.JWTSecret)
	orderHandler := handler.NewOrderHandler(orderService)
	wsHandler := handler.NewWebSocketHandler(orderService, wsService)
	wsManager := websocket.NewWebSocketManager(hub)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, response.NewSuccessResponse(gin.H{
			"status": "ok",
		}))
	})

	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		auth := api.Group("")
		auth.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
		{
			auth.POST("/orders", orderHandler.CreateOrder)
			auth.GET("/orders", orderHandler.GetUserOrders)
			auth.GET("/orders/:id", orderHandler.GetOrder)
			auth.GET("/orders/:id/history", orderHandler.GetOrderHistory)

			admin := auth.Group("")
			admin.Use(middleware.RoleMiddleware("admin", "driver"))
			{
				admin.GET("/admin/orders", orderHandler.GetAllOrders)
				admin.PUT("/orders/:id/status", wsHandler.UpdateOrderStatusAndNotify)
			}
		}
	}

	r.GET("/ws/orders/:id",
		middleware.JWTAuthMiddleware(cfg.JWTSecret),
		wsManager.HandleWebSocket,
	)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
