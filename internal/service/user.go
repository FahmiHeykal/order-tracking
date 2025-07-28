package service

import (
	"order-tracking/internal/dto"
	"order-tracking/internal/model"
	"order-tracking/internal/repository"
	"order-tracking/pkg/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(req dto.RegisterRequest) (*model.User, error) {
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     model.RoleUser,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(req dto.LoginRequest) (*model.User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}
