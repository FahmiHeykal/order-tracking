package handler

import (
	"net/http"
	"order-tracking/internal/dto"
	"order-tracking/internal/service"
	"order-tracking/pkg/response"
	"order-tracking/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *service.UserService
	jwtSecret   string
}

func NewAuthHandler(userService *service.UserService, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
		return
	}

	user, err := h.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
		return
	}

	res := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(res))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(err.Error()))
		return
	}

	user, err := h.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid credentials"))
		return
	}

	token, err := utils.GenerateJWTToken(user.ID, string(user.Role), h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{
		"token": token,
		"user": dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  string(user.Role),
		},
	}))
}
