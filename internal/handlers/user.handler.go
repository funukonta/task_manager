package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *handlers_Users) GetUsers(c *gin.Context) {
	result, err := h.services.GetUsers()
	if err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Data: result}).Send(c)
}

func (h *handlers_Users) GetUserById(c *gin.Context) {
	user := models.UserModel{}
	if err := c.ShouldBindUri(&user); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	result, err := h.services.GetUserById(user.ID)
	if err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Data: result}).Send(c)
}

func (h *handlers_Users) EditUser(c *gin.Context) {
	user := models.UserModel{}
	var err error

	if err = c.ShouldBind(&user); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}
	id := c.Param("id")
	user.ID, err = strconv.Atoi(id)
	if err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	if err := h.services.EditUser(&user); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Message: fmt.Sprintf("Berhasil update user id: %d", user.ID)}).Send(c)
}

func (h *handlers_Users) DeleteUser(c *gin.Context) {
	user := models.UserModel{}
	if err := c.ShouldBindUri(&user); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	if err := h.services.DeleteUser(user.ID); err != nil {
		pkg.Responses(http.StatusBadRequest, &pkg.BodJson{Message: err.Error()}).Send(c)
		return
	}

	pkg.Responses(http.StatusOK, &pkg.BodJson{Message: fmt.Sprintf("Berhasil delete user id: %d", user.ID)}).Send(c)
}
