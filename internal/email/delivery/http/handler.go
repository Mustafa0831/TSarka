package http

import (
	"github.com/Mustafa0831/TSarka/internal/email"
	"github.com/labstack/echo/v4"
)

//Handler ...
type Handler struct {
	emailService email.EmailService
}

//NewHandler ...
func NewHandler(service email.EmailService) *Handler {
	return &Handler{service}
}

//Init ...
func (h *Handler) Init(api *echo.Group) {
	emailAPI := api.Group("/email")

	emailAPI.POST("/check", h.emailCheck)
	emailAPI.POST("/iin", h.iinCheck)
}
