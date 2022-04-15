package http

import (
	"github.com/Mustafa0831/TSarka/internal/counter"
	"github.com/labstack/echo/v4"
)

//Handler ...
type Handler struct {
	counterService *counter.Service
}

//NewHandler ...
func NewHandler(service *counter.Service) *Handler {
	return &Handler{service}
}

//Init ...
func (h *Handler) Init(api *echo.Group) {
	counterAPI := api.Group("counter")

	counterAPI.POST("/add/:num", h.counterAdd)
	counterAPI.POST("/sub/:num", h.counterSub)
	counterAPI.GET("/val", h.counterVal)
}
