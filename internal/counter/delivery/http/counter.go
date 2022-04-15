package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) counterAdd(c echo.Context) error {
	num := c.Param("num")
	if err := h.counterService.SetIncrement(c.Request().Context(), num); err != nil {
		return c.String(http.StatusInternalServerError,"SetIncrement")
	}
	num,err:=h.counterService.Get(c.Request().Context())
	if err !=nil {
		return c.String(http.StatusInternalServerError, "Get")
	}
	return c.String(http.StatusOK,num)
}

func (h *Handler) counterSub(c echo.Context) error{
	num := c.Param("num")
	if err:=h.counterService.SetDecrement(c.Request().Context(),num);err!=nil {
		return c.String(http.StatusInternalServerError, "SetDecrement")
	}
	num, err:= h.counterService.Get(c.Request().Context())
	if err!= nil {
		return c.String(http.StatusInternalServerError, "Get")
	}
	return c.String(http.StatusOK, num)
}

func (h *Handler) counterVal(c echo.Context) error{
	num,err := h.counterService.Get(c.Request().Context())
	if err!=nil{
		return c.String(http.StatusInternalServerError,"Get")
	}
	return c.String(http.StatusOK, num)
}
