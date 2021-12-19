package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/veremchukvv/stonks-test/internal/service"
)

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
		auth.PATCH("/update", h.updateUser)
		auth.DELETE("/delete", h.deleteUser)
		auth.GET("/oauth/google", h.oauthGoogle)
		auth.GET("/oauth/vk", h.oauthVK)
		auth.GET("/callback/google", h.callbackGoogle)
		auth.GET("/callback/vk", h.callbackVK)
		auth.POST("/signin", h.signin)
		auth.POST("/signout", h.signout)
	}

	//api := e.Group("/api/v1")
	//{
	//	profile := api.Group("/profile")
	//	{
	//		profile.POST("/", h.createProfile)
	//		profile.PUT("/:id", h.modifyProfile)
	//		profile.GET("/:id", h.getProfile)
	//		profile.DELETE("/:id", h.deleteProfile)
	//	}

	portfolio := e.Group("/api/v1/portfolio")
	{
		portfolio.POST("/", h.createPortfolio)
		portfolio.GET("/", h.getAllPortfolios)
		portfolio.PATCH("/:id", h.modifyPortfolio)
		portfolio.GET("/:id", h.getPortfolioDeals)
		portfolio.GET("/closed/:id", h.getPortfolioClosedDeals)
		portfolio.DELETE("/:id", h.deletePortfolio)
	}

	deal := e.Group("/api/v1/deals")
	{
		deal.GET("/:id", h.getOneDeal)
		deal.POST("/:id", h.closeDeal)
		deal.DELETE("/:id", h.deleteDeal)
	}

	closedDeal := e.Group("/api/v1/closed")
	{
		closedDeal.GET("/:id", h.getOneClosedDeal)
		closedDeal.DELETE("/:id", h.deleteClosedDeal)
	}

	market := e.Group("/api/v1/stockmarket")
	{
		market.GET("/", h.getAllStocks)
		market.GET("/:id", h.getOneStock)
		market.POST("/deal", h.makeDeal)
		market.DELETE("/deal/:id", h.deleteDeal)

	}

	return e
}
