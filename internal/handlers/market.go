package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"net/http"
	"strconv"
)

func (h *Handler) getAllStocks(c echo.Context) error {
	//log := logging.FromContext(h.ctx)

	s, err := h.services.MarketService.GetAllStocks(context.Background())

	if err != nil {
		return c.JSON(500, "can't get stocks")
	}

	return c.JSON(200, s)
}

func (h *Handler) getOneStock(c echo.Context) error {
	log := logging.FromContext(h.ctx)

	stockId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}

	s, err := h.services.MarketService.GetOneStock(context.Background(), stockId)

	if err != nil {
		return c.JSON(500, "can't get stock")
	}

	return c.JSON(200, s)
}

func (h *Handler) makeDeal(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) deleteDeal(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}
