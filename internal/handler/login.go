package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) login(c echo.Context) error {
	return c.String(http.StatusOK, "LOGIN OK")
}

func (h *Handler) logout(c echo.Context) error {
	return c.String(http.StatusOK, "LOGOUT OK")
}