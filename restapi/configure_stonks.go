// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/veremchukvv/stonks-test/restapi/operations"
	"github.com/veremchukvv/stonks-test/restapi/operations/portfolios"
	"github.com/veremchukvv/stonks-test/restapi/operations/profile"
	"github.com/veremchukvv/stonks-test/restapi/operations/stocks"
)

//go:generate swagger generate server --target ../../stonks-test --name Stonks --spec ../api/swagger.yml --principal interface{}

func configureFlags(api *operations.StonksAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.StonksAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ProfileGetProfileHandler == nil {
		api.ProfileGetProfileHandler = profile.GetProfileHandlerFunc(func(params profile.GetProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.GetProfile has not yet been implemented")
		})
	}
	if api.StocksPostStocksHandler == nil {
		api.StocksPostStocksHandler = stocks.PostStocksHandlerFunc(func(params stocks.PostStocksParams) middleware.Responder {
			return middleware.NotImplemented("operation stocks.PostStocks has not yet been implemented")
		})
	}
	if api.PortfoliosAddNewPortfolioHandler == nil {
		api.PortfoliosAddNewPortfolioHandler = portfolios.AddNewPortfolioHandlerFunc(func(params portfolios.AddNewPortfolioParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.AddNewPortfolio has not yet been implemented")
		})
	}
	if api.ProfileCreateProfileHandler == nil {
		api.ProfileCreateProfileHandler = profile.CreateProfileHandlerFunc(func(params profile.CreateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.CreateProfile has not yet been implemented")
		})
	}
	if api.PortfoliosDeletePortfolioHandler == nil {
		api.PortfoliosDeletePortfolioHandler = portfolios.DeletePortfolioHandlerFunc(func(params portfolios.DeletePortfolioParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.DeletePortfolio has not yet been implemented")
		})
	}
	if api.PortfoliosGetPortfolioByIDHandler == nil {
		api.PortfoliosGetPortfolioByIDHandler = portfolios.GetPortfolioByIDHandlerFunc(func(params portfolios.GetPortfolioByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.GetPortfolioByID has not yet been implemented")
		})
	}
	if api.StocksGetStocksHandler == nil {
		api.StocksGetStocksHandler = stocks.GetStocksHandlerFunc(func(params stocks.GetStocksParams) middleware.Responder {
			return middleware.NotImplemented("operation stocks.GetStocks has not yet been implemented")
		})
	}
	if api.ProfileLoginUserHandler == nil {
		api.ProfileLoginUserHandler = profile.LoginUserHandlerFunc(func(params profile.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.LoginUser has not yet been implemented")
		})
	}
	if api.ProfileLogoutUserHandler == nil {
		api.ProfileLogoutUserHandler = profile.LogoutUserHandlerFunc(func(params profile.LogoutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.LogoutUser has not yet been implemented")
		})
	}
	if api.PortfoliosUpdatePortfolioHandler == nil {
		api.PortfoliosUpdatePortfolioHandler = portfolios.UpdatePortfolioHandlerFunc(func(params portfolios.UpdatePortfolioParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.UpdatePortfolio has not yet been implemented")
		})
	}
	if api.ProfileUpdateProfileHandler == nil {
		api.ProfileUpdateProfileHandler = profile.UpdateProfileHandlerFunc(func(params profile.UpdateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.UpdateProfile has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
