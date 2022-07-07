// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetEventEndpointHandlerFunc turns a function with the right signature into a get event endpoint handler
type GetEventEndpointHandlerFunc func(GetEventEndpointParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEventEndpointHandlerFunc) Handle(params GetEventEndpointParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetEventEndpointHandler interface for that can handle valid get event endpoint params
type GetEventEndpointHandler interface {
	Handle(GetEventEndpointParams, interface{}) middleware.Responder
}

// NewGetEventEndpoint creates a new http.Handler for the get event endpoint operation
func NewGetEventEndpoint(ctx *middleware.Context, handler GetEventEndpointHandler) *GetEventEndpoint {
	return &GetEventEndpoint{Context: ctx, Handler: handler}
}

/*GetEventEndpoint swagger:route GET /di/v1/models/{model_id}/events/{event_type}/{endpoint_id} Events getEventEndpoint

Get endpoint description.

Get a specific endpoint description.

*/
type GetEventEndpoint struct {
	Context *middleware.Context
	Handler GetEventEndpointHandler
}

func (o *GetEventEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEventEndpointParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}