package handler

import "github.com/labstack/echo/v4"

type Handler struct {

}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	auth := e.Group("/user")
	{
		auth.POST("/login", h.login)
		auth.POST("/logout", h.logout)
	}

	/*api := e.Group("/api/v1")
	{
		profile := api.Group("/profile")
		{
			profile.POST("/")
			profile.PUT("/:id")
			profile.GET("/:id")
			profile.DELETE("/:id")
		}
		portfolio := profile.Group("/portfolio")
		{
			portfolio.POST("/")
			portfolio.PUT("/:id")
			portfolio.GET("/:id")
			portfolio.DELETE("/:id")
		}
		market := api.Group("/stocks")
		{
			market.GET("/")
			market.GET("/:id")
			market.POST("/:id/deal")
		}
	}*/

	return e
}
