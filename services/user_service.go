package services

import (
	"github.com/Yalm/go-psql/models"
	"github.com/Yalm/go-psql/repositories"
)

type UserService interface {
	FindAll() *[]models.User
}

type userService struct {
	userRepositoy repositories.UserRepository
}

func NewUserService(userRepositoy repositories.UserRepository) UserService {
	return &userService{userRepositoy}
}

func (ctx *userService) FindAll() *[]models.User {
	return ctx.userRepositoy.FindAll()
}
