package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/veremchukvv/stonks-test/internal/oauth"
	"log"
	"net/http"
)

func (h *Handler) signup(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented in MVP. Only signing with OAuth available")
}

func (h *Handler) oauthGoogle(c echo.Context) error {
	cfg := oauth.GetOauthConfig()
	state := oauth.GetRandomState()
	url := cfg.AuthCodeURL(state)
	log.Print(url)
	http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

func (h *Handler) oauthVK(c echo.Context) error {
	cfg := oauth.GetOauthConfig()
	state := oauth.GetRandomState()
	url := cfg.AuthCodeURL(state)
	log.Print(url)
	http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

func (h *Handler) callbackGoogle(c echo.Context) error {
	content, err := oauth.GetUserInfo(context.Background(), oauth.GetRandomState(), c.Request().FormValue("state"), c.Request().FormValue("code"), oauth.GetOauthConfig() )
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return nil
	}

	fmt.Fprintf(c.Response(), "Content: %s\n", content)
	return nil
}

func (h *Handler) callbackVK(c echo.Context) error {
	content, err := oauth.GetUserInfo(context.Background(), oauth.GetRandomState(), c.Request().FormValue("state"), c.Request().FormValue("code"), oauth.GetOauthConfig() )
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return nil
	}

	fmt.Fprintf(c.Response(), "Content: %s\n", content)
	return nil
}

func (h *Handler) signin(c echo.Context) error {
	return c.String(http.StatusOK, "signin OK")
}

func (h *Handler) signout(c echo.Context) error {
	return c.String(http.StatusOK, "signout OK")
}