package handlers

import (
	"net/http"

	"github.com/funukonta/task_manager/internal/models"
	"github.com/funukonta/task_manager/internal/services"
	"github.com/funukonta/task_manager/pkg"
	"github.com/gin-gonic/gin"
)

type handlers_Users struct {
	services services.Service_User
}

func New_HanlderUser(services services.Service_User) *handlers_Users {
	return &handlers_Users{services: services}
}

func (h *handlers_Users) RegisterUser(c *gin.Context) {
	reqJson := new(models.UserModel)
	if err := c.ShouldBind(reqJson); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err}).Send(c)
		return
	}

	created, err := h.services.RegisterUser(reqJson)
	if err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Data: created}).Send(c)
}
