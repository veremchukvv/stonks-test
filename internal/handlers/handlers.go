package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/veremchukvv/stonks-test/internal/config"
	"github.com/veremchukvv/stonks-test/internal/service"

	"github.com/veremchukvv/stonks-test/pkg/logging"
)

type Handler struct {
	ctx      context.Context
	services *service.Services
	cfg      *config.Config
}

func NewHandlers(ctx context.Context, services *service.Services, cfg *config.Config) *Handler {
	return &Handler{
		ctx,
		services,
		cfg,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	log := logging.NewLogger(false, "console")
	log.Info(h.cfg.Server.CORS)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     h.cfg.Server.CORS,
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

	profile := e.Group("/profile")
	{
		profile.POST("/", h.createProfile)
		profile.PUT("/:id", h.modifyProfile)
		profile.GET("/:id", h.getProfile)
		profile.DELETE("/:id", h.deleteProfile)
	}

	portfolio := e.Group("/api/v1/portfolios")
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

	e.Static("/api/", "../../api")

	return e
}
