package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/oauth"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"net/http"
	"time"
)

func (h *Handler) signup(c echo.Context) error {
	log := logging.FromContext(h.ctx)
	var user models.User
	err := c.Bind(&user)
	log.Error(err)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "Unmarshalling data error"}`))
		return nil
	}
	createdUser, err := h.services.UserService.CreateUser(c.Request().Context(), &user)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "Creating User error"}`))
		return nil
	}
	log.Infof("Created user %d", createdUser.Id)
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *Handler) signin(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "Unmarshalling data error"}`))
		return nil
	}
	token, err := h.services.UserService.GenerateToken(h.ctx, user.Email, user.Password)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write([]byte(`{"error": "Authentication failure"}`))
		return nil
	}
	c.SetCookie(&http.Cookie{Name: "jwt", Value: token})
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) oauthGoogle(c echo.Context) error {
	cfg := oauth.GetOauthConfig()
	state := oauth.GetRandomState()
	url := cfg.AuthCodeURL(state)
	http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

func (h *Handler) oauthVK(c echo.Context) error {
	cfg := oauth.GetOauthConfig()
	state := oauth.GetRandomState()
	url := cfg.AuthCodeURL(state)
	http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

func (h *Handler) callbackGoogle(c echo.Context) error {
	content, err := oauth.GetUserInfo(context.Background(), oauth.GetRandomState(), c.Request().FormValue("state"), c.Request().FormValue("code"), oauth.GetOauthConfig())
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return nil
	}
	fmt.Fprintf(c.Response(), "Content: %s\n", content)
	return nil
}

func (h *Handler) callbackVK(c echo.Context) error {
	content, err := oauth.GetUserInfo(context.Background(), oauth.GetRandomState(), c.Request().FormValue("state"), c.Request().FormValue("code"), oauth.GetOauthConfig())
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return nil
	}
	fmt.Fprintf(c.Response(), "Content: %s\n", content)
	c.SetCookie(&http.Cookie{Name: "lol", Value: "lollol", Expires: time.Now().Add(10 * time.Minute)})

	return c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/")
}

func (h *Handler) signout(c echo.Context) error {
	return c.String(http.StatusOK, "signout OK")
}
