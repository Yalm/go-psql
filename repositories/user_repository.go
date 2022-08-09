package repositories

import (
	"github.com/Yalm/go-psql/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() *[]models.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ctx *userRepository) FindAll() *[]models.User {
	result := &[]models.User{}

	ctx.db.Model(&models.User{}).Find(&result)

	return result
}
