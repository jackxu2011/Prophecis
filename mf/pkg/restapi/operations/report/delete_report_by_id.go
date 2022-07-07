// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReportByIDHandlerFunc turns a function with the right signature into a delete report by Id handler
type DeleteReportByIDHandlerFunc func(DeleteReportByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReportByIDHandlerFunc) Handle(params DeleteReportByIDParams) middleware.Responder {
	return fn(params)
}

// DeleteReportByIDHandler interface for that can handle valid delete report by Id params
type DeleteReportByIDHandler interface {
	Handle(DeleteReportByIDParams) middleware.Responder
}

// NewDeleteReportByID creates a new http.Handler for the delete report by Id operation
func NewDeleteReportByID(ctx *middleware.Context, handler DeleteReportByIDHandler) *DeleteReportByID {
	return &DeleteReportByID{Context: ctx, Handler: handler}
}

/*DeleteReportByID swagger:route DELETE /mf/v1/report/{reportId} report deleteReportById

delete report by id

delete report

*/
type DeleteReportByID struct {
	Context *middleware.Context
	Handler DeleteReportByIDHandler
}

func (o *DeleteReportByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReportByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}