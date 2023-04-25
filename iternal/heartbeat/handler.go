package heartbeat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tz/pkg/handlers"
)

type handler struct{}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *gin.RouterGroup) {
	router.GET("/heartbeat", h.heartBeat)
}

func (h *handler) heartBeat(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}
