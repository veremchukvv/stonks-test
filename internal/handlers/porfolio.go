package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/pkg/logging"
)

func (h *Handler) createPortfolio(c echo.Context) error {
	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	var newPortfolio models.Portfolio

	err = c.Bind(&newPortfolio)
	if err != nil {
		return c.JSON(500, "Unmarshalling data error")
	}
	createdPortfolio, err := h.services.PortfolioService.CreatePortfolio(context.Background(), token.Value, &newPortfolio)
	if err != nil {
		return c.JSON(500, "Error on create portfolio")
	}
	return c.JSON(200, createdPortfolio)
}

func (h *Handler) getAllPortfolios(c echo.Context) error {
	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}
	p, err := h.services.PortfolioService.GetAllPortfolios(context.Background(), token.Value)
	if err != nil {
		return c.JSON(500, "can't get portfolios")
	}
	if p == nil {
		return c.JSON(200, []string{})
	}
	return c.JSON(200, p)
}

func (h *Handler) modifyPortfolio(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) getPortfolioDeals(c echo.Context) error {
	type response struct {
		PortfolioResp *models.OnePortfolioResp
		DealResp      []*models.DealResp
	}

	log := logging.FromContext(h.ctx)
	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}
	portfolioId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}
	portfolio, stocks, err := h.services.PortfolioService.GetPortfolioDeals(context.Background(), token.Value, portfolioId)
	if err != nil {
		return c.JSON(500, "Can't get portfolio info")
	}

	httpResponse := &response{
		portfolio,
		stocks,
	}

	return c.JSON(200, httpResponse)
}

func (h *Handler) getPortfolioClosedDeals(c echo.Context) error {
	log := logging.FromContext(h.ctx)

	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}
	portfolioId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}
	closedDeals, err := h.services.PortfolioService.GetPortfolioClosedDeals(context.Background(), token.Value, portfolioId)
	if err != nil {
		return c.JSON(500, "Can't get portfolio closed deals info")
	}

	return c.JSON(200, closedDeals)
}

func (h *Handler) deletePortfolio(c echo.Context) error {
	log := logging.FromContext(h.ctx)
	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}
	portfolioId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}

	err = h.services.PortfolioService.DeletePortfolio(context.Background(), token.Value, portfolioId)
	if err != nil {
		log.Infof("can't delete portfolio with ID %d", portfolioId)
		return c.JSON(500, "can't delete portfolio")
	}
	// return c.Redirect(http.StatusOK, "http://localhost:3000/")
	return c.JSON(200, "portfolio deleted")
}
