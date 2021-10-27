package handlers

import (
	"github.com/labstack/echo/v4"
)


//const apiVersion = "v1"

type Handler struct {
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	auth := e.Group("/users")
	{
		auth.POST("/signup", h.signup)
		auth.POST("/oauth", h.oauth)
		auth.POST("/callback", h.callback)
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
