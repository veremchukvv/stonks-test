package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/veremchukvv/stonks-test/internal/models"
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
	log := logging.FromContext(h.ctx)

	var req models.Deal

	err := c.Bind(&req)
	if err != nil {
		log.Infof("error on get params from HTTP POST request %v", err)
		return c.JSON(500, "error on get params from HTTP POST request")
	}

	cookie, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.Infof("not logined %v", err)
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	d, err := h.services.MarketService.CreateDeal(context.Background(), cookie.Value, req.StockID, req.StockAmount, req.PortfolioID)
	if err != nil {
		log.Infof("error on creating deal %v", err)
		return c.JSON(500, "error on creating deal")
	}

	return c.JSON(200, d)
}

func (h *Handler) deleteDeal(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}
