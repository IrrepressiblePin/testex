package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tz/iternal/user/entities"
	"tz/pkg/handlers"
)

type handler struct{}

var u = []entities.User{
	{Name: "Oleg", Age: 10},
	{Name: "Alex", Age: 20},
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *gin.RouterGroup) {
	router.GET("/user", h.findAll)
}

func (h *handler) findAll(c *gin.Context) {
	c.JSON(http.StatusOK, u)
}
