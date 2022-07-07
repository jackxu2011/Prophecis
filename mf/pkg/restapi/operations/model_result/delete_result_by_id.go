// Code generated by go-swagger; DO NOT EDIT.

package model_result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteResultByIDHandlerFunc turns a function with the right signature into a delete result by Id handler
type DeleteResultByIDHandlerFunc func(DeleteResultByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteResultByIDHandlerFunc) Handle(params DeleteResultByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteResultByIDHandler interface for that can handle valid delete result by Id params
type DeleteResultByIDHandler interface {
	Handle(DeleteResultByIDParams) middleware.Responder
}

// NewDeleteResultByID creates a new http.Handler for the delete result by Id operation
func NewDeleteResultByID(ctx *middleware.Context, handler DeleteResultByIDHandler) *DeleteResultByID {
	return &DeleteResultByID{Context: ctx, Handler: handler}
}

/*DeleteResultByID swagger:route DELETE /mf/v1/result/id/{resultId} model_result deleteResultById

delete result by id

delete result

*/
type DeleteResultByID struct {
	Context *middleware.Context
	Handler DeleteResultByIDHandler
}

func (o *DeleteResultByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteResultByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}