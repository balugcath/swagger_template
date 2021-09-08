// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateHandlerFunc turns a function with the right signature into a update handler
type UpdateHandlerFunc func(UpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateHandlerFunc) Handle(params UpdateParams) middleware.Responder {
	return fn(params)
}

// UpdateHandler interface for that can handle valid update params
type UpdateHandler interface {
	Handle(UpdateParams) middleware.Responder
}

// NewUpdate creates a new http.Handler for the update operation
func NewUpdate(ctx *middleware.Context, handler UpdateHandler) *Update {
	return &Update{Context: ctx, Handler: handler}
}

/* Update swagger:route PUT /{id} todos update

Update update API

*/
type Update struct {
	Context *middleware.Context
	Handler UpdateHandler
}

func (o *Update) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
