package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vuluu2k/remember_fullstack/server/model"
)

type UserService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindById(ctx, uid)

	return u, err
}

func (s *UserService) SignUp(ctx context.Context, u *model.User) error {
	panic("Method not implemented")
}
