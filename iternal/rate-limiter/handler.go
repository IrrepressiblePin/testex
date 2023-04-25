package rate_limiter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tz/iternal/rate-limiter/dto"
	"tz/pkg/handlers"
)

type handler struct {
	service *Service
}

func NewHandler(service *Service) handlers.Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	router.PATCH("/rate", h.Update)
}

func (h *handler) Update(c *gin.Context) {
	var data dto.UpdateDTO

	if err := c.ShouldBind(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "validation error")
		return
	}

	if data.Second == 0 || data.Request == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "rate error")
		return
	}

	h.service.Update(data)
	c.Status(http.StatusOK)
}
