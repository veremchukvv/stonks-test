package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) createPortfolio(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")

}

func (h *Handler) getAllPortfolios(c echo.Context) error {
	//log := logging.FromContext(h.ctx)

	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) modifyPortfolio(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) getPortfolio(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}

func (h *Handler) deletePortfolio(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented yet")
}
