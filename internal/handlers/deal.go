package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"net/http"
	"strconv"
)

func (h *Handler) getOneDeal(c echo.Context) error {
	log := logging.FromContext(h.ctx)

	dealId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}

	cookie, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.Infof("not logined %v", err)
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	d, err := h.services.DealService.GetOneDeal(context.Background(), cookie.Value, dealId)
	if err != nil {
		return c.JSON(500, "can't get deal info")
	}

	return c.JSON(200, d)
}

func (h *Handler) closeDeal(c echo.Context) error {
	log := logging.FromContext(h.ctx)

	dealId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}

	cookie, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.Infof("not logined %v", err)
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	err = h.services.DealService.CloseDeal(context.Background(), cookie.Value, dealId)
	if err != nil {
		return c.JSON(500, "can't close deal")
	}

	return nil
}

func (h *Handler) deleteDeal(c echo.Context) error {
	log := logging.FromContext(h.ctx)

	dealId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Info("can't parse URL params")
		return c.JSON(500, "can't parse URL params")
	}

	cookie, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			log.Infof("not logined %v", err)
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	err = h.services.DealService.DeleteDeal(context.Background(), cookie.Value, dealId)
	if err != nil {
		return c.JSON(500, "can't delete deal")
	}

	return nil
}



