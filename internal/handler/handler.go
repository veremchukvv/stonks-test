package handler

import "github.com/labstack/echo/v4"

type Handler struct {

}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()
	return router
}
