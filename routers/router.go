package routers

import (
	"net/http"

	"github.com/Yalm/go-psql/controllers"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

func New(router *mux.Router, container *dig.Container) {
	container.Invoke(func(userController controllers.UserController) {
		router.HandleFunc("/users", userController.FindAll).Methods(http.MethodPost)
	})
}
