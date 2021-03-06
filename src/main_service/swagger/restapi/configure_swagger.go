// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"

	"swagger_template/src/main_service/swagger/restapi/operations"
	"swagger_template/src/main_service/swagger/restapi/operations/todos"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name TodosAPI -inpkg

/* TodosAPI  */
type TodosAPI interface {
	Add(ctx context.Context, params todos.AddParams) middleware.Responder

	Destroy(ctx context.Context, params todos.DestroyParams) middleware.Responder

	Find(ctx context.Context, params todos.FindParams) middleware.Responder

	Update(ctx context.Context, params todos.UpdateParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	TodosAPI
	Logger func(string, ...interface{})
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// Authorizer is used to authorize a request after the Auth function was called using the "Auth*" functions
	// and the principal was stored in the context in the "AuthKey" context value.
	Authorizer func(*http.Request) error

	// Authenticator to use for all APIKey authentication
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// Authenticator to use for all Bearer authentication
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// Authenticator to use for all Basic authentication
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	h, _, err := HandlerAPI(c)
	return h, err
}

// HandlerAPI returns an http.Handler given the handler configuration
// and the corresponding *Swagger instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.SwaggerAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewSwaggerAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger

	if c.APIKeyAuthenticator != nil {
		api.APIKeyAuthenticator = c.APIKeyAuthenticator
	}
	if c.BasicAuthenticator != nil {
		api.BasicAuthenticator = c.BasicAuthenticator
	}
	if c.BearerAuthenticator != nil {
		api.BearerAuthenticator = c.BearerAuthenticator
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.TodosAddHandler = todos.AddHandlerFunc(func(params todos.AddParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.TodosAPI.Add(ctx, params)
	})
	api.TodosDestroyHandler = todos.DestroyHandlerFunc(func(params todos.DestroyParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.TodosAPI.Destroy(ctx, params)
	})
	api.TodosFindHandler = todos.FindHandlerFunc(func(params todos.FindParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.TodosAPI.Find(ctx, params)
	})
	api.TodosUpdateHandler = todos.UpdateHandlerFunc(func(params todos.UpdateParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		return c.TodosAPI.Update(ctx, params)
	})
	api.ServerShutdown = func() {}
	return api.Serve(c.InnerMiddleware), api, nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}

// authorizer is a helper function to implement the runtime.Authorizer interface.
type authorizer func(*http.Request) error

func (a authorizer) Authorize(req *http.Request, principal interface{}) error {
	if a == nil {
		return nil
	}
	ctx := storeAuth(req.Context(), principal)
	return a(req.WithContext(ctx))
}

func storeAuth(ctx context.Context, principal interface{}) context.Context {
	return context.WithValue(ctx, AuthKey, principal)
}
