package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/oauth"
	"github.com/veremchukvv/stonks-test/internal/repository/pg"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"net/http"
)

func (h *Handler) signup(c echo.Context) error {
	log := logging.FromContext(h.ctx)
	var user models.User
	err := c.Bind(&user)
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
	content, err := oauth.GetUserInfo(context.Background(), oauth.GetRandomState(), c.Request().FormValue("state"),
		c.Request().FormValue("code"), oauth.GetOauthConfig())
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return nil
	}
	fmt.Fprintf(c.Response(), "Content: %s\n", content)
	return nil
}

func (h *Handler) callbackVK(c echo.Context) error {

	type VKContent struct {
		First_name string `json:"first_name"`
		Id         int    `json:"id"`
		Last_name  string `json:"last_name"`
	}

	type VKResponse struct {
		Response []VKContent `json:"response"`
	}

	log := logging.FromContext(h.ctx)

	content, err := oauth.GetUserInfo(h.ctx, oauth.GetRandomState(), c.Request().FormValue("state"),
		c.Request().FormValue("code"), oauth.GetOauthConfig())
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return nil
	}

	var input VKResponse
	err = json.Unmarshal(content, &input)
	if err != nil {
		log.Info(err)
	}

	_, err = h.services.UserService.GetVKUserByID(h.ctx, input.Response[0].Id)
	if err != nil {
		if errors.Is(err, pg.ErrVkUserNotFound) {
			log.Info(err)
			log.Info("trying to create new VK user")

			newVKUser := &models.VKUser{
				VKId:     input.Response[0].Id,
				Name:     input.Response[0].First_name,
				Lastname: input.Response[0].Last_name,
			}

			_, err := h.services.UserService.CreateVKUser(h.ctx, newVKUser)
			if err != nil {
				log.Errorf("Can't create VK user: %v", err)
				return err
			}
			log.Infof("Created VK user with id: %d", newVKUser.VKId)
		} else {
			log.Errorf("Other VK error: %v", err)
			return err
		}
	}

	token, err := h.services.UserService.GenerateVKToken(h.ctx, input.Response[0].Id)

	c.SetCookie(&http.Cookie{Name: "jwt", Value: token})

	log.Infof("Successfull login for VK user: %d", input.Response[0].Id)

	return c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/")
}

func (h *Handler) signout(c echo.Context) error {
	return c.String(http.StatusOK, "signout OK")
}
