package router

import (
	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/stallone/internal/api/handlers/healthcheck"
)

type Router struct {
	Echo *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()

	setupHealthRoutes(e)

	return &Router{
		Echo: e,
	}
}

func setupHealthRoutes(e *echo.Echo) {
	e.GET("/healthcheck", healthcheck.NewHealthCheckHandler().Handle)
}
