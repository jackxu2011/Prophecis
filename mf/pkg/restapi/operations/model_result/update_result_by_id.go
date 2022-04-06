// Code generated by go-swagger; DO NOT EDIT.

package model_result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateResultByIDHandlerFunc turns a function with the right signature into a update result by Id handler
type UpdateResultByIDHandlerFunc func(UpdateResultByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateResultByIDHandlerFunc) Handle(params UpdateResultByIDParams) middleware.Responder {
	return fn(params)
}

// UpdateResultByIDHandler interface for that can handle valid update result by Id params
type UpdateResultByIDHandler interface {
	Handle(UpdateResultByIDParams) middleware.Responder
}

// NewUpdateResultByID creates a new http.Handler for the update result by Id operation
func NewUpdateResultByID(ctx *middleware.Context, handler UpdateResultByIDHandler) *UpdateResultByID {
	return &UpdateResultByID{Context: ctx, Handler: handler}
}

/*UpdateResultByID swagger:route PUT /mf/v1/result/id/{resultId} model_result updateResultById

update result

update result

*/
type UpdateResultByID struct {
	Context *middleware.Context
	Handler UpdateResultByIDHandler
}

func (o *UpdateResultByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateResultByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
