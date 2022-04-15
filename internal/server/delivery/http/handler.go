package http

import (
	"net/http"

	"github.com/Mustafa0831/TSarka/internal/counter"
	counterHTTP "github.com/Mustafa0831/TSarka/internal/counter/delivery/http"
	"github.com/Mustafa0831/TSarka/internal/email"
	emailHTTP "github.com/Mustafa0831/TSarka/internal/email/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//Handler ...
type Handler struct {
	counterService *counter.Service
	emailService   *email.Service
}

//NewHandler ...
func NewHandler(counterService *counter.Service, emailService *email.Service) *Handler {
	return &Handler{
		counterService,
		emailService,
	}
}

//Init ...
func (h *Handler) Init() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	router.GET("/ping", func(e echo.Context) error {
		return e.JSON(http.StatusOK, nil)
	})

	counterGroup := router.Group("/rest")

	counterHandler := counterHTTP.NewHandler(h.counterService)

	counterHandler.Init(counterGroup)

	emailHandler := emailHTTP.NewHandler(h.emailService)

	emailHandler.Init(counterGroup)

	return router
}
