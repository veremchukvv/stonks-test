package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) signup(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) signin(c echo.Context) error {
	return c.String(http.StatusOK, "signin OK")
}

func (h *Handler) signout(c echo.Context) error {
	return c.String(http.StatusOK, "signout OK")
}