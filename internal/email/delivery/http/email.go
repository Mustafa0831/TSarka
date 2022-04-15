package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//CheckRequest ...
type CheckRequest struct {
	Text string
}

func (h *Handler) emailCheck(c echo.Context) error {
	emailRequest := new(CheckRequest)
	if err := c.Bind(emailRequest); err != nil {
		return err
	}

	email, err := h.emailService.FindEmailFromText(emailRequest.Text)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, email)
}

func (h *Handler) iinCheck(c echo.Context) error {
	checkRequest := new(CheckRequest)
	if err := c.Bind(checkRequest); err != nil {
		return err
	}

	iin, err := h.emailService.FindIinFromText(checkRequest.Text)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, iin)
}
