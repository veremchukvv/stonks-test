package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) getAllStocks(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) getOneStock(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) makeDeal(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) deleteDeal(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

