package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) createProfile(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) modifyProfile(c echo.Context) {

}

func (h *Handler) getProfile(c echo.Context) {

}

func (h *Handler) deleteProfile(c echo.Context) {

}
