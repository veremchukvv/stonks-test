package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/veremchukvv/stonks-test/internal/models"
	"github.com/veremchukvv/stonks-test/internal/oauth"
	"github.com/veremchukvv/stonks-test/internal/repository/pg"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"golang.org/x/oauth2"
)

func (h *Handler) signup(c echo.Context) error {
	log := logging.FromContext(h.ctx)
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(500, "Unmarshalling data error")
	}
	createdUser, err := h.services.UserService.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(500, "Creating User error")
	}
	log.Infof("Created user with ID: %d", createdUser.Id)
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *Handler) signin(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(500, "Error on parsing request")
	}
	token, err := h.services.UserService.GenerateToken(h.ctx, user.Email, user.Password)
	if err != nil {
		return c.JSON(401, "Authentication failure")
	}
	c.SetCookie(&http.Cookie{Name: "jwt", Value: token, HttpOnly: true, Path: "/"})
	return c.JSON(http.StatusOK, "Login successful")
}

func (h *Handler) user(c echo.Context) error {
	log := logging.FromContext(h.ctx)
	cookie, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	u, err := h.services.UserService.GetUser(c.Request().Context(), cookie.Value)
	if err != nil {
		log.Infof("Error on query user: %v", err)
		return err
	}
	log.Info(u)
	if u != nil {
		return c.JSON(200, u)
	}
	return nil
}

func (h *Handler) oauthGoogle(c echo.Context) error {
	cfg := oauth.GetOauthGoogleConfig()
	state := oauth.GetRandomState()
	url := cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

func (h *Handler) oauthVK(c echo.Context) error {
	cfg := oauth.GetOauthVKConfig()
	state := oauth.GetRandomState()
	url := cfg.AuthCodeURL(state)
	http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

func (h *Handler) updateUser(c echo.Context) error {
	log := logging.FromContext(h.ctx)

	var u models.User

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(500, "Unmarshalling data error")
	}

	cookie, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	uu, err := h.services.UserService.UpdateUser(c.Request().Context(), &u, cookie.Value)
	if err != nil {
		log.Infof("Error on update user: %v", err)
		return err
	}
	if uu != nil {
		return c.JSON(200, uu)
	}
	return nil
}

func (h *Handler) deleteUser(c echo.Context) error {
	cookie, err := c.Request().Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(401, "not logined")
		}
		return c.JSON(500, "can't parse cookie")
	}

	err = h.services.UserService.DeleteUser(c.Request().Context(), cookie.Value)
	if err != nil {
		return c.JSON(500, "error on delete user")
	}
	c.SetCookie(&http.Cookie{Name: "jwt", Value: "", HttpOnly: true, Path: "/", Expires: time.Now().Add(-time.Hour)})
	// return c.Redirect(http.StatusOK, "http://localhost:3000/")
	return c.JSON(200, "User deleted")
}

func (h *Handler) callbackGoogle(c echo.Context) error {
	type GoogleContent struct {
		ID        string `json:"id"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Email     string `json:"email"`
	}

	log := logging.FromContext(h.ctx)

	var err error
	content, err := oauth.GetUserGoogleInfo(context.Background(), oauth.GetRandomState(), c.Request().FormValue("state"),
		c.Request().FormValue("code"), oauth.GetOauthGoogleConfig())
	if err != nil {
		log.Info(err)
	}

	var input GoogleContent
	err = json.Unmarshal(content, &input)
	if err != nil {
		log.Info(err)
	}

	var gu *models.User
	gu, err = h.services.UserService.GetGoogleUserByID(h.ctx, input.ID)
	if err != nil {
		if errors.Is(err, pg.ErrGoogleUserNotFound) {
			log.Info("trying to create new Google user")

			newGoogleUser := &models.User{
				GoogleId: input.ID,
				Name:     input.FirstName,
				Lastname: input.LastName,
				Email:    input.Email,
			}

			var ngu *models.User
			ngu, err = h.services.UserService.CreateGoogleUser(h.ctx, newGoogleUser)
			if err != nil {
				log.Errorf("Can't create Google user: %v", err)
				return err
			}
			log.Infof("Created Google user with id: %d", ngu.Id)

			var token string
			token, err = h.services.UserService.GenerateGoogleToken(ngu.Id)
			if err != nil {
				log.Info("error on generating google token")
			}
			c.SetCookie(&http.Cookie{Name: "jwt", Value: token, HttpOnly: true, Path: "/"})
			log.Infof("Successful login for Google user: %d", ngu.Id)
		} else {
			log.Errorf("Other Google error: %v", err)
			return err
		}
	}

	token, err := h.services.UserService.GenerateGoogleToken(gu.Id)
	if err != nil {
		log.Info("error on generating google token")
	}

	c.SetCookie(&http.Cookie{Name: "jwt", Value: token, HttpOnly: true, Path: "/"})

	log.Infof("Successful login for Google user: %d", gu.Id)

	return c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/")
}

func (h *Handler) callbackVK(c echo.Context) error {
	type VKContent struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	type VKResponse struct {
		Response []VKContent `json:"response"`
	}

	log := logging.FromContext(h.ctx)

	var err error
	content, err := oauth.GetUserVKInfo(h.ctx, oauth.GetRandomState(), c.Request().FormValue("state"),
		c.Request().FormValue("code"), oauth.GetOauthVKConfig())
	if err != nil {
		http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return nil
	}

	var input VKResponse
	err = json.Unmarshal(content, &input)
	if err != nil {
		log.Info(err)
	}

	_, err = h.services.UserService.GetVKUserByID(h.ctx, input.Response[0].ID)
	if err != nil {
		if errors.Is(err, pg.ErrVkUserNotFound) {
			log.Info(err)
			log.Info("trying to create new VK user")

			newVKUser := &models.User{
				Id:       input.Response[0].ID,
				Name:     input.Response[0].FirstName,
				Lastname: input.Response[0].LastName,
			}

			_, err = h.services.UserService.CreateVKUser(h.ctx, newVKUser)
			if err != nil {
				log.Errorf("Can't create VK user: %v", err)
				return err
			}
			log.Infof("Created VK user with id: %d", newVKUser.Id)
		} else {
			log.Errorf("Other VK error: %v", err)
			return err
		}
	}

	token, err := h.services.UserService.GenerateVKToken(input.Response[0].ID)
	if err != nil {
		log.Errorf("Can't generate VK token: %v", err)
		return err
	}

	c.SetCookie(&http.Cookie{Name: "jwt", Value: token, HttpOnly: true, Path: "/"})

	log.Infof("Successful login for VK user: %d", input.Response[0].ID)

	return c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/")
}

func (h *Handler) signout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "jwt", Value: "", HttpOnly: true, Path: "/", Expires: time.Now().Add(-time.Hour)})
	return c.JSON(200, "See you next time!")
	// return c.Redirect(http.StatusOK, "http://localhost:3000/")
}
