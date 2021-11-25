package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) getAllStocks(c echo.Context) error {
	s, err := h.services.MarketService.GetAllStocks(context.Background())

	if err != nil {
		return c.JSON(500, "can't get stocks")
	}
	return c.JSON(200, s)
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

