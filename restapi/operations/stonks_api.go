// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/veremchukvv/stonks-test/restapi/operations/portfolios"
	"github.com/veremchukvv/stonks-test/restapi/operations/profile"
	"github.com/veremchukvv/stonks-test/restapi/operations/stocks"
)

// NewStonksAPI creates a new Stonks instance
func NewStonksAPI(spec *loads.Document) *StonksAPI {
	return &StonksAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ProfileGetProfileHandler: profile.GetProfileHandlerFunc(func(params profile.GetProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.GetProfile has not yet been implemented")
		}),
		StocksPostStocksHandler: stocks.PostStocksHandlerFunc(func(params stocks.PostStocksParams) middleware.Responder {
			return middleware.NotImplemented("operation stocks.PostStocks has not yet been implemented")
		}),
		PortfoliosAddNewPortfolioHandler: portfolios.AddNewPortfolioHandlerFunc(func(params portfolios.AddNewPortfolioParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.AddNewPortfolio has not yet been implemented")
		}),
		ProfileCreateProfileHandler: profile.CreateProfileHandlerFunc(func(params profile.CreateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.CreateProfile has not yet been implemented")
		}),
		PortfoliosDeletePortfolioHandler: portfolios.DeletePortfolioHandlerFunc(func(params portfolios.DeletePortfolioParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.DeletePortfolio has not yet been implemented")
		}),
		PortfoliosGetPortfolioByIDHandler: portfolios.GetPortfolioByIDHandlerFunc(func(params portfolios.GetPortfolioByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.GetPortfolioByID has not yet been implemented")
		}),
		StocksGetStocksHandler: stocks.GetStocksHandlerFunc(func(params stocks.GetStocksParams) middleware.Responder {
			return middleware.NotImplemented("operation stocks.GetStocks has not yet been implemented")
		}),
		ProfileLoginUserHandler: profile.LoginUserHandlerFunc(func(params profile.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.LoginUser has not yet been implemented")
		}),
		ProfileLogoutUserHandler: profile.LogoutUserHandlerFunc(func(params profile.LogoutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.LogoutUser has not yet been implemented")
		}),
		PortfoliosUpdatePortfolioHandler: portfolios.UpdatePortfolioHandlerFunc(func(params portfolios.UpdatePortfolioParams) middleware.Responder {
			return middleware.NotImplemented("operation portfolios.UpdatePortfolio has not yet been implemented")
		}),
		ProfileUpdateProfileHandler: profile.UpdateProfileHandlerFunc(func(params profile.UpdateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.UpdateProfile has not yet been implemented")
		}),
	}
}

/*StonksAPI This is a fun Stonks service. You can find out more about us at [https://github.com/veremchukvv/stonks-test](https://github.com/veremchukvv/stonks-test). */
type StonksAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// ProfileGetProfileHandler sets the operation handler for the get profile operation
	ProfileGetProfileHandler profile.GetProfileHandler
	// StocksPostStocksHandler sets the operation handler for the post stocks operation
	StocksPostStocksHandler stocks.PostStocksHandler
	// PortfoliosAddNewPortfolioHandler sets the operation handler for the add new portfolio operation
	PortfoliosAddNewPortfolioHandler portfolios.AddNewPortfolioHandler
	// ProfileCreateProfileHandler sets the operation handler for the create profile operation
	ProfileCreateProfileHandler profile.CreateProfileHandler
	// PortfoliosDeletePortfolioHandler sets the operation handler for the delete portfolio operation
	PortfoliosDeletePortfolioHandler portfolios.DeletePortfolioHandler
	// PortfoliosGetPortfolioByIDHandler sets the operation handler for the get portfolio by Id operation
	PortfoliosGetPortfolioByIDHandler portfolios.GetPortfolioByIDHandler
	// StocksGetStocksHandler sets the operation handler for the get stocks operation
	StocksGetStocksHandler stocks.GetStocksHandler
	// ProfileLoginUserHandler sets the operation handler for the login user operation
	ProfileLoginUserHandler profile.LoginUserHandler
	// ProfileLogoutUserHandler sets the operation handler for the logout user operation
	ProfileLogoutUserHandler profile.LogoutUserHandler
	// PortfoliosUpdatePortfolioHandler sets the operation handler for the update portfolio operation
	PortfoliosUpdatePortfolioHandler portfolios.UpdatePortfolioHandler
	// ProfileUpdateProfileHandler sets the operation handler for the update profile operation
	ProfileUpdateProfileHandler profile.UpdateProfileHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *StonksAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *StonksAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *StonksAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *StonksAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *StonksAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *StonksAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *StonksAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *StonksAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *StonksAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the StonksAPI
func (o *StonksAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ProfileGetProfileHandler == nil {
		unregistered = append(unregistered, "profile.GetProfileHandler")
	}
	if o.StocksPostStocksHandler == nil {
		unregistered = append(unregistered, "stocks.PostStocksHandler")
	}
	if o.PortfoliosAddNewPortfolioHandler == nil {
		unregistered = append(unregistered, "portfolios.AddNewPortfolioHandler")
	}
	if o.ProfileCreateProfileHandler == nil {
		unregistered = append(unregistered, "profile.CreateProfileHandler")
	}
	if o.PortfoliosDeletePortfolioHandler == nil {
		unregistered = append(unregistered, "portfolios.DeletePortfolioHandler")
	}
	if o.PortfoliosGetPortfolioByIDHandler == nil {
		unregistered = append(unregistered, "portfolios.GetPortfolioByIDHandler")
	}
	if o.StocksGetStocksHandler == nil {
		unregistered = append(unregistered, "stocks.GetStocksHandler")
	}
	if o.ProfileLoginUserHandler == nil {
		unregistered = append(unregistered, "profile.LoginUserHandler")
	}
	if o.ProfileLogoutUserHandler == nil {
		unregistered = append(unregistered, "profile.LogoutUserHandler")
	}
	if o.PortfoliosUpdatePortfolioHandler == nil {
		unregistered = append(unregistered, "portfolios.UpdatePortfolioHandler")
	}
	if o.ProfileUpdateProfileHandler == nil {
		unregistered = append(unregistered, "profile.UpdateProfileHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *StonksAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *StonksAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *StonksAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *StonksAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *StonksAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *StonksAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the stonks API
func (o *StonksAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *StonksAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profile"] = profile.NewGetProfile(o.context, o.ProfileGetProfileHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/stocks"] = stocks.NewPostStocks(o.context, o.StocksPostStocksHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/portfolios"] = portfolios.NewAddNewPortfolio(o.context, o.PortfoliosAddNewPortfolioHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/profile"] = profile.NewCreateProfile(o.context, o.ProfileCreateProfileHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/portfolios/{portfolioId}"] = portfolios.NewDeletePortfolio(o.context, o.PortfoliosDeletePortfolioHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/portfolios/{portfolioId}"] = portfolios.NewGetPortfolioByID(o.context, o.PortfoliosGetPortfolioByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/stocks"] = stocks.NewGetStocks(o.context, o.StocksGetStocksHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = profile.NewLoginUser(o.context, o.ProfileLoginUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/logout"] = profile.NewLogoutUser(o.context, o.ProfileLogoutUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/portfolios"] = portfolios.NewUpdatePortfolio(o.context, o.PortfoliosUpdatePortfolioHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/profile"] = profile.NewUpdateProfile(o.context, o.ProfileUpdateProfileHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *StonksAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *StonksAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *StonksAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *StonksAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *StonksAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
