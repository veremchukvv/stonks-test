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

func (h *Handler) createPortfolio(c echo.Context) error {
	//log := logging.FromContext(h.ctx)
	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			c.Response().WriteHeader(http.StatusUnauthorized)
			c.Response().Write([]byte(`{"error": "not logined"}`))
			return nil
		}
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "can't parse cookie'"}`))
		return nil
	}
	var newPortfolio models.Portfolio
	c.Bind(&newPortfolio)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "Unmarshalling data error"}`))
		return nil
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
			c.Response().WriteHeader(http.StatusUnauthorized)
			c.Response().Write([]byte(`{"error": "not logined"}`))
			return nil
		}
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "can't parse cookie'"}`))
		return nil
	}
	p, err := h.services.PortfolioService.GetAllPortfolios(context.Background(), token.Value)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "can't get portfolios'"}`))
	}
	if p == nil {
		return c.JSON(200, []string{})
	}
	return c.JSON(200, p)
}

func (h *Handler) modifyPortfolio(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) getPortfolio(c echo.Context) error {

	type response struct {
		PortfolioResp *models.OnePortfolioResp
		StocksResp    []*models.StockResp
	}

	log := logging.FromContext(h.ctx)
	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			c.Response().WriteHeader(http.StatusUnauthorized)
			c.Response().Write([]byte(`{"error": "not logined"}`))
			return nil
		}
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "can't parse cookie'"}`))
		return nil
	}
	portfolioId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}
	portfolio, stocks, err := h.services.PortfolioService.GetOnePortfolio(context.Background(), token.Value, portfolioId)
	if err != nil {
		return c.JSON(500, "Can't get portfolio info")
	}

	httpResponse := &response{
		portfolio,
		stocks,
	}

	return c.JSON(200, httpResponse)
}

func (h *Handler) deletePortfolio(c echo.Context) error {
	log := logging.FromContext(h.ctx)
	token, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			c.Response().WriteHeader(http.StatusUnauthorized)
			c.Response().Write([]byte(`{"error": "not logined"}`))
			return nil
		}
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "can't parse cookie'"}`))
		return nil
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

	return c.Redirect(http.StatusOK, "http://localhost:3000/")
}
