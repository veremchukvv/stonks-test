package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/veremchukvv/stonks-test/internal/service"
)

//const apiVersion = "v1"

type Handler struct {
	ctx      context.Context
	services *service.Services
}

func NewHandlers(ctx context.Context, services *service.Services) *Handler {
	return &Handler{
		ctx,
		services,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowCredentials: true,
	}))

	auth := e.Group("/users")
	{
		auth.POST("/signup", h.signup)
		auth.GET("/user", h.user)
		auth.GET("/oauth/google", h.oauthGoogle)
		auth.GET("/oauth/vk", h.oauthVK)
		auth.GET("/callback/google", h.callbackGoogle)
		auth.GET("/callback/vk", h.callbackVK)
		auth.POST("/signin", h.signin)
		auth.POST("/signout", h.signout)
	}

	api := e.Group("/api/v1")
	{
		profile := api.Group("/profile")
		{
			profile.POST("/", h.createProfile)
			profile.PUT("/:id", h.modifyProfile)
			profile.GET("/:id", h.getProfile)
			profile.DELETE("/:id", h.deleteProfile)
		}
		portfolio := profile.Group("/portfolio")
		{
			portfolio.POST("/", h.createPortfolio)
			portfolio.GET("/", h.getAllPortfolios)
			portfolio.PUT("/:id", h.modifyPortfolio)
			portfolio.GET("/:id", h.getPortfolio)
			portfolio.DELETE("/:id", h.deletePortfolio)
		}
		//market := api.Group("/stocks")
		//{
		//	market.GET("/")
		//	market.GET("/:id")
		//	market.POST("/:id/deal")
		//}
	}

	return e
}
