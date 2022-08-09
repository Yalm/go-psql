package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Yalm/go-psql/services"
)

type UserController interface {
	FindAll(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{userService}
}

func (ctx *userController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ctx.userService.FindAll())
}
